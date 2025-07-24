/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package md5

import (
	"crypto/md5"
	"crypto/subtle"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// MD5 implements the MD5 hashing algorithm
type MD5 struct {
	config *types.Config
	codec  interfaces.Codec
}

func (c *MD5) Type() string {
	return types.TypeMD5.String()
}

type ConfigValidator struct {
	SaltLength int
}

func (v ConfigValidator) Validate(config *types.Config) interface{} {
	if config.SaltLength < 8 {
		return fmt.Errorf("salt length must be at least 8 bytes")
	}
	return nil
}

// NewMD5Crypto creates a new MD5 crypto instance
func NewMD5Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = types.DefaultConfig()
	}
	validator := &ConfigValidator{}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid md5 config: %v", err)
	}
	return &MD5{
		config: config,
		codec:  codec.NewCodec(types.TypeMD5),
	}, nil
}

// Hash implements the hash method
func (c *MD5) Hash(password string) (string, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, string(salt))
}

// HashWithSalt implements the hash with salt method
func (c *MD5) HashWithSalt(password, salt string) (string, error) {
	hash := md5.Sum([]byte(password + salt))
	return c.codec.Encode([]byte(salt), hash[:]), nil
}

// Verify implements the verify method
func (c *MD5) Verify(parts *types.HashParts, password string) error {
	if parts.Algorithm != types.TypeMD5 {
		return errors.ErrAlgorithmMismatch
	}

	newHash := md5.Sum([]byte(password + string(parts.Salt)))
	if subtle.ConstantTimeCompare(newHash[:], parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}
	return nil
}
