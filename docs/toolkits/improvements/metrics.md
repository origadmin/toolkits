# Metrics Module Improvements & Todo List

This document outlines identified areas for improvement and pending tasks within the `metrics` module.

---

## 1. Lack of Concrete Implementations for Metrics Interfaces

### Problem Description

The `metrics` module currently defines `Metrics` and `Recorder` interfaces but provides no concrete implementations.
This means the package is not directly usable for collecting and recording metrics without external code providing an
implementation.

### Current State Analysis

- **Relevant Files**: `metrics/metrics.go`, `metrics/recorder.go`
- **Observation**: Only interfaces and data structures (`MetricData`, `MetricType`) are defined. No structs implement
  these interfaces within the package.

### Proposed Solution(s)

Provide at least one concrete implementation for the `Metrics` and `Recorder` interfaces. A common approach would be to
integrate with a popular metrics system like Prometheus (using `github.com/prometheus/client_golang`) or OpenTelemetry.

### Expected Impact

- **Usability**: The `metrics` module will become directly functional for collecting and recording metrics.
- **Completeness**: Provide a ready-to-use solution for metrics management.

### Verification Plan

- **Unit Tests**: Develop comprehensive unit tests for the implemented metrics solution to verify data collection,
  aggregation, and exposure.

---

## 2. Missing `README.md` Documentation

### Problem Description

The `metrics` module lacks a `README.md` file, which is crucial for providing an overview of the package's purpose, its
functionalities, and how to use it.

### Current State Analysis

- **Observation**: No `README.md` file was found within the `metrics` directory.

### Proposed Solution(s)

Create a `README.md` file that includes:

- A brief overview of the `metrics` package as a metrics abstraction.
- Explanation of the `Metrics` and `Recorder` interfaces and their methods.
- Explanation of the `MetricData` struct and `MetricType` enumeration.
- Instructions on how to implement these interfaces.
- If a concrete implementation is provided, details on its configuration and usage.

### Expected Impact

- **Improved Usability**: Developers can quickly understand and use the `metrics` package.
- **Better Maintainability**: Clear documentation reduces the learning curve for new contributors.

### Verification Plan

- **Documentation Review**: Ensure the `README.md` is clear, concise, and accurate.

---

## 3. Lack of Comprehensive Unit Tests

### Problem Description

Since there are no concrete implementations of the `Metrics` and `Recorder` interfaces, there are no unit tests for the
`metrics` module. This means that once an implementation is added, its correctness and robustness will not be verified.

### Current State Analysis

- **Observation**: No `_test.go` files were found within the `metrics` directory.

### Proposed Solution(s)

Once concrete implementations of the `Metrics` and `Recorder` interfaces are provided, develop comprehensive unit tests
to verify their functionality. Tests should cover:

- Correctness of metric increments, gauges, and histograms.
- Concurrent access and thread safety.
- Proper labeling and data aggregation.
- Integration with the chosen metrics backend (if applicable).

### Expected Impact

- **Increased Reliability**: Ensure the correctness and robustness of the metrics collection functionality.
- **Easier Maintenance**: Future changes can be made with confidence, knowing that existing functionality is protected
  by tests.

### Verification Plan

- **Unit Tests**: Develop and execute unit tests using Go's testing framework.

---

## 4. Overlapping Responsibilities of `Metrics` and `Recorder` Interfaces

### Problem Description

The `Metrics` and `Recorder` interfaces have somewhat overlapping responsibilities, which could lead to confusion
regarding their intended use and potentially redundant method definitions.

### Current State Analysis

- **Relevant Files**: `metrics/metrics.go`, `metrics/recorder.go`
- **Observation**: `Metrics` has `Observe` (takes `MetricData`) and `Log` (takes individual parameters). `Recorder` has
  specific methods for each metric type.

### Proposed Solution(s)

Refine the responsibilities of `Metrics` and `Recorder` interfaces. For example:

- `Metrics` could be a higher-level interface for general metric operations (e.g., `Enabled()`, `Disable()`,
  `Observe(MetricData)`).
- `Recorder` could be a lower-level interface that concrete implementations use to interact with the metrics backend, or
  it could be removed if `Metrics.Observe` is sufficient.
- The `Log` method in `Metrics` could be removed in favor of `Observe(MetricData)` to promote a single, structured data
  input.

### Expected Impact

- **Improved API Clarity**: Make the purpose of each interface more distinct.
- **Reduced Redundancy**: Eliminate overlapping method definitions.
- **Better Maintainability**: Easier to understand and extend the metrics system.

### Verification Plan

- **API Design Review**: Review the refined interface definitions.
- **Refactoring**: Adjust existing code (if any) to conform to the new interface structure.

---

## 5. Hardcoded `MetricLabelNames` and Lack of Extensibility

### Problem Description

The `metricLabelNames` map is hardcoded, making it inflexible for adding new metric types or modifying existing label
sets without directly changing the source code.

### Current State Analysis

- **Relevant File**: `metrics/metrics.go`
- **Relevant Code Snippet**:
    ```go
    var metricLabelNames = map[MetricType][]string{ /* ... */ }
    ```

### Proposed Solution(s)

Introduce a registration mechanism for metric types and their associated labels. This could involve:

- A `RegisterMetricType(MetricType, []string)` function.
- Allowing implementations to define their own label sets.

### Expected Impact

- **Increased Flexibility**: Easily add or modify metric types and labels without code changes.
- **Improved Extensibility**: Support dynamic metric definitions.

### Verification Plan

- **API Design Review**: Review the new registration mechanism.
- **Unit Tests**: Test the registration and retrieval of metric labels.

---

## 6. Inconsistent Latency Type (`float64` vs. `int64`)

### Problem Description

There is an inconsistency in the type used for latency metrics: `MetricData.Latency` is `float64`, `Metrics.Log` takes
`float64`, but `Recorder.RequestDurationSeconds` and `Recorder.SummaryLatencyLog` take `int64`.

### Current State Analysis

- **Relevant Files**: `metrics/data.go`, `metrics/metrics.go`, `metrics/recorder.go`
- **Observation**: Mixed usage of `float64` and `int64` for latency.

### Proposed Solution(s)

Standardize the type for latency metrics. It is generally recommended to use `time.Duration` for internal representation
and convert to `float64` (seconds) or `int64` (milliseconds/nanoseconds) only when exposing or logging the metric.

### Expected Impact

- **Improved Consistency**: Reduce potential for type conversion errors and confusion.
- **Accuracy**: Ensure precise representation of time durations.

### Verification Plan

- **Code Review**: Ensure all latency-related fields and parameters use a consistent type.
- **Unit Tests**: Verify correct handling of latency values.
