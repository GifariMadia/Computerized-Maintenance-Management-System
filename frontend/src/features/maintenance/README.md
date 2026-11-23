# Feature Folder: Modul 3 — Breakdown, PM, Inspeksi & Spare Part (Track 4)

- **Branch**: `feat/module-3-maintenance-ops`
- **PIC**: Member #4 (TBD) / Kolaborasi Tim (lihat `docs/TEAM_WORKFLOW.md`)
- **URS**: `4.1–4.12`
- **Komponen yang dibangun di folder ini**:
  - `BreakdownReportForm.tsx` (validasi status aset, tampilkan ringkasan Kartu Mesin, prioritas & eskalasi K3)
  - `WorkOrderBoard.tsx` (filter status `OPEN`, `IN_PROGRESS`, `WAITING_PART`, `PENDING_VERIFICATION`, `REWORK`, `CLOSED`, penugasan & ETA history)
  - `PMAnnualScheduleView.tsx` & `PMChecklistForm.tsx` (rencana tahunan, reschedule dengan alasan, konfirmasi WO & checklist)
  - `InspectionExecutionForm.tsx` (parameter ukur + indikator otomatis Normal/Peringatan/Kritis + tombol buat permintaan korektif)
  - `SparePartRequestPanel.tsx` (cek stok + waktu pembaruan gudang, request/issue/return, posisi pemasangan komponen)
  - `ShiftHandoverLog.tsx` & `SupervisorVerificationPanel.tsx` (serah-terima lintas shift, approve/rework & tutup WO)
