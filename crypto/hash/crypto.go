/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"fmt"

	"github.com/goexts/generic/settings"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type Crypto interface {
	Type() types.Type
	Hash(password string) (string, error)
	HashWithSalt(password, salt string) (string, error)
	Verify(hashed, password string) error
}

type crypto struct {
	algorithm types.Type
	codec     interfaces.Codec
	crypto    interfaces.Cryptographic
	cryptos   map[types.Type]interfaces.Cryptographic
}

func (c *crypto) Type() types.Type {
	return c.algorithm
}

func (c *crypto) Hash(password string) (string, error) {
	return c.crypto.Hash(password)
}

func (c *crypto) HashWithSalt(password, salt string) (string, error) {
	return c.crypto.HashWithSalt(password, salt)
}

func (c *crypto) Verify(hashed, password string) error {
	// Decode the hash value
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}

	// Get algorithm instance from cache or create new one
	alg, exists := c.cryptos[parts.Algorithm]
	if !exists {
		algorithm, exists := algorithms[parts.Algorithm]
		if !exists {
			return fmt.Errorf("unsupported algorithm: %s", parts.Algorithm)
		}

		// Create cryptographic instance and cache it
		var err error
		cfg := &types.Config{}
		if algorithm.defaultConfig != nil {
			cfg = algorithm.defaultConfig()
		}
		alg, err = algorithm.creator(cfg)
		if err != nil {
			return err
		}
		c.cryptos[parts.Algorithm] = alg
	}

	return alg.Verify(parts, password)
}

// NewCrypto creates a new cryptographic instance
func NewCrypto(alg types.Type, opts ...types.Option) (Crypto, error) {
	// Get algorithm creator and default config
	algorithm, exists := algorithms[alg]
	if !exists {
		return nil, fmt.Errorf("unsupported algorithm: %s", alg)
	}
	cfg := &types.Config{}
	if algorithm.defaultConfig != nil {
		cfg = algorithm.defaultConfig()
	}

	// Apply default config if not set
	cfg = settings.Apply(cfg, opts)
	cryptographic, err := algorithm.creator(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create cryptographic instance: %v", err)
	}
	// Create cryptographic instance
	return &crypto{
		algorithm: alg,
		crypto:    cryptographic,
		codec:     core.NewCodec(alg),
		cryptos:   make(map[types.Type]interfaces.Cryptographic),
	}, nil
}

// RegisterAlgorithm registers a new hash algorithm
func RegisterAlgorithm(t types.Type, creator AlgorithmCreator, defaultConfig AlgorithmConfig) {
	algorithms[t] = algorithm{
		creator:       creator,
		defaultConfig: defaultConfig,
	}
}

// Verify verifies a password
func Verify(hashed, password string) error {
	return defaultCrypto.Verify(hashed, password)
}

func Generate(password string) (string, error) {
	return defaultCrypto.Hash(password)
}

func GenerateWithSalt(password, salt string) (string, error) {
	return defaultCrypto.HashWithSalt(password, salt)
}
