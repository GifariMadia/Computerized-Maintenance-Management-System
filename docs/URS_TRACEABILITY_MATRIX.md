# URS Traceability Matrix — CMMS PT Bintang Toedjoe Site Pulogadung

Pemetaan menyeluruh setiap klausul User Requirement Specification (URS) ke **Track**, **Branch**, **PIC**, **Tabel PostgreSQL**, dan **Endpoint REST API**.

---

## 1. Akses Pengguna dan Pengendalian Data (Track 1 — Gifari / IT)

| URS | Tahapan | GMP | Rank | Branch | Tabel Utama | Endpoint Utama |
| :---: | :--- | :---: | :---: | :--- | :--- | :--- |
| **1.1** | Autentikasi & sesi (lockout 3x gagal, timeout) | Y | I | `feat/core-auth-gxp` | `users`, `auth_login_logs` | `POST /api/v1/auth/login`, `POST /api/v1/auth/logout` |
| **1.2** | Hak akses & cakupan data (Peran × Divisi × Area) | Y | I | `feat/core-auth-gxp` | `user_scopes`, `divisions`, `plants_areas` | `GET /api/v1/core/divisions`, `GET/PUT /api/v1/core/users/{id}/scopes` |
| **1.3** | Validasi & penyimpanan (anti catatan ganda) | Y | I | `feat/core-auth-gxp` | `idempotency_keys` | Middleware validasi & `Idempotency-Key` pada seluruh `POST/PUT/PATCH` |
| **1.4** | Audit trail & koreksi (immutable ALCOA) | Y | I | `feat/core-auth-gxp` | `audit_trails` | `GET /api/v1/core/audit-trails` |

---

## 2. Modul 1 — Dashboard & Analitik (Track 2 — Rainer / Teknik Industri)

| URS | Tahapan | GMP | Rank | Branch | Tabel / View Utama | Endpoint Utama |
| :---: | :--- | :---: | :---: | :--- | :--- | :--- |
| **2.1** | Pemilihan cakupan & multi-filter dashboard | Y | I | `feat/module-1-dashboard-analytics` | `user_scopes`, `assets`, `work_orders` | `GET /api/v1/dashboard/overview?division=&from=&to=` |
| **2.2** | Pengambilan, pemeriksaan kualitas & kesegaran data | Y | I | `feat/module-1-dashboard-analytics` + Core | `spare_parts.inventory_synced_at`, `meta.data_stale` | `POST /api/v1/dashboard/refresh` |
| **2.3** | Ringkasan aset & pekerjaan (drill-down WO) | N | I | `feat/module-1-dashboard-analytics` | `vw_asset_overview`, `work_orders`, `pm_schedules` | `GET /api/v1/dashboard/overview` |
| **2.4** | Analisis breakdown, overlap downtime, MTTR & MTBF | Y | I | `feat/module-1-dashboard-analytics` | `breakdown_reports`, `work_orders`, `asset_operational_logs` | `GET /api/v1/dashboard/breakdown-downtime` |
| **2.5** | Analisis biaya (Planned, Unplanned, Breakdown vs Budget) | N | I | `feat/module-1-dashboard-analytics` | `work_orders`, `wo_material_transactions`, `maintenance_budgets` | `GET /api/v1/dashboard/costs` |
| **2.6** | Analisis kebutuhan rencana, teoretis & pemakaian spare part | N | I | `feat/module-1-dashboard-analytics` | `pm_plans`, `spare_parts`, `wo_material_transactions` | `GET /api/v1/dashboard/spare-parts` |
| **2.7** | Definisi, target, ambang & keterlacakan versi KPI | Y | I | `feat/module-1-dashboard-analytics` | `kpi_definitions` | `GET/POST /api/v1/dashboard/kpi` |
| **2.8** | Evaluasi & perbandingan antar-aset → Kartu Mesin | N | I | `feat/module-1-dashboard-analytics` | `assets`, `work_orders` | `GET /api/v1/dashboard/asset-evaluation` |
| **2.9** | Evaluasi biaya pemeliharaan terhadap output operasional | N | D | `feat/module-1-dashboard-analytics` | `asset_operational_logs`, `work_orders` | `GET /api/v1/dashboard/cost-to-output` |

---

## 3. Modul 2 — Data Aset & Kartu Mesin (Track 3 — Lovi / Teknik Elektro)

| URS | Tahapan | GMP | Rank | Branch | Tabel Utama | Endpoint Utama |
| :---: | :--- | :---: | :---: | :--- | :--- | :--- |
| **3.1** | Daftar, filter, pencarian & QR/barcode aset | Y | I | `feat/module-2-assets-machine-card` | `assets` | `GET /api/v1/assets`, `GET /api/v1/assets/qr/{code}` |
| **3.2** | Pendaftaran aset baru & Asset ID unik | Y | I | `feat/module-2-assets-machine-card` | `assets`, `asset_documents` | `POST /api/v1/assets` |
| **3.3** | Spesifikasi Mesin dinamis & dokumen teknis | Y | I | `feat/module-2-assets-machine-card` | `assets.technical_specs`, `asset_documents` | `GET /api/v1/assets/{id}`, `POST /api/v1/assets/{id}/documents` |
| **3.4** | Kartu Mesin kronologis & riwayat komponen | Y | I | `feat/module-2-assets-machine-card` | `assets`, `work_orders`, `asset_component_positions` | `GET /api/v1/assets/{id}/machine-card` |
| **3.5** | Perubahan data aset (approval + optimistic locking) | Y | I | `feat/module-2-assets-machine-card` | `assets.version`, `asset_change_requests`, `audit_trails` | `PATCH /api/v1/assets/{id}`, `POST /api/v1/assets/{id}/change-requests` |
| **3.6** | Penonaktifan aset (cek WO terbuka & PM aktif) | Y | I | `feat/module-2-assets-machine-card` | `assets`, `asset_change_requests`, `work_orders`, `pm_schedules` | `POST /api/v1/assets/{id}/deactivate` |

