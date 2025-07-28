/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package argon2

import (
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/argon2"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/validator"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

var (
	TypeArgon2 = types.Type{
		Name: constants.ARGON2,
	}
)

// Argon2 implements the Argon2 hashing algorithm
type Argon2 struct {
	p       types.Type
	params  *params
	config  *types.Config
	keyFunc func(password []byte, salt []byte, time uint32, memory uint32, threads uint8, keyLen uint32) []byte
}

func (c *Argon2) Salt() ([]byte, error) {
	return rand.RandomBytes(c.config.SaltLength)
}

func (c *Argon2) Type() types.Type {
	return TypeArgon2
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

func NewDefaultArgon2(config *types.Config) (interfaces.Cryptographic, error) {
	return NewArgon2(TypeArgon2, config)
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
	//paramsMap, err := codecPkg.DecodeParams(config.ParamConfig)
	//if err != nil {
	//	return nil, fmt.Errorf("invalid argon2 param config: %v", err)
	//}
	//params, err := parseParams(paramsMap)
	//if err != nil {
	//	return nil, fmt.Errorf("invalid argon2 param config: %v", err)
	//}
	v := validator.WithParams(&params{})
	if err := v.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid argon2 config: %v", err)
	}
	keyFunc := KeyFunc(p)
	p.Underlying = ""
	if keyFunc == nil {
		return nil, fmt.Errorf("unsupported argon2 type: %s", p.Name)
	}

	return &Argon2{
		p:       p,
		params:  v.Params(),
		keyFunc: keyFunc,
		config:  config,
	}, nil
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
	hash := c.keyFunc(
		[]byte(password),
		salt,
		c.params.TimeCost,
		c.params.MemoryCost,
		c.params.Threads,
		c.params.KeyLength,
	)
	return &types.HashParts{
		Algorithm: c.p.String(),
		Params:    c.params.ToMap(),
		Hash:      hash,
		Salt:      salt,
	}, nil
}

// Verify implements the verify method
func (c *Argon2) Verify(parts *types.HashParts, password string) error {
	if parts.Algorithm != c.p.Name {
		return errors.ErrAlgorithmMismatch
	}
	algorithm, err := types.ParseAlgorithm(parts.Algorithm)
	if err != nil {
		return err
	}
	keyFunc := KeyFunc(algorithm)
	if keyFunc == nil {
		return errors.ErrAlgorithmMismatch
	}
	// Parse parameters
	params, err := parseParams(parts.Params)
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

func KeyFunc(p types.Type) func(password []byte, salt []byte, time uint32, memory uint32, threads uint8, keyLen uint32) []byte {
	switch p.Name {
	case constants.ARGON2, constants.ARGON2id:
		return argon2.IDKey
	case constants.ARGON2i:
		return argon2.Key
	default:
		return nil
	}
}
