build:
	@golangci-lint fmt
	@go mod tidy
	@go build -o ./tmp/main ./cmd/.

# lint should happen on ci/cd in future.
# for speed of builds leaving separate for now.
lint:
	@golangci-lint run
#	@go vet ./...
#	@staticcheck ./...
#	@revive --config .revive.toml ./..

vulchk:
	@govulncheck ./...

test:
	@go test ./... -v
#   NOTE integration tests require the server to be running.
#   @go test -v ./internal/test_integration/...

exec:
	@./tmp/main

run:
	@golangci-lint fmt
	@go mod tidy
	@go run ./cmd/main.go

watch:
	@air ./cmd/main.go

.PHONY: build test exec run watch