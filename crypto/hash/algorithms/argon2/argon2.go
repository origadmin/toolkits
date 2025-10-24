/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package argon2

import (
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/argon2"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/validator"
	"github.com/origadmin/toolkits/crypto/rand"
)

type KeyFunc func(password []byte, salt []byte, time uint32, memory uint32, threads uint8, keyLen uint32) []byte

// Argon2 implements the Argon2 hashing algorithm
type Argon2 struct {
	algSpec types.Spec
	params  *Params
	config  *types.Config
	keyFunc KeyFunc
}

var (
	argon2Spec = types.Spec{Name: types.ARGON2}
	argon2i    = types.Spec{Name: types.ARGON2i}
	argon2id   = types.Spec{Name: types.ARGON2id}
)

func (c *Argon2) Spec() types.Spec {
	return c.algSpec
}

// DefaultConfig returns the default configuration for Argon2
func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: types.DefaultSaltLength,
	}
}

// Hash implements the hash method
func (c *Argon2) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *Argon2) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	hashBytes := c.keyFunc(
		[]byte(password),
		salt,
		c.params.TimeCost,
		c.params.MemoryCost,
		c.params.Threads,
		c.params.KeyLength,
	)
	return types.NewHashPartsFull(c.algSpec, hashBytes, salt, c.params.ToMap()), nil
}

// Verify implements the verify method
func (c *Argon2) Verify(parts *types.HashParts, password string) error {
	// parts.Algorithm is already of type types.Spec, so no need to parse it again.
	// We can directly use parts.Algorithm for comparison and passing to ParseKeyFunc.
	keyFunc := ParseKeyFunc(parts.Algorithm)
	if keyFunc == nil {
		return fmt.Errorf("unsupported argon2 type: %s", parts.Algorithm.String())
	}

	// Parse parameters
	params, err := FromMap(parts.Params)
	if err != nil {
		return err
	}
	hash := keyFunc(
		[]byte(password),
		parts.Salt,
		params.TimeCost,
		params.MemoryCost,
		params.Threads,
		params.KeyLength,
	)

	if subtle.ConstantTimeCompare(hash, parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}

	return nil
}

func ParseKeyFunc(algSpec types.Spec) KeyFunc {
	switch algSpec.Name {
	case types.ARGON2id:
		return argon2.IDKey
	case types.ARGON2i:
		return argon2.Key
	default:
		return nil
	}
}

func NewDefaultArgon2(config *types.Config) (interfaces.Cryptographic, error) {
	return NewArgon2(argon2Spec, config)
}

// NewArgon2 creates a new Argon2 crypto instance
func NewArgon2(algSpec types.Spec, config *types.Config) (interfaces.Cryptographic, error) {
	// Ensure algorithm-specific default config is applied when caller passes nil.
	if config == nil {
		config = DefaultConfig()
	}

	v, err := validator.ValidateParams(config, DefaultParams())
	if err != nil {
		return nil, fmt.Errorf("invalid argon2 config: %v", err)
	}

	keyFunc := ParseKeyFunc(algSpec)
	algSpec.Underlying = ""
	if keyFunc == nil {
		return nil, fmt.Errorf("unsupported argon2 type: %s", algSpec.String())
	}

	return &Argon2{
		algSpec: algSpec,
		params:  v.Params,
		keyFunc: keyFunc,
		config:  v.Config,
	}, nil
}

func NewArgon2i(cfg *types.Config) (interfaces.Cryptographic, error) {
	return NewArgon2(argon2i, cfg)
}

func NewArgon2id(cfg *types.Config) (interfaces.Cryptographic, error) {
	return NewArgon2(argon2id, cfg)
}
