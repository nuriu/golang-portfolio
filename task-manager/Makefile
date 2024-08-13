build: generate
	@go build -o bin/tm cmd/api/main.go

test:
	@go test -v ./...

run: build
	@./bin/tm

generate:
	@swag init -d cmd/api,internal/http/handlers,internal/http/models -g main.go --parseDependency true --parseInternal true

cu: generate
	@docker compose up

cd:
	@docker compose down
