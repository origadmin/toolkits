# Gemini Agent Guidelines for this Project

This document serves as a specific guideline for the Gemini agent interacting with this project.

## Project Information (as of 2025-08-01)

### Project Root Directory
$pwd

### Go Version
`go1.13 windows/amd64`

### Monorepo Structure
This project is a Go monorepo managed with Go Workspaces. It includes the following sub-modules:
*   `runtime/`
*   `toolkits/`
*   `contrib/`
*   `tools/origen/`

### Go Workspaces Configuration
The `go.work` file in the root directory defines the modules:
```
go 1.22

use (
	./runtime
	./toolkits
	./contrib
	./tools/origen
	.
)
```

### Unified Build and Development Tools
*   **Makefile:** A root `Makefile` (`./Makefile`) is used to orchestrate common development tasks (build, test, lint, generate, etc.).
*   **golangci-lint:** Configured via `.golangci.yml` for unified code quality checks.
*   **pre-commit hooks:** Configured via `.pre-commit-config.yaml` for automated checks before commits.

### Version Management
This monorepo adopts a **unified versioning strategy**, where the entire monorepo shares a single version number based on Git tags in the root. Releases are automated using `GoReleaser` configured in `.goreleaser.yaml` and triggered by GitHub Actions (`.github/workflows/release.yml`).

### Documentation Structure
Project documentation is organized under the `docs/` directory in the root. Module-specific analysis and improvement documents are located under `docs/<module_name>/` (e.g., `docs/runtime/`, `docs/toolkits/`). General project guidelines are in `docs/PROJECT_GUIDELINES.md`.

### GitHub Templates
Standardized GitHub templates (e.g., issue templates, PR templates) are managed as a `git subtree` under `docs/templates/github/`, sourced from `https://github.com/origadmin/.github`.