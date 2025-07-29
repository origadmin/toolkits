/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"fmt"

	"github.com/goexts/generic/settings"

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

	// Decode the hash value
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}

	// Get algorithm instance from global factory
	factory := getFactory()

	cryptographic, err := factory.create(parts.Algorithm)
	if err != nil {
		return err
	}

	// Wrap the created algorithm with a safe verifier before using it
	safeAlg := &safeVerifier{wrapped: cryptographic}

	return safeAlg.Verify(parts, password)
}

// NewCrypto creates a new cryptographic instance
func NewCrypto(algName string, opts ...types.Option) (Crypto, error) {
	algType, err := types.ParseType(algName)
	if err != nil {
		return nil, err
	}
	algorithm, exists := algorithmMap[algType.Name]
	if !exists {
		return nil, fmt.Errorf("unsupported algorithm: %s", algType)
	}

	// Apply opts to default config, create instance for HASH
	cfg := settings.Apply(algorithm.defaultConfig(), opts)
	cryptographic, err := algorithm.creator(algType, cfg)
	if err != nil {
		return nil, err
	}

	// Create cryptographic instance
	return &crypto{
		algImpl: cryptographic,
		codec:   codec.NewCodec(),
	}, nil
}

// RegisterAlgorithm registers a new hash algorithm
func RegisterAlgorithm(algType types.Type, creator AlgorithmCreator, defaultConfig AlgorithmConfig) {
	algorithmMap[algType.Name] = algorithm{
		creator:       creator,
		defaultConfig: defaultConfig,
	}
}

// safeVerifier wraps a Cryptographic implementation to add nil checks for HashParts
type safeVerifier struct {
	wrapped interfaces.Cryptographic
}

func (s *safeVerifier) Type() types.Type {
	return s.wrapped.Type()
}

func (s *safeVerifier) Hash(password string) (*types.HashParts, error) {
	return s.wrapped.Hash(password)
}

func (s *safeVerifier) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	return s.wrapped.HashWithSalt(password, salt)
}

func (s *safeVerifier) Verify(parts *types.HashParts, password string) error {
	// Perform nil checks before delegating to the wrapped verifier
	if parts == nil || parts.Hash == nil || parts.Salt == nil {
		return errors.ErrInvalidHashParts
	}
	return s.wrapped.Verify(parts, password)
}