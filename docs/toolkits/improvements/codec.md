# Codec Module Improvements & Todo List

This document outlines identified areas for improvement and pending tasks within the `codec` module.

---

## 1. Unimplemented Encoding Functions

### Problem Description

The `toml` and `yaml` sub-packages currently lack implementation for their respective encoding functions (`EncodeTOML`
and `EncodeYAML`). These functions return empty strings and `nil` errors, which is misleading and incomplete.

### Current State Analysis

- **Relevant Files**:
    - `codec/toml/toml.go`: `EncodeTOML` function.
    - `codec/yaml/yaml.go`: `EncodeYAML` function.
- **Observation**: Both functions contain `// TODO: Implement ... encoding` comments and provide no actual encoding
  logic.

### Proposed Solution(s)

Implement the `EncodeTOML` and `EncodeYAML` functions to correctly serialize Go data structures into TOML and YAML
strings, respectively, leveraging their underlying libraries (`github.com/BurntSushi/toml` and `gopkg.in/yaml.v3`).

### Expected Impact

- **Completeness**: The `codec` module will provide full encoding/decoding capabilities for all advertised formats.
- **Usability**: Users will be able to serialize data to TOML and YAML using a consistent API.

### Verification Plan

- **Unit Tests**: Add comprehensive unit tests for `EncodeTOML` and `EncodeYAML` to verify correct serialization of
  various data types, including structs, maps, and slices, and handle edge cases.

---

## 2. Redundant JSON Parsing in `DecodeJSON`

### Problem Description

The `DecodeJSON` function in `codec/json/json.go` performs redundant parsing of the input JSON string, potentially
leading to unnecessary performance overhead.

### Current State Analysis

- **Relevant File**: `codec/json/json.go`
- **Relevant Code Snippet**:
    ```go
    func DecodeJSON(json string, v interface{}) (Value, error) {
        var val Value
        err := sonic.UnmarshalString(json, v) // First parse
        if err != nil {
            return val, err
        }
        val = DefaultJson.Get(json) // Second parse (implicit in Get)
        return val, nil
    }
    ```
- **Observation**: The `json` string is first unmarshaled into `v` using `sonic`, and then `DefaultJson.Get(json)` (
  which uses `jsoniter.Any`) implicitly parses the same string again to create the `Value` object.

### Proposed Solution(s)

Refactor `DecodeJSON` to avoid redundant parsing. This could involve:

1. If `jsoniter.Any` provides a way to construct a `Value` from an already unmarshaled Go interface, use that.
2. If not, and `Value` is primarily for dynamic access, consider if `sonic`'s own dynamic access features (if any) could
   be leveraged, or if the `Value` type is strictly necessary for this function's return.

### Expected Impact

- **Performance**: Reduce CPU cycles and memory allocations by eliminating redundant parsing.
- **Efficiency**: Improve the overall efficiency of JSON decoding operations.

### Verification Plan

- **Unit Tests**: Ensure existing tests for `DecodeJSON` still pass.
- **Performance Benchmarks**: Add benchmarks to measure the performance improvement after refactoring.

---

## 3. Unused `github.com/goexts/generic` Dependency

### Problem Description

The `codec` module's `go.mod` file lists `github.com/goexts/generic` as a dependency, but its usage is not immediately
apparent in the provided code snippets.

### Current State Analysis

- **Relevant File**: `codec/go.mod`
- **Observation**: `github.com/goexts/generic` is listed as a direct dependency.
- **Code Review**: In `codec/json/json.go`, `codec/toml/toml.go`, `codec/yaml/yaml.go`, and `codec/ini/ini.go`, the
  `Value` type is aliased from `generic.Any`.

### Proposed Solution(s)

Confirm if `github.com/goexts/generic` is indeed used. If it is only used for aliasing `generic.Any`, and `jsoniter.Any`
or similar types can serve the same purpose without an external dependency, consider removing
`github.com/goexts/generic` to reduce the dependency footprint. If `generic.Any` provides unique, necessary
functionality, document its specific use case.

### Expected Impact

- **Reduced Dependency Footprint**: If unused, removing the dependency will simplify the project's dependency graph.
- **Improved Clarity**: Clarify the purpose of the `Value` type and its underlying implementation.

### Verification Plan

- **Code Review**: Thoroughly review all code within the `codec` module to confirm the usage of
  `github.com/goexts/generic`.
- **Dependency Check**: After removal (if applicable), run `go mod tidy` to ensure no build issues arise.
