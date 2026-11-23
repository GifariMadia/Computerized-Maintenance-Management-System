export default function MaintenanceModulePage() {
  return (
    <div className="space-y-4">
      <div className="bg-white border border-slate-200 rounded-xl p-6">
        <p className="text-xs font-semibold text-emerald-700">
          Track 4 · Branch: feat/module-3-maintenance-ops · PIC: Member #4 / Kolaborasi Tim
        </p>
        <h2 className="text-xl font-bold mt-1">
          Modul 3 — Breakdown, Preventive Maintenance, Inspeksi &amp; Spare Part
        </h2>
        <p className="text-sm text-slate-600 mt-2">
          Area pengembangan untuk Pelaporan Breakdown &amp; Eskalasi K3 (URS 4.1),
          Pembuatan WO &amp; Prediksi Selesai (URS 4.2), Diagnosis &amp; Pencatatan
          Waktu/Downtime (URS 4.3), Rencana &amp; Pelaksanaan PM (URS 4.4–4.5),
          Inspeksi &amp; Temuan Abnormal (URS 4.6–4.7), Permintaan/Pemakaian Spare
          Part (URS 4.8–4.9), Serah-Terima Lintas Shift (URS 4.10), serta Verifikasi,
          Rework &amp; Penutupan WO (URS 4.11–4.12).
        </p>
        <p className="text-xs text-slate-500 mt-4">
          Letakkan komponen UI modul ini di <code>frontend/src/features/maintenance/</code>.
        </p>
      </div>
    </div>
  );
}
