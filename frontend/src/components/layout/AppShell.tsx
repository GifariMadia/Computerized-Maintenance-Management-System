import Link from "next/link";
import type { ReactNode } from "react";

const NAV_ITEMS = [
  { href: "/", label: "Ringkasan Sistem", tag: "Home" },
  { href: "/dashboard", label: "Modul 1: Dashboard & Analitik", tag: "Rainer" },
  { href: "/assets", label: "Modul 2: Data Aset & Kartu Mesin", tag: "Lovi" },
  { href: "/maintenance", label: "Modul 3: Breakdown, PM & Inspeksi", tag: "Track 4" },
  { href: "/admin", label: "Core: Auth, RBAC & Audit Trail", tag: "Gifari" },
];

export function AppShell({ children }: { children: ReactNode }) {
  return (
    <div className="min-h-screen flex flex-col lg:flex-row bg-slate-50 text-slate-900">
      <aside className="w-full lg:w-72 bg-slate-900 text-slate-100 flex-shrink-0">
        <div className="p-5 border-b border-slate-800">
          <p className="text-xs uppercase tracking-wider text-emerald-400 font-semibold">
            PT Bintang Toedjoe · Site Pulogadung
          </p>
          <h1 className="text-lg font-bold mt-1">CMMS Engineering</h1>
          <p className="text-xs text-slate-400 mt-1">
            Maintenance · Utility · Building · Spare Part
          </p>
        </div>
        <nav className="p-3 space-y-1">
          {NAV_ITEMS.map((item) => (
            <Link
              key={item.href}
              href={item.href}
              className="flex items-center justify-between px-3 py-2.5 rounded-lg text-sm hover:bg-slate-800 transition-colors"
            >
              <span>{item.label}</span>
              <span className="text-[11px] px-2 py-0.5 rounded bg-slate-800 text-slate-300">
                {item.tag}
              </span>
            </Link>
          ))}
        </nav>
      </aside>
      <main className="flex-1 p-5 lg:p-8 max-w-6xl w-full mx-auto">
        {children}
      </main>
    </div>
  );
}
