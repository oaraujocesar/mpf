generate:
	docker compose exec server sqlc generate
migration-create:
	migrate create -ext sql -dir database/migrations -seq $(name)
migrate-up:
	docker compose exec server migrate --path=database/migrations --database "postgres://postgres:postgres@database:5432/mpf?sslmode=disable" --verbose up
migrate-down:
	docker compose exec server migrate --path=database/migrations --database "postgres://postgres:postgres@database:5432/mpf?sslmode=disable" --verbose down
test:
	go test -v -cover ./...

.PHONY: migration-create migrate-up migrate-down test
