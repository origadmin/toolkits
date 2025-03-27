/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package scrypt

import (
	"fmt"

	"golang.org/x/crypto/scrypt"

	"github.com/origadmin/toolkits/crypto/hash/base"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// ScryptCrypto implements the Scrypt hashing algorithm
type ScryptCrypto struct {
	config *types.Config
}

// NewScryptCrypto creates a new Scrypt crypto instance
func NewScryptCrypto(config *types.Config) (interfaces.Cryptographic, error) {
	return &ScryptCrypto{
		config: config,
	}, nil
}

// Hash implements the hash method
func (c *ScryptCrypto) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *ScryptCrypto) HashWithSalt(password, salt string) (string, error) {
	hash, err := scrypt.Key([]byte(password), []byte(salt), c.config.Scrypt.N, c.config.Scrypt.R, c.config.Scrypt.P, c.config.Scrypt.KeyLen)
	if err != nil {
		return "", err
	}
	encoder := NewScryptHashEncoder()
	return encoder.Encode([]byte(salt), hash), nil
}

// Verify implements the verify method
func (c *ScryptCrypto) Verify(hashed, password string) error {
	encoder := NewScryptHashEncoder()
	parts, err := encoder.Decode(hashed)
	if err != nil {
		return err
	}

	if parts.Algorithm != types.TypeScrypt {
		return fmt.Errorf("algorithm mismatch")
	}

	hash, err := scrypt.Key([]byte(password), parts.Salt, c.config.Scrypt.N, c.config.Scrypt.R, c.config.Scrypt.P, c.config.Scrypt.KeyLen)
	if err != nil {
		return err
	}

	if string(hash) != string(parts.Hash) {
		return fmt.Errorf("password not match")
	}

	return nil
}

// ScryptHashEncoder implements the hash encoder interface
type ScryptHashEncoder struct {
	*base.BaseHashCodec
}

// NewScryptHashEncoder creates a new Scrypt hash encoder
func NewScryptHashEncoder() interfaces.HashCodec {
	return &ScryptHashEncoder{
		BaseHashCodec: base.NewBaseHashCodec(types.TypeScrypt).(*base.BaseHashCodec),
	}
}
