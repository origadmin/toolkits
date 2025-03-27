/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package sha256

import (
	"crypto/sha256"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/base"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// SHA256Crypto implements the SHA256 hashing algorithm
type SHA256Crypto struct {
	config *types.Config
}

// NewSHA256Crypto creates a new SHA256 crypto instance
func NewSHA256Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return &SHA256Crypto{
		config: config,
	}, nil
}

// Hash implements the hash method
func (c *SHA256Crypto) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *SHA256Crypto) HashWithSalt(password, salt string) (string, error) {
	hash := sha256.Sum256([]byte(password + salt))
	encoder := NewSHA256HashEncoder()
	return encoder.Encode([]byte(salt), hash[:]), nil
}

// Verify implements the verify method
func (c *SHA256Crypto) Verify(hashed, password string) error {
	encoder := NewSHA256HashEncoder()
	parts, err := encoder.Decode(hashed)
	if err != nil {
		return err
	}

	if parts.Algorithm != types.TypeSHA256 {
		return fmt.Errorf("algorithm mismatch")
	}

	newHash := sha256.Sum256([]byte(password + string(parts.Salt)))
	if string(newHash[:]) != string(parts.Hash) {
		return fmt.Errorf("password not match")
	}

	return nil
}

// SHA256HashEncoder implements the hash encoder interface
type SHA256HashEncoder struct {
	*base.BaseHashCodec
}

// NewSHA256HashEncoder creates a new SHA256 hash encoder
func NewSHA256HashEncoder() interfaces.HashCodec {
	return &SHA256HashEncoder{
		BaseHashCodec: base.NewBaseHashCodec(types.TypeSHA256).(*base.BaseHashCodec),
	}
}
