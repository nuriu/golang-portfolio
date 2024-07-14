build: generate
	@go build -o bin/rest-api cmd/app/main.go

test:
	@go test -v ./...

run: build
	@./bin/rest-api

generate:
	@swag init -d cmd/app,api/handlers,api/models -g main.go --parseDependency true --parseInternal true
