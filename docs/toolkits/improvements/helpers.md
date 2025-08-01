# Helpers Module Improvements & Todo List

This document outlines identified areas for improvement and pending tasks within the `helpers` module.

---

## 1. Lack of Comprehensive Unit Tests

### Problem Description

The `helpers` module currently lacks comprehensive unit tests for its functions. This makes it difficult to ensure the
correctness and robustness of the utility functions, especially in various edge cases and across different environments.

### Current State Analysis

- **Relevant Files**: `helpers/discovery.go`, `helpers/intbytes.go`, `helpers/time_duration.go`
- **Observation**: No `_test.go` files were found within the `helpers` directory.

### Proposed Solution(s)

Add comprehensive unit tests for all public functions in `discovery.go`, `intbytes.go`, and `time_duration.go`. Tests
should cover:

- Basic functionality.
- Edge cases (e.g., empty inputs, zero values, large numbers, network configurations).
- Correctness of conversions.
- Error handling scenarios.

### Expected Impact

- **Increased Reliability**: Ensure the correctness and robustness of helper functions.
- **Easier Maintenance**: Future changes can be made with confidence, knowing that existing functionality is protected
  by tests.

### Verification Plan

- **Unit Tests**: Develop and execute unit tests using Go's testing framework.

---

## 2. Missing `README.md` Documentation

### Problem Description

The `helpers` module lacks a `README.md` file, which is crucial for providing an overview of the package's purpose, its
functionalities, and how to use it.

### Current State Analysis

- **Observation**: No `README.md` file was found within the `helpers` directory.

### Proposed Solution(s)

Create a `README.md` file that includes:

- A brief overview of the `helpers` package and its sub-components.
- A list of its public functions and their purposes.
- Usage examples for common scenarios.
- Any special considerations (e.g., units for time conversions, endianness for byte conversions).

### Expected Impact

- **Improved Usability**: Developers can quickly understand and use the `helpers` package.
- **Better Maintainability**: Clear documentation reduces the learning curve for new contributors.

### Verification Plan

- **Documentation Review**: Ensure the `README.md` is clear, concise, and accurate.

---

## 3. Unit Inconsistency in `time_duration.go`

### Problem Description

The `Int64ToDuration` and `DurationToInt64` functions in `time_duration.go` use inconsistent units, leading to potential
confusion and incorrect calculations.

### Current State Analysis

- **Relevant File**: `helpers/time_duration.go`
- **Relevant Code Snippets**:
    ```go
    func Int64ToDuration(seconds int64) time.Duration {
        return time.Duration(seconds) * 1e6 // Interprets seconds as microseconds
    }

    func DurationToInt64(duration time.Duration) int64 {
        return duration.Milliseconds() // Returns milliseconds
    }
    ```
- **Observation**: `Int64ToDuration` converts an `int64` (presumably seconds) to `time.Duration` by multiplying by
  `1e6` (microseconds), while `DurationToInt64` converts `time.Duration` to `int64` in milliseconds.

### Proposed Solution(s)

Standardize the units for these conversion functions. It is recommended to use `time.Second` or `time.Millisecond`
explicitly for clarity. For example:

```go
// Int64SecondsToDuration converts seconds (int64) to time.Duration.
func Int64SecondsToDuration(seconds int64) time.Duration {
    return time.Duration(seconds) * time.Second
}

// DurationToInt64Seconds converts time.Duration to seconds (int64).
func DurationToInt64Seconds(duration time.Duration) int64 {
    return int64(duration.Seconds())
}

// Int64MillisecondsToDuration converts milliseconds (int64) to time.Duration.
func Int64MillisecondsToDuration(milliseconds int64) time.Duration {
    return time.Duration(milliseconds) * time.Millisecond
}

// DurationToInt64Milliseconds converts time.Duration to milliseconds (int64).
func DurationToInt64Milliseconds(duration time.Duration) int64 {
    return duration.Milliseconds()
}
```

### Expected Impact

- **Increased Accuracy**: Prevent incorrect time calculations.
- **Improved Clarity**: Make the unit of conversion explicit.

### Verification Plan

- **Unit Tests**: Add tests to verify the correctness of the conversions with the chosen units.

---

## 4. Lack of Error Handling in `intbytes.go`

### Problem Description

The `BytesToUint64` function in `intbytes.go` does not handle cases where the input byte slice length is not a multiple
of 8, potentially leading to silent data truncation or unexpected behavior.

### Current State Analysis

- **Relevant File**: `helpers/intbytes.go`
- **Relevant Code Snippet**:
    ```go
    func BytesToUint64(buf []byte) []uint64 {
        size := len(buf)
        if size == 0 {
            return nil
        }
        ints := make([]uint64, size/8)
        for i := 0; i < size/8; i++ {
            ints[i] = binary.BigEndian.Uint64(buf[i*8 : i*8+8])
        }
        return ints
    }
    ```
- **Observation**: If `len(buf)` is not divisible by 8, the last few bytes will be ignored without any error or warning.

### Proposed Solution(s)

Modify `BytesToUint64` to return an error if the input byte slice length is not a multiple of 8. This will make the
function more robust and prevent silent data loss.

### Expected Impact

- **Improved Robustness**: Prevent silent data truncation.
- **Clearer Error Reporting**: Provide explicit errors for invalid input.

### Verification Plan

- **Unit Tests**: Add tests to verify that `BytesToUint64` returns an error for invalid input lengths.

---

## 5. `discovery.go` - `ServiceDiscoveryEndpoint` Logic Complexity

### Problem Description

The `ServiceDiscoveryEndpoint` function in `discovery.go` has complex logic for parsing and reconstructing endpoints,
especially when dealing with existing `endpoint` strings. Its `Deprecated` status suggests it might be replaced, but its
current complexity could be a source of bugs.

### Current State Analysis

- **Relevant File**: `helpers/discovery.go`
- **Observation**: The function involves string splitting, joining, and conditional logic to manipulate endpoint URLs.

### Proposed Solution(s)

Given its `Deprecated` status, prioritize its eventual removal once `ServiceEndpoint` (or its new location in
`runtime.service`) fully replaces its functionality. If it must remain for compatibility, consider simplifying its logic
or adding more robust parsing/validation.

### Expected Impact

- **Reduced Complexity**: Simplify the codebase by removing or refactoring complex, deprecated logic.
- **Improved Maintainability**: Make the code easier to understand and less prone to errors.

### Verification Plan

- **Code Review**: Ensure the function is eventually removed or its logic is simplified.
- **Migration Plan**: Document the migration path for users from `ServiceDiscoveryEndpoint` to its replacement.
