/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package hmac256

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// HMAC256 implements the HMAC-SHA256 hashing algorithm
type HMAC256 struct {
	config *types.Config
	codec  interfaces.Codec
}

func (c *HMAC256) Type() string {
	return types.TypeHMAC256.String()
}

type ConfigValidator struct {
	SaltLength int
}

func (v ConfigValidator) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return fmt.Errorf("salt length must be at least 8 bytes")
	}
	return nil
}

// NewHMAC256Crypto creates a new HMAC256 crypto instance
func NewHMAC256Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	validator := &ConfigValidator{}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid hmac256 config: %v", err)
	}
	return &HMAC256{
		config: config,
		codec:  core.NewCodec(types.TypeHMAC256),
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}

// Hash implements the hash method
func (c *HMAC256) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, string(salt))
}

// HashWithSalt implements the hash with salt method
func (c *HMAC256) HashWithSalt(password, salt string) (string, error) {
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(password))
	hash := h.Sum(nil)
	return c.codec.Encode([]byte(salt), hash, ""), nil
}

// Verify implements the verify method
func (c *HMAC256) Verify(hashed, password string) error {
	parts, err := c.codec.Decode(hashed)
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
