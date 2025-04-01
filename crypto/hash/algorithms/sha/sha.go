/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package sha

import (
	"crypto/subtle"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// SHA implements the SHA hashing algorithm
type SHA struct {
	config   *types.Config
	codec    interfaces.Codec
	hashHash core.Hash
}

func (c *SHA) Type() string {
	return c.hashHash.String()
}

type ConfigValidator struct {
}

func (v ConfigValidator) Validate(config *types.Config) interface{} {
	if config.SaltLength < 8 {
		return core.ErrSaltLengthTooShort
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
	hashHash, err := core.ParseHash(hashType.String())
	if err != nil {
		return nil, err
	}

	return &SHA{
		config:   config,
		codec:    core.NewCodec(hashType),
		hashHash: hashHash,
	}, nil
}

func NewSha1Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha1, config)
}

func NewSha256Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha256, config)
}

func NewSha512Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHACrypto(types.TypeSha512, config)
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}

// Hash implements the hash method
func (c *SHA) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
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
func (c *SHA) Verify(hashed, password string) error {
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}

	if parts.Algorithm.String() != c.hashHash.String() {
		return core.ErrAlgorithmMismatch
	}
	newHash := c.hashHash.New().Sum([]byte(password + string(parts.Salt)))
	if subtle.ConstantTimeCompare(newHash, parts.Hash) != 1 {
		return core.ErrPasswordNotMatch
	}

	return nil
}
