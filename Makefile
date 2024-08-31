.PHONY: test lint image

TESTS ?= ./...

test:
	@go test -count 1 -v $(TESTS)

# Test with coverage (CI)
testCI:
	@go test ./... --cover -v

# Checks code with golangci-lint linters
lint:
	@golangci-lint run
	@hadolint Dockerfile

# Run the api
run:
	@go run ./cmd/rd-clone-api

# Migrate the DB
migrate:
	@go run ./cmd/migrate

# Create Docker image
image:
	@docker build -t reddit-clone .