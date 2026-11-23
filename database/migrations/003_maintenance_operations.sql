-- ==============================================================================
-- Migration 003: Modul 3 — Breakdown, Preventive Maintenance, Inspection & Spare Part
-- Track 4 (PIC: Member #4 / Kolaborasi Tim) | URS: 4.1 – 4.12
-- ==============================================================================

-- URS 4.4 & 4.5: Rencana Tahunan & Jadwal Preventive Maintenance
CREATE TABLE IF NOT EXISTS pm_plans (
    id                      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asset_id                UUID NOT NULL REFERENCES assets(id) ON DELETE RESTRICT,
    plan_year               SMALLINT NOT NULL,
    title                   VARCHAR(200) NOT NULL,
    schedule_basis          VARCHAR(24) NOT NULL CHECK (schedule_basis IN ('CALENDAR', 'OPERATING_HOURS')),
    interval_value          NUMERIC(10, 2) NOT NULL,
    interval_unit           VARCHAR(20) NOT NULL CHECK (interval_unit IN ('DAYS', 'WEEKS', 'MONTHS', 'HOURS')),
    work_instruction        TEXT,
    safety_requirements     TEXT,
    checklist_template      JSONB NOT NULL DEFAULT '[]'::jsonb,
    standard_materials      JSONB NOT NULL DEFAULT '[]'::jsonb, -- [{part_code, qty, unit}]
    is_active               BOOLEAN NOT NULL DEFAULT TRUE,
    created_by              UUID REFERENCES users(id),
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS pm_schedules (
    id                    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    pm_plan_id            UUID NOT NULL REFERENCES pm_plans(id) ON DELETE RESTRICT,
    asset_id              UUID NOT NULL REFERENCES assets(id) ON DELETE RESTRICT,
    original_planned_date DATE NOT NULL,
    planned_date          DATE NOT NULL,
    target_op_hours       NUMERIC(12, 2),
    reschedule_reason     TEXT,
    rescheduled_by        UUID REFERENCES users(id),
    rescheduled_at        TIMESTAMPTZ,
    status                VARCHAR(24) NOT NULL DEFAULT 'SCHEDULED' CHECK (status IN (
                              'SCHEDULED', 'IN_PROGRESS', 'COMPLETED', 'RESCHEDULED', 'SKIPPED'
                          )),
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- URS 4.6: Template & Parameter Standar Inspeksi (Domain input: Lovi - Teknik Elektro)
CREATE TABLE IF NOT EXISTS inspection_templates (
    id               UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asset_id         UUID REFERENCES assets(id),
    asset_type       VARCHAR(64),
    inspection_type  VARCHAR(80) NOT NULL,
    frequency        VARCHAR(32) NOT NULL CHECK (frequency IN ('DAILY', 'WEEKLY', 'MONTHLY', 'QUARTERLY', 'CUSTOM')),
    -- [{param_name, unit, normal_min, normal_max, warning_min, warning_max, critical_min, critical_max}]
    parameters       JSONB NOT NULL DEFAULT '[]'::jsonb,
    is_active        BOOLEAN NOT NULL DEFAULT TRUE,
    created_by       UUID REFERENCES users(id),
    created_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- URS 4.1: Pelaporan Breakdown
CREATE TABLE IF NOT EXISTS breakdown_reports (
    id                   UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    report_no            VARCHAR(64) UNIQUE NOT NULL,
    asset_id             UUID NOT NULL REFERENCES assets(id) ON DELETE RESTRICT,
    work_request_ref     VARCHAR(64),
    symptom              TEXT NOT NULL,
    problem_category     VARCHAR(100),
    operating_condition  VARCHAR(32) NOT NULL CHECK (operating_condition IN (
                             'STOPPED', 'LIMITED_OPERATION', 'STILL_RUNNING'
                         )),
    impact_description   TEXT,
    priority             VARCHAR(16) NOT NULL CHECK (priority IN ('EMERGENCY', 'HIGH', 'MEDIUM', 'LOW')),
    is_safety_hazard     BOOLEAN NOT NULL DEFAULT FALSE,
    k3_escalated_at      TIMESTAMPTZ,
    reported_by          UUID NOT NULL REFERENCES users(id),
    occurred_at          TIMESTAMPTZ NOT NULL, -- Waktu kejadian aktual (ALCOA Contemporaneous)
    recorded_at          TIMESTAMPTZ NOT NULL DEFAULT NOW(), -- Waktu input sistem
    downtime_started_at  TIMESTAMPTZ,
    status               VARCHAR(24) NOT NULL DEFAULT 'OPEN' CHECK (status IN ('OPEN', 'IN_PROGRESS', 'RESOLVED', 'CANCELLED'))
);

-- URS 4.2, 4.3, 4.10, 4.11, 4.12: Work Orders (Terpadu untuk Breakdown, PM, Inspection)
CREATE TABLE IF NOT EXISTS work_orders (
    id                      UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    wo_number               VARCHAR(64) UNIQUE NOT NULL,
    wo_type                 VARCHAR(24) NOT NULL CHECK (wo_type IN ('BREAKDOWN', 'PREVENTIVE', 'INSPECTION', 'CORRECTIVE')),
    asset_id                UUID NOT NULL REFERENCES assets(id) ON DELETE RESTRICT,
    division_id             UUID NOT NULL REFERENCES divisions(id),
    breakdown_report_id     UUID REFERENCES breakdown_reports(id),
    pm_schedule_id          UUID REFERENCES pm_schedules(id),
    inspection_template_id  UUID REFERENCES inspection_templates(id),
    parent_finding_ref      UUID,
    external_wo_ref         VARCHAR(80),
    -- URS 4.10b: Status WO terpisah dari kondisi aset dan hasil pekerjaan
    status                  VARCHAR(32) NOT NULL DEFAULT 'OPEN' CHECK (status IN (
                                'OPEN', 'IN_PROGRESS', 'WAITING_PART',
                                'PENDING_VERIFICATION', 'REWORK', 'CLOSED', 'CANCELLED'
                            )),
    work_result             VARCHAR(32) CHECK (work_result IN (
                                'COMPLETED', 'PARTIALLY_COMPLETED', 'NEEDS_FOLLOW_UP'
                            )),
    priority                VARCHAR(16) NOT NULL DEFAULT 'MEDIUM' CHECK (priority IN ('EMERGENCY', 'HIGH', 'MEDIUM', 'LOW')),
    assigned_to             UUID REFERENCES users(id),
    assigned_by             UUID REFERENCES users(id),
    predicted_completion_at TIMESTAMPTZ,
    -- Diagnosis & tindakan (URS 4.3)
    diagnosis               TEXT,
    failure_category        VARCHAR(100),
    root_cause              TEXT,
    action_taken            TEXT,
    checklist_results       JSONB DEFAULT '[]'::jsonb,
    inspection_readings     JSONB DEFAULT '[]'::jsonb,
    -- Pencatatan waktu & downtime (URS 4.3b-c & ALCOA)
    occurred_at             TIMESTAMPTZ,
    work_started_at         TIMESTAMPTZ,
    waiting_minutes         INTEGER NOT NULL DEFAULT 0,
    active_repair_minutes   INTEGER NOT NULL DEFAULT 0,
    asset_ready_at          TIMESTAMPTZ,
    downtime_type           VARCHAR(20) CHECK (downtime_type IN ('PLANNED', 'UNPLANNED')),
    labor_cost              NUMERIC(14, 2) NOT NULL DEFAULT 0,
    contractor_cost         NUMERIC(14, 2) NOT NULL DEFAULT 0,
    other_cost              NUMERIC(14, 2) NOT NULL DEFAULT 0,
    -- Kontingensi input manual (URS 4.1 BCP)
    manual_log_ref          VARCHAR(100),
    operational_date        DATE NOT NULL DEFAULT CURRENT_DATE,
    origin_shift            VARCHAR(20),
    -- Verifikasi & penutupan (URS 4.11 & 4.12)
    submitted_by            UUID REFERENCES users(id),
    submitted_at            TIMESTAMPTZ,
    verified_by             UUID REFERENCES users(id),
    verified_at             TIMESTAMPTZ,
    verifier_notes          TEXT,
    closed_at               TIMESTAMPTZ,
    version                 INTEGER NOT NULL DEFAULT 1,
    created_by              UUID NOT NULL REFERENCES users(id),
    created_at              TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at              TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_wo_asset ON work_orders (asset_id);
CREATE INDEX IF NOT EXISTS idx_wo_status ON work_orders (status);
CREATE INDEX IF NOT EXISTS idx_wo_type ON work_orders (wo_type);
CREATE INDEX IF NOT EXISTS idx_wo_division ON work_orders (division_id);
CREATE INDEX IF NOT EXISTS idx_wo_operational_date ON work_orders (operational_date);

-- URS 4.2c: Riwayat perubahan prediksi selesai breakdown
CREATE TABLE IF NOT EXISTS wo_eta_history (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    work_order_id     UUID NOT NULL REFERENCES work_orders(id) ON DELETE CASCADE,
    previous_eta      TIMESTAMPTZ,
    new_eta           TIMESTAMPTZ NOT NULL,
    reason            TEXT NOT NULL,
    updated_by        UUID NOT NULL REFERENCES users(id),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- URS 4.10c & ALCOA 1.1/1.3: Catatan kerja & serah-terima lintas shift
CREATE TABLE IF NOT EXISTS wo_shift_logs (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    work_order_id     UUID NOT NULL REFERENCES work_orders(id) ON DELETE CASCADE,
    operational_date  DATE NOT NULL,
    shift_code        VARCHAR(20) NOT NULL,
    performed_by      UUID NOT NULL REFERENCES users(id),
    recorded_by       UUID NOT NULL REFERENCES users(id), -- ALCOA 1.1: pisahkan pelaksana vs penginput
    action_summary    TEXT NOT NULL,
    blockers          TEXT,
    waiting_parts_note TEXT,
    next_steps        TEXT,
    started_at        TIMESTAMPTZ NOT NULL,
    ended_at          TIMESTAMPTZ,
    recorded_at       TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- URS 4.8 & 4.9: Referensi Spare Part (sinkron dari sistem persediaan perusahaan)
CREATE TABLE IF NOT EXISTS spare_parts (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    part_code         VARCHAR(80) UNIQUE NOT NULL,
    description       VARCHAR(250) NOT NULL,
    unit              VARCHAR(32) NOT NULL,
    warehouse_loc     VARCHAR(120),
    available_stock   NUMERIC(14, 3) NOT NULL DEFAULT 0,
    reorder_point     NUMERIC(14, 3),
    unit_cost         NUMERIC(14, 2) NOT NULL DEFAULT 0,
    inventory_synced_at TIMESTAMPTZ, -- URS 2.2c & 4.8a: penanda keterkinian stok
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- URS 4.8 & 4.9: Transaksi material per WO (issued, returned, actual used — terrekonsiliasi)
CREATE TABLE IF NOT EXISTS wo_material_transactions (
    id                   UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    work_order_id        UUID NOT NULL REFERENCES work_orders(id) ON DELETE RESTRICT,
    asset_id             UUID NOT NULL REFERENCES assets(id) ON DELETE RESTRICT,
    spare_part_id        UUID NOT NULL REFERENCES spare_parts(id) ON DELETE RESTRICT,
    external_tx_no       VARCHAR(100) UNIQUE, -- Cegah duplikasi biaya saat rekonsiliasi ulang (URS 4.9b)
    tx_type              VARCHAR(20) NOT NULL CHECK (tx_type IN ('REQUEST', 'ISSUE', 'RETURN')),
    quantity             NUMERIC(14, 3) NOT NULL CHECK (quantity > 0),
    unit_cost_snapshot   NUMERIC(14, 2) NOT NULL DEFAULT 0,
    warehouse_ref        VARCHAR(120),
    component_position_id UUID REFERENCES asset_component_positions(id),
    reconciled           BOOLEAN NOT NULL DEFAULT FALSE,
    transacted_by        UUID REFERENCES users(id),
    transacted_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);
