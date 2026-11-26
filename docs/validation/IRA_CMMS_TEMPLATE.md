# Initial Risk Assessment (IRA) — CMMS PT Bintang Toedjoe Site Pulogadung

> **Referensi Form Resmi:** Lampiran 6 – `WI-QO-QO-1046.03` (PT Bintang Toedjoe, A Kalbe Company)  
> **File Word Siap Cetak / FUPD:** [IRA_CMMS_PT_Bintang_Toedjoe_Pulogadung.docx](IRA_CMMS_PT_Bintang_Toedjoe_Pulogadung.docx)  
> **Doc No.:** `IRA-CMMS-PGD-001` *(Draft Rev. 00 — ganti dengan nomor FUPD resmi)*  
> **Software Name:** Computerized Maintenance Management System (CMMS)

---

## 1. Lembar Persetujuan (Approval Sheet)

| Peran | Nama & Jabatan | Sign | Date |
| :--- | :--- | :---: | :---: |
| **Prepared by** | Petrus Juan Teranggani *(Maintenance Supervisor Site Pulogadung)* | | |
| **Prepared by** | Didit Nashrulli Hakim *(Utility Supervisor Site Pulogadung)* | | |
| **Prepared by** | Bernadus Asto Wicaksono *(Building Maintenance Supervisor Site Pulogadung)* | | |
| **Reviewed by** | Ferdinand Natanael S *(Robotic & IoT Supervisor — Operational Technology)* | | |
| **Reviewed by** | Vany Okky Safitri *(Quality Assurance Supervisor)* | | |
| **Reviewed by** | Bharestu Ressy Rafsanzani *(Operational Technology Manager)* | | |
| **Reviewed by** | Yudhi Dharmawan *(Engineering Manager)* | | |
| **Approved by** | Yudi *(Quality Assurance Manager)* | | |

---

## 2. System Information

| Field | Isi Khusus CMMS |
| :--- | :--- |
| **System / Application Name** | Computerized Maintenance Management System (CMMS) — Department Engineering Site Pulogadung |
| **System Description** | Aplikasi berbasis website terpusat untuk mengelola master data aset & Kartu Mesin (Modul 2), aktivitas pemeliharaan Breakdown Maintenance, Preventive Maintenance (PM), Inspection & penggunaan spare part (Modul 3), serta Dashboard & Analitik KPI pemeliharaan (Modul 1) untuk empat divisi Department Engineering (**Maintenance**, **Utility**, **Building Maintenance**, dan **Spare Part**) PT Bintang Toedjoe Site Pulogadung. Dibangun secara custom menggunakan Next.js (TypeScript), Go (REST API), PostgreSQL (Single Operational DB), dan Docker pada OS Linux (WIB) dengan pengendalian akses (RBAC), audit trail immutable, serta prinsip ALCOA. |
| **Software Version / Release** | `v1.0` (Initial Development Release — Custom Web Application) |
| **Process Owner** | Department Engineering PT Bintang Toedjoe Site Pulogadung (Maintenance, Utility, Building Maintenance, Spare Part) |
| **Developer / Vendor** | Internal Development Team (IT / Operational Technology & Engineering — PT Bintang Toedjoe) |
| **Document Reference** | URS Computerized Maintenance Management System (CMMS) — Tanggal 12 November 2026 (Rev. 00) |

---

## 3. GxP Assessment

- **Assessment Type:** `[x] New system` | `[ ] Existing system`

### GxP Impact Assessment

