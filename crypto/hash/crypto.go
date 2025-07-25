/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"github.com/origadmin/toolkits/crypto/hash/codec"
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

	// Get algorithm instance from cache or create new one
	cryptographic, err := c.factory.create(parts.Algorithm)
	if err != nil {
		return err
	}
	return cryptographic.Verify(parts, password)
}

// NewCrypto creates a new cryptographic instance
func NewCrypto(alg string, opts ...types.Option) (Crypto, error) {
	factory := &algorithmFactory{
		cryptos: make(map[types.Type]interfaces.Cryptographic),
	}
	algorithm, err := types.ParseAlgorithm(alg)
	if err != nil {
		return nil, err
	}

	cryptographic, err := factory.create(algorithm, opts...)
	if err != nil {
		return nil, err
	}

	// Create cryptographic instance
	return &crypto{
		crypto:  cryptographic,
		codec:   codec.NewCodec(algorithm),
		factory: factory,
	}, nil
}

// RegisterAlgorithm registers a new hash algorithm
func RegisterAlgorithm(t types.Type, creator AlgorithmCreator, defaultConfig AlgorithmConfig) {
	algorithms[t.Name] = algorithm{
		creator:       creator,
		defaultConfig: defaultConfig,
	}
}
