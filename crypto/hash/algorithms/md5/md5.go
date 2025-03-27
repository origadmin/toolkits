/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package md5

import (
	"crypto/md5"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/base"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// MD5Crypto implements the MD5 hashing algorithm
type MD5Crypto struct {
	config *types.Config
}

// NewMD5Crypto creates a new MD5 crypto instance
func NewMD5Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return &MD5Crypto{
		config: config,
	}, nil
}

// Hash implements the hash method
func (c *MD5Crypto) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *MD5Crypto) HashWithSalt(password, salt string) (string, error) {
	hash := md5.Sum([]byte(password + salt))
	encoder := NewMD5HashEncoder()
	return encoder.Encode([]byte(salt), hash[:]), nil
}

// Verify implements the verify method
func (c *MD5Crypto) Verify(hashed, password string) error {
	encoder := NewMD5HashEncoder()
	parts, err := encoder.Decode(hashed)
	if err != nil {
		return err
	}

	if parts.Algorithm != types.TypeMD5 {
		return fmt.Errorf("algorithm mismatch")
	}

	newHash := md5.Sum([]byte(password + string(parts.Salt)))
	if string(newHash[:]) != string(parts.Hash) {
		return fmt.Errorf("password not match")
	}

	return nil
}

// MD5HashEncoder implements the hash encoder interface
type MD5HashEncoder struct {
	*base.BaseHashCodec
}

// NewMD5HashEncoder creates a new MD5 hash encoder
func NewMD5HashEncoder() interfaces.HashCodec {
	return &MD5HashEncoder{
		BaseHashCodec: base.NewBaseHashCodec(types.TypeMD5).(*base.BaseHashCodec),
	}
}
