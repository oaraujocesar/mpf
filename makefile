generate:
	docker compose exec server sqlc generate
migration-create:
	docker compose exec server migrate create -ext sql -dir database/migrations -seq $(name)
migrate-up:
	docker compose exec server migrate --path=database/migrations --database "postgres://postgres:postgres@database:5432/mpf?sslmode=disable" --verbose up
migrate-down:
	docker compose exec server migrate --path=database/migrations --database "postgres://postgres:postgres@database:5432/mpf?sslmode=disable" --verbose down
test:
	docker compose exec server go test -v -cover ./...

.PHONY: migration-create migrate-up migrate-down test generate
