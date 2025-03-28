/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package sha256

import (
	"crypto/sha256"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// Sha256 implements the Sha256 hashing algorithm
type Sha256 struct {
	config *types.Config
	codec  interfaces.Codec
}

func (c *Sha256) Type() string {
	return types.TypeSha256.String()
}

// NewSHA256Crypto creates a new Sha256 crypto instance
func NewSHA256Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return &Sha256{
		config: config,
		codec:  core.NewCodec(types.TypeSha256),
	}, nil
}

// Hash implements the hash method
func (c *Sha256) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, string(salt))
}

// HashWithSalt implements the hash with salt method
func (c *Sha256) HashWithSalt(password, salt string) (string, error) {
	hash := sha256.Sum256([]byte(password + salt))
	return c.codec.Encode([]byte(salt), hash[:]), nil
}

// Verify implements the verify method
func (c *Sha256) Verify(hashed, password string) error {
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}

	if parts.Algorithm != types.TypeSha256 {
		return fmt.Errorf("algorithm mismatch")
	}

	newHash := sha256.Sum256([]byte(password + string(parts.Salt)))
	if string(newHash[:]) != string(parts.Hash) {
		return fmt.Errorf("password not match")
	}

	return nil
}
