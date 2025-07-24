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

- [x] **Status: Completed**
- [x] **Issue: Multiple design flaws including security risks, poor performance,
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

- [ ] **Status: To Do (Current Task)**
- [ ] **Issue 1: Lack of Thread Safety (Concurrency Bug)**

**Analysis:**
The `algorithmFactory` in `factory.go` uses a `cryptos` map to cache hash algorithm instances. Access to this map (read and write) is not protected by any concurrency control mechanism. In a multi-goroutine environment, concurrent access to this map can lead to data races, resulting in program crashes or unpredictable behavior.

**Current State Analysis:**
*   **Relevant File:** `factory.go`
*   **Relevant Code Snippet:**
    ```go
    type algorithmFactory struct {
    	cryptos map[types.Type]interfaces.Cryptographic
    }

    func (f *algorithmFactory) create(algType types.Type, opts ...types.Option) (interfaces.Cryptographic, error) {
    	if alg, exists := f.cryptos[algType]; exists { // Read operation
    		return alg, nil
    	}
        // ...
    	f.cryptos[algType] = alg // Write operation
    	return alg, nil
    }
    ```
*   **Observation:** The `cryptos` map is a non-concurrent-safe type. The `exists` check and the subsequent assignment (`f.cryptos[algType] = alg`) are not atomic, making them vulnerable to data races under concurrent access.

**Proposed Solution(s):**
Use `sync.RWMutex` to protect the `cryptos` map. Read operations will acquire a read lock, and write operations will acquire a write lock.

**Alternative Solutions Considered:**
*   **`sync.Mutex`:** Simpler to implement, but less performant in read-heavy scenarios as it blocks all other reads during a read operation. `RWMutex` is preferred for read-mostly caches.
*   **`sync.Once` + Lazy Initialization:** Not suitable here as `algorithmFactory` manages multiple distinct algorithm instances, not just a single one. Each instance might be requested at different times.

**Comparison & Evaluation of Alternatives:**

| Feature           | `sync.Mutex` | `sync.RWMutex` (Chosen) |
| :---------------- | :----------- | :---------------------- |
| **Concurrency**   | Read-blocking| Concurrent reads allowed|
| **Performance**   | Lower        | Higher (read-heavy)     |
| **Complexity**    | Simple       | Slightly more complex   |

**Chosen Solution Details:**
*   **File:** `factory.go`
*   **Modification:**
    1.  Add `mu sync.RWMutex` field to `algorithmFactory` struct.
    2.  In `create` method, acquire `f.mu.RLock()` before reading `f.cryptos` and `f.mu.RUnlock()` after.
    3.  Acquire `f.mu.Lock()` before writing to `f.cryptos` and `f.mu.Unlock()` after.

**Expected Impact:**
*   **Thread Safety:** Resolves data race issues in `algorithmFactory`, ensuring stability and correctness under concurrent access.
*   **Performance:** Improves concurrency performance in read-heavy scenarios compared to `sync.Mutex`.

**Verification Plan:**
*   **Unit Tests:** Add concurrent tests for `algorithmFactory.create` to verify no data races occur under heavy load.

- [ ] **Issue 2: Panicking in `init()` (Robustness Issue)**

**Analysis:**
The `init()` function in `hash.go` panics if `NewCrypto` fails during the initialization of the `defaultCrypto` global instance. While `panic` terminates the program immediately, a more robust library should log errors and allow for graceful failure or recovery if possible. Furthermore, public functions relying on `defaultCrypto` (e.g., `Generate`, `Verify`) might attempt to dereference a `nil` pointer if initialization fails, leading to runtime errors.

