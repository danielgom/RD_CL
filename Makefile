.PHONY: test lint image

# Test with coverage (dev/local environment)
testL: genM
	@go test ./... --cover -v

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