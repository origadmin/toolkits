/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package sha1

import (
	"crypto/sha1"
	"crypto/subtle"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// Sha1 implements the SHA1 hashing algorithm
type Sha1 struct {
	config *types.Config
	codec  interfaces.Codec
}

func (c *Sha1) Type() string {
	return types.TypeSha1.String()
}

type ConfigValidator struct {
}

func (v ConfigValidator) Validate(config *types.Config) interface{} {
	if config.SaltLength < 8 {
		return core.ErrSaltLengthTooShort
	}
	return nil
}

// NewSHA1Crypto creates a new SHA1 crypto instance
func NewSHA1Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	validator := &ConfigValidator{}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid sha1 config: %v", err)
	}
	return &Sha1{
		config: config,
		codec:  core.NewCodec(types.TypeSha1),
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}

// Hash implements the hash method
func (c *Sha1) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, string(salt))
}

// HashWithSalt implements the hash with salt method
func (c *Sha1) HashWithSalt(password, salt string) (string, error) {
	hash := sha1.Sum([]byte(password + salt))
	return c.codec.Encode([]byte(salt), hash[:]), nil
}

// Verify implements the verify method
func (c *Sha1) Verify(hashed, password string) error {
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}

	if parts.Algorithm != types.TypeSha1 {
		return core.ErrAlgorithmMismatch
	}

	newHash := sha1.Sum([]byte(password + string(parts.Salt)))
	if subtle.ConstantTimeCompare(newHash[:], parts.Hash) != 1 {
		return core.ErrPasswordNotMatch
	}

	return nil
}
