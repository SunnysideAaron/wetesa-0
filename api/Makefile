# build (compile) code
# TODO why does this end in "."?
build:
	@go build -o ./tmp/main ./cmd/.

test:
	@go test ./... -v

exec:
	@./tmp/main

watch:
	@air ./cmd/main.go

.PHONY: build test exec watch