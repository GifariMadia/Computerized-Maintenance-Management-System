import Link from "next/link";

const MODULES = [
  {
    track: "Track 1 · Core Platform & GxP",
    branch: "feat/core-auth-gxp",
    owner: "Gifari (IT)",
    href: "/admin",
    urs: "URS 1.1–1.4, 2.2, 4.2, ALCOA",
    summary:
      "Autentikasi & lockout 3x gagal, RBAC multi-divisi/area, Audit Trail immutable (ALCOA), Idempotency, Docker Linux & Backup/Restore.",
  },
  {
    track: "Track 2 · Modul 1 Dashboard & Analitik",
    branch: "feat/module-1-dashboard-analytics",
    owner: "Rainer (Teknik Industri)",
    href: "/dashboard",
    urs: "URS 2.1–2.9, 5.1–5.5",
    summary:
      "Asset Overview, MTTR/MTBF/PM Compliance, overlap downtime lintas shift, analisis biaya Planned vs Unplanned, Spare Part & Laporan.",
  },
  {
    track: "Track 3 · Modul 2 Data Aset & Kartu Mesin",
    branch: "feat/module-2-assets-machine-card",
    owner: "Lovi (Teknik Elektro)",
    href: "/assets",
    urs: "URS 3.1–3.6 (+ Standar Inspeksi 4.6)",
    summary:
      "Master Data Aset (4 divisi), QR/Barcode, Spesifikasi Mesin dinamis (daya, tegangan, kapasitas), Kartu Mesin & Penonaktifan Aset.",
  },
  {
    track: "Track 4 · Modul 3 Maintenance Operations",
    branch: "feat/module-3-maintenance-ops",
    owner: "Member #4 (TBD / Kolaborasi Tim)",
    href: "/maintenance",
    urs: "URS 4.1–4.12",
    summary:
      "Pelaporan Breakdown & eskalasi K3, Work Order & ETA, Jadwal PM Tahunan, Eksekusi Inspeksi, Spare Part (Waiting Part) & Verifikasi.",
  },
];

export default function HomePage() {
  return (
    <div className="space-y-6">
      <header className="bg-white border border-slate-200 rounded-xl p-6 shadow-xs">
        <span className="inline-block text-xs font-semibold px-2.5 py-1 rounded-full bg-emerald-50 text-emerald-700 border border-emerald-200">
          Base Architecture · Next.js + Go + PostgreSQL + Docker (Linux / WIB)
        </span>
        <h2 className="text-2xl font-bold mt-3">
          Computerized Maintenance Management System (CMMS)
        </h2>
        <p className="text-slate-600 mt-1 text-sm">
          Department Engineering — PT Bintang Toedjoe Site Pulogadung. Satu basis
          data operasional terpusat untuk Divisi Maintenance, Utility, Building
          Maintenance, dan Spare Part dengan kepatuhan GxP / ALCOA.
        </p>
      </header>

      <section className="grid grid-cols-1 md:grid-cols-2 gap-4">
        {MODULES.map((m) => (
          <Link
            key={m.branch}
            href={m.href}
            className="block bg-white border border-slate-200 rounded-xl p-5 hover:border-emerald-500 transition-colors shadow-xs"
          >
            <div className="flex items-center justify-between gap-2">
              <span className="text-xs font-semibold text-emerald-700">
                {m.track}
              </span>
              <code className="text-[11px] px-2 py-0.5 rounded bg-slate-100 text-slate-700">
                {m.branch}
              </code>
            </div>
            <h3 className="font-semibold text-base mt-2">{m.owner}</h3>
            <p className="text-xs text-slate-500 mt-0.5">{m.urs}</p>
            <p className="text-sm text-slate-600 mt-3">{m.summary}</p>
          </Link>
        ))}
      </section>
    </div>
  );
}
