build: generate
	@go build -o bin/tm cmd/api/main.go

test:
	@go test -v ./...

cover:
	@go test -coverprofile=bin/coverage.out ./...
	@go tool cover -html=bin/coverage.out -o=bin/coverage.html

run: build
	@./bin/tm

generate:
	@swag init -d cmd/api,internal/http/handlers,internal/http/models -g main.go --parseDependency true --parseInternal true

cu: generate
	@docker compose up

cud: generate
	@docker compose up -d

cd:
	@docker compose down
