# Arsitektur dan Batas Modul

- **Monorepo modular monolith**: satu Go module, satu frontend Vite, satu database.
- `cmd/api`: composition root dan graceful shutdown.
- `internal/platform`: code infrastructure generik, tidak memiliki domain.
- `internal/modules/<feature>`: domain + service + repository interface atau implementasi.
- `internal/transport/httpapi`: handler, routing, middleware, status mapping.
- `db/migrations`: schema source of truth; `db/queries`: named queries untuk sqlc.
- `api/openapi/openapi.yaml`: API contract source of truth. Generate clients dari sini.
- `web/src/features`: UI per fitur. `web/src/app`: router dan providers. `web/src/api`: transport client.

Arah dependency: `transport -> modules -> repository -> PostgreSQL`. Jangan melakukan import HTTP pada domain service, jangan query database langsung dari React, dan jangan membiarkan frontend memutuskan authorization.

## Sample tasks module

Untuk menjaga starter benar-benar sederhana dan dapat dipelajari, modul tasks menggunakan SQL pgx langsung serta kode API client kecil. Config `sqlc.yaml`, `oapi-codegen.yaml`, dan `orval.config.ts` disiapkan untuk migrasi ke code generation, **tetapi hasil codegen belum menjadi dependensi runtime**. Setelah memutuskan kontrak, gunakan output codegen sebagai satu-satunya API layer dan jangan mempertahankan implementasi SQL/typed client duplikat dalam production.

## Keamanan

Demo API sengaja **tidak menggunakan autentikasi** karena hanya data contoh untuk development. Go hanya mendaftarkan endpoint tasks ketika `DEMO_MODE=true` dan `APP_ENV=development` (validated during startup; HTTP_ADDR also must be loopback). Jangan expose development server ke internet. Untuk production: implementasikan OIDC login, server-side sessions, CSRF protection, RBAC/resource permissions, audit log, rate limiting, backups, secure headers dan threat modeling sebelum membuka resource bisnis.