**Current State Analysis:**
*   **Relevant File:** `hash.go`
*   **Relevant Code Snippet:**
    ```go
    var (
    	defaultCrypto Crypto
    )

    func init() {
    	// ... initialization logic ...
    	if defaultCrypto == nil {
    		cryptographic, err := NewCrypto(core.DefaultType)
    		if err != nil {
    			panic(err) // Panics here if NewCrypto fails
    		}
    		defaultCrypto = cryptographic
    	}
    }

    func Verify(hashed, password string) error {
    	return defaultCrypto.Verify(hashed, password) // Potential nil pointer dereference
    }

    func Generate(password string) (string, error) {
    	return defaultCrypto.Hash(password) // Potential nil pointer dereference
    }
    ```
*   **Observation:** The `init()` function's `panic` behavior is brittle. Public functions do not check for a `nil` `defaultCrypto` before use.

**Proposed Solution(s):**
1.  Modify `init()` to log errors instead of panicking when `NewCrypto` fails.
2.  Modify public functions (`Generate`, `Verify`, `GenerateWithSalt`) to explicitly check if `defaultCrypto` is `nil` and return an error if it is, guiding the user to handle initialization failures.

**Alternative Solutions Considered:**
*   **Force User Initialization:** Remove `defaultCrypto` initialization from `init()` and require users to explicitly call an initialization function. This increases user burden and might lead to forgotten initialization. The current `init()` provides convenience that should be preserved while improving robustness.

**Comparison & Evaluation of Alternatives:**

| Feature           | `panic` (Current) | Log Error & Return Error (Chosen) |
| :---------------- | :---------------- | :-------------------------------- |
| **Program Behavior**| Immediate crash   | Logs error, program continues     |
| **Robustness**    | Poor              | Good                              |
| **User Experience**| Poor              | Good (errors can be caught)       |
| **API Clarity**   | Poor (implicit crash) | Good (explicit error return)      |

**Chosen Solution Details:**
*   **File:** `hash.go`
*   **Modification:**
    1.  Import `log` package.
    2.  Modify `init()` function:
        ```go
        func init() {
            alg := os.Getenv(ENV)
            if alg == "" {
                alg = core.DefaultType
            }
            t := types.ParseType(alg)
            var err error
            if t != types.TypeUnknown {
                defaultCrypto, err = NewCrypto(t)
            }
            if defaultCrypto == nil || err != nil { // Check both defaultCrypto and potential error from NewCrypto
                if err != nil {
                    log.Printf("Failed to initialize default crypto with type %s: %v", t, err)
                }
                // Try to initialize with default type if previous attempt failed or was unknown
                defaultCrypto, err = NewCrypto(core.DefaultType)
                if err != nil {
                    log.Printf("Failed to initialize default crypto with default type %s: %v", core.DefaultType, err)
                    // At this point, defaultCrypto will remain nil.
                    // Public functions will handle this by returning an error.
                }
            }
        }
        ```
    3.  Modify `Verify`, `Generate`, `GenerateWithSalt` functions:
        ```go
        func Verify(hashed, password string) error {
            if defaultCrypto == nil {
                return errors.New("hash: default crypto not initialized")
            }
            return defaultCrypto.Verify(hashed, password)
        }

        func Generate(password string) (string, error) {
            if defaultCrypto == nil {
                return "", errors.New("hash: default crypto not initialized")
            }
            return defaultCrypto.Hash(password)
        }

        func GenerateWithSalt(password, salt string) (string, error) {
            if defaultCrypto == nil {
                return "", errors.New("hash: default crypto not initialized")
            }
            return defaultCrypto.HashWithSalt(password, salt)
        }
        ```

**Expected Impact:**
*   **Robustness:** Improves the library's fault tolerance, preventing program crashes due to initialization failures.
*   **Maintainability:** Clear error return mechanisms allow callers to better handle initialization failures.
*   **User Experience:** Provides explicit errors instead of panics, making debugging and error handling easier for users.

**Verification Plan:**
*   **Unit Tests:**
    *   Write test cases to simulate `NewCrypto` initialization failures and verify that `init()` correctly logs errors without panicking.
    *   Write test cases to verify that calling `Generate`, `Verify`, and `GenerateWithSalt` when `defaultCrypto` is uninitialized returns the expected error.