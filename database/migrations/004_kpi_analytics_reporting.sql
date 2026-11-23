-- ==============================================================================
-- Migration 004: Modul 1 — Master KPI, Operasional, Anggaran, Analitik & Pelaporan
-- Track 2 (PIC: Rainer - Teknik Industri) | URS: 2.1 – 2.9, 5.1 – 5.5
-- ==============================================================================

-- URS 2.7: Master Definisi KPI, Rumus, Target & Ambang yang dapat ditelusuri per periode berlaku
CREATE TABLE IF NOT EXISTS kpi_definitions (
    id                    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    kpi_code              VARCHAR(40) NOT NULL, -- e.g. MTTR, MTBF, AVAILABILITY, PM_COMPLIANCE
    name                  VARCHAR(150) NOT NULL,
    formula_description   TEXT NOT NULL,
    numerator_rule        TEXT NOT NULL,
    denominator_rule      TEXT NOT NULL,
    unit                  VARCHAR(32) NOT NULL,
    include_waiting_time  BOOLEAN NOT NULL DEFAULT FALSE, -- URS 2.7a: perlakuan waktu tunggu MTTR
    target_value          NUMERIC(14, 4) NOT NULL,
    warning_threshold     NUMERIC(14, 4),
    critical_threshold    NUMERIC(14, 4),
    effective_from        DATE NOT NULL,
    effective_to          DATE,
    version               INTEGER NOT NULL DEFAULT 1,
    approved_by           UUID REFERENCES users(id),
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (kpi_code, version)
);

-- Seed definisi KPI awal (Rainer - Teknik Industri akan memvalidasi rumus & target bersama Engineering Manager)
INSERT INTO kpi_definitions (
    kpi_code, name, formula_description, numerator_rule, denominator_rule, unit,
    include_waiting_time, target_value, warning_threshold, critical_threshold, effective_from, version
) VALUES
    ('MTTR', 'Mean Time To Repair', 'Total Active Repair Time / Jumlah Kejadian Breakdown', 'SUM(active_repair_minutes) / 60', 'COUNT(breakdown_wo_closed)', 'hours', FALSE, 2.0, 3.0, 4.0, '2026-01-01', 1),
    ('MTBF', 'Mean Time Between Failures', '(Total Operating Hours - Total Downtime) / Jumlah Kegagalan', 'SUM(operating_hours) - SUM(downtime_hours)', 'COUNT(breakdown_wo_closed)', 'hours', FALSE, 168.0, 120.0, 72.0, '2026-01-01', 1),
    ('PM_COMPLIANCE', 'PM Schedule Compliance', '(PM Selesai Tepat Waktu / Total PM Terjadwal) * 100%', 'COUNT(pm_completed_on_time)', 'COUNT(pm_scheduled)', '%', FALSE, 95.0, 90.0, 85.0, '2026-01-01', 1)
ON CONFLICT (kpi_code, version) DO NOTHING;

-- URS 2.4c & 2.9: Data Jam Operasi & Output Operasional untuk MTBF, Availability & Cost-to-Output
CREATE TABLE IF NOT EXISTS asset_operational_logs (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asset_id          UUID NOT NULL REFERENCES assets(id) ON DELETE RESTRICT,
    operational_date  DATE NOT NULL,
    shift_code        VARCHAR(20) NOT NULL,
    planned_hours     NUMERIC(8, 2) NOT NULL DEFAULT 0,
    operating_hours   NUMERIC(8, 2) NOT NULL DEFAULT 0,
    output_quantity   NUMERIC(14, 3),
    output_unit       VARCHAR(32),
    source_system     VARCHAR(80) NOT NULL DEFAULT 'MANUAL',
    synced_at         TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (asset_id, operational_date, shift_code)
);

-- URS 2.5c: Data Anggaran Pemeliharaan untuk perbandingan Aktual vs Anggaran
CREATE TABLE IF NOT EXISTS maintenance_budgets (
    id             UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    fiscal_year    SMALLINT NOT NULL,
    fiscal_month   SMALLINT NOT NULL CHECK (fiscal_month BETWEEN 1 AND 12),
    division_id    UUID NOT NULL REFERENCES divisions(id),
    category       VARCHAR(80) NOT NULL, -- PLANNED, UNPLANNED, SPARE_PART, CONTRACTOR
    budget_amount  NUMERIC(16, 2) NOT NULL DEFAULT 0,
    currency       VARCHAR(8) NOT NULL DEFAULT 'IDR',
    created_by     UUID REFERENCES users(id),
    updated_at     TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (fiscal_year, fiscal_month, division_id, category)
);

-- URS 5.4 & 5.5: Pelaporan Shift Otomatis & Pengiriman Susulan
CREATE TABLE IF NOT EXISTS shift_reports (
    id                    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    operational_date      DATE NOT NULL,
    shift_code            VARCHAR(20) NOT NULL,
    division_id           UUID NOT NULL REFERENCES divisions(id),
    pic_user_id           UUID NOT NULL REFERENCES users(id),
    is_nil_activity       BOOLEAN NOT NULL DEFAULT FALSE,
    status                VARCHAR(24) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'COMPLETE', 'SENT', 'DELIVERY_FAILED')),
    missing_fields        JSONB DEFAULT '[]'::jsonb,
    completed_at          TIMESTAMPTZ,
    sent_at               TIMESTAMPTZ,
    delivery_attempts     SMALLINT NOT NULL DEFAULT 0,
    last_delivery_error   TEXT,
    late_completion       BOOLEAN NOT NULL DEFAULT FALSE, -- URS 5.5c: bedakan telat isi vs gangguan sistem
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (operational_date, shift_code, division_id, pic_user_id)
);

-- URS 5.2c & 5.3: Riwayat Ekspor & Revisi Laporan
CREATE TABLE IF NOT EXISTS report_exports (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    report_code       VARCHAR(64) NOT NULL,
    version           INTEGER NOT NULL DEFAULT 1,
    previous_export_id UUID REFERENCES report_exports(id),
    revision_reason   TEXT,
    export_format     VARCHAR(16) NOT NULL CHECK (export_format IN ('PDF', 'EXCEL', 'CSV')),
    applied_filters   JSONB NOT NULL DEFAULT '{}'::jsonb,
    status            VARCHAR(20) NOT NULL CHECK (status IN ('SUCCESS', 'FAILED')),
    error_message     TEXT,
    generated_by      UUID NOT NULL REFERENCES users(id),
    generated_at      TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Analytical View dasar untuk Kartu Mesin (URS 3.4) & Dashboard (URS 2.3)
CREATE OR REPLACE VIEW vw_asset_overview AS
SELECT
    d.code AS division_code,
    d.name AS division_name,
    a.status,
    a.criticality,
    a.asset_type,
    COUNT(a.id) AS total_assets
FROM assets a
JOIN divisions d ON d.id = a.division_id
GROUP BY d.code, d.name, a.status, a.criticality, a.asset_type;
