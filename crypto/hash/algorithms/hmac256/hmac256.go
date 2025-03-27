/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package hmac256

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/base"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// HMAC256Crypto implements the HMAC-SHA256 hashing algorithm
type HMAC256Crypto struct {
	config *types.Config
}

// NewHMAC256Crypto creates a new HMAC256 crypto instance
func NewHMAC256Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return &HMAC256Crypto{
		config: config,
	}, nil
}

// Hash implements the hash method
func (c *HMAC256Crypto) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *HMAC256Crypto) HashWithSalt(password, salt string) (string, error) {
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(password))
	hash := h.Sum(nil)
	encoder := NewHMAC256HashEncoder()
	return encoder.Encode([]byte(salt), hash), nil
}

// Verify implements the verify method
func (c *HMAC256Crypto) Verify(hashed, password string) error {
	encoder := NewHMAC256HashEncoder()
	parts, err := encoder.Decode(hashed)
	if err != nil {
		return err
	}

	if parts.Algorithm != types.TypeHMAC256 {
		return fmt.Errorf("algorithm mismatch")
	}

	h := hmac.New(sha256.New, parts.Salt)
	h.Write([]byte(password))
	newHash := h.Sum(nil)
	if string(newHash) != string(parts.Hash) {
		return fmt.Errorf("password not match")
	}

	return nil
}

// HMAC256HashEncoder implements the hash encoder interface
type HMAC256HashEncoder struct {
	*base.BaseHashCodec
}

// NewHMAC256HashEncoder creates a new HMAC256 hash encoder
func NewHMAC256HashEncoder() interfaces.HashCodec {
	return &HMAC256HashEncoder{
		BaseHashCodec: base.NewBaseHashCodec(types.TypeHMAC256).(*base.BaseHashCodec),
	}
}
