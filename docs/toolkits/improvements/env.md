# Env Module Improvements & Todo List

This document outlines identified areas for improvement and pending tasks within the `env` module.

---

## 1. Lack of Comprehensive Unit Tests

### Problem Description

The `env` module currently lacks comprehensive unit tests for its functions and interfaces. This makes it difficult to
ensure the correctness and robustness of the environment variable operations, especially in various edge cases.

### Current State Analysis

- **Relevant File**: `env/env.go`
- **Observation**: No `_test.go` files were found within the `env` directory.

### Proposed Solution(s)

Add comprehensive unit tests for all public functions (`SetEnv`, `GetEnv`, `LookupEnv`, `Var`, `WithPrefix`) and the
`Env` interface methods. Tests should cover:

- Basic functionality.
- Edge cases (e.g., empty strings, special characters).
- Concurrent access (if applicable).
- Correctness of `key` transformation in `SetEnv`.
- Behavior of `WithPrefix`.

### Expected Impact

- **Increased Reliability**: Ensure the correctness and robustness of environment variable operations.
- **Easier Maintenance**: Future changes can be made with confidence, knowing that existing functionality is protected
  by tests.

### Verification Plan

- **Unit Tests**: Develop and execute unit tests using Go's testing framework.

---

## 2. Missing `README.md` Documentation

### Problem Description

The `env` module lacks a `README.md` file, which is crucial for providing an overview of the package's purpose, its
functionalities, and how to use it.

### Current State Analysis

- **Observation**: No `README.md` file was found within the `env` directory.

### Proposed Solution(s)

Create a `README.md` file that includes:

- A brief overview of the `env` package.
- A list of its public functions and their purposes.
- Usage examples for common scenarios.
- Any special considerations (e.g., `SetEnv` converting keys to uppercase).

### Expected Impact

- **Improved Usability**: Developers can quickly understand and use the `env` package.
- **Better Maintainability**: Clear documentation reduces the learning curve for new contributors.

### Verification Plan

- **Documentation Review**: Ensure the `README.md` is clear, concise, and accurate.

---

## 3. `SetEnv` Key Transformation Behavior

### Problem Description

The `SetEnv` function automatically converts the provided key to uppercase before setting the environment variable.
While this might be intended for consistency, it could lead to unexpected behavior if the caller expects the original
casing to be preserved.

### Current State Analysis

- **Relevant File**: `env/env.go`
- **Relevant Code Snippet**:
    ```go
    func SetEnv(key, value string) error {
        return os.Setenv(strings.ToUpper(key), value)
    }
    ```

### Proposed Solution(s)

- **Clarify Documentation**: Explicitly document this behavior in the `README.md` and Go Doc comments for `SetEnv`.
- **Optional: Provide Alternative**: If there's a use case for preserving original casing, consider adding an
  alternative function (e.g., `SetEnvRaw`) that does not perform the uppercase conversion.

### Expected Impact

- **Reduced Ambiguity**: Prevent unexpected behavior for users.
- **Improved API Clarity**: Clearly communicate the function's side effects.

### Verification Plan

- **Documentation Review**: Verify the updated documentation.
- **Unit Tests**: Add tests to confirm the uppercase conversion behavior.

---

## 4. `Var` and `prefixedVar` Naming

### Problem Description

The names `Var` and `prefixedVar` might not be immediately intuitive for functions that construct environment variable
names. `Var` could be confused with a variable, and `prefixedVar` is an internal function.

### Current State Analysis

- **Relevant File**: `env/env.go`
- **Relevant Code Snippets**:
    ```go
    func Var(vv ...string) string { ... }
    func prefixedVar(prefix string, vv ...string) string { ... }
    ```

### Proposed Solution(s)

Consider renaming these functions to be more descriptive of their purpose, such as `BuildEnvKey` or `FormatEnvKey` for
`Var`, and `buildPrefixedEnvKey` for `prefixedVar`.

### Expected Impact

- **Improved Readability**: Make the code easier to understand for new and existing contributors.
- **Enhanced Clarity**: Clearly convey the function's role in constructing environment variable names.

### Verification Plan

- **Code Review**: Ensure new names are consistently applied and improve clarity.
- **Refactoring**: Perform a safe refactoring of the function names.

---

## 5. `WithPrefix` Panics on Empty Prefix

### Problem Description

The `WithPrefix` function panics if an empty string is provided as a prefix. In library functions, it is generally
preferable to return an error rather than panicking, allowing callers to handle the error gracefully.

### Current State Analysis

- **Relevant File**: `env/env.go`
- **Relevant Code Snippet**:
    ```go
    func WithPrefix(prefix string) Env {
        if prefix == "" {
            panic("prefix cannot be empty")
        }
        return &env{
            prefix: prefix,
        }
    }
    ```

### Proposed Solution(s)

Modify `WithPrefix` to return an error (e.g., `(Env, error)`) when the `prefix` is empty, allowing the caller to handle
the invalid input gracefully.

### Expected Impact

- **Improved Robustness**: Prevent application crashes due to invalid input.
- **Better Error Handling**: Enable callers to handle invalid prefix scenarios programmatically.

### Verification Plan

- **Unit Tests**: Add tests to verify that `WithPrefix` returns an error for empty prefixes and does not panic.

---

## 6. Evaluation of `Env` Interface Design

### Problem Description

The `Env` interface and its `env` implementation introduce a level of abstraction. While useful for extensibility (e.g.,
mocking for tests, different backends), for a simple utility package, it might add unnecessary complexity if there are
no immediate plans for such extensions.

### Current State Analysis

- **Relevant File**: `env/env.go`
- **Observation**: The package defines an `Env` interface and a concrete `env` struct.

### Proposed Solution(s)

Evaluate the long-term vision for the `env` package. If there are no concrete plans to support multiple `Env`
implementations or extensive mocking, consider simplifying the package by removing the interface and directly exposing
the utility functions. If the interface is deemed necessary for future extensibility or testing, ensure its purpose is
clearly documented.

### Expected Impact

- **Simplified Codebase (if applicable)**: Reduce complexity if the abstraction is not currently needed.
- **Clearer Intent**: Ensure the design aligns with the package's current and future goals.

### Verification Plan

- **Design Review**: Conduct a review to determine the necessity of the interface.
- **Refactoring (if applicable)**: Perform a safe refactoring if the interface is removed.
