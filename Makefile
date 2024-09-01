test:
	docker compose exec server go test -v -coverprofile=c.out -cover ./...
test-coverage:
	docker compose exec server go test -v -coverprofile=c.out -cover ./... && go tool cover -html="c.out"

.PHONY: generate migration-create migrate-up migrate-down migrate-up-test migrate-down-test test test-coverage
