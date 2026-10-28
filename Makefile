SHELL := /bin/sh
DB_URL ?= postgres://office:office_dev_only@localhost:5432/office?sslmode=disable

.PHONY: help db-up db-down db-logs migrate-up migrate-down api web generate-sqlc generate-api-go generate-api-web test build fmt check
help:
	@echo 'db-up | migrate-up | api | web | test | build | fmt | check'

db-up:
	docker compose up -d db

db-down:
	docker compose down

db-logs:
	docker compose logs -f db

# Install goose: go install github.com/pressly/goose/v3/cmd/goose@latest
migrate-up:
	goose -dir db/migrations postgres "$(DB_URL)" up
migrate-down:
	goose -dir db/migrations postgres "$(DB_URL)" down

api:
	go run ./cmd/api
web:
	cd web && npm run dev

# Optional code generation: do not manually edit generated output.
generate-sqlc:
	sqlc generate
# Install oapi-codegen v2 before running this target.
generate-api-go:
	oapi-codegen -config oapi-codegen.yaml api/openapi/openapi.yaml
# Install web deps first.
generate-api-web:
	cd web && npm run generate:api

test:
	go test ./...
	cd web && npm run test -- --run

build:
	mkdir -p bin
	go build -o bin/office-api ./cmd/api
	cd web && npm run build

fmt:
	gofmt -w cmd internal

check:
	go vet ./...
	go test ./...
	cd web && npm run typecheck && npm run test -- --run && npm run build
