/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"fmt"

	"github.com/goexts/generic/configure"

	"github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type Crypto interface {
	Type() types.Type
	Hash(password string) (string, error)
	HashWithSalt(password string, salt []byte) (string, error)
	Verify(hashed, password string) error
}

type crypto struct {
	codec   interfaces.Codec
	// algImpl is the cryptographic implementation used for hashing, wrapped with cachedVerifier
	algImpl interfaces.Cryptographic
}

func (c *crypto) Type() types.Type {
	return c.algImpl.Type()
}

func (c *crypto) Hash(password string) (string, error) {
	hashParts, err := c.algImpl.Hash(password)
	if err != nil {
		return "", err
	}
	return c.codec.Encode(hashParts)
}

func (c *crypto) HashWithSalt(password string, salt []byte) (string, error) {
	hashParts, err := c.algImpl.HashWithSalt(password, salt)
	if err != nil {
		return "", err
	}
	return c.codec.Encode(hashParts)
}

func (c *crypto) Verify(hashed, password string) error {
	// Check for nil or empty hashed string early
	if hashed == "" {
		return errors.ErrInvalidHash
	}

	// Decode the hash value to get algorithm type and parts
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}

	// Perform nil checks for HashParts directly here
	if parts == nil || parts.Hash == nil || parts.Salt == nil {
		return errors.ErrInvalidHashParts
	}

	// Get algorithm instance from global factory based on the decoded algorithm
	factory := getFactory()
	cryptographic, err := factory.create(parts.Algorithm) // Pass types.Type directly
	if err != nil {
		return err
	}

	// Wrap the created algorithm with a cached verifier for performance
	cachedAlg := NewCachedVerifier(cryptographic)

	return cachedAlg.Verify(parts, password)
}

// NewCrypto creates a new cryptographic instance
func NewCrypto(algName string, opts ...types.Option) (Crypto, error) {
	// 1. Parse the algorithm name string into a structured Type
	algType, err := types.ParseType(algName)
	if err != nil {
		return nil, err
	}

	// 2. Look up the algorithm entry to get its specific resolver
	algEntry, exists := algorithmMap[algType.Name]
	if !exists {
		return nil, fmt.Errorf("unsupported algorithm: %s", algType.String())
	}

	// 3. Resolve the parsed Type to its canonical form using the algorithm's specific resolver
	resolvedAlgType, err := algEntry.resolver.ResolveType(algType) // Use algEntry.resolver
	if err != nil {
		return nil, fmt.Errorf("failed to resolve algorithm type %s: %w", algType.String(), err)
	}

	// 4. Apply options to default config and create the cryptographic instance
	cfg := configure.Apply(algEntry.defaultConfig(), opts)
	cryptographic, err := algEntry.creator(resolvedAlgType, cfg) // Pass the resolved type
	if err != nil {
		return nil, err
	}

	// Wrap the cryptographic implementation with a cached verifier for hashing operations
	finalAlg := NewCachedVerifier(cryptographic)

	// Create cryptographic instance
	return &crypto{
		algImpl: finalAlg,
		codec:   codec.NewCodec(),
	},
	nil
}

// RegisterAlgorithm registers a new hash algorithm
func RegisterAlgorithm(algType types.Type, creator interfaces.AlgorithmCreator, defaultConfig interfaces.AlgorithmConfig) {
	algorithmMap[algType.Name] = algorithm{
		algType:       algType,
		creator:       creator,
		defaultConfig: defaultConfig,
		resolver:      defaultTypeResolver,
	}
}

// Removed safeVerifier type and its methods as its logic is now in crypto.Verify
