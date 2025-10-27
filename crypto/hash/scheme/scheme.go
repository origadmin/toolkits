/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package scheme defines the core interfaces for hash algorithm implementations.
package scheme

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Scheme defines the interface for a specific, configured cryptographic hash algorithm.
// An instance of a Scheme is expected to be immutable once created.
type Scheme interface {
	// Spec returns the unique, structured specification of this scheme.
	Spec() types.Spec

	// Hash generates a hash for the given password, creating a new salt internally.
	Hash(password string) (*types.HashParts, error)

	// HashWithSalt generates a hash for the given password with a specified salt.
	HashWithSalt(password string, salt []byte) (*types.HashParts, error)

	// Verify checks if the given password matches the information stored in HashParts.
	// The implementation should rely on the configuration of the Scheme instance it is called on,
	// which is expected to be created by the factory with the correct parameters from the original hash.
	Verify(parts *types.HashParts, password string) error
}

// Factory defines the contract for creating Scheme instances.
// This is the sole entry point for creating any Scheme, used by both NewCrypto and Verify.
type Factory interface {
	Config() *types.Config
	// Create uses a unified Config object to create a Scheme instance.
	Create(spec types.Spec, cfg *types.Config) (Scheme, error)
}

// AlgorithmCreator is a function that creates an instance of Scheme given its type and configuration.
type AlgorithmCreator func(algSpec types.Spec, cfg *types.Config) (Scheme, error)

// Create implements the Factory interface.
func (c AlgorithmCreator) Create(spec types.Spec, cfg *types.Config) (Scheme, error) {
	return c(spec, cfg)
}

// AlgorithmConfig defines a function that returns a Config object for a specific algorithm.
type AlgorithmConfig func() *types.Config

// Config implements the AlgorithmConfig interface.
func (c AlgorithmConfig) Config() *types.Config {
	return c()
}
