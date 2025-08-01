# Project Guidelines

## 1. Project Overview
This document outlines the guidelines and conventions for developing within the `framework` monorepo. Our goal is to maintain a consistent, high-quality, and easily maintainable codebase.

### Core Technologies
*   **Go Language:** Primary programming language.
*   **Monorepo:** All core components (`runtime`, `toolkits`, `contrib`, `tools/origen`) are managed within a single Git repository.
*   **Go Workspaces:** Utilized for seamless local development across multiple Go modules within the monorepo.

## 2. Development Environment Setup

### Go Version
Ensure you are using **Go 1.23** or later.
```bash
go version
```

### Recommended IDE
*   **IntelliJ IDEA (with Go Plugin):** Provides excellent Go language support, refactoring, and debugging capabilities.

### Essential Tools
Install the following tools globally or ensure they are accessible in your PATH:
*   **Git:** For version control.
*   **Protoc:** Protocol Buffer compiler.
*   **Buf:** For Protobuf linting, formatting, and breaking change detection.
*   **Go tools:**
    ```bash
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
    ```

## 3. Code Style & Conventions

### Go Language
*   **Formatting:** Always run `go fmt ./...` before committing.
*   **Linting:** Use `go vet ./...` and consider additional linters (e.g., `golangci-lint`) to ensure code quality.
*   **Naming:** Follow Go's official naming conventions (e.g., `CamelCase` for exported names, `camelCase` for unexported names).

### Protobuf
*   Follow the [Buf Style Guide](https://docs.buf.build/style-guide/rules).

## 4. Git Workflow

### Branching Strategy
We use a Gitflow-like branching strategy:
*   `main`: Production-ready code.
*   `develop`: Integration branch for new features.
*   `feature/*`: For new features.
*   `bugfix/*`: For bug fixes.
*   `release/*`: For preparing new releases.
*   `hotfix/*`: For urgent production fixes.

### Commit Message Convention
All commit messages must follow the [Angular Commit Message Convention](https://github.com/angular/angular/blob/main/CONTRIBUTING.md#commit-message-format).
Example:
```
feat(scope): Add new feature X

This commit introduces feature X which does Y and Z.
BREAKING CHANGE: Old API is no longer supported.
```

### Pull Request (PR) Process
1.  Create a new branch from `develop`.
2.  Implement changes and commit frequently with descriptive messages.
3.  Ensure all tests pass and code is linted.
4.  Create a Pull Request to `develop`.
5.  Request reviews from at least two team members.
6.  Address feedback and iterate.
7.  Once approved, squash and merge the PR.

## 5. Monorepo Specifics

### Go Workspaces
The `go.work` file in the root directory defines the modules within this monorepo.
To work with the monorepo:
```bash
go work init
go work use ./runtime ./toolkits ./contrib ./tools/origen .
```
This allows `go build`, `go test`, `go run`, and `go mod tidy` to operate across all defined modules without needing `replace` directives in individual `go.mod` files for inter-module dependencies.

### Submodule Management (`git subtree`)
*   **Adding/Updating:** Use `git subtree add` and `git subtree pull` for managing external repositories as subtrees.
*   **Pushing Changes Back:** Refer to `MONOREPO_MIGRATION_PLAN.md` for detailed instructions on pushing changes from a subtree back to its original repository, especially concerning `go.mod` handling.

### Unified Build Process
A root `Makefile` will be used to orchestrate builds, tests, and other common tasks across all modules.

## 6. Testing
*   **Unit Tests:** Write unit tests for all new code. Run with `go test ./...`.
*   **Integration Tests:** For interactions between components.
*   **Test Coverage:** Aim for high test coverage.

## 7. Documentation
*   Maintain clear and concise documentation for all code, APIs, and features.
*   Update `README.md` files for each module and the root monorepo as needed.
