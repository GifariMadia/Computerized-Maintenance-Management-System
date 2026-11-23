#!/usr/bin/env bash
# ==============================================================================
# CMMS PostgreSQL Automated Daily Backup Script (Linux)
# URS 4.2: Backup otomatis harian, tersimpan terpisah, RPO <= 4 jam, retensi 5 tahun
# Contoh cron Linux (tiap 4 jam):
#   0 */4 * * * /path/to/repo/infra/scripts/backup-postgres.sh >> /var/log/cmms-backup.log 2>&1
# ==============================================================================
set -euo pipefail

export TZ="Asia/Jakarta"
TIMESTAMP="$(date +%Y%m%d_%H%M%S)"
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(cd "${SCRIPT_DIR}/../.." && pwd)"
BACKUP_DIR="${BACKUP_DIR:-${ROOT_DIR}/backups}"

if [[ -f "${ROOT_DIR}/.env" ]]; then
  set -a
  # shellcheck disable=SC1091
  source "${ROOT_DIR}/.env"
  set +a
fi

DB_NAME="${POSTGRES_DB:-cmms_db}"
DB_USER="${POSTGRES_USER:-cmms_app}"
OUT_FILE="${BACKUP_DIR}/cmms_${DB_NAME}_${TIMESTAMP}.sql.gz"

mkdir -p "${BACKUP_DIR}"
echo "[$(date '+%Y-%m-%d %H:%M:%S WIB')] Starting PostgreSQL backup -> ${OUT_FILE}"

docker compose -f "${ROOT_DIR}/docker-compose.yml" exec -T postgres \
  pg_dump -U "${DB_USER}" -d "${DB_NAME}" --clean --if-exists --format=plain \
  | gzip -9 > "${OUT_FILE}"

sha256sum "${OUT_FILE}" > "${OUT_FILE}.sha256"
echo "[$(date '+%Y-%m-%d %H:%M:%S WIB')] Backup completed successfully (checksum saved)."
