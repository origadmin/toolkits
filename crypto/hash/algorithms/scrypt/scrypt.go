/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package scrypt

import (
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/scrypt"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// Scrypt implements the Scrypt hashing algorithm
type Scrypt struct {
	p      types.Type
	params *Params
	config *types.Config
}

func (c *Scrypt) Type() types.Type {
	return c.p
}

type ConfigValidator struct {
	params *Params
}

func (v ConfigValidator) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return errors.ErrSaltLengthTooShort
	}
	// N must be > 1 and a power of 2
	if v.params.N <= 1 || v.params.N&(v.params.N-1) != 0 {
		return fmt.Errorf("N must be > 1 and a power of 2")
	}
	if v.params.R <= 0 {
		return fmt.Errorf("R must be > 0")
	}
	if v.params.P <= 0 {
		return fmt.Errorf("P must be > 0")
	}
	if v.params.KeyLen <= 0 {
		return fmt.Errorf("key length must be > 0")
	}

	return nil
}

// NewScrypt creates a new Scrypt crypto instance
func NewScrypt(config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}

	if config.ParamConfig == "" {
		config.ParamConfig = DefaultParams().String()
	}
	params, err := parseParams(config.ParamConfig)
	if err != nil {
		return nil, fmt.Errorf("invalid scrypt param config: %v", err)
	}

	validator := &ConfigValidator{
		params: params,
	}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid scrypt config: %v", err)
	}
	return &Scrypt{
		p:      types.Type{Name: constants.SCRYPT},
		params: params,
		config: config,
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength:  16,
		ParamConfig: DefaultParams().String(),
	}
}

// Hash implements the hash method
func (c *Scrypt) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *Scrypt) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	hash, err := scrypt.Key([]byte(password), salt, c.params.N, c.params.R, c.params.P, c.params.KeyLen)
	if err != nil {
		return nil, err
	}
	return types.NewHashPartsWithParams(c.Type(), salt, hash, c.params.ToMap()), nil
}

// Verify implements the verify method
func (c *Scrypt) Verify(parts *types.HashParts, password string) error {
	if parts.Algorithm.Name != constants.SCRYPT {
		return errors.ErrAlgorithmMismatch
	}
	// Parse parameters
	params, err := parseParams(parts.Params)
	if err != nil {
		return err
	}
	hash, err := scrypt.Key([]byte(password), parts.Salt, params.N, params.R, params.P, params.KeyLen)
	if err != nil {
		return err
	}
	if subtle.ConstantTimeCompare(hash, parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}

	return nil
}
