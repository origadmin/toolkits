/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package bcrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// Bcrypt implements the Bcrypt hashing algorithm
type Bcrypt struct {
	config *types.Config
	codec  interfaces.Codec
}

func (c *Bcrypt) Type() string {
	return types.TypeBcrypt.String()
}

type ConfigValidator struct {
}

func (v ConfigValidator) Validate(config *types.Config) interface{} {
	if config.SaltLength < 8 {
		return core.ErrSaltLengthTooShort
	}
	if config.Cost < 4 || config.Cost > 31 {
		return core.ErrCostOutOfRange
	}
	return nil
}

// NewBcryptCrypto creates a new Bcrypt crypto instance
func NewBcryptCrypto(config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	validator := &ConfigValidator{}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid bcrypt config: %v", err)
	}
	return &Bcrypt{
		config: config,
		codec:  core.NewCodec(types.TypeBcrypt),
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
		Cost:       10,
	}
}

// Hash implements the hash method
func (c *Bcrypt) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, string(salt))
}

// HashWithSalt implements the hash with salt method
func (c *Bcrypt) HashWithSalt(password, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+salt), c.config.Cost)
	if err != nil {
		return "", err
	}
	return c.codec.Encode([]byte(salt), hash), nil
}

// Verify implements the verify method
func (c *Bcrypt) Verify(hashed, password string) error {
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}
	if parts.Algorithm != types.TypeBcrypt {
		return fmt.Errorf("algorithm mismatch")
	}
	err = bcrypt.CompareHashAndPassword(parts.Hash, []byte(password+string(parts.Salt)))
	if err != nil {
		return fmt.Errorf("password not match")
	}
	return nil
}
