export default function DashboardModulePage() {
  return (
    <div className="space-y-4">
      <div className="bg-white border border-slate-200 rounded-xl p-6">
        <p className="text-xs font-semibold text-emerald-700">
          Track 2 · Branch: feat/module-1-dashboard-analytics · PIC: Rainer (Teknik Industri)
        </p>
        <h2 className="text-xl font-bold mt-1">
          Modul 1 — Dashboard, Analitik KPI & Pelaporan
        </h2>
        <p className="text-sm text-slate-600 mt-2">
          Area pengembangan untuk Asset Overview (URS 2.3), Breakdown &amp; Downtime
          Analysis + MTTR/MTBF (URS 2.4, 2.7), Maintenance Cost &amp; Budget (URS 2.5,
          2.9), Analisis Spare Part &amp; Reorder Point (URS 2.6), serta Laporan
          Fleksibel &amp; Shift Report (URS 5.1–5.5).
        </p>
        <p className="text-xs text-slate-500 mt-4">
          Letakkan komponen UI modul ini di <code>frontend/src/features/dashboard/</code>.
        </p>
      </div>
    </div>
  );
}
