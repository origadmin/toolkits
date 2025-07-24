/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package sha

import (
	"crypto/subtle"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// SHA implements the SHA hashing algorithm
type SHA struct {
	config   *types.Config
	codec    interfaces.Codec
	hashHash stdhash.Hash
}

func (c *SHA) Type() string {
	return c.codec.Type().String()
}

type ConfigValidator struct {
}

func (v ConfigValidator) Validate(config *types.Config) interface{} {
	if config.SaltLength < 8 {
		return errors.ErrSaltLengthTooShort
	}
	return nil
}

// NewSHACrypto creates a new SHA crypto instance
func NewSHACrypto(hashType types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	validator := &ConfigValidator{}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid sha config: %v", err)
	}
	hashHash, err := stdhash.ParseHash(hashType.String())
	if err != nil {
		return nil, err
	}

	return &SHA{
		config:   config,
		codec:    codec.NewCodec(hashType),
		hashHash: hashHash,
	}, nil
}

func NewSha1Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha1, config)
}

func NewSha224Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha224, config)
}

func NewSha256Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha256, config)
}

func NewSha512Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha512, config)
}

func NewSha3224(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha3224, config)
}

func NewSha3256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha3256, config)
}

func NewSha384Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha384, config)
}

func NewTypeSha3512(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha3512, config)
}

func NewSha3512(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha3512, config)
}
func NewSha3512224(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha3512224, config)
}

func NewSha3512256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha3512256, config)
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}

// Hash implements the hash method
func (c *SHA) Hash(password string) (string, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, string(salt))
}

// HashWithSalt implements the hash with salt method
func (c *SHA) HashWithSalt(password, salt string) (string, error) {
	newHash := c.hashHash.New().Sum([]byte(password + salt))
	return c.codec.Encode([]byte(salt), newHash[:]), nil
}

// Verify implements the verify method
func (c *SHA) Verify(parts *types.HashParts, password string) error {
	if !parts.Algorithm.Is(c.codec.Type()) {
		return errors.ErrAlgorithmMismatch
	}
	newHash := c.hashHash.New().Sum([]byte(password + string(parts.Salt)))
	if subtle.ConstantTimeCompare(newHash, parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}

	return nil
}
