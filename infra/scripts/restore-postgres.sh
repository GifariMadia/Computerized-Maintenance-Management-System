#!/usr/bin/env bash
# ==============================================================================
# CMMS PostgreSQL Restore & Verification Script (Linux)
# URS 4.2: Pengujian pemulihan data berkala (RTO <= 8 jam)
# Usage: ./infra/scripts/restore-postgres.sh backups/cmms_cmms_db_YYYYMMDD_HHMMSS.sql.gz
# ==============================================================================
set -euo pipefail

export TZ="Asia/Jakarta"
BACKUP_FILE="${1:-}"
if [[ -z "${BACKUP_FILE}" || ! -f "${BACKUP_FILE}" ]]; then
  echo "Usage: $0 <path-to-backup.sql.gz>" >&2
  exit 1
fi

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/../.." && pwd)"

if [[ -f "${ROOT_DIR}/.env" ]]; then
  set -a
  # shellcheck disable=SC1091
  source "${ROOT_DIR}/.env"
  set +a
fi

if [[ -f "${BACKUP_FILE}.sha256" ]]; then
  echo "[$(date '+%Y-%m-%d %H:%M:%S WIB')] Verifying SHA-256 checksum..."
  sha256sum -c "${BACKUP_FILE}.sha256"
fi

DB_NAME="${POSTGRES_DB:-cmms_db}"
DB_USER="${POSTGRES_USER:-cmms_app}"

echo "[$(date '+%Y-%m-%d %H:%M:%S WIB')] Restoring ${BACKUP_FILE} into ${DB_NAME}..."
gunzip -c "${BACKUP_FILE}" | docker compose -f "${ROOT_DIR}/docker-compose.yml" exec -T postgres \
  psql -v ON_ERROR_STOP=1 -U "${DB_USER}" -d "${DB_NAME}"

echo "[$(date '+%Y-%m-%d %H:%M:%S WIB')] Restore completed. Verify Asset <-> WO <-> Machine Card integrity."
