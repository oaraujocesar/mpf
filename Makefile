test:
	docker compose exec server go test -v -coverprofile=c.out -cover ./...
test-coverage:
	docker compose exec server go test -v -coverprofile=c.out -cover ./... && go tool cover -html="c.out"
compose-up:
	docker compose up -d
generate:
	docker compose exec server sqlc generate

.PHONY: test test-coverage compose-up generate