---

## 4. Modul 3 — Breakdown, PM, Inspeksi & Spare Part (Track 4 — Member #4 / Kolaborasi)

| URS | Tahapan | GMP | Rank | Branch | Tabel Utama | Endpoint Utama |
| :---: | :--- | :---: | :---: | :--- | :--- | :--- |
| **4.1** | Pelaporan, prioritas breakdown & eskalasi K3 | Y | I | `feat/module-3-maintenance-ops` | `breakdown_reports`, `assets` | `POST /api/v1/breakdowns` |
| **4.2** | Pembuatan WO, penugasan ulang & prediksi selesai | Y | I | `feat/module-3-maintenance-ops` | `work_orders`, `wo_eta_history` | `POST /api/v1/work-orders`, `PATCH /api/v1/work-orders/{id}/eta` |
| **4.3** | Diagnosis, perbaikan, waktu tunggu/aktif & downtime | Y | I | `feat/module-3-maintenance-ops` | `work_orders` | `PATCH /api/v1/work-orders/{id}/execution` |
| **4.4** | Rencana PM tahunan & riwayat reschedule | Y | I | `feat/module-3-maintenance-ops` | `pm_plans`, `pm_schedules` | `GET/POST /api/v1/pm-plans`, `PATCH /api/v1/pm-schedules/{id}/reschedule` |
| **4.5** | Persiapan, checklist & pelaksanaan PM | Y | I | `feat/module-3-maintenance-ops` | `pm_schedules`, `work_orders` | `POST /api/v1/pm-schedules/{id}/start` |
| **4.6** | Perencanaan, template & parameter inspeksi | Y | I | `feat/module-3-maintenance-ops` *(input Lovi)* | `inspection_templates`, `work_orders` | `GET/POST /api/v1/inspections/templates` |
| **4.7** | Hasil inspeksi (Normal/Peringatan/Kritis) & temuan | Y | I | `feat/module-3-maintenance-ops` | `work_orders.inspection_readings` | `POST /api/v1/inspections/{wo_id}/readings` |
| **4.8** | Kebutuhan, cek stok & permintaan spare part (`WAITING_PART`) | Y | I | `feat/module-3-maintenance-ops` | `spare_parts`, `wo_material_transactions` | `POST /api/v1/work-orders/{id}/materials/request` |
| **4.9** | Pemakaian, pengembalian, biaya & ganti komponen | Y | I | `feat/module-3-maintenance-ops` | `wo_material_transactions`, `asset_component_positions` | `POST /api/v1/work-orders/{id}/materials/reconcile` |
| **4.10** | Progres, hasil & serah-terima lintas shift | Y | I | `feat/module-3-maintenance-ops` | `work_orders`, `wo_shift_logs` | `POST /api/v1/work-orders/{id}/shift-logs` |
| **4.11** | Pengajuan, verifikasi Supervisor & `REWORK` | Y | I | `feat/module-3-maintenance-ops` | `work_orders`, `audit_trails` | `POST /api/v1/work-orders/{id}/submit`, `POST /api/v1/work-orders/{id}/verify` |
| **4.12** | Penutupan WO atomik & update Kartu Mesin + jadwal PM | Y | I | `feat/module-3-maintenance-ops` | `work_orders`, `assets`, `pm_schedules`, `audit_trails` | `POST /api/v1/work-orders/{id}/close` |

---

## 5. Pelaporan, Penelusuran & Non-Fungsional

| URS | Tahapan | GMP | Rank | PIC Utama | Komponen / Script |
| :---: | :--- | :---: | :---: | :--- | :--- |
| **5.1** | Laporan fleksibel multi-filter + drill-down | Y | I | Rainer | `backend/internal/modules/dashboard` |
| **5.2** | Ekspor PDF / Excel / CSV + cetak Kartu Mesin | N | I | Gifari + Rainer + Lovi | `report_exports`, export engine |
| **5.3** | Koreksi & revisi laporan berversi | Y | I | Gifari + Rainer | `report_exports.version`, `audit_trails` |
| **5.4** | Pemeriksaan & pelaporan shift otomatis | Y | D | Rainer + Gifari | `shift_reports` |
| **5.5** | Pelengkapan & pengiriman laporan susulan | Y | D | Rainer + Gifari | `shift_reports.late_completion` |
| **NF 4.1** | Business Contingency Plan (input manual + stale indicator) | Y | I | Gifari | `work_orders.manual_log_ref`, `meta.data_stale` |
| **NF 4.2** | Backup harian (RPO ≤ 4 jam, RTO ≤ 8 jam, retensi 5 thn) | Y | I | Gifari | `infra/scripts/backup-postgres.sh`, `restore-postgres.sh` |
| **NF 4.3** | Performa load ≤ 3 detik & validasi 4 divisi | Y | I | Seluruh Tim | Docker Linux + indexing PostgreSQL + UAT |
