#!make
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

default_migration_cmd := -dir database/migrations -table ${MIGRATION_TABLE_NAME} postgres "host=${DB_HOST} port=${DB_PORT} user=${DB_USERNAME} password=${DB_PASSWORD} dbname=${DB_DATABASE} sslmode=disable"

module-new:
	touch domain/$$name.go && mkdir -p $$name/entities && mkdir -p $$name/delivery/http && mkdir -p $$name/delivery/grpc && mkdir -p $$name/repositories && mkdir -p $$name/usecases
	
run:
	go run main.go

dev:
	CompileDaemon -build="go build -buildvcs=false -o go-starter-template" -command="./go-starter-template"

init-schema:
	go run cmd/command/command.go init-schema

build:
	go build -o go-starter-template main.go

migrate-up:
	goose $(default_migration_cmd) up

migrate-up-by-one:
	goose $(default_migration_cmd) up-by-one

migrate-up-to:
	goose $(default_migration_cmd) up-to $$version

migrate-down:
	goose $(default_migration_cmd) down

migrate-down-to:
	goose $(default_migration_cmd) down-to $$version

migrate-redo:
	goose $(default_migration_cmd) redo

migrate-reset:
	goose $(default_migration_cmd) reset

migrate-status:
	goose $(default_migration_cmd) status

migrate-version:
	goose $(default_migration_cmd) version

migrate-create:
	goose $(default_migration_cmd) create $$name sql

migrate-fix:
	goose $(default_migration_cmd) fix
	