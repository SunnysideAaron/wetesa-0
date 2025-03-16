# build (compile) code
# TODO why does this end in "."?
build:
	@go build -o ./tmp/main ./cmd/.

test:
	@go test ./... -v

exec:
	@./tmp/main

run:
	@go run ./cmd/main.go

watch:
	@air ./cmd/main.go

.PHONY: build test exec run watch