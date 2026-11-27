.PHONY: help env up down logs ps db-up migrate backup restore backend-dev backend-test frontend-dev frontend-build

SHELL := /bin/bash

help: ## Tampilkan daftar perintah tersedia
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-18s\033[0m %s\n", $$1, $$2}'

env: ## Salin .env.example menjadi .env jika belum ada
	@if [ ! -f .env ]; then cp .env.example .env && echo "Created .env from .env.example"; else echo ".env already exists"; fi

up: env ## Jalankan seluruh stack (PostgreSQL, MinIO, Backend Go, Frontend Next.js)
	docker compose up -d --build

down: ## Hentikan seluruh container
	docker compose down

logs: ## Lihat log seluruh container secara realtime
	docker compose logs -f

ps: ## Cek status container
	docker compose ps

db-up: env ## Jalankan hanya PostgreSQL, MinIO, dan inisialisasi bucket (untuk dev lokal Go & Next.js)
	docker compose -f docker-compose.yml -f docker-compose.dev.yml up -d postgres minio minio-init

migrate: ## Jalankan semua file migrasi SQL secara berurutan ke PostgreSQL container
	@set -a && source .env && set +a && \
	for f in database/migrations/*.sql; do \
		echo "Applying $$f..."; \
		docker compose exec -T postgres psql -v ON_ERROR_STOP=1 -U "$${POSTGRES_USER:-cmms_app}" -d "$${POSTGRES_DB:-cmms_db}" < "$$f"; \
	done

backup: ## Backup database PostgreSQL terkompresi (URS 4.2)
	./infra/scripts/backup-postgres.sh

restore: ## Restore database dari file backup (gunakan: make restore FILE=backups/xxx.sql.gz)
	./infra/scripts/restore-postgres.sh "$(FILE)"

backend-dev: ## Jalankan backend Go secara lokal (tanpa rebuild container)
	cd backend && TZ=Asia/Jakarta go run ./cmd/api

backend-test: ## Jalankan unit test backend Go
	cd backend && TZ=Asia/Jakarta go test -v ./...

frontend-dev: ## Jalankan frontend Next.js secara lokal
	cd frontend && npm run dev

frontend-build: ## Build production frontend Next.js
	cd frontend && npm run build
