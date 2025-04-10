/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
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
	codec   interfaces.Codec
	crypto  interfaces.Cryptographic
	factory internalFactory
}

func (c *crypto) Type() types.Type {
	return c.codec.Type()
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

	algType, configName := core.AlgorithmTypeHash(parts.Algorithm)
	// Get algorithm instance from cache or create new one
	cryptographic, err := c.factory.create(algType, types.WithName(configName))
	if err != nil {
		return err
	}
	return cryptographic.Verify(parts, password)
}

// NewCrypto creates a new cryptographic instance
func NewCrypto(alg types.Type, opts ...types.Option) (Crypto, error) {
	algType, configName := core.AlgorithmTypeHash(alg)

	factory := &algorithmFactory{
		cryptos: make(map[types.Type]interfaces.Cryptographic),
	}
	if configName != "" {
		opts = append(opts, types.WithName(configName))
	}

	cryptographic, err := factory.create(algType, opts...)
	if err != nil {
		return nil, err
	}

	// Create cryptographic instance
	return &crypto{
		crypto:  cryptographic,
		codec:   core.NewCodec(alg),
		factory: factory,
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
