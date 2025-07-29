/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package argon2

import (
	"crypto/subtle"
	"fmt"

	"github.com/goexts/generic"
	"golang.org/x/crypto/argon2"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/validator"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

type KeyFunc func(password []byte, salt []byte, time uint32, memory uint32, threads uint8, keyLen uint32) []byte

// Argon2 implements the Argon2 hashing algorithm
type Argon2 struct {
	p       types.Type
	params  *Params
	config  *types.Config
	keyFunc KeyFunc
}

var (
	argon2Type = types.Type{Name: constants.ARGON2}
	argon2i    = types.Type{Name: constants.ARGON2i}
	argon2id   = types.Type{Name: constants.ARGON2id}
)

func (c *Argon2) Type() types.Type {
	return c.p
}

// DefaultConfig returns the default configuration for Argon2
func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength:  16, // Default salt length
		ParamConfig: DefaultParams().String(),
	}
}

// ConfigValidator validates the Argon2 configuration
func ConfigValidator(config *types.Config) error {
	if config.SaltLength < 8 {
		return errors.ErrSaltLengthTooShort
	}
	return nil
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
	return types.NewHashPartsFull(c.p, hashBytes, salt, c.params.ToMap()), nil
}

// Verify implements the verify method
func (c *Argon2) Verify(parts *types.HashParts, password string) error {
	algorithm, err := types.ParseType(parts.Algorithm)
	if err != nil {
		return err
	}
	keyFunc := ParseKeyFunc(algorithm)
	if keyFunc == nil {
		return fmt.Errorf("unsupported argon2 type: %s", algorithm.String())
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

func ParseKeyFunc(p types.Type) KeyFunc {
	switch p.Name {
	case constants.ARGON2id:
		return argon2.IDKey
	case constants.ARGON2i:
		return argon2.Key
	default:
		return nil
	}
}

func NewDefaultArgon2(config *types.Config) (interfaces.Cryptographic, error) {
	return NewArgon2(argon2Type, config)
}

// NewArgon2 creates a new Argon2 crypto instance
func NewArgon2(p types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	// Use default config if provided config is nil
	if config == nil {
		config = DefaultConfig()
	}

	if config.ParamConfig == "" {
		config.ParamConfig = DefaultParams().String()
	}
	v := validator.WithParams(&Params{})
	if err := v.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid argon2 config: %v", err)
	}
	p = generic.Must(ResolveType(p))
	keyFunc := ParseKeyFunc(p)
	p.Underlying = ""
	if keyFunc == nil {
		return nil, fmt.Errorf("unsupported argon2 type: %s", p.String())
	}

	return &Argon2{
		p:       p,
		params:  v.Params(),
		keyFunc: keyFunc,
		config:  config,
	}, nil
}

func NewArgon2i(cfg *types.Config) (interfaces.Cryptographic, error) {
	return NewArgon2(argon2i, cfg)
}

func NewArgon2id(cfg *types.Config) (interfaces.Cryptographic, error) {
	return NewArgon2(argon2id, cfg)
}
