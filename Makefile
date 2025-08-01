# Root Makefile for the Monorepo

.PHONY: all build test lint fmt generate clean init deps help

# Define subdirectories containing Go modules
MODULES := runtime toolkits contrib tools/origen

# Default target
all: build test lint

# Build all modules
build:
	@echo "Building all Go modules..."
	go build ./...

# Run tests for all modules
test:
	@echo "Running tests for all Go modules..."
	go test ./...

# Run golangci-lint for all modules
lint:
	@echo "Running golangci-lint..."
	golangci-lint run ./...

# Format all Go code
fmt:
	@echo "Formatting all Go code..."
	go fmt ./...

# Generate code (Protobuf, Go generate, etc.)
generate:
	@echo "Generating code for all modules..."
	# Call generate targets in submodules
	$(foreach module,$(MODULES),$(MAKE) -C $(module) generate;)
	# Add any other top-level go generate commands here
	go generate ./...

# Clean build artifacts
clean:
	@echo "Cleaning build artifacts..."
	go clean ./...
	rm -rf bin dist # Example: remove common build output directories

# Initialize development environment
init:
	@echo "Initializing development environment..."
	# Install common Go tools
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
	go install github.com/go-kratos/kratos/cmd/protoc-gen-go-http/v2@latest
	go install github.com/google/gnostic/cmd/protoc-gen-openapi@latest
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/envoyproxy/protoc-gen-validate@latest
	go install github.com/bufbuild/buf/cmd/buf@latest
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-lint@latest
	go install github.com/bufbuild/buf/cmd/protoc-gen-buf-breaking@latest
	# Install golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.59.1 # Example version
	# Install pre-commit
	pip install pre-commit
	pre-commit install

# Update Protobuf dependencies
deps:
	@echo "Updating Protobuf dependencies..."
	# Call deps targets in submodules or run top-level buf commands
	$(foreach module,$(MODULES),$(MAKE) -C $(module) deps;)

# Show help
help:
	@echo ""
	@echo "Usage:"
	@echo "  make [target]"
	@echo ""
	@echo "Targets:"
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
		helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
