-- ==============================================================================
-- CMMS Database Initialization (PostgreSQL 16 on Linux)
-- URS 2.1: Waktu operasional WIB (Asia/Jakarta)
-- ==============================================================================
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

DO $$
BEGIN
    EXECUTE format('ALTER DATABASE %I SET timezone TO ''Asia/Jakarta''', current_database());
END
$$;
