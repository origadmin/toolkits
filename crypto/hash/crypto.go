/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Crypto defines the interface for a cryptographic instance.
// It can be used to hash new passwords with its configured scheme and verify any hash.
type Crypto interface {
	Spec() types.Spec
	Hash(password string) (string, error)
	HashWithSalt(password string, salt []byte) (string, error)
	Verify(hashed, password string) error
}

// crypto is the internal implementation of the Crypto interface.
type crypto struct {
	algImpl scheme.Scheme
}

var globalCodec = codec.NewCodec()

func (c *crypto) Spec() types.Spec {
	return c.algImpl.Spec()
}

func (c *crypto) Hash(password string) (string, error) {
	hashParts, err := c.algImpl.Hash(password)
	if err != nil {
		return "", err
	}
	return globalCodec.Encode(hashParts)
}

func (c *crypto) HashWithSalt(password string, salt []byte) (string, error) {
	hashParts, err := c.algImpl.HashWithSalt(password, salt)
	if err != nil {
		return "", err
	}
	return globalCodec.Encode(hashParts)
}

// Verify provides a convenient way to verify a hash using the crypto instance.
// It delegates the call to the package-level Verify function, which uses the default factory.
func (c *crypto) Verify(hashed, password string) error {
	return Verify(hashed, password)
}
