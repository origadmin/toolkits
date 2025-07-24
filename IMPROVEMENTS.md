# Project-Wide Improvements & Todo List

This document serves as a central tracker for planned improvements, refactoring
efforts, and bug fixes across all modules of the project.

---

## Standard Template for Optimization Analysis

When proposing or documenting an optimization or significant refactoring, please
use the following structure to ensure clarity, completeness, and traceability of
decisions.

### 1. Problem Description

Clearly state the issue or area identified for improvement. What is the current
limitation, bug, or inefficiency?

### 2. Current State Analysis

Describe how the system or component currently works. Detail the existing
implementation, relevant code snippets, and specific observations that highlight
the problem.

### 3. Proposed Solution(s)

Outline the primary approach(es) to address the problem. Provide a high-level
overview of the intended changes.

### 4. Alternative Solutions Considered

List and briefly describe other approaches that were evaluated but not chosen.
This section is crucial for documenting the thought process and avoiding
revisiting discarded ideas.

### 5. Comparison & Evaluation of Alternatives

Provide a detailed comparison of the proposed solution against the alternatives.
This is where quantitative (e.g., performance metrics, memory usage) and
qualitative (e.g., maintainability, complexity, security implications) arguments
should be made. A comparison table is highly recommended for clarity.

### 6. Chosen Solution Details

Provide a detailed plan for the selected solution, including specific code
changes, new components, or refactoring steps.

### 7. Expected Impact

Describe the anticipated benefits of the implemented solution (e.g., performance
improvement, security enhancement, reduced memory footprint, improved
maintainability, clearer API).

### 8. Verification Plan

Outline how the changes will be tested and verified to ensure correctness,
performance gains, and absence of regressions.

---

## Crypto Module

This section outlines critical security vulnerabilities and areas for
improvement found in the `crypto` toolkit. It covers the `aes`, `rand`, and
`hash` sub-packages.

### 1. `crypto/aes` Package

- [x] **Status: Completed**
- [x] **Issue: Predictable IV in AES CBC Mode (Critical Security
  Vulnerability)**

**Analysis:**
The original `EncryptCBC` implementation used a static and predictable
Initialization Vector (IV). This is a critical vulnerability in CBC mode.

**Remediation Plan:**

1. **Preserve Compatibility:** The existing `EncryptCBC` and `DecryptCBC`
   functions were preserved for compatibility with external services (e.g., WeChat
   Pay).
2. **Add Security Warnings:** The legacy functions were marked as `Deprecated`
   with clear warnings against their use in new development.
3. **Introduce Secure Alternative:** New, secure functions (`EncryptGCM`,
   `DecryptGCM`) using the modern AES-GCM mode were added as the recommended
   standard for all new encryption needs.

### 2. `crypto/rand` Package

- [ ] **Status: In Progress (Current Task)**
- [ ] **Issue: Multiple design flaws including security risks, poor performance,
  and a mutable API.**

**Analysis:**
The original `rand` package suffered from several critical issues:

1. **Security:** It used the non-cryptographically secure `math/rand` package.
2. **API Design:** It exposed mutable global variables (e.g., `rand.All`),
   allowing external packages to change their state, leading to unpredictable
   behavior. It also used a `sync.Pool` for small objects, which added unnecessary
   complexity.
3. **Performance:** The initial plan to fix the security issue involved
   repeated calls to `crypto/rand`, which is inefficient. The character set
   construction was also happening at runtime.

**Remediation Plan (Final Design):**

1. **Adopt a Secure & Immutable API (Constructor Pattern):**
    * **Decision:** Eliminate all mutable global state. The final design uses
      the **Constructor Pattern** to provide a robust and predictable API.
    * **Rationale:** Initial ideas of using public variables (`rand.All`) or
      getters for private instances (`rand.All()`) were rejected because they still
      exposed a modifiable global object, which is not thread-safe and leads to
      unpredictable behavior. The chosen pattern ensures complete safety and
      encapsulation by returning an interface.
    * **Implementation:**
        * A new interface `RandGenerator` will be defined, exposing
          `RandBytes`, `RandString`, and `Read` methods.
        * The concrete implementation struct `randGenerator` will be
          unexported.
        * `NewRand(Kind)` and `CustomRand(charset string)` functions will act
          as constructors, returning the `RandGenerator` interface type.
        * The public API is based on immutable `const` values of type `Kind`
          (e.g., `rand.KindAlphanumeric`). These act as safe configurations for the
          constructors.
        * Naming: `NewRand` and `CustomRand` are standard Go idioms.
          `RandGenerator` clearly defines the role.

2. **Implement High-Performance Generation:**
    * **Decision:** Pre-compute all character set variations via an `init()`
      function and cache them in a private `map`.
    * **Rationale & Alternatives Considered:** The goal was to minimize
      runtime memory allocations.

   | Feature             | `switch` Statement                       | `array/slice` Lookup        | `init()` + `map` (Chosen) |
   |:--------------------|:-----------------------------------------|:----------------------------|:--------------------------|
   | **Lookup Speed**    | Extremely Fast                           | Extremely Fast              | Very Fast                 |
   | **Memory Alloc.**   | Per-call allocation                      | Per-call allocation         | Only at startup           |
   | **Maintainability** | Complex logic, error-prone for new kinds | Brittle, hard to extend     | Clean, easy to extend     |
   | **Robustness**      | Risk of runtime errors if not exhaustive | Risk of index out of bounds | Highly robust             |


   * **`switch` statement:** This was rejected. While a `switch` has a
     very fast lookup (due to compiler optimizations into a jump table for small,
     contiguous integer cases), it would require **runtime string concatenation** for
     combined sets (e.g., `KindAlphanumeric`). This repeated memory allocation makes
     it significantly slower overall than a one-time computation.
   * **`array/slice` lookup:** This was rejected. While offering the
     fastest lookup (direct memory access), it is too **brittle and hard to
     maintain**. The array size would need to be hardcoded, and adding a new `Kind`
     could easily cause an index-out-of-bounds panic, making the code fragile.
   * **Winning Approach (`init` + `map`):** This approach has zero
     runtime allocation when creating a new generator (after the initial `init`
     phase), making it the highest-performance and most robust solution. It leverages
     Go's `init` function for one-time setup and `map` for efficient, flexible
     lookups.
     * **Implementation:** The core `RandBytes` method was rewritten to read
       all required random bytes from `crypto/rand` in a single call, then use
       efficient modulo arithmetic to map bytes to the chosen character set.

3. **Simplify and Clarify:**
    * **Decision:** Delete the `generate.go` file and provide clean, top-level
      convenience functions.
    * **Implementation:** New functions `RandomString(n int)` and
      `RandomBytes(n int)` are provided for the most common use cases, using a safe
      default (Alphanumeric).

### 3. `crypto/hash` Package

- [ ] **Status: To Do**
- [ ] **Issue 1: Lack of Thread Safety (Concurrency Bug)**
- [ ] **Issue 2: Panicking in `init()` (Robustness Issue)**

**Analysis & Plan:** See original analysis. The plan remains to add a
`sync.Mutex` to the factory and to replace the `panic` in the `init` function
with graceful error handling.
