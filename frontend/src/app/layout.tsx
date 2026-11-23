import type { Metadata } from "next";
import type { ReactNode } from "react";
import { AppShell } from "@/components/layout/AppShell";
import "./globals.css";

export const metadata: Metadata = {
  title: "CMMS — PT Bintang Toedjoe Site Pulogadung",
  description:
    "Computerized Maintenance Management System Department Engineering (Maintenance, Utility, Building Maintenance, Spare Part)",
};

export default function RootLayout({ children }: { children: ReactNode }) {
  return (
    <html lang="id">
      <body>
        <AppShell>{children}</AppShell>
      </body>
    </html>
  );
}
