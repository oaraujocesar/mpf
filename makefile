migration: 
	migrate --path=internal/database/migrations --database "postgres://postgres:postgres@localhost:5432/mpf?sslmode=disable" --verbose up