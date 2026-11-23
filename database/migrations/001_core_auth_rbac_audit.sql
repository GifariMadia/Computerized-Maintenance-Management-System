-- ==============================================================================
-- Migration 001: Core Auth, RBAC, Organizational Scopes & Immutable Audit Trail
-- Track 1 (PIC: Gifari - IT) | URS: 1.1, 1.2, 1.3, 1.4, ALCOA 1.1–1.5
-- ==============================================================================

CREATE TABLE IF NOT EXISTS divisions (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code        VARCHAR(32) UNIQUE NOT NULL,
    name        VARCHAR(120) NOT NULL,
    description TEXT,
    is_active   BOOLEAN NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- Seed 4 Divisi Engineering PT Bintang Toedjoe Site Pulogadung (URS Ruang Lingkup)
INSERT INTO divisions (code, name, description) VALUES
    ('MAINTENANCE', 'Maintenance', 'Pemeliharaan, perbaikan, dan pemeriksaan mesin/peralatan produksi'),
    ('UTILITY', 'Utility', 'Pemeliharaan dan pemeriksaan peralatan serta sistem penunjang operasional'),
    ('BUILDING', 'Building Maintenance', 'Pemeliharaan, perbaikan, dan pemeriksaan fasilitas/sarana bangunan'),
    ('SPARE_PART', 'Spare Part', 'Pengelolaan informasi material, kebutuhan, persediaan, dan referensi penggunaan spare part')
ON CONFLICT (code) DO NOTHING;

CREATE TABLE IF NOT EXISTS plants_areas (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    code        VARCHAR(32) UNIQUE NOT NULL,
    name        VARCHAR(120) NOT NULL,
    site        VARCHAR(100) NOT NULL DEFAULT 'Site Pulogadung',
    parent_id   UUID REFERENCES plants_areas(id),
    is_active   BOOLEAN NOT NULL DEFAULT TRUE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS users (
    id                    UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    employee_id           VARCHAR(50) UNIQUE NOT NULL,
    username              VARCHAR(80) UNIQUE NOT NULL,
    full_name             VARCHAR(150) NOT NULL,
    email                 VARCHAR(180) UNIQUE NOT NULL,
    password_hash         TEXT NOT NULL,
    role                  VARCHAR(32) NOT NULL CHECK (role IN (
                              'ADMINISTRATOR',
                              'TECHNICIAN',
                              'INSPECTOR',
                              'SUPERVISOR',
                              'MANAGER',
                              'SPARE_PART_STAFF'
                          )),
    is_active             BOOLEAN NOT NULL DEFAULT TRUE,
    failed_login_attempts SMALLINT NOT NULL DEFAULT 0,
    locked_until          TIMESTAMPTZ,
    last_login_at         TIMESTAMPTZ,
    created_at            TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at            TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- URS 1.2b: Pengguna dengan peran sama dapat memiliki cakupan divisi/area berbeda
CREATE TABLE IF NOT EXISTS user_scopes (
    id          UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id     UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    division_id UUID REFERENCES divisions(id),
    area_id     UUID REFERENCES plants_areas(id),
    can_approve BOOLEAN NOT NULL DEFAULT FALSE,
    can_verify  BOOLEAN NOT NULL DEFAULT FALSE,
    can_export  BOOLEAN NOT NULL DEFAULT FALSE,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (user_id, division_id, area_id)
);

-- URS 1.1b: Pencatatan seluruh percobaan login (sukses maupun gagal)
CREATE TABLE IF NOT EXISTS auth_login_logs (
    id             UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    user_id        UUID REFERENCES users(id),
    username_input VARCHAR(80) NOT NULL,
    succeeded      BOOLEAN NOT NULL,
    failure_reason VARCHAR(120),
    ip_address     INET,
    user_agent     TEXT,
    occurred_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- URS 1.3c & 4.12d: Idempotency keys untuk mencegah transaksi/catatan ganda saat retry
CREATE TABLE IF NOT EXISTS idempotency_keys (
    key             VARCHAR(128) PRIMARY KEY,
    user_id         UUID REFERENCES users(id),
    endpoint        VARCHAR(200) NOT NULL,
    request_hash    VARCHAR(64) NOT NULL,
    response_status SMALLINT,
    response_body   JSONB,
    created_at      TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at      TIMESTAMPTZ NOT NULL DEFAULT (NOW() + INTERVAL '24 hours')
);

-- URS 1.4 & ALCOA (Attributable, Legible, Contemporaneous, Original, Accurate)
-- Tabel ini append-only: perubahan & penghapusan dilarang oleh aplikasi maupun trigger.
CREATE TABLE IF NOT EXISTS audit_trails (
    id           BIGSERIAL PRIMARY KEY,
    actor_id     UUID REFERENCES users(id),
    actor_name   VARCHAR(150) NOT NULL,
    actor_role   VARCHAR(32) NOT NULL,
    action       VARCHAR(40) NOT NULL CHECK (action IN (
                     'CREATE', 'UPDATE', 'CANCEL', 'APPROVE', 'REJECT',
                     'VERIFY', 'REWORK', 'CLOSE', 'DEACTIVATE', 'EXPORT',
                     'ACCESS_DENIED', 'CORRECTION'
                 )),
    entity_type  VARCHAR(80) NOT NULL,
    entity_id    VARCHAR(100) NOT NULL,
    reference_no VARCHAR(100),
    reason       TEXT,
    before_data  JSONB,
    after_data   JSONB,
    ip_address   INET,
    created_at   TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_audit_trails_entity ON audit_trails (entity_type, entity_id);
CREATE INDEX IF NOT EXISTS idx_audit_trails_ref ON audit_trails (reference_no);
CREATE INDEX IF NOT EXISTS idx_audit_trails_created ON audit_trails (created_at DESC);

CREATE OR REPLACE FUNCTION prevent_audit_trail_mutation()
RETURNS TRIGGER AS $$
BEGIN
    RAISE EXCEPTION 'GxP/ALCOA violation: audit_trails records are immutable and cannot be updated or deleted';
END;
$$ LANGUAGE plpgsql;

DROP TRIGGER IF EXISTS trg_audit_trails_immutable ON audit_trails;
CREATE TRIGGER trg_audit_trails_immutable
BEFORE UPDATE OR DELETE ON audit_trails
FOR EACH ROW EXECUTE FUNCTION prevent_audit_trail_mutation();
