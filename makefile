generate:
	docker compose exec server sqlc generate
migration-create:
	docker compose exec server migrate create -ext sql -dir database/migrations -seq $(name)
migrate-up:
	docker compose exec server migrate --path=database/migrations --database "postgres://postgres:postgres@database:5432/mpf?sslmode=disable" --verbose up
migrate-down:
	docker compose exec server migrate --path=database/migrations --database "postgres://postgres:postgres@database:5432/mpf?sslmode=disable" --verbose down
migrate-up-test:
	docker compose exec server migrate --path=database/migrations --database "postgres://postgres:postgres@database:5432/mpf_test?sslmode=disable" --verbose up
migrate-down-test:
	docker compose exec server migrate --path=database/migrations --database "postgres://postgres:postgres@database:5432/mpf_test?sslmode=disable" --verbose down
create-test-db:
	docker compose exec database psql -U postgres -c "CREATE DATABASE mpf_test"
test:
	docker compose exec server go test -v -coverprofile=c.out -cover ./...
test-coverage:
	docker compose exec server go test -v -coverprofile=c.out -cover ./... && go tool cover -html="c.out"

.PHONY: generate migration-create migrate-up migrate-down migrate-up-test migrate-down-test test test-coverage
