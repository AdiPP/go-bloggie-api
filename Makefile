#!make
ifneq (,$(wildcard ./.env))
    include .env
    export
endif

init-schema:
	goose -dir ./pkg/database/migrations sqlite3 ./foo.db up

init-seeder:
	goose -dir ./pkg/database/seeders sqlite3 ./foo.db up