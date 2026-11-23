export default function AssetsModulePage() {
  return (
    <div className="space-y-4">
      <div className="bg-white border border-slate-200 rounded-xl p-6">
        <p className="text-xs font-semibold text-emerald-700">
          Track 3 · Branch: feat/module-2-assets-machine-card · PIC: Lovi (Teknik Elektro)
        </p>
        <h2 className="text-xl font-bold mt-1">
          Modul 2 — Master Data Aset, Spesifikasi Mesin &amp; Kartu Mesin
        </h2>
        <p className="text-sm text-slate-600 mt-2">
          Area pengembangan untuk Daftar &amp; Pencarian Aset / QR Scan (URS 3.1),
          Pendaftaran Aset Unik (URS 3.2), Spesifikasi Mesin Dinamis &amp; Lampiran
          Teknis (URS 3.3), Kartu Mesin Kronologis &amp; Riwayat Komponen (URS 3.4),
          Perubahan Data Kritis dengan Persetujuan (URS 3.5), dan Penonaktifan Aset
          (URS 3.6).
        </p>
        <p className="text-xs text-slate-500 mt-4">
          Letakkan komponen UI modul ini di <code>frontend/src/features/assets/</code>.
        </p>
      </div>
    </div>
  );
}
