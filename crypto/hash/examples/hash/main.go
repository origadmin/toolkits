/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package main provides an example of using the hash package
package main

import (
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func main() {
	// The hash is generated using the SHA-256 algorithm
	sha256Crypto, err := hash.NewCrypto(types.SHA256)
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

	// Verify the SHA-256 hash
	err = sha256Crypto.Verify(sha256Hash, "myPassword123")
	if err != nil {
		fmt.Println("SHA-256 verification failed:", err)
	} else {
		fmt.Println("SHA-256 verification succeeded")
	}

	// Upgrade to the Argon2 algorithm
	argon2Crypto, err := hash.NewCrypto(types.ARGON2)
	if err != nil {
		fmt.Println("Failed to create Argon2 crypto:", err)
		return
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

	// Verify the old hash
	err = argon2Crypto.Verify(sha256Hash, "myPassword123")
	if err != nil {
		fmt.Println("Argon2 SHA-256 verification failed:", err)
	} else {
		fmt.Println("Argon2 SHA-256 verification succeeded")
	}

	err = hash.Verify(sha256Hash, "myPassword123")
	if err != nil {
		fmt.Println("Global SHA-256 verification failed:", err)
	} else {
		fmt.Println("Global SHA-256 verification succeeded")
	}
}