| No | Pertanyaan Evaluasi | Jawaban | Justifikasi Singkat (Konteks URS CMMS) |
| :-: | :--- | :---: | :--- |
| 1 | Involved in creation, modification, transfer or storage of data related to regulatory submission? | **No** | Catatan pemeliharaan internal Engineering, bukan dossier registrasi/submission regulatori langsung. |
| 2 | Automate or control production, packaging, storage or distribution of material and product? | **No** | CMMS bukan sistem otomasi/PLC/MES pengontrol proses produksi. |
| 3 | Used for managing product or material sample? | **No** | Tidak mengelola sampel produk/material QC. |
| 4 | Used to control / monitor utility system (HVAC, purified water, compressed air, etc.)? | **No** | Tidak berfungsi sebagai BMS/SCADA kontrol real-time; pemeliharaan aset utilitas dicakup di No. 5 & 6. |
| 5 | Used for GxP support functions such as calibration, maintenance and qualification of equipment / instrument? | **Yes** | Mengelola Breakdown, PM, Inspection, dan Kartu Mesin peralatan produksi, utilitas, dan fasilitas bangunan (URS Modul 2 & 3). |
| 6 | Support the maintenance of GxP system? | **Yes** | Mendukung pemeliharaan peralatan/sistem GxP dan keterlacakan riwayat penggantian komponen. |
| 7 | Used for calculations or algorithms that might affect product quality, patient safety, and data integrity? | **Yes** | Menghitung downtime/overlap, MTTR, MTBF, PM Compliance, jadwal PM berikutnya, target ganti komponen, dan evaluasi ambang Normal/Peringatan/Kritis inspeksi (URS 2.4, 2.7, 4.7, 4.9). |
| 8 | Generate analytical for testing of finished product or material? | **No** | Tidak menghasilkan data pengujian laboratorium/QC. |
| 9 | Used for product disposition? | **No** | Tidak melakukan release/reject batch produk. |
| 10 | Used to provide training and maintain personnel training records? | **No** | Bukan LMS/sistem catatan pelatihan. |
| 11 | Used for employee identification and access? | **No** | Menggunakan RBAC internal terintegrasi dengan layanan autentikasi perusahaan, bukan sistem HR/access door. |
| 12 | Used for quality trending, quality metrics and/or product review? | **No** | Analitik difokuskan pada performa & biaya pemeliharaan aset, bukan Product Quality Review (PQR) langsung. |
| 13 | Create, modify, maintain, archive, retrieve, or transmit data regarding product quality, patient safety, and data integrity? | **Yes** | Menyimpan & mengarsipkan Electronic Record (Kartu Mesin, WO, hasil PM/Inspeksi, Audit Trail) dengan retensi 5 tahun dan prinsip ALCOA (URS 1.4, 4.2). |
| 14 | Handle the quality management system (CAPA, change control, deviations, audits, risk management, recall)? | **No** | Bukan sistem eQMS. |
| 15 | Handle market complaint or adverse event reporting? | **No** | Tidak menangani keluhan pasar / farmakovigilans. |
| 16 | Create, revise, distribute and store documents related to GxP compliance? | **Yes** | Menyimpan dokumen teknis/manual/datasheet aset, Kartu Mesin, dan laporan maintenance berversi (URS 3.2, 3.3, 5.2, 5.3). |

- **Integrated System / Application:**
  1. Layanan Autentikasi Perusahaan (verifikasi akun, kebijakan sesi & lockout 3x gagal).
  2. Sistem Persediaan / Inventory Perusahaan (kode part, stok, gudang, pengeluaran & pengembalian material berbasis Nomor WO).
  3. Penyimpanan Dokumen (MinIO / Object Storage — foto aset, manual, datasheet, gambar teknis & lampiran WO).
  4. Layanan Notifikasi & Pelaporan (penugasan, eskalasi K3, verifikasi, rework, penutupan WO & laporan shift).
- **Conclusion:** **`[x] GxP`** | `[ ] non-GxP, the assessment stops here`

---

## 4. Categories of Software and Hardware

| Kategori | Pilihan Terpilih | Justifikasi |
| :--- | :--- | :--- |
| **Software Category (GAMP 5)** | `[x] 5 – Custom Application` | Dibangun custom dengan Next.js (TypeScript), Go, dan PostgreSQL sesuai URS PT Bintang Toedjoe. |
| **Hardware Category** | `[x] Category 1 – Standard Hardware Components` | Server Linux standar (acuan awal 4 vCPU, 8 GB RAM, SSD) serta PC/tablet/smartphone standar (URS 2.1). |
| **System Type** | `[x] Closed system` | Lingkungan sistem dikendalikan secara internal oleh PT Bintang Toedjoe (HTTPS, jaringan & autentikasi perusahaan). |
| **Criticality** | `[x] Medium (Indirect impact)` | Dampak tidak langsung terhadap kualitas produk/keselamatan pasien melalui pemeliharaan aset GxP, dan berdampak pada integritas data pemeliharaan. *(Catatan: diskusikan dengan QA bila perlu dinaikkan ke High untuk aset kritis).* |
| **ERES Applicability** | `[x] Electronic Record` & `[x] Electronic Signature` | Kartu Mesin, WO, hasil PM/Inspeksi & Audit Trail adalah Electronic Record; verifikasi WO & approval penonaktifan/perubahan aset kritis melibatkan persetujuan elektronik. |
| **Validation Approach** | `[x] Need computerized system validation` | Wajib CSV penuh (GAMP 5 Cat. 5 + GxP + 21 CFR Part 11). |

---

## 5. 21 CFR Part 11 Applicability (ERES Assessment)

### A. Electronic Record

| Pertanyaan | Jawaban | Catatan Implementasi CMMS |
| :--- | :---: | :--- |
| Able to view the data result? | **Yes** | Dashboard, Kartu Mesin, detail WO, hasil PM & Inspeksi (URS 2.3, 3.4, 4.10). |
| Able to produce printout? | **Yes** | Ekspor PDF/Excel/CSV laporan & cetak/PDF Kartu Mesin beserta periode & waktu pembuatan WIB (URS 5.2). |
| Able to manage recipe or method? | **Yes** | Mengelola template checklist PM, standar parameter & batas Normal/Peringatan/Kritis inspeksi, serta Master Rumus KPI *(bukan recipe proses produksi)*. |
| Able to process data? | **Yes** | Kalkulasi downtime/overlap, MTTR, MTBF, PM Compliance, biaya, kebutuhan teoretis spare part, jadwal PM berikutnya. |
| Able to store data? | **Yes** | PostgreSQL operasional + MinIO Object Storage dengan backup harian & retensi 5 tahun (URS 2.2, 4.2). |
| How does the data use? | **Using electronic records in lieu of paper records** | Menggantikan catatan kertas/terpisah menjadi sistem terpusat; catatan manual hanya digunakan sebagai BCP saat sistem down (URS 4.1). |
| Access controls to restrict use to authorized personnel only? | **Yes** | Autentikasi, lockout 3x gagal, session timeout, dan RBAC Peran × Divisi × Area (URS 1.1–1.2). |
| Audit trail mechanism (who, what, when, why)? | **Yes** | Tabel `audit_trails` append-only/immutable mencatat pelaku, aksi, nilai sebelum/sesudah, timestamp WIB, dan alasan (URS 1.4, ALCOA). |

