export default function AdminCoreModulePage() {
  return (
    <div className="space-y-4">
      <div className="bg-white border border-slate-200 rounded-xl p-6">
        <p className="text-xs font-semibold text-emerald-700">
          Track 1 · Branch: feat/core-auth-gxp · PIC: Gifari (IT)
        </p>
        <h2 className="text-xl font-bold mt-1">
          Core Platform — Autentikasi, RBAC Multi-Cakupan &amp; Audit Trail (ALCOA)
        </h2>
        <p className="text-sm text-slate-600 mt-2">
          Area pengembangan untuk Manajemen Akun &amp; Sesi / Lockout 3x Gagal (URS
          1.1), Matriks Hak Akses per Peran × Divisi × Area (URS 1.2), Validasi &amp;
          Idempotency (URS 1.3), Audit Trail Immutable &amp; Penelusuran Koreksi (URS
          1.4 / ALCOA), serta Status Sinkronisasi Persediaan (URS 2.2).
        </p>
        <p className="text-xs text-slate-500 mt-4">
          Letakkan komponen UI modul ini di <code>frontend/src/features/core/</code>.
        </p>
      </div>
    </div>
  );
}
