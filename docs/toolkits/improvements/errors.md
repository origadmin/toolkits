# Errors Module Improvements & Todo List

This document outlines identified areas for improvement and pending tasks within the `errors` module.

---

## 1. Incomplete JSON Serialization in `Error.Error()`

### Problem Description

The `Error()` method of the `Error` struct currently uses `fmt.Sprintf` to manually construct a JSON string. This
approach only includes `ID`, `Code`, and `Detail` fields, omitting other important fields like `Cause`, `Metadata`,
`Timestamp`, `StackTrace`, and `Any`.

### Current State Analysis

- **Relevant File**: `errors/errors.go`
- **Relevant Code Snippet**:
    ```go
    func (e *Error) Error() string {
        // ...
        return fmt.Sprintf(`{"id":"%s","code":%d,"detail":"%s"}`, e.ID, e.Code, e.Detail)
    }
    ```
- **Observation**: Other fields in the `Error` struct are tagged with `json:"..."` but are not included in the `Error()`
  method's output.

### Proposed Solution(s)

Refactor the `Error()` method to use the `encoding/json` package to properly marshal the entire `Error` struct into a
JSON string. This will ensure all serializable fields are included, providing a more complete and consistent error
representation.

### Expected Impact

- **Completeness**: All relevant error details will be included in the string representation.
- **Consistency**: The `Error()` method's output will align with the `json` tags defined in the `Error` struct.
- **Improved Debuggability**: More comprehensive error information will be available for logging and debugging.

### Verification Plan

- **Unit Tests**: Add or modify tests to verify that the JSON output from `Error()` includes all expected fields and is
  valid JSON.

---

## 2. Redundant `Error.As` Method Implementation

### Problem Description

The `Error` struct implements its own `As` method, which is redundant given that Go's standard library `errors.As`
function already provides this functionality and works correctly with errors that implement the `Unwrap` method.

### Current State Analysis

- **Relevant File**: `errors/errors.go`
- **Relevant Code Snippet**:
    ```go
    func (e *Error) As(target interface{}) bool {
        // ... manual reflection-based implementation ...
    }
    ```
- **Observation**: The `Error` struct already implements `Unwrap() error`, which is sufficient for `errors.As` to
  traverse error chains.

### Proposed Solution(s)

Remove the custom `As` method from the `Error` struct. Rely entirely on the standard library's `errors.As` function.

### Expected Impact

- **Code Simplification**: Reduce unnecessary code.
- **Improved Maintainability**: Align with Go's idiomatic error handling practices.
- **Reduced Potential for Bugs**: Eliminate the risk of inconsistencies between custom and standard `As` behavior.

### Verification Plan

- **Unit Tests**: Ensure existing tests that rely on `errors.As` (from the standard library) continue to pass after the
  removal of the custom `As` method.

---

## 3. Dependency on `github.com/pkg/errors`

### Problem Description

The `errors` module currently depends on `github.com/pkg/errors`. While `pkg/errors` was a valuable library before Go
1.13, many of its features (like error wrapping and unwrapping) are now part of the standard library.

### Current State Analysis

- **Relevant File**: `errors/go.mod`
- **Observation**: `github.com/pkg/errors v0.9.1` is listed as a direct dependency.

### Proposed Solution(s)

Evaluate the usage of `pkg/errors` within the module. If its functionalities can be fully replaced by Go's standard
library error handling features (e.g., `fmt.Errorf("... %w", err)`, `errors.Is`, `errors.As`), consider migrating away
from `pkg/errors` to reduce external dependencies.

### Expected Impact

- **Reduced Dependency Footprint**: Simplify the project's dependency graph.
- **Modernization**: Align the codebase with current Go best practices for error handling.

### Verification Plan

- **Code Review**: Identify all uses of `pkg/errors` and determine if they can be replaced by standard library
  equivalents.
- **Unit Tests**: Ensure all existing tests continue to pass after migration.

---

## 4. Inconsistent Naming for Multi-Error Type

### Problem Description

The `README.md` refers to a `ThreadSafeMultiError` type, but the actual implementation uses `*multierror.Error` from the
`github.com/hashicorp/go-multierror` package directly.

### Current State Analysis

- **Relevant File**: `errors/errors.go` and `errors/README.md`
- **Relevant Code Snippet**:
    ```go
    // In errors.go
    func ThreadSafe(err error) *multierror.Error { ... }

    // In README.md
    // The `ThreadSafeMultiError` type (returned by `ThreadSafe(err error)`) ...
    ```

### Proposed Solution(s)

Align the documentation with the code. Either:

1. Update the `README.md` to explicitly state that `ThreadSafe` returns `*multierror.Error`.
2. If `ThreadSafeMultiError` is intended to be a custom wrapper or alias, define it explicitly in the code.

### Expected Impact

- **Improved Clarity**: Eliminate confusion between documentation and code.
- **Accuracy**: Ensure the documentation accurately reflects the implementation.

### Verification Plan

- **Documentation Review**: Verify the updated `README.md`.

---

## 5. Default `Code` Value in `FromError`

### Problem Description

When `FromError` converts a non-`*Error` type error, it assigns a `Code` of `0` to the newly created `*Error` instance.
A `Code` of `0` might conflict with valid application-specific error codes or be ambiguous.

### Current State Analysis

- **Relevant File**: `errors/errors.go`
- **Relevant Code Snippet**:
    ```go
    func FromError(err error) *Error {
        // ...
        return New("error.unknown", 0, err.Error())
    }
    ```

### Proposed Solution(s)

Consider using a distinct, non-zero default error code for unknown errors (e.g., a negative number, or a very large
positive number outside the range of common HTTP/RPC status codes) to avoid potential conflicts and improve clarity.

### Expected Impact

- **Reduced Ambiguity**: Clearly distinguish unknown errors from errors with a specific `Code` of `0`.
- **Improved Error Classification**: Facilitate more precise error handling based on error codes.

### Verification Plan

- **Unit Tests**: Add tests to verify the default `Code` value when converting generic errors.
