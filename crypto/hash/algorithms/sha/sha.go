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
	"github.com/origadmin/toolkits/crypto/hash/internal/validator"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// SHA implements the SHA hashing algorithm
type SHA struct {
	p        types.Type
	config   *types.Config
	hashHash stdhash.Hash
}

var (
	sha1Type       = types.NewType(constants.SHA1)
	sha256Type     = types.NewType(constants.SHA256)
	sha224Type     = types.NewType(constants.SHA224)
	sha384Type     = types.NewType(constants.SHA384)
	sha512Type     = types.NewType(constants.SHA512)
	sha512_224Type = types.NewType(constants.SHA512_224)
	sha512_256Type = types.NewType(constants.SHA512_256)
	sha3_224Type   = types.NewType(constants.SHA3_224)
	sha3_256Type   = types.NewType(constants.SHA3_256)
	sha3_384Type   = types.NewType(constants.SHA3_384)
	sha3_512Type   = types.NewType(constants.SHA3_512)
)

func (c *SHA) Type() types.Type {
	return c.p
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
	hashBytes := newHash.Sum(nil)
	return types.NewHashPartsWithHashSalt(c.p, hashBytes, salt), nil
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

// NewSHA creates a new SHA crypto instance
func NewSHA(p types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}

	v := validator.WithParams(&Params{})
	if err := v.Validate(config); err != nil {
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
	return NewSHA(sha1Type, config)
}

func NewSha224(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha224Type, config)
}

func NewSha256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha256Type, config)
}

func NewSha512(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha512Type, config)
}

func NewSha3224(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha3_224Type, config)
}

func NewSha3256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha3_256Type, config)
}

func NewSha3384(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha3_384Type, config)
}

func NewSha384(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha384Type, config)
}

func NewSha3512(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha3_512Type, config)
}
func NewSha3512224(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha512_224Type, config)
}

func NewSha3512256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha512_256Type, config)
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}
