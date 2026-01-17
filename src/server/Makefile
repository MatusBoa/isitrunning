DC=docker compose -f $(DC_FILE)

include .env

build:
	$(DC) build
start:
	$(DC) up -d

server-migration: start
	@test -n "$(name)" || (echo "Usage: make migration name=<migration_name>" && exit 1)
	$(DC) exec -e "GOOSE_DRIVER=postgres" server-ws goose create -dir internal/database/pgsql/migrations $(name) sql
server-migrate: start
	$(DC) exec -e "GOOSE_DRIVER=postgres" server-ws goose -dir internal/database/pgsql/migrations "${DB_DSN}" up
server-jwt-keys: start
	@test -n "$(name)" || (echo "Usage: make api-jwt-tokens name=jwt" && exit 1)
	$(DC) exec server-ws openssl genpkey -algorithm RSA -outform PEM -out keys/$(name)-private.pem -pkeyopt rsa_keygen_bits:2048
	$(DC) exec server-ws openssl rsa -in keys/$(name)-private.pem -pubout -out keys/$(name)-public.pem
