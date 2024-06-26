run:
	@./bin/api


build:
	@go build -o bin/api

test:
	@go test -v ./...
