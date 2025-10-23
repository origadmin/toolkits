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
	Hash(password string) (*types.HashParts, error)
	// HashWithSalt generates a hash for the given password with the specified salt
	HashWithSalt(password string, salt []byte) (*types.HashParts, error)
	// Verify checks if the given hashed password matches the plaintext password
	Verify(parts *types.HashParts, oldPassword string) error
}

// AlgorithmConfig defines a function type that returns a default configuration for an algorithm.
type AlgorithmConfig func() *types.Config

// AlgorithmCreator defines a function type that creates a new Cryptographic instance.
// It accepts the resolved algorithm Type and a configuration.
type AlgorithmCreator func(algType types.Type, cfg *types.Config) (Cryptographic, error)
