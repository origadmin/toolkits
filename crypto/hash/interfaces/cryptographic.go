/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package interfaces

// Cryptographic defines the interface for cryptographic operations
type Cryptographic interface {
	Type() string
	// Hash generates a hash for the given password
	Hash(password string) (string, error)
	// HashWithSalt generates a hash for the given password with the specified salt
	HashWithSalt(password, salt string) (string, error)
	// Verify checks if the given hashed password matches the plaintext password
	Verify(hashed, password string) error
}
