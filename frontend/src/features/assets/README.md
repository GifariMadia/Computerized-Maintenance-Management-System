# Feature Folder: Modul 2 — Data Aset, Spesifikasi Mesin & Kartu Mesin (Track 3)

- **Branch**: `feat/module-2-assets-machine-card`
- **PIC**: Lovi (Teknik Elektro)
- **URS**: `3.1–3.6`
- **Komponen yang dibangun di folder ini**:
  - `AssetTableAndFilter.tsx` (pencarian kode/nama, filter 4 divisi, lokasi, kategori, criticality, status, tombol scan QR)
  - `AssetRegistrationForm.tsx` (form pendaftaran aset + dynamic `technical_specs` per jenis aset + upload manual/datasheet/foto)
  - `MachineSpecView.tsx` (tampilan Spesifikasi Mesin & viewer lampiran teknis)
  - `MachineCardTimeline.tsx` (riwayat kronologis Breakdown/PM/Inspection terverifikasi, riwayat posisi komponen & cetak PDF Kartu Mesin)
  - `AssetDeactivationModal.tsx` (cek WO terbuka & PM aktif + input alasan & tanggal efektif + approval Manager)
