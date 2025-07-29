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
	algType  types.Type
	config   *types.Config
	hashHash stdhash.Hash
}

var (
	sha1AlgType       = types.NewType(constants.SHA1)
	sha256AlgType     = types.NewType(constants.SHA256)
	sha224AlgType     = types.NewType(constants.SHA224)
	sha384AlgType     = types.NewType(constants.SHA384)
	sha512AlgType     = types.NewType(constants.SHA512)
	sha512_224AlgType = types.NewType(constants.SHA512_224)
	sha512_256AlgType = types.NewType(constants.SHA512_256)
	sha3_224AlgType   = types.NewType(constants.SHA3_224)
	sha3_256AlgType   = types.NewType(constants.SHA3_256)
	sha3_384AlgType   = types.NewType(constants.SHA3_384)
	sha3_512AlgType   = types.NewType(constants.SHA3_512)
)

func (c *SHA) Type() types.Type {
	return c.algType
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
	return types.NewHashPartsWithHashSalt(c.algType, hashBytes, salt), nil
}

// Verify implements the verify method
func (c *SHA) Verify(parts *types.HashParts, password string) error {
	algType, err := types.ParseType(parts.Algorithm)
	if err != nil {
		return err
	}
	hashHash, err := stdhash.ParseHash(algType.Name)
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
func NewSHA(algType types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}

	v := validator.WithParams(&Params{})
	if err := v.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid sha config: %v", err)
	}
	hashHash, err := stdhash.ParseHash(algType.Name)
	if err != nil {
		return nil, err
	}

	return &SHA{
		algType:  algType,
		config:   config,
		hashHash: hashHash,
	}, nil
}

func NewSha1(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha1AlgType, config)
}

func NewSha224(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha224AlgType, config)
}

func NewSha256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha256AlgType, config)
}

func NewSha512(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha512AlgType, config)
}

func NewSha3224(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha3_224AlgType, config)
}

func NewSha3256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha3_256AlgType, config)
}

func NewSha3384(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha3_384AlgType, config)
}

func NewSha384(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha384AlgType, config)
}

func NewSha3512(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha3_512AlgType, config)
}
func NewSha3512224(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha512_224AlgType, config)
}

func NewSha3512256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewSHA(sha512_256AlgType, config)
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}