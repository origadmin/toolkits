# Project-Wide Improvements & Todo List

This document serves as a central tracker for planned improvements, refactoring efforts, and bug fixes across all modules of the project.

---

## Crypto Module

This section outlines critical security vulnerabilities and areas for improvement found in the `crypto` toolkit. It covers the `aes`, `rand`, and `hash` sub-packages.

### 1. `crypto/aes` Package

- [x] **Status: Completed**
- [x] **Issue: Predictable IV in AES CBC Mode (Critical Security Vulnerability)**

**Analysis:**
The original `EncryptCBC` implementation used a static and predictable Initialization Vector (IV), derived directly from the encryption key (`key[:blockSize]`). In CBC mode, the IV must be random and unpredictable for each encryption. Using a static IV means that identical plaintext blocks will produce identical ciphertext blocks, leaking information about the data's structure and making the encryption vulnerable to chosen-plaintext attacks.

**Remediation Plan:**
1.  **Preserve Compatibility:** The existing `EncryptCBC` and `DecryptCBC` functions will be preserved to maintain compatibility with external services like WeChat Pay, which rely on this specific, non-standard implementation.
2.  **Add Security Warnings:** The documentation for the legacy CBC functions has been updated with `Deprecated` tags and clear warnings, advising against their use in any new development due to the security risks.
3.  **Introduce Secure Alternative:** New, secure encryption functions (`EncryptGCM`, `DecryptGCM`) have been added. These functions use the AES-GCM (Galois/Counter Mode), which is a modern, authenticated encryption (AEAD) mode. It provides strong confidentiality and, crucially, integrity protection, and it correctly manages random nonces.
4.  **Recommendation:** All new cryptographic requirements within the project should use the new `EncryptGCM` and `DecryptGCM` functions.

### 2. `crypto/rand` Package

- [ ] **Status: In Progress (Current Task)**
- [ ] **Issue: Use of Non-Cryptographically Secure Random Number Generator (Critical Security Vulnerability)**

**Analysis:**
The `rand` package currently uses `math/rand/v2` for generating random data, including for the `GenerateSalt` function. `math/rand` is a pseudo-random number generator (PRNG) and is **not** suitable for security-sensitive contexts. Its output is predictable, meaning that salts or other values generated with it are not truly random, undermining the security of password hashing and other cryptographic operations.

**Remediation Plan:**
1.  **Replace RNG:** Replace all calls to `math/rand/v2` with `crypto/rand`. The `crypto/rand` package is specifically designed for cryptographic operations and sources entropy from the underlying operating system.
2.  **Refactor `Rand` struct:** Update the `RandBytes`, `RandString`, and `Read` methods to use `crypto/rand` for generating random values from the specified character sets.

### 3. `crypto/hash` Package

- [ ] **Status: To Do**
- [ ] **Issue 1: Lack of Thread Safety (Concurrency Bug)**

**Analysis:**
The `algorithmFactory` uses a map `cryptos` to cache algorithm instances. This map is read from and written to in the `create` method without any synchronization mechanism. If multiple goroutines call a hash function concurrently, it can lead to a race condition when they all try to initialize and cache an algorithm instance at the same time, potentially causing a panic or data corruption.

**Remediation Plan:**
1.  **Add a Mutex:** Introduce a `sync.Mutex` to the `algorithmFactory` struct.
2.  **Protect Map Access:** Lock the mutex before checking for or writing to the `cryptos` map and unlock it immediately after.

- [ ] **Issue 2: Panicking in `init()` (Robustness Issue)**

**Analysis:**
The package's `init()` function will `panic` if it fails to create a default crypto instance. While this ensures the default instance is always valid, it can make the library fragile and cause an entire application to crash if there's a configuration issue.

**Remediation Plan:**
1.  **Graceful Error Handling:** Modify the `init` function to handle the error without panicking. It can log the error and leave the `defaultCrypto` instance as `nil`.
2.  **Check at Usage Time:** Update the functions that use `defaultCrypto` (e.g., `Verify`, `Generate`) to check if it is `nil` and return an error if it hasn't been initialized, guiding the user to properly configure the package.