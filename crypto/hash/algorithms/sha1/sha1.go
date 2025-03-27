/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package sha1

import (
	"crypto/sha1"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/base"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// SHA1Crypto implements the SHA1 hashing algorithm
type SHA1Crypto struct {
	config *types.Config
}

// NewSHA1Crypto creates a new SHA1 crypto instance
func NewSHA1Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return &SHA1Crypto{
		config: config,
	}, nil
}

// Hash implements the hash method
func (c *SHA1Crypto) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *SHA1Crypto) HashWithSalt(password, salt string) (string, error) {
	hash := sha1.Sum([]byte(password + salt))
	encoder := NewSHA1HashEncoder()
	return encoder.Encode([]byte(salt), hash[:]), nil
}

// Verify implements the verify method
func (c *SHA1Crypto) Verify(hashed, password string) error {
	encoder := NewSHA1HashEncoder()
	parts, err := encoder.Decode(hashed)
	if err != nil {
		return err
	}

	if parts.Algorithm != types.TypeSHA1 {
		return fmt.Errorf("algorithm mismatch")
	}

	newHash := sha1.Sum([]byte(password + string(parts.Salt)))
	if string(newHash[:]) != string(parts.Hash) {
		return fmt.Errorf("password not match")
	}

	return nil
}

// SHA1HashEncoder implements the hash encoder interface
type SHA1HashEncoder struct {
	*base.BaseHashCodec
}

// NewSHA1HashEncoder creates a new SHA1 hash encoder
func NewSHA1HashEncoder() interfaces.HashCodec {
	return &SHA1HashEncoder{
		BaseHashCodec: base.NewBaseHashCodec(types.TypeSHA1).(*base.BaseHashCodec),
	}
}
