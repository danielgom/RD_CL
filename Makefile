.PHONY: test lint image

test: ## make test TESTS="-run Method/Suite ./internal/file"
	@go test -count 1 -v -race $(if $(TESTS),$(TESTS),./...)

# Test with coverage (CI)
testCI:
	@go test ./... --cover -v

# Checks code with golangci-lint linters
lint:
	@golangci-lint run --timeout 3m --fix

# Run the api
run:
	@go run ./cmd/rd-clone-api

# Migrate the DB
migrate:
	@go run ./cmd/migrate

# Create Docker image
image:
	@docker build -t reddit-clone .