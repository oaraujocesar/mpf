migration-create:
	migrate create -ext sql -dir database/migrations -seq $(name)
migrate-up: 
	migrate --path=database/migrations --database "postgres://postgres:postgres@database:5432/mpf?sslmode=disable" --verbose up
migrate-down: 
	migrate --path=database/migrations --database "postgres://postgres:postgres@database:5432/mpf?sslmode=disable" --verbose down

.PHONY: migration-create migrate-up migrate-down