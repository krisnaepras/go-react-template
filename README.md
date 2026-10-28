# Office Starter — Go Native + React + Vite + PostgreSQL

Aplikasi starter reusable untuk web internal kantor. **Berisi source code nyata**, bukan hanya rancangan folder: Go `net/http` API, React SPA, PostgreSQL, migration, CRUD contoh, OpenAPI, unit tests, Dockerfiles, CI, dan instruksi AI agent.

> **Status: development starter, bukan aplikasi production siap pakai.** Demo tasks sengaja anonim dan hanya didaftarkan saat `APP_ENV=development` dan `DEMO_MODE=true` **dan API hanya mendengarkan di loopback**. Jangan gunakan untuk menyimpan data kantor nyata sebelum menambahkan SSO/OIDC, authorization, CSRF, dan audit.

## Requirements

- Go **1.27+**, Node.js **24+**, npm, Docker + Docker Compose
- `goose` CLI untuk migration (install: `go install github.com/pressly/goose/v3/cmd/goose@latest`)
- Koneksi internet pertama kali untuk mengunduh Go modules dan npm packages

## Quick start (Linux/macOS)

```bash
cp .env.example .env
# Shell tidak otomatis memuat file .env; export env hanya dalam session lokal ini.
set -a; . ./.env; set +a
docker compose up -d db
# Install goose jika belum ada; pastikan `goose` ada di PATH.
goose -dir db/migrations postgres "$DATABASE_URL" up
# terminal 1:
go run ./cmd/api
# terminal 2:
cd web && npm install && npm run dev
```

- UI: http://127.0.0.1:5173
- API liveness: http://127.0.0.1:8080/healthz
- API readiness: http://127.0.0.1:8080/readyz
- API contoh: http://127.0.0.1:8080/api/v1/tasks
- `DEMO_MODE` tidak dapat dijalankan dengan `HTTP_ADDR=:8080`; hanya loopback diperbolehkan. Untuk production, set `DEMO_MODE=false` dan alamat bind yang sesuai deployment.
- Vite meneruskan `/api` ke Go API pada local dev. Tidak perlu CORS atau token di Vite.

**Windows:** jalankan environment variables dari PowerShell secara manual atau gunakan bash pada WSL. Perintah `set -a` khusus shell POSIX.

## Penting: lockfiles dan code generation

Arsip ini dibuat dalam lingkungan tanpa akses ke npm registry / Go module proxy. Karena itu:
- Dependensi **belum dapat diunduh/di-compile di lingkungan pembuatan template**, sehingga hasil build end-to-end belum terverifikasi.
- **`web/package-lock.json` belum tersedia.** Jalankan `npm install` pertama kali, review dependency tree, lalu commit hasil `package-lock.json`. Setelah itu gunakan `npm ci` di CI/Docker untuk reproducible builds.
- **`go.sum` belum tersedia.** Jalankan `go mod tidy` (butuh internet) dan commit file `go.sum` yang dihasilkan.
- Fitur template dasar menggunakan pgx SQL langsung dan API client manual. Konfigurasi `sqlc`, `oapi-codegen`, dan Orval tersedia sebagai langkah **opsional** untuk mengadopsi codegen. Jangan mengaktifkan codegen tanpa mengubah adapter yang memakai kode hasil generate, untuk menghindari dua sumber implementasi.
- Review versi paket sebelum production dan jalankan tests setelah resolusi dependensi.

## Commands

```bash
make db-up                 # start PostgreSQL only
make migrate-up            # needs goose and DB_URL (or uses demo default)
make api                   # start Go
make web                   # start Vite after npm install
make check                 # tests, vet, FE typecheck/build
make generate-sqlc         # optional, needs sqlc
make generate-api-go       # optional, needs oapi-codegen
make generate-api-web      # optional, needs Orval dependencies
```

## Struktur singkat

```text
cmd/api/                       Go entrypoint
internal/platform/             config, DB pool, JSON HTTP utilities
internal/modules/tasks/        domain, service, repository, tests
internal/transport/httpapi/    routes, handlers, middleware
api/openapi/                   OpenAPI source contract
db/migrations/                 goose Up/Down
db/queries/                    example sqlc queries
web/src/app/                   TanStack Router & layout
web/src/features/              React features per domain
web/src/api/                   API transport adapter
deploy/                         Dockerfiles + Caddy SPA proxy
.github/workflows/             CI
AGENTS.md                      AI coding agent rules
```

## Deployment

`compose.yaml` hanya berisi PostgreSQL untuk development lokal. Untuk deployment multi-container, `deploy/Dockerfile.api`, `deploy/Dockerfile.web`, dan `deploy/Caddyfile` adalah contoh awal. Sediakan service bernama `api` di network Caddy, jalankan migration melalui job terpisah **sebelum** merilis API, dan simpan secrets di secret manager. Jangan menggunakan database password contoh di production.

Sebelum deployment production, baca `docs/production-checklist.md`. Panduan arsitektur lebih lengkap ada di `docs/blueprint.md` dan `docs/architecture.md`.

## Upgrade template ke proyek baru

1. Gunakan repo sebagai GitHub Template / copy folder tanpa `.git`.
2. Ubah module path `example.com/office-template` secara menyeluruh sebelum `go mod tidy`.
3. Ganti nama `web/package.json`, branding, API metadata, dan env.
4. Buat modul domain baru mengikuti pola `internal/modules/tasks`.
5. Tambahkan SSO/OIDC, authorization, dan audit **sebelum** data nyata.
6. Buat lockfiles, jalankan tests, dan gunakan CI sebagai release gate.

Sumber: https://go.dev/doc/modules/layout · https://go.dev/doc/go1.22 · https://vite.dev/guide/ · https://docs.sqlc.dev/ · https://www.postgresql.org/docs/ · https://spec.openapis.org/oas/latest.html

### Helper untuk membuat proyek baru

Jalankan pada salinan template baru (sebelum mulai coding):

```bash
python3 scripts/init_project.py --module github.com/acme/hr-platform --title "HR Platform"
```

Helper mengganti module path dan branding umum. Review semua nama, OpenAPI, konfigurasi, credentials, serta environment sebelum dipakai. Script hanya bisa dijalankan sekali pada starter yang belum diinisialisasi.
