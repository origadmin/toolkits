GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)

.PHONY: init
# Install general Go tools if needed
init:
	# Add any general Go tools installation here if required for the toolkit

.PHONY: tidy
# Run go mod tidy
tidy:
	go mod tidy

.PHONY: test
# Run tests
test:
	go test ./...

.PHONY: clean
# Clean Go build cache
clean:
	go clean

.PHONY: all
# Default target: init, tidy, and test
all: init tidy test

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
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