### B. Electronic Signature

| Pertanyaan | Jawaban | Catatan Implementasi CMMS |
| :--- | :---: | :--- |
| Require electronic signatures and handwritten signatures executed to electronic records? | **Yes** | Persetujuan penutupan WO oleh Supervisor (URS 4.12) serta perubahan/penonaktifan aset kritis oleh Manager Engineering (URS 3.5–3.6). |

---

## 6. Summary of Non-Compliant (Pre-Go-Live Gaps) & Conclusion

| Summary of the Non-Compliant (Item Wajib Dipenuhi Sebelum Go-Live) | Justification / Mitigasi CSV |
| :--- | :--- |
| 1. **Electronic Signature Manifestation & Re-authentication (§11.50 / §11.200)** — Persetujuan WO & aset kritis wajib menampilkan nama lengkap, timestamp WIB, makna persetujuan, dan re-autentikasi bila ditetapkan sebagai E-Signature penuh. | Konfirmasi cakupan E-Signature dengan QA saat penyusunan FS/DS (apakah re-input password saat Approve/Close WO wajib), lalu verifikasi di OQ. |
| 2. **Audit Trail Immutable & SOP Review (§11.10(e) & ALCOA)** — Proteksi database trigger + aplikasi terhadap ubah/hapus audit trail dan prosedur review berkala. | Sudah dirancang di `001_core_auth_rbac_audit.sql` (`trg_audit_trails_immutable`); wajib diverifikasi pada uji negatif OQ + pengesahan SOP review QA. |
| 3. **Access Control & Lockout 3x Gagal (URS 1.1–1.2)** — Integrasi layanan autentikasi perusahaan, timeout sesi, dan isolasi cakupan Divisi/Area. | Diimplementasikan pada Track 1 (`feat/core-auth-gxp` — Gifari) dan diuji pada OQ keamanan. |
| 4. **Integrasi Inventory & Penanda `Data Belum Terkini` (URS 2.2, 4.8–4.9)** — Rekonsiliasi transaksi anti-duplikasi biaya dan penanganan putus koneksi persediaan. | Verifikasi integrasi di lingkungan staging/OQ bersama Divisi Spare Part. |
| 5. **Backup/Restore, Retensi 5 Tahun & BCP (URS 4.1–4.2)** — Backup otomatis (RPO ≤ 4 jam), uji restore (RTO ≤ 8 jam), dan form manual sementara saat down. | Skrip dasar tersedia di `infra/scripts/`; wajib lulus uji drill restore sebelum Go-Live dan didukung SOP BCP resmi. |

### Conclusion
Aplikasi **CMMS Department Engineering PT Bintang Toedjoe Site Pulogadung** adalah **New System**, **GxP**, **GAMP 5 Category 5 (Custom Application)**, **Hardware Category 1**, **Closed System**, **Criticality Medium (Indirect Impact)**, dan **21 CFR Part 11 Applicable (Electronic Record & Electronic Signature)**, sehingga **wajib menjalani Computerized System Validation (CSV)** lengkap (URS, FS/DS, FMEA/Risk Assessment, IQ, OQ, PQ/UAT 4 divisi, RTM, dan VSR) sebelum Go-Live.

---

## 7. Checklist Diskusi Final dengan QA (Sebelum Submit ke FUPD)
1. **Nomor Dokumen**: Ganti placeholder `IRA-CMMS-PGD-001 (Draft Rev. 00)` dengan nomor dokumen resmi dari Document Control / FUPD.
2. **Konfirmasi Criticality**: Pastikan bersama QA Supervisor (*Vany Okky Safitri*) apakah **Medium (Indirect Impact)** sudah sesuai atau perlu **High** untuk aset produksi/utilitas kritis GxP.
3. **Konfirmasi Electronic Signature**: Putuskan bersama QA apakah persetujuan Supervisor/Manager di CMMS menggunakan **Electronic Signature penuh** (re-autentikasi password/PIN di setiap aksi Approve/Close) atau **logged-in approval + audit trail** dengan tanda tangan FUPD pada dokumen laporan cetak.
4. **Konfirmasi Baris 33 (Recipe/Method) & Baris 17 (GxP Documents)**: Verifikasi apakah QA sepakat mencentang **Yes** (pendekatan konservatif untuk template PM/inspeksi/rumus KPI & Kartu Mesin) atau menginginkan **No** dengan justifikasi di Table 3.
