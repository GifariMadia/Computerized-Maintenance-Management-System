# Computerized-Maintenance-Management-System

**Computerized Maintenance Management System (CMMS)** berbasis web untuk **Department Engineering PT Bintang Toedjoe Site Pulogadung**, mencakup empat divisi operasional (**Maintenance**, **Utility**, **Building Maintenance**, dan **Spare Part**) dalam satu basis data terpusat dengan kepatuhan **GxP / ALCOA**.

---

## Tech Stack (Linux Target Environment)

- **Frontend**: Next.js 15 (App Router) + TypeScript + Tailwind CSS
- **Backend**: Go 1.24 (Modular Monolith, REST API JSON)
- **Database**: PostgreSQL 16 (Single Operational DB, `TZ=Asia/Jakarta` / WIB)
- **Document Storage**: MinIO (S3-compatible untuk foto aset, manual, datasheet, gambar teknis)
- **Container & Ops**: Docker, Docker Compose, Linux shell scripts (automated backup/restore)

---

## Pembagian Modul, Branch & Penanggung Jawab Tim

| Track | Git Branch | Modul & Cakupan URS | PIC & Background |
| :---: | :--- | :--- | :--- |
| **1** | `feat/core-auth-gxp` | **Core Platform**: Auth & Lockout 3x, RBAC Multi-Divisi/Area, Immutable Audit Trail (ALCOA), Idempotency, Sync Inventory, Docker Linux & Backup (`URS 1.1–1.4, 2.2, 4.2`) | **Gifari** *(IT — Tech Lead)* |
| **2** | `feat/module-1-dashboard-analytics` | **Modul 1 — Dashboard & Analitik**: Asset Overview, MTTR/MTBF/PM Compliance, Overlap Downtime Lintas Shift, Cost & Budget, Analisis Spare Part & Laporan (`URS 2.1–2.9, 5.1–5.5`) | **Rainer** *(Teknik Industri)* |
| **3** | `feat/module-2-assets-machine-card` | **Modul 2 — Data Aset & Kartu Mesin**: Master Aset 4 Divisi, QR/Barcode, Spesifikasi Mesin Dinamis, Kartu Mesin & Komponen, Approval Perubahan & Deaktivasi (`URS 3.1–3.6`) | **Lovi** *(Teknik Elektro)* |
| **4** | `feat/module-3-maintenance-ops` | **Modul 3 — Maintenance Operations**: Breakdown & Eskalasi K3, WO & ETA, Jadwal & Eksekusi PM, Inspeksi & Temuan, Spare Part (`Waiting Part`), Shift Handover, Verifikasi/Rework/Close (`URS 4.1–4.12`) | **Anggota #4 (TBD)** *(atau kolaborasi 3 orang — lihat `docs/TEAM_WORKFLOW.md`)* |

- Branch **`main`**: Release produksi tervalidasi.
- Branch **`develop`**: Integrasi gabungan dari 4 feature branch via Pull Request.

---

## Struktur Direktori Utama

```text
├── docker-compose.yml            # Orchestration: PostgreSQL, MinIO, Go API, Next.js
├── Makefile                      # Perintah cepat Linux (make up, make migrate, make backup)
├── .env.example                  # Template konfigurasi environment (WIB, DB, Auth)
├── database/
│   ├── init/                     # Ekstensi PostgreSQL & zona waktu Asia/Jakarta
│   └── migrations/               # 001 (Core), 002 (Aset), 003 (Ops), 004 (KPI/Report)
├── backend/                      # Go Modular Monolith API (/api/v1/...)
│   ├── cmd/api/main.go
│   └── internal/modules/         # core/ | dashboard/ | assets/ | maintenance/
├── frontend/                     # Next.js App Router + TypeScript
│   └── src/
│       ├── app/                  # Routes: /dashboard, /assets, /maintenance, /admin
│       └── features/             # Folder terisolasi per modul/branch
├── infra/scripts/                # Backup & restore PostgreSQL otomatis (URS 4.2)
└── docs/                         # Dokumentasi arsitektur, pembagian tugas & URS matrix
```

---

## Quickstart (Linux + Docker)

```bash
cp .env.example .env
make up
make migrate
```

- Frontend: `http://localhost:3000`
- Backend Health Check: `http://localhost:8080/api/v1/health`

---

## Dokumentasi Kerja Tim

- [Arsitektur Sistem & GxP/ALCOA](docs/ARCHITECTURE.md)
- [Template Kerja Bareng, Pembagian Tugas, RACI & Sprint Roadmap](docs/TEAM_WORKFLOW.md)
- [URS Traceability Matrix (URS 1.1 – 5.5)](docs/URS_TRACEABILITY_MATRIX.md)
- [Panduan Onboarding Linux & Git untuk Tim](docs/ONBOARDING_LINUX.md)
