# Crypto Toolkit

## Overview

The Crypto Toolkit is a powerful library designed to provide cryptographic functionalities, including hashing, encryption, and decryption. This toolkit is built with performance and security in mind, making it an essential component for any application that requires secure data handling.

## Hashing

One of the standout features of the Crypto Toolkit is its robust hashing capabilities. Hashing is a fundamental aspect of cryptography, used to ensure data integrity and authenticity. The toolkit supports various hashing algorithms, allowing developers to choose the most suitable one for their needs.

### Key Features of Hashing:

- **Speed and Efficiency**: The hashing functions are optimized for performance, ensuring that even large datasets can be processed quickly without compromising security.
- **Multiple Algorithms**: The toolkit supports a variety of hashing algorithms, including SHA-256, SHA-512, Argon2, Bcrypt, Scrypt, and others, providing flexibility for different use cases. This allows for easy upgrades to stronger algorithms without modifying the password verification logic.
- **Algorithm Compatibility**: When users need to enhance security or switch algorithms, the toolkit ensures compatibility with existing hashed passwords. This means you can seamlessly transition to a new hashing algorithm without needing to rehash all stored passwords.
- **Built-in Salt Management**: The toolkit includes built-in salt management, eliminating the need for users to manage salts manually. This enhances security by ensuring that each hash is unique, even for identical inputs.
- **Collision Resistance**: The hashing algorithms are designed to minimize the risk of collisions, ensuring that each unique input produces a unique hash output.
- **Secure Data Integrity**: By using cryptographic hashes, you can verify the integrity of your data, ensuring that it has not been altered or tampered with during transmission or storage.
- **Customizable Parameters**: Each hashing algorithm allows for customizable parameters such as salt length, iteration count, and memory cost, enabling developers to fine-tune the security and performance of their hashing operations.
- **Caching Mechanism**: The toolkit includes a caching mechanism to optimize the verification process, reducing the computational overhead for frequently verified hashes.

### Supported Algorithms

The Crypto Toolkit supports a wide range of hashing algorithms, including but not limited to:

- **Argon2**: A modern, memory-hard hashing algorithm designed to resist GPU and ASIC attacks.
- **Bcrypt**: A widely-used hashing algorithm that is resistant to brute-force attacks.
- **Scrypt**: A memory-intensive hashing algorithm that is designed to be computationally expensive.
- **SHA-256/SHA-512**: Secure Hash Algorithms that are widely used for data integrity checks.
- **PBKDF2**: A key derivation function that is commonly used for password hashing.
- **HMAC**: A keyed-hash message authentication code that provides both data integrity and authenticity.

### Compatibility and Salt Management

The Crypto Toolkit is designed to be highly compatible with existing hashed passwords. When upgrading to a stronger hashing algorithm, the toolkit can still verify passwords that were hashed with the older algorithm. This ensures a smooth transition without requiring users to reset their passwords.
Additionally, the toolkit provides built-in salt management, which automatically generates and manages salts for each hash. This eliminates the need for developers to manually handle salts, reducing the risk of security vulnerabilities. However, the toolkit also supports custom salt generation, allowing developers to use their own salt values if needed.

### Example of Compatibility and Salt Management

```go
package main

import (
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func main() {
	// The hash is generated using the SHA-256 algorithm
	sha256Crypto, err := hash.NewCrypto(types.TypeSha256)
	if err != nil {
		fmt.Println("Failed to create SHA-256 crypto:", err)
		return
	}

	sha256Hash, err := sha256Crypto.Hash("myPassword123")
	if err != nil {
		fmt.Println("Failed to hash with SHA-256:", err)
		return
	}
	fmt.Println("SHA-256 Hash:", sha256Hash)

	// Upgrade to the Argon2 algorithm
	argon2Crypto, err := hash.NewCrypto(types.TypeArgon2)
	if err != nil {
		fmt.Println("Failed to create Argon2 crypto:", err)
		return
	}

	// Verify the old SHA-256 hash
	err = argon2Crypto.Verify(sha256Hash, "myPassword123")
	if err != nil {
		fmt.Println("SHA-256 verification failed:", err)
	} else {
		fmt.Println("SHA-256 verification succeeded")
	}

	// Use Argon2 to generate a new hash
	argon2Hash, err := argon2Crypto.Hash("myPassword123")
	if err != nil {
		fmt.Println("Failed to hash with Argon2:", err)
		return
	}
	fmt.Println("Argon2 Hash:", argon2Hash)

	// Verify the Argon2 hash
	err = argon2Crypto.Verify(argon2Hash, "myPassword123")
	if err != nil {
		fmt.Println("Argon2 verification failed:", err)
	} else {
		fmt.Println("Argon2 verification succeeded")
	}
}
```
