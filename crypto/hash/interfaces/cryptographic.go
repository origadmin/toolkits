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
	// Hash generates a hash for the given password
	Hash(password string) (string, error)
	// HashWithSalt generates a hash for the given password with the specified salt
	HashWithSalt(password, salt string) (string, error)
	// Verify checks if the given hashed password matches the plaintext password
	Verify(parts *types.HashParts, password string) error
}
