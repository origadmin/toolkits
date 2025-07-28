/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package sha

import (
	"crypto/subtle"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// SHA implements the SHA hashing algorithm
type SHA struct {
	p        types.Type
	config   *types.Config
	hashHash stdhash.Hash
}

func (c *SHA) Type() types.Type {
	return c.p
}

type ConfigValidator struct {
}

func (v ConfigValidator) Validate(config *types.Config) interface{} {
	if config.SaltLength < 8 {
		return fmt.Errorf("salt length must be at least 8 bytes")
	}
	return nil
}

// NewSHA creates a new SHA crypto instance
func NewSHA(p types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	validator := &ConfigValidator{}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid sha config: %v", err)
	}
	hashHash, err := stdhash.ParseHash(p.Name)
	if err != nil {
		return nil, err
	}

	return &SHA{
		p:        p,
		config:   config,
		hashHash: hashHash,
	}, nil
}

func NewSha1(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(types.Type{
		Name: constants.SHA1,
	}, config)
}

func NewSha224(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(types.Type{
		Name: constants.SHA224,
	}, config)
}

func NewSha256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(types.Type{
		Name: constants.SHA256,
	}, config)
}

func NewSha512(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(types.Type{
		Name: constants.SHA512,
	}, config)
}

func NewSha3224(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(types.Type{
		Name: constants.SHA3_224,
	}, config)
}

func NewSha3256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(types.Type{
		Name: constants.SHA3_256,
	}, config)
}

func NewSha384(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(types.Type{
		Name: constants.SHA384,
	}, config)
}

func NewSha3512(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(types.Type{
		Name: constants.SHA3_512,
	}, config)
}
func NewSha3512224(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(types.Type{
		Name: constants.SHA3_224,
	}, config)
}

func NewSha3512256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(types.Type{
		Name: constants.SHA3_512_256,
	}, config)
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}

// Hash implements the hash method
func (c *SHA) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *SHA) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	newHash := c.hashHash.New()
	newHash.Write([]byte(password))
	newHash.Write(salt)
	return types.NewHashParts(c.p, salt, newHash.Sum(nil)), nil
}

// Verify implements the verify method
func (c *SHA) Verify(parts *types.HashParts, password string) error {
	hashHash, err := stdhash.ParseHash(parts.Algorithm)
	if err != nil {
		return err
	}
	newHash := hashHash.New()
	newHash.Write([]byte(password))
	newHash.Write(parts.Salt)
	if subtle.ConstantTimeCompare(newHash.Sum(nil), parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}
	return nil
}
