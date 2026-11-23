## Ringkasan Perubahan
- **Track / Modul**: `[ ] Core & GxP (Gifari)` / `[ ] Modul 1 Dashboard (Rainer)` / `[ ] Modul 2 Aset (Lovi)` / `[ ] Modul 3 Maintenance Ops (Track 4)`
- **Referensi URS**: (contoh: `URS 3.2`, `URS 4.1`, `ALCOA 1.3`)
- **Deskripsi Singkat**:

## Checklist Kualitas & GxP / ALCOA
- [ ] Perubahan hanya menyentuh folder modul sendiri (atau sudah koordinasi jika menyentuh `shared/` / `database/migrations/`).
- [ ] Semua timestamp operasional menggunakan WIB (`Asia/Jakarta`) dan memisahkan `occurred_at` vs `recorded_at` bila relevan.
- [ ] Operasi mutasi penting (Create/Update/Approve/Verify/Cancel/Deactivate) tercatat ke `audit_trails`.
- [ ] Validasi input & pesan error jelas bagi pengguna (Maintenance, Utility, Building Maintenance, Spare Part).
- [ ] `go test ./...` (backend) dan `npm run build` / lint (frontend) berhasil tanpa error.
- [ ] Tidak ada credential/secret (`.env`) yang ikut ter-commit.

## Bukti Uji (Screenshot / API Response)
<!-- Lampirkan contoh response JSON atau screenshot tampilan desktop/tablet -->
