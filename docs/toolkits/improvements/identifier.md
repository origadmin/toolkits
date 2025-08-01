# Identifier Module Improvements & Todo List

This document outlines identified areas for improvement and pending tasks within the `identifier` module.

---

## 1. Redundant Interface Methods (`GenerateString/Number`, `ValidateString/Number`)

### Problem Description

The `StringGenerator` and `NumberGenerator` interfaces contain redundant methods (`GenerateString()`/`GenerateNumber()`
and `ValidateString()`/`ValidateNumber()`) that are aliases for `Generate()` and `Validate()` respectively. This
redundancy adds unnecessary verbosity to the interfaces.

### Current State Analysis

- **Relevant Files**: `identifier/identifier.go` (where interfaces are defined)
- **Relevant Code Snippets**:
    ```go
    type StringGenerator interface {
        // ...
        Generate() string
        GenerateString() string // Alias for Generate()
        Validate(id string) bool
        ValidateString(id string) bool // Alias for Validate()
        // ...
    }

    type NumberGenerator interface {
        // ...
        Generate() int64
        GenerateNumber() int64 // Alias for Generate()
        Validate(id int64) bool
        ValidateNumber(id int64) bool // Alias for Validate()
        // ...
    }
    ```

### Proposed Solution(s)

Remove the redundant `GenerateString()`, `GenerateNumber()`, `ValidateString()`, and `ValidateNumber()` methods from
their respective interfaces. Only keep `Generate()` and `Validate()`.

### Expected Impact

- **Code Simplification**: Reduce interface verbosity and improve clarity.
- **Improved Readability**: Make the interfaces more concise and easier to understand.

### Verification Plan

- **Code Refactoring**: Update all implementations of these interfaces to use the non-aliased methods.
- **Unit Tests**: Ensure all existing tests continue to pass after the refactoring.

---

## 2. Global State Modification in Test `init()` Functions

### Problem Description

Several test files within the `identifier` module use `init()` functions to set global default generators (
`identifier.SetDefaultString()` or `identifier.SetDefaultNumber()`). This practice can lead to test instability and
unpredictable behavior due to global state modification and uncertain `init()` execution order.

### Current State Analysis

- **Relevant Files**: `identifier/ksuid/ksuid_test.go`, `identifier/shortid/shortid_test.go`,
  `identifier/snowflake/snowflake_test.go`, `identifier/sonyflake/sonyflake_test.go`, `identifier/ulid/ulid_test.go`,
  `identifier/uuid/uuid_test.go`, `identifier/xid/xid_test.go`
- **Relevant Code Snippet Example** (from `ksuid_test.go`):
    ```go
    func init() {
        identifier.SetDefaultString(New())
    }
    ```

### Proposed Solution(s)

Refactor the tests to avoid modifying global state in `init()` functions. Instead, each test function should:

1. Explicitly set up its required default generator at the beginning of the test.
2. Use `t.Cleanup()` to reset the global state to its original value after the test completes, ensuring test isolation.

### Expected Impact

- **Improved Test Stability**: Tests will be more reliable and less prone to flakiness.
- **Better Test Isolation**: Each test will run in a clean environment, preventing side effects from other tests.
- **Easier Debugging**: Issues will be easier to pinpoint as they won't be caused by unexpected global state changes.

### Verification Plan

- **Unit Tests**: Rerun all tests after refactoring to ensure they pass consistently.

---

## 3. Unused `github.com/goexts/generic` Dependency

### Problem Description

The `identifier` module's `go.mod` file lists `github.com/goexts/generic` as a dependency, but its usage is not apparent
in the provided code snippets.

### Current State Analysis

- **Relevant File**: `identifier/go.mod`
- **Observation**: `github.com/goexts/generic v0.2.4` is listed as a direct dependency.

### Proposed Solution(s)

Confirm if `github.com/goexts/generic` is indeed used within the `identifier` module. If it is not, remove it from
`go.mod` to reduce the project's dependency footprint.

### Expected Impact

- **Reduced Dependency Footprint**: Simplify the project's dependency graph.
- **Improved Build Times**: Potentially reduce build times by eliminating unnecessary dependency resolution.

### Verification Plan

- **Code Review**: Thoroughly review all code within the `identifier` module to confirm the usage of
  `github.com/goexts/generic`.
- **Dependency Check**: After removal (if applicable), run `go mod tidy` to ensure no build issues arise.
