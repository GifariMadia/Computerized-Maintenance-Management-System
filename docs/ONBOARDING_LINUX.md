# Panduan Onboarding Lingkungan Pengembangan Linux

Panduan praktis ini ditujukan bagi seluruh anggota tim (**Gifari**, **Rainer**, **Lovi**, dan **Member #4**) agar bisa menjalankan CMMS secara seragam di Linux (Ubuntu/Debian atau WSL2 Ubuntu di Windows).

---

## 1. Prasyarat Software (Linux)

Pastikan tools berikut terpasang di sistem Linux Anda:
- **Git** (`git --version`)
- **Docker Engine & Docker Compose Plugin** (`docker compose version`)
- **Go 1.24+** *(opsional jika hanya menjalankan via Docker, wajib jika run backend native)*
- **Node.js 22 LTS & npm** *(opsional jika hanya menjalankan via Docker, wajib jika run frontend native)*
- **Make** (`sudo apt-get install -y make`)

---

## 2. Clone Repository & Pindah ke Branch Masing-Masing

```bash
git clone https://github.com/GifariMadia/Computerized-Maintenance-Management-System.git
cd Computerized-Maintenance-Management-System

# Ambil semua branch remote
git fetch --all
```

Pilih branch sesuai peran Anda:
- **Gifari (IT)**: `git checkout feat/core-auth-gxp`
- **Rainer (Teknik Industri)**: `git checkout feat/module-1-dashboard-analytics`
- **Lovi (Teknik Elektro)**: `git checkout feat/module-2-assets-machine-card`
- **Member #4 / Kolaborasi Ops**: `git checkout feat/module-3-maintenance-ops`

---

## 3. Menjalankan Aplikasi via Docker Compose

```bash
# 1. Buat file .env dari template
make env

# 2. Jalankan PostgreSQL, MinIO, Backend Go, dan Frontend Next.js
make up

# 3. Terapkan seluruh migrasi tabel CMMS (001 - 004)
make migrate

# 4. Cek status container
make ps
```

Akses layanan lokal:
- **Frontend Next.js**: `http://localhost:3000`
- **Backend Go API Health**: `http://localhost:8080/api/v1/health`
- **MinIO Console (Storage Dokumen)**: `http://localhost:9001`
- **PostgreSQL**: `localhost:5432` (db: `cmms_db`, user: `cmms_app`)

---

## 4. Mode Hybrid Development (Cepat untuk Coding)

Jika ingin hot-reload saat coding Frontend/Backend:
```bash
# Jalankan hanya database & MinIO di Docker
make db-up
make migrate

# Terminal 1: Jalankan Backend Go
make backend-dev

# Terminal 2: Jalankan Frontend Next.js
cd frontend && npm install && npm run dev
```

---

## 5. Alur Commit & Pull Request yang Aman

```bash
# Pastikan berada di branch modul Anda (JANGAN di main!)
git status

# Simpan perubahan dengan format pesan commit konvensional
git add .
git commit -m "feat(assets): add dynamic technical specs validation for utility equipment [URS 3.3]"

# Push ke branch Anda
git push origin HEAD
```

Setelah push, buka GitHub dan buat **Pull Request** dengan target **`develop`** (bukan `main`), lalu isi checklist PR Template (termasuk verifikasi GxP/ALCOA).
