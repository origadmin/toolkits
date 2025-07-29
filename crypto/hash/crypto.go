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
	HashWithSalt(password string, salt []byte) (string, error)
	Verify(hashed, password string) error
}

type crypto struct {
	codec   interfaces.Codec
	crypto  interfaces.Cryptographic
	factory internalFactory
}

func (c *crypto) Type() types.Type {
	return c.crypto.Type()
}

func (c *crypto) Hash(password string) (string, error) {
	hashParts, err := c.crypto.Hash(password)
	if err != nil {
		return "", err
	}
	return c.codec.Encode(hashParts)
}

func (c *crypto) HashWithSalt(password string, salt []byte) (string, error) {
	hashParts, err := c.crypto.HashWithSalt(password, salt)
	if err != nil {
		return "", err
	}
	return c.codec.Encode(hashParts)
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
func NewCrypto(cryptoType string, opts ...types.Option) (Crypto, error) {
	factory := &algorithmFactory{
		cryptos: make(map[string]interfaces.Cryptographic),
	}

	cryptographic, err := factory.create(cryptoType, opts...)
	if err != nil {
		return nil, err
	}

	// Create cryptographic instance
	return &crypto{
		crypto:  cryptographic,
		codec:   codec.NewCodec(),
		factory: factory,
	}, nil
}

// RegisterAlgorithm registers a new hash algorithm
func RegisterAlgorithm(t types.Type, creator AlgorithmCreator, defaultConfig AlgorithmConfig) {
	algorithmMap[t.Name] = algorithm{
		creator:       creator,
		defaultConfig: defaultConfig,
	}
}
