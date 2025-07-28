/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package pbkdf2

import (
	"crypto/subtle"
	"fmt"
	"hash"

	"golang.org/x/crypto/pbkdf2"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// PBKDF2 implements the PBKDF2 hashing algorithm
type PBKDF2 struct {
	params *Params
	config *types.Config
	codec  interfaces.Codec
}

func (c *PBKDF2) Type() types.Type {
	return c.codec.Type()
}

type ConfigValidator struct {
	params *Params
}

func (v ConfigValidator) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return errors.ErrSaltLengthTooShort
	}
	if v.params.Iterations < 1000 {
		return fmt.Errorf("iterations must be at least 1000")
	}
	if v.params.KeyLength < 8 {
		return fmt.Errorf("key length must be at least 8 bytes")
	}
	if v.params.HashType == "" {
		return fmt.Errorf("hash type must be specified")
	}
	_, err := stdhash.ParseHash(v.params.HashType)
	if err != nil {
		return err
	}
	return nil
}

// NewPBKDF2 creates a new PBKDF2 crypto instance
func NewPBKDF2(p types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}

	// If a specific underlying hash is requested (e.g., pbkdf2-sha256),
	// ensure it's used in the params if not already set.
	if p.Underlying != "" {
		defaultParams := DefaultParams()
		if defaultParams.HashType == "" || defaultParams.HashType == stdhash.SHA256.String() { // Default to SHA256 if not specified
			defaultParams.HashType = p.Underlying
		}
		if config.ParamConfig == "" {
			config.ParamConfig = defaultParams.String()
		}
	}

	params, err := parseParams(config.ParamConfig)
	if err != nil {
		return nil, fmt.Errorf("invalid pbkdf2 param config: %v", err)
	}

	validator := &ConfigValidator{
		params: params,
	}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid pbkdf2 config: %v", err)
	}
	return &PBKDF2{
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
func (c *PBKDF2) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

func (c *PBKDF2) hashFromName(name string) (func() hash.Hash, error) {
	parseHash, err := stdhash.ParseHash(name)
	if err != nil {
		return nil, err
	}
	return parseHash.New, nil
}

// HashWithSalt implements the hash with salt method
func (c *PBKDF2) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	hashHash, err := c.hashFromName(c.params.HashType)
	if err != nil {
		return nil, err
	}
	newHash := pbkdf2.Key([]byte(password), salt, c.params.Iterations, int(c.params.KeyLength), hashHash)
	return types.NewHashPartsWithParams(c.Type(), salt, newHash, c.params.ToMap()), nil
}

// Verify implements the verify method
func (c *PBKDF2) Verify(parts *types.HashParts, password string) error {
	if !parts.Algorithm.Is(c.Type()) {
		return errors.ErrAlgorithmMismatch
	}

	// Parse parameters
	params, err := parseParams(parts.Params)
	if err != nil {
		return err
	}

	// The hash function is recreated based on the hash type being parsed
	hashHash, err := c.hashFromName(params.HashType)
	if err != nil {
		return err
	}

	newHash := pbkdf2.Key([]byte(password), parts.Salt, params.Iterations, int(params.KeyLength), hashHash)
	if subtle.ConstantTimeCompare(newHash, parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}
	return nil
}
