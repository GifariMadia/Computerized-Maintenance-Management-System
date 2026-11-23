# Feature Folder: Core Auth, RBAC & Audit Trail (Track 1)

- **Branch**: `feat/core-auth-gxp`
- **PIC**: Gifari (IT)
- **URS**: `1.1`, `1.2`, `1.3`, `1.4`, `ALCOA 1.1–1.5`
- **Komponen yang dibangun di folder ini**:
  - `LoginForm.tsx` (handle lockout 3x gagal & pesan session timeout)
  - `UserScopeMatrix.tsx` (atur Peran × Divisi × Area × Approve/Verify/Export)
  - `AuditTrailTable.tsx` (filter berdasarkan `Asset ID`, `WO No`, pelaku, aksi, alasan, diff before/after)
