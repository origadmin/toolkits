/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package scheme implements the functions, types, and interfaces for the module.
package scheme

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Scheme defines the interface for cryptographic operations
type Scheme interface {
	Spec() types.Spec
	// Hash generates a hash for the given password
	Hash(password string) (*types.HashParts, error)
	// HashWithSalt generates a hash for the given password with the specified salt
	HashWithSalt(password string, salt []byte) (*types.HashParts, error)
	// Verify checks if the given hashed password matches the plaintext password
	Verify(parts *types.HashParts, oldPassword string) error
}

// AlgorithmConfig defines a function type that returns a default configuration for an algorithm.
type AlgorithmConfig func() *types.Config

// AlgorithmCreator defines a function type that creates a new Scheme instance.
// It accepts the resolved algorithm Spec and a configuration.
type AlgorithmCreator func(algSpec types.Spec, cfg *types.Config) (Scheme, error)
