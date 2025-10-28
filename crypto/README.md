# Crypto Toolkit

## Overview

The Crypto Toolkit provides a unified and extensible framework for cryptographic operations, with a primary focus on password hashing and verification.

## Hashing (`hash` package)

The `hash` package offers a consistent interface to work with a variety of hashing algorithms, from standard ones like SHA-256 and bcrypt to custom, user-defined implementations.

### Key Features

- **Unified Interface**: The `Crypto` interface provides simple `Hash` and `Verify` methods, abstracting away the complexities of each underlying algorithm.
- **Extensible by Design**: New algorithms can be easily integrated by implementing the `scheme.Scheme` and `scheme.Factory` interfaces and registering them.
- **Automatic Algorithm Detection**: The library automatically identifies the algorithm from the encoded hash string during verification, allowing for seamless algorithm upgrades. You can hash a password with SHA-256, and later verify it with a `Crypto` instance configured to use Argon2.
- **Tunable Parameters**: Algorithm-specific parameters, such as bcrypt's cost or Argon2's memory usage, can be configured at creation time using option functions (e.g., `bcrypt.WithCost`).
- **Built-in Salt Management**: Secure salt generation is handled automatically, but the library also supports hashing with a user-provided salt for specific use cases.
- **Verification Caching**: Repeated verification attempts for the same hash and password are automatically cached to reduce computational overhead.

### Basic Usage

This example demonstrates how to create a hash with a standard algorithm and verify it.

```go
package main

import (
	"fmt"
	"log"

	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func main() {
	password := "my-secret-password"

	// 1. Create a crypto instance with a specific algorithm and options.
	// Here, we use bcrypt with a custom cost.
	c, err := hash.NewCrypto(types.BCRYPT, bcrypt.WithCost(10))
	if err != nil {
		log.Fatalf("Failed to create crypto: %v", err)
	}

	// 2. Hash the password. The result is a single, encoded string that contains
	//    the algorithm name, parameters, salt, and the hash itself.
	hashed, err := c.Hash(password)
	if err != nil {
		log.Fatalf("Hashing failed: %v", err)
	}
	fmt.Println("Bcrypt Hash:", hashed)

	// 3. Verify the password. The `Verify` method automatically detects the algorithm
	//    from the hashed string and uses the correct logic to compare.
	if err := c.Verify(hashed, password); err != nil {
		fmt.Println("Verification failed!")
	} else {
		fmt.Println("Verification successful!")
	}
}
```

### Advanced Usage

#### Algorithm Upgrade

The `hash` package makes it easy to upgrade your hashing algorithms without forcing users to reset their passwords. The `Verify` method can check passwords against hashes created by any registered algorithm, regardless of the `Crypto` instance's default algorithm.

```go
func demonstrateUpgrade() {
	password := "password-to-migrate"

	// 1. Assume you have an old hash created with SHA-256.
	sha256Crypto, _ := hash.NewCrypto(types.SHA256)
	oldHash, _ := sha256Crypto.Hash(password)
	fmt.Println("Old SHA-256 Hash:", oldHash)

	// 2. Your application is now configured to use the more secure Argon2 algorithm.
	argon2Crypto, _ := hash.NewCrypto(types.ARGON2)

	// 3. A user logs in. You can still verify their password against the old SHA-256 hash.
	if err := argon2Crypto.Verify(oldHash, password); err == nil {
		fmt.Println("Verification of old hash successful!")

		// 4. (Recommended) Since verification was successful, create a new hash with the
		//    upgraded algorithm and store it in your database for future logins.
		newHash, _ := argon2Crypto.Hash(password)
		fmt.Println("New Argon2 Hash:", newHash)
		// database.UpdateUserPasswordHash(userID, newHash)
	}
}
```

#### Custom Algorithm

The framework is fully extensible. You can add your own hashing algorithm by implementing the `scheme.Scheme` and `scheme.Factory` interfaces.

See `examples/example_test.go` for a complete, working example of how to define and register a custom algorithm.

## Random Data Generation (`rand` package)

The `rand` package provides cryptographically secure, high-performance random string and byte generation. It is designed for scenarios requiring strong randomness, such as generating passwords, tokens, or cryptographic keys.

### Key Features

- **Cryptographically Secure**: Utilizes `crypto/rand.Reader` for all random byte generation, ensuring high-quality, unpredictable randomness suitable for security-sensitive applications.
- **Modulo Bias Elimination**: Employs rejection sampling to guarantee a statistically uniform distribution of characters from the chosen character set, even when the character set size is not a power of two.
- **High Performance**: Optimized for speed by using an internal buffer to minimize expensive system calls to `crypto/rand.Reader`.
- **Lazy Initialization for Global Functions**: Package-level convenience functions (`RandomString`, `RandomBytes`, etc.) use `sync.Once` to lazily initialize and reuse `Generator` instances, reducing startup overhead and resource consumption.
- **Flexible Character Sets**: Supports predefined character sets (digits, lowercase, uppercase, symbols, alphanumeric, all with symbols) and custom character sets.
- **Efficient `io.Reader` Implementation**: The `Read` method provides the most performant way to fill pre-allocated byte slices, avoiding internal allocations and making it ideal for high-throughput scenarios.

### Basic Usage

#### Generating Alphanumeric Random Strings

```go
package main

import (
	"fmt"
	"log"

	"github.com/origadmin/toolkits/crypto/rand"
)

func main() {
	// Generate a 16-character alphanumeric random string
	randomStr, err := rand.RandomString(16)
	if err != nil {
		log.Fatalf("Failed to generate random string: %v", err)
	}
	fmt.Printf("Alphanumeric Random String: %s\n", randomStr)

	// Generate 32 alphanumeric random bytes
	randomBytes, err := rand.RandomBytes(32)
	if err != nil {
		log.Fatalf("Failed to generate random bytes: %v", err)
	}
	fmt.Printf("Alphanumeric Random Bytes: %x\n", randomBytes)
}
```

#### Generating Random Strings with Symbols

```go
package main

import (
	"fmt"
	"log"

	"github.com/origadmin/toolkits/crypto/rand"
)

func main() {
	// Generate a 20-character random string including symbols
	randomStrWithSymbols, err := rand.RandomStringWithSymbols(20)
	if err != nil {
		log.Fatalf("Failed to generate random string with symbols: %v", err)
	}
	fmt.Printf("Random String with Symbols: %s\n", randomStrWithSymbols)
}
```

#### Using a Custom Character Set

```go
package main

import (
	"fmt"
	"log"

	"github.com/origadmin/toolkits/crypto/rand"
)

func main() {
	// Create a generator with a custom character set (e.g., only hex characters)
	hexCharset := "0123456789abcdef"
	hexGenerator := rand.NewGeneratorWithCharset(hexCharset)

	// Generate a 64-character hex string
	hexString, err := hexGenerator.RandString(64)
	if err != nil {
		log.Fatalf("Failed to generate hex string: %v", err)
	}
	fmt.Printf("Custom Hex String: %s\n", hexString)
}
```

#### High-Performance Byte Generation with `Read`

```go
package main

import (
	"fmt"
	"log"

	"github.com/origadmin/toolkits/crypto/rand"
)

func main() {
	// Create a generator for alphanumeric characters
	gen := rand.NewGenerator(rand.KindAlphanumeric)

	// Pre-allocate a byte slice
	buffer := make([]byte, 128)

	// Fill the buffer with random bytes using the Read method
	n, err := gen.Read(buffer)
	if err != nil {
		log.Fatalf("Failed to read random bytes: %v", err)
	}
	fmt.Printf("Read %d random bytes: %s\n", n, string(buffer))
}
```
