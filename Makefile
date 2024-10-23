SHELL := /bin/bash
BASEDIR = $(shell pwd)
SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
VERSION := v1.0.0
ENV := dev
VersionDir = github.com/origadmin/toolkits/version

# gitHash The current commit id is the same as the gitCommit result
gitHash = $(shell git rev-parse HEAD)
# gitBranch The current branch name
gitBranch = $(shell git rev-parse --abbrev-ref HEAD)
# gitTag The latest tag name
gitTag = $(shell \
			if [ "`git describe --tags --abbrev=0 2>/dev/null`" != "" ]; then \
				git describe --tags --abbrev=0; \
			else \
				git log --pretty=format:'%h' -n 1; \
			fi)
# Same as the previous gitHash
gitCommit = $(shell git log --pretty=format:'%H' -n 1)
# gitTreeState The state of git tree, clean or dirty
gitTreeState = $(shell if git status | grep -q 'clean'; then echo clean; else echo dirty; fi)
# buildDate The time of build
buildDate = $(shell TZ=Asia/Shanghai date +%FT%T%z)
# buildDate = $(shell TZ=Asia/Shanghai date +%F\ %T%z | tr 'T' ' ')

ifeq ($(ENV), dev)
    BUILD_FLAGS = -race
endif

ifeq ($(ENV), pro)
    LDFLAGS = -w
endif


LDFLAGS += -X "${VersionDir}.gitTag=${gitTag}"
LDFLAGS += -X "${VersionDir}.buildDate=${buildDate}"
LDFLAGS += -X "${VersionDir}.gitCommit=${gitCommit}"
LDFLAGS += -X "${VersionDir}.gitTreeState=${gitTreeState}"
LDFLAGS += -X "${VersionDir}.gitBranch=${gitBranch}"
LDFLAGS += -X "${VersionDir}.version=${VERSION}"

.PHONY: all
all: lint version

.PHONY: version
version:
	go run -v -ldflags '$(LDFLAGS)' $(BUILD_FLAGS) -gcflags=all="-N -l" ./cmd/version

.PHONY: lint
lint:
	go fmt ./...
	go vet ./...
	goimports -w .

 cover:
	go test ./... -v -short -coverprofile .coverage.txt
	go tool cover -func .coverage.txt

# Example Delete all.sa,.sb,.sc,... extensions from a specified directory (including subdirectories).
# Hidden files of.sz (beginning with. Followed by _ or.) Delete all
.PHONY: clean
clean:
	rm -f east_money || true
	find . -name "[._]*.s[a-w][a-z]" | xargs -i rm -f {} || true
	rm -rf ./log || true

help:
	@echo "make build   - compile the source code"
	@echo "make clean   - remove binary file and vim swp files"
	@echo "make lint    - run go tool 'fmt', 'vet', 'goimports', 'golangci-lint' "

.DEFAULT_GOAL := help
