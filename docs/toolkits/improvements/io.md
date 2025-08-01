# IO Module Improvements & Todo List

This document outlines identified areas for improvement and pending tasks within the `io` module.

---

## 1. Incomplete and Unportable Unit Tests

### Problem Description

The `io` module's unit tests (`io_test.go`) are incomplete, covering only a small fraction of the module's functions.
Additionally, the existing `TestDeleteFile` uses hardcoded Windows paths, making it unportable and difficult to run on
other operating systems.

### Current State Analysis

- **Relevant File**: `io/io_test.go`
- **Observation**: Only `TestDeleteFile` is present, with a `TODO` for adding more test cases. The test uses a literal
  Windows path (e.g., `D:\workspace\project\sugoitech\admin\data\upload\01HVVE64TWS98M1ZFYAYFXWHHM`).

### Proposed Solution(s)

1. **Expand Test Coverage**: Add comprehensive unit tests for all public functions in `io.go`, including `NewFile`,
   `SaveFile`, `ReadFile`, `GetFileExtension`, `GetFileNameWithoutExtension`, `GenerateFileName`, and all `IsXxx`
   functions.
2. **Improve Test Portability**: Use `t.TempDir()` to create temporary directories and files for testing file
   operations. This ensures tests are isolated, clean up after themselves, and run correctly across different operating
   systems.
3. **Test Edge Cases**: Include tests for edge cases such as empty filenames, non-existent files, and various file
   extensions for `IsXxx` functions.

### Expected Impact

- **Increased Reliability**: Ensure the correctness and robustness of file operations and type detection.
- **Improved Portability**: Tests will run consistently across different development environments.
- **Easier Maintenance**: Future changes can be made with confidence, knowing that existing functionality is protected
  by tests.

### Verification Plan

- **Unit Tests**: Develop and execute comprehensive unit tests using Go's testing framework.

---

## 2. Potential Memory Issues with Large Files in `File` Struct

### Problem Description

The `File` struct stores the entire file content as a `[]byte` slice. This design can lead to significant memory
consumption and potential out-of-memory (OOM) errors when dealing with large files.

### Current State Analysis

- **Relevant File**: `io/io.go`
- **Relevant Code Snippet**:
    ```go
    type File struct {
        Name    string `json:"name"`
        Size    int64  `json:"size"`
        Content []byte `json:"content"` // Stores entire file content
    }

    func NewFile(file *multipart.FileHeader) (*File, error) { /* ... reads all content into buf.Bytes() */ }
    func ReadFile(path string) (*File, error) { /* ... reads all content into os.ReadFile() */ }
    ```

### Proposed Solution(s)

Refactor the `File` struct and related functions to support streaming or lazy loading of file content, especially for
large files. Options include:

1. **Modify `File` struct**: Change `Content []byte` to an `io.Reader` or `io.ReadCloser` interface, allowing content to
   be read on demand.
2. **Provide alternative functions**: Offer `NewFileStream` or `ReadFileStream` that return an `io.ReadCloser` instead
   of a `*File` with full content.
3. **Add size limits**: Implement checks to prevent loading excessively large files into memory.

### Expected Impact

- **Reduced Memory Footprint**: Prevent OOM errors and improve performance when handling large files.
- **Improved Scalability**: Enable the module to process files of arbitrary size more efficiently.

### Verification Plan

- **Unit Tests**: Add tests for streaming behavior and memory usage with large dummy files.
- **Performance Benchmarks**: Measure memory consumption and processing time for large files.

---

## 3. Redundant and Deprecated `IsXxx` Functions

### Problem Description

The `io` module contains numerous `IsXxx` functions that are redundant (e.g., `IsImage` and `IsImageFile` performing the
same check) and many of the older versions are marked as `Deprecated`. This increases code duplication and can lead to
confusion.

### Current State Analysis

- **Relevant File**: `io/io.go`
- **Observation**: Many functions like `IsImage`, `IsVideo`, `IsText`, etc., are duplicated with `File` suffix
  versions (`IsImageFile`, `IsVideoFile`, `IsTextFile`) and marked as `Deprecated`.

### Proposed Solution(s)

1. **Remove Deprecated Functions**: Once a reasonable transition period has passed, remove all functions explicitly
   marked as `Deprecated`.
2. **Consolidate Logic**: Ensure that the remaining `IsXxxFile` functions are the single source of truth for file type
   detection.

### Expected Impact

- **Code Simplification**: Reduce the overall code size and complexity.
- **Improved Maintainability**: Easier to update and manage file type detection logic.
- **Clearer API**: Prevent confusion for users about which function to use.

### Verification Plan

- **Code Review**: Ensure all deprecated functions are removed and their usage is updated.
- **Unit Tests**: Verify that the remaining `IsXxxFile` functions work correctly.

---

## 4. Unused `github.com/goexts/generic` Dependency

### Problem Description

The `io` module's `go.mod` file lists `github.com/goexts/generic` as a dependency, but its usage is not apparent in the
provided code snippets.

### Current State Analysis

- **Relevant File**: `io/go.mod`
- **Observation**: `github.com/goexts/generic v0.3.0` is listed as a direct dependency.

### Proposed Solution(s)

Confirm if `github.com/goexts/generic` is indeed used within the `io` module. If it is not, remove it from `go.mod` to
reduce the project's dependency footprint.

### Expected Impact

- **Reduced Dependency Footprint**: Simplify the project's dependency graph.
- **Improved Build Times**: Potentially reduce build times by eliminating unnecessary dependency resolution.

### Verification Plan

- **Code Review**: Thoroughly review all code within the `io` module to confirm the usage of
  `github.com/goexts/generic`.
- **Dependency Check**: After removal (if applicable), run `go mod tidy` to ensure no build issues arise.
