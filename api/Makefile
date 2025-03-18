# TODO why does this end in "."?
build:
	@go fmt ./...
	@go mod tidy
	@go build -o ./tmp/main ./cmd/.

test:
	@go test ./... -v

exec:
	@./tmp/main

run:
	@go fmt ./...
	@go mod tidy
	@go run ./cmd/main.go

watch:
	@air ./cmd/main.go

.PHONY: build test exec run watch