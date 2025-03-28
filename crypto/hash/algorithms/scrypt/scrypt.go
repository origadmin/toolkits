/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package scrypt

import (
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/scrypt"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// Scrypt implements the Scrypt hashing algorithm
type Scrypt struct {
	config *types.Config
	codec  interfaces.Codec
}

type ConfigValidator struct {
}

func (v ConfigValidator) Validate(config *types.Config) interface{} {
	if config.SaltLength < 8 {
		return fmt.Errorf("salt length must be at least 8 bytes")
	}
	return nil
}

// NewScryptCrypto creates a new Scrypt crypto instance
func NewScryptCrypto(config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = types.DefaultConfig()
	}
	validator := &ConfigValidator{}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid scrypt config: %v", err)
	}
	return &Scrypt{
		config: config,
		codec:  core.NewCodec(types.TypeScrypt),
	}, nil
}

// Hash implements the hash method
func (c *Scrypt) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *Scrypt) HashWithSalt(password, salt string) (string, error) {
	hash, err := scrypt.Key([]byte(password), []byte(salt), c.config.Scrypt.N, c.config.Scrypt.R, c.config.Scrypt.P, c.config.Scrypt.KeyLen)
	if err != nil {
		return "", err
	}
	return c.codec.Encode([]byte(salt), hash), nil
}

// Verify implements the verify method
func (c *Scrypt) Verify(hashed, password string) error {
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}

	if parts.Algorithm != types.TypeScrypt {
		return core.ErrAlgorithmMismatch
	}

	hash, err := scrypt.Key([]byte(password), parts.Salt, c.config.Scrypt.N, c.config.Scrypt.R, c.config.Scrypt.P, c.config.Scrypt.KeyLen)
	if err != nil {
		return err
	}

	if subtle.ConstantTimeCompare(hash, parts.Hash) != 1 {
		return core.ErrPasswordNotMatch
	}

	return nil
}
