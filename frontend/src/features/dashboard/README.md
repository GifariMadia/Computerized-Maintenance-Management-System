# Feature Folder: Modul 1 — Dashboard, KPI & Pelaporan (Track 2)

- **Branch**: `feat/module-1-dashboard-analytics`
- **PIC**: Rainer (Teknik Industri)
- **URS**: `2.1–2.9`, `5.1–5.5`
- **Komponen yang dibangun di folder ini**:
  - `DashboardFilterBar.tsx` (filter gabungan periode, divisi, area, line, kategori, jenis WO, status, shift)
  - `AssetOverviewCards.tsx` & `WorkOrderSummary.tsx` (dengan drill-down ke daftar WO)
  - `ReliabilityKpiPanel.tsx` (MTTR, MTBF, Availability, PM Compliance + target/threshold status + data stale banner)
  - `BreakdownDowntimeCharts.tsx` (Pareto penyebab dominan, tren downtime, top bad-actor assets)
  - `CostAndSparePartCharts.tsx` (Planned vs Unplanned vs Budget, Theoretical vs Actual spare part, Reorder alert)
  - `FlexibleReportBuilder.tsx` (ekspor PDF/Excel/CSV & jejak revisi laporan)
