/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package bcrypt

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"

	"github.com/origadmin/toolkits/crypto/hash/base"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// BcryptCrypto implements the Bcrypt hashing algorithm
type BcryptCrypto struct {
	config *types.Config
	codec  interfaces.HashCodec
}

// NewBcryptCrypto creates a new Bcrypt crypto instance
func NewBcryptCrypto(config *types.Config) (interfaces.Cryptographic, error) {
	return &BcryptCrypto{
		config: config,
		codec:  base.GetCodec(types.TypeBcrypt),
	}, nil
}

// Hash implements the hash method
func (c *BcryptCrypto) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *BcryptCrypto) HashWithSalt(password, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+salt), c.config.Cost)
	if err != nil {
		return "", err
	}
	return c.codec.Encode([]byte(salt), hash), nil
}

// Verify implements the verify method
func (c *BcryptCrypto) Verify(hashed, password string) error {
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
