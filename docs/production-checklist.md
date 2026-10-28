# Checklist sebelum dipakai di lingkungan kantor

- [ ] Implementasikan OIDC/SSO + logout, session invalidation, secure HttpOnly cookies.
- [ ] Tambahkan authorization deny-by-default dan tes antar-user/tenant di setiap service.
- [ ] Tambahkan CSRF protection untuk cookie-based auth + pemeriksaan Origin.
- [ ] Pastikan DEMO_MODE=false dan endpoint demo tidak diregistrasi.
- [ ] Ganti contoh credential/database; gunakan secret manager/env injected.
- [ ] Tutup port database dari publik; TLS/HTTPS reverse proxy.
- [ ] Setup migrations di pipeline terkontrol sebelum rollout aplikasi.
- [ ] Backup terjadwal + simulasi restore; strategi point-in-time recovery bila dibutuhkan.
- [ ] Observability: structured logs, error alerts, metrics, tracing bila diperlukan.
- [ ] Batasi request size, rate limits per actor, timeout DB dan outbound calls.
- [ ] Gunakan npm ci + committed lockfile setelah npm install pertama.
- [ ] Pin/review Docker images dan package updates; jalankan Dependabot/Renovate.
- [ ] Go tests + race checks + FE typecheck/build + E2E smoke test di CI.
- [ ] Test staging termasuk rollback dan failover DB sebelum production.
- [ ] Tinjau isu privasi dan data residency sebelum menambah AI Assistant/browser tools.
