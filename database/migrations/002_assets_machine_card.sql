-- ==============================================================================
-- Migration 002: Modul 2 — Master Data Aset, Spesifikasi Mesin & Kartu Mesin
-- Track 3 (PIC: Lovi - Teknik Elektro) | URS: 3.1, 3.2, 3.3, 3.4, 3.5, 3.6
-- ==============================================================================

CREATE TABLE IF NOT EXISTS assets (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asset_code        VARCHAR(64) UNIQUE NOT NULL,
    qr_code_value     VARCHAR(128) UNIQUE NOT NULL,
    name              VARCHAR(200) NOT NULL,
    asset_type        VARCHAR(64) NOT NULL, -- e.g. PRODUCTION_MACHINE, UTILITY_SYSTEM, BUILDING_FACILITY
    category          VARCHAR(100) NOT NULL,
    division_id       UUID NOT NULL REFERENCES divisions(id),
    area_id           UUID REFERENCES plants_areas(id),
    location_detail   VARCHAR(200),
    line_name         VARCHAR(100),
    manufacturer      VARCHAR(150),
    model             VARCHAR(150),
    serial_number     VARCHAR(120) UNIQUE,
    manufacture_year  SMALLINT,
    criticality       VARCHAR(20) NOT NULL CHECK (criticality IN ('CRITICAL', 'HIGH', 'MEDIUM', 'LOW')),
    status            VARCHAR(24) NOT NULL DEFAULT 'ACTIVE' CHECK (status IN (
                          'ACTIVE', 'BREAKDOWN', 'UNDER_MAINTENANCE', 'STANDBY', 'DEACTIVATED'
                      )),
    operating_hours   NUMERIC(12, 2) NOT NULL DEFAULT 0,
    is_gxp_relevant   BOOLEAN NOT NULL DEFAULT FALSE,
    -- URS 3.3b: Parameter teknis dinamis per jenis aset (dimensi, kapasitas, daya, tegangan, dll)
    technical_specs   JSONB NOT NULL DEFAULT '{}'::jsonb,
    deactivated_at    TIMESTAMPTZ,
    deactivation_note TEXT,
    -- URS 3.5c: Optimistic locking mencegah concurrent overwrite
    version           INTEGER NOT NULL DEFAULT 1,
    created_by        UUID REFERENCES users(id),
    updated_by        UUID REFERENCES users(id),
    created_at        TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at        TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_assets_division ON assets (division_id);
CREATE INDEX IF NOT EXISTS idx_assets_status ON assets (status);
CREATE INDEX IF NOT EXISTS idx_assets_criticality ON assets (criticality);
CREATE INDEX IF NOT EXISTS idx_assets_specs_gin ON assets USING GIN (technical_specs);

-- URS 3.2b & 3.3c: Foto kondisi aset, manual, datasheet, dan gambar teknis
CREATE TABLE IF NOT EXISTS asset_documents (
    id            UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asset_id      UUID NOT NULL REFERENCES assets(id) ON DELETE RESTRICT,
    doc_type      VARCHAR(40) NOT NULL CHECK (doc_type IN (
                      'PHOTO', 'MANUAL', 'DATASHEET', 'TECHNICAL_DRAWING', 'CERTIFICATE', 'OTHER'
                  )),
    title         VARCHAR(200) NOT NULL,
    storage_key   TEXT NOT NULL,
    mime_type     VARCHAR(120) NOT NULL,
    size_bytes    BIGINT NOT NULL,
    uploaded_by   UUID REFERENCES users(id),
    uploaded_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- URS 3.5b & 3.6b: Persetujuan perubahan data kritis (criticality, status, divisi, lokasi) & penonaktifan aset
CREATE TABLE IF NOT EXISTS asset_change_requests (
    id                UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asset_id          UUID NOT NULL REFERENCES assets(id) ON DELETE RESTRICT,
    request_type      VARCHAR(32) NOT NULL CHECK (request_type IN ('CRITICAL_CHANGE', 'DEACTIVATION', 'REACTIVATION')),
    proposed_changes  JSONB NOT NULL,
    effective_date    DATE,
    reason            TEXT NOT NULL,
    status            VARCHAR(24) NOT NULL DEFAULT 'PENDING' CHECK (status IN ('PENDING', 'APPROVED', 'REJECTED')),
    requested_by      UUID NOT NULL REFERENCES users(id),
    requested_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    reviewed_by       UUID REFERENCES users(id),
    reviewed_at       TIMESTAMPTZ,
    review_notes      TEXT
);

-- URS 3.4c & 4.9c-d: Master posisi komponen & riwayat penggantian komponen pada Kartu Mesin
CREATE TABLE IF NOT EXISTS asset_component_positions (
    id                        UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    asset_id                  UUID NOT NULL REFERENCES assets(id) ON DELETE RESTRICT,
    position_code             VARCHAR(64) NOT NULL,
    position_name             VARCHAR(150) NOT NULL,
    part_code                 VARCHAR(80),
    interval_basis            VARCHAR(20) CHECK (interval_basis IN ('CALENDAR_DAYS', 'OPERATING_HOURS')),
    replacement_interval      NUMERIC(12, 2),
    last_replaced_at          TIMESTAMPTZ,
    last_replaced_op_hours    NUMERIC(12, 2),
    next_target_date          DATE,
    next_target_op_hours      NUMERIC(12, 2),
    UNIQUE (asset_id, position_code)
);
