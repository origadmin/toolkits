/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package interfaces

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Cryptographic defines the interface for cryptographic operations
type Cryptographic interface {
	Type() types.Type
	// Salt generates a random salt
	Salt() ([]byte, error)
	// Hash generates a hash for the given password
	Hash(password string) (*types.HashParts, error)
	// HashWithSalt generates a hash for the given password with the specified salt
	HashWithSalt(password string, salt []byte) (*types.HashParts, error)
	// Verify checks if the given hashed password matches the plaintext password
	Verify(parts *types.HashParts, oldPassword string) error
}
