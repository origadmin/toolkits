/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package sha

import (
	"crypto/subtle"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/validator"
	"github.com/origadmin/toolkits/crypto/rand"
)

// SHA implements the SHA hashing algorithm
type SHA struct {
	algSpec  types.Spec
	config   *types.Config
	hashHash stdhash.Hash
}

var (
	sha1AlgSpec       = types.New(types.SHA1)
	sha256AlgSpec     = types.New(types.SHA256)
	sha224AlgSpec     = types.New(types.SHA224)
	sha384AlgSpec     = types.New(types.SHA384)
	sha512AlgSpec     = types.New(types.SHA512)
	sha512_224AlgSpec = types.New(types.SHA512_224)
	sha512_256AlgSpec = types.New(types.SHA512_256)
	sha3_224AlgSpec   = types.New(types.SHA3_224)
	sha3_256AlgSpec   = types.New(types.SHA3_256)
	sha3_384AlgSpec   = types.New(types.SHA3_384)
	sha3_512AlgSpec   = types.New(types.SHA3_512)
)

func (c *SHA) Spec() types.Spec {
	return c.algSpec
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
	if len(salt) > 0 {
		newHash.Write(salt)
	}
	hashBytes := newHash.Sum(nil)
	return types.NewHashPartsWithHashSalt(c.algSpec, hashBytes, salt), nil
}

// Verify implements the verify method
func (c *SHA) Verify(parts *types.HashParts, password string) error {
	// parts.Algorithm is already of type types.Spec. Use its Name field for stdhash.ParseHash.
	hashHash, err := stdhash.ParseHash(parts.Algorithm.Name)
	if err != nil {
		return err
	}
	newHash := hashHash.New()
	newHash.Write([]byte(password))
	if len(parts.Salt) > 0 {
		newHash.Write(parts.Salt)
	}
	if subtle.ConstantTimeCompare(newHash.Sum(nil), parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}
	return nil
}

// NewSHA creates a new SHA crypto instance
func NewSHA(algSpec types.Spec, config *types.Config) (scheme.Scheme, error) {
	// Ensure algorithm-specific default config is applied when caller passes nil.
	if config == nil {
		config = DefaultConfig()
	}

	v, err := validator.ValidateParams(config, DefaultParams())
	if err != nil {
		return nil, fmt.Errorf("invalid sha config: %v", err)
	}

	hashHash, err := stdhash.ParseHash(algSpec.Name)
	if err != nil {
		return nil, err
	}

	return &SHA{
		algSpec:  algSpec,
		config:   v.Config,
		hashHash: hashHash,
	}, nil
}

func NewSha1(config *types.Config) (scheme.Scheme, error) {
	return NewSHA(sha1AlgSpec, config)
}

func NewSha224(config *types.Config) (scheme.Scheme, error) {
	return NewSHA(sha224AlgSpec, config)
}

func NewSha256(config *types.Config) (scheme.Scheme, error) {
	return NewSHA(sha256AlgSpec, config)
}

func NewSha512(config *types.Config) (scheme.Scheme, error) {
	return NewSHA(sha512AlgSpec, config)
}

func NewSha3224(config *types.Config) (scheme.Scheme, error) {
	return NewSHA(sha3_224AlgSpec, config)
}

func NewSha3256(config *types.Config) (scheme.Scheme, error) {
	return NewSHA(sha3_256AlgSpec, config)
}

func NewSha3384(config *types.Config) (scheme.Scheme, error) {
	return NewSHA(sha3_384AlgSpec, config)
}

func NewSha384(config *types.Config) (scheme.Scheme, error) {
	return NewSHA(sha384AlgSpec, config)
}

func NewSha3512(config *types.Config) (scheme.Scheme, error) {
	return NewSHA(sha3_512AlgSpec, config)
}
func NewSha3512224(config *types.Config) (scheme.Scheme, error) {
	return NewSHA(sha512_224AlgSpec, config)
}

func NewSha3512256(config *types.Config) (scheme.Scheme, error) {
	return NewSHA(sha512_256AlgSpec, config)
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}
