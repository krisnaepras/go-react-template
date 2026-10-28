# Petunjuk AI Coding Agent (Template Go Native + React + Vite)

## Architecture contract
- Gunakan Go `net/http` + `http.ServeMux`, jangan tambahkan Gin/Chi tanpa keputusan tertulis.
- Satu Go module di root, backend di `cmd/`, `internal/platform`, `internal/modules`, `internal/transport`.
- Feature-first frontend di `web/src/features`, shared UI di `web/src/components` bila benar-benar reusable.
- Arah dependensi: HTTP handler -> service -> repository. Service tidak boleh import HTTP/framework.
- PostgreSQL schema hanya boleh diubah dengan migration versi baru; gunakan parameterized queries.
- OpenAPI di `api/openapi/openapi.yaml` menjadi kontrak publik. Update bila mengubah endpoint.
- Jangan edit file generated secara manual. Jalankan generator dan review diff.
- Jangan menambah abstraksi generik atau dependency hanya untuk menghemat sedikit baris kode.

## Workflow setiap perubahan
1. Baca README dan file modul yang terkait. Batasi scope.
2. Nyatakan perubahan schema, API contract, permission, dan test yang akan terdampak.
3. Tambahkan migration jika perlu dan update OpenAPI sebelum implementasi endpoint.
4. Tambahkan aturan bisnis dan unit tests, lalu handler dan UI.
5. Jalankan gofmt, go test, go vet, FE typecheck, build, tests; laporkan yang gagal/tidak dijalankan.
6. Jangan klaim production-ready bila SSO, authorization, CSRF, audit, dan backup belum ada.

## Security rules
- `DEMO_MODE=true` hanya untuk development dan loopback, tanpa data real. Jangan menambah endpoint anonim baru.
- Jangan mengakses data user/tenant berdasarkan ID dari body tanpa memeriksa identitas dari session.
- Jangan menyimpan token dalam localStorage. Jangan log credentials, session, body, atau PII.
- Authorization selalu di backend, di level resource/service; frontend guard hanya UX.
- Batasi payload, validasi input backend, dan hindari SQL string concatenation.
- Jangan mengubah migration lama yang sudah diterapkan ke environment bersama.
- Jangan otomatis menjalankan migrasi destructive atau melakukan reset database produksi.

## Naming
- Prefer nama eksplisit (CreateTask, UpdateTask) daripada BaseHandler/GenericService.
- Uji failure paths dan not-found, bukan hanya happy path.
- Untuk fitur baru, tiru pola Tasks, lalu hapus demo setelah auth production siap.
