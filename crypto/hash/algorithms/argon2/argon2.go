/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package argon2

import (
	"crypto/subtle"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// Params represents parameters for Argon2 algorithm
type Params struct {
	TimeCost   uint32
	MemoryCost uint32
	Threads    uint8
	KeyLength  uint32
}

var (
	TypeArgon2 = types.Type{Name: constants.ARGON2}
)

// parseParams parses Argon2 parameters from string
func parseParams(params string) (result Params, err error) {
	// Handle empty string case
	if params == "" {
		return result, nil
	}

	kv, err := codecPkg.ParseParams(params)
	if err != nil {
		return result, err
	}
	// Parse time cost
	if v, ok := kv["t"]; ok {
		timeCost, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return result, fmt.Errorf("invalid time cost: %v", err)
		}
		result.TimeCost = uint32(timeCost)
	}

	// Parse memory cost
	if v, ok := kv["m"]; ok {
		memoryCost, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return result, fmt.Errorf("invalid memory cost: %v", err)
		}
		result.MemoryCost = uint32(memoryCost)
	}

	// Parse threads
	if v, ok := kv["p"]; ok {
		threads, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return result, fmt.Errorf("invalid threads: %v", err)
		}
		result.Threads = uint8(threads)
	}

	// Parse key length
	if v, ok := kv["k"]; ok {
		keyLength, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return result, fmt.Errorf("invalid key length: %v", err)
		}
		result.KeyLength = uint32(keyLength)
	}

	return result, nil
}

// String returns the string representation of parameters
func (p Params) String() string {
	var parts []string
	if p.TimeCost > 0 {
		parts = append(parts, fmt.Sprintf("t:%d", p.TimeCost))
	}
	if p.MemoryCost > 0 {
		parts = append(parts, fmt.Sprintf("m:%d", p.MemoryCost))
	}
	if p.Threads > 0 {
		parts = append(parts, fmt.Sprintf("p:%d", p.Threads))
	}
	if p.KeyLength > 0 {
		parts = append(parts, fmt.Sprintf("k:%d", p.KeyLength))
	}
	return strings.Join(parts, ",")
}

// Argon2 implements the Argon2 hashing algorithm
type Argon2 struct {
	params Params
	config *types.Config
	codec  interfaces.Codec
}

func (c *Argon2) Type() types.Type {
	return c.codec.Type()
}

// ConfigValidator implements the config validator for Argon2
type ConfigValidator struct {
	params Params
}

// Validate validates the Argon2 configuration
func (v *ConfigValidator) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return errors.ErrSaltLengthTooShort
	}
	if v.params.TimeCost < 1 {
		return fmt.Errorf("invalid time cost: %d", v.params.TimeCost)
	}
	if v.params.MemoryCost < 1 {
		return fmt.Errorf("invalid memory cost: %d", v.params.MemoryCost)
	}
	if v.params.Threads < 1 {
		return fmt.Errorf("invalid threads: %d", v.params.Threads)
	}
	if v.params.KeyLength < 4 || v.params.KeyLength > 1024 {
		return fmt.Errorf("invalid key length: %d, must be between 4 and 1024", v.params.KeyLength)
	}
	return nil
}

// DefaultConfig returns the default configuration for Argon2
func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength:  16, // Default salt length
		ParamConfig: DefaultParams().String(),
	}
}

func DefaultParams() Params {
	return Params{
		TimeCost:   constants.DefaultTimeCost,   // Default time cost
		MemoryCost: constants.DefaultMemoryCost, // Default memory cost (64MB)
		Threads:    constants.DefaultThreads,    // Default threads
		KeyLength:  32,                          // Default key length
	}
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
	params, err := parseParams(config.ParamConfig)
	if err != nil {
		return nil, fmt.Errorf("invalid argon2 param config: %v", err)
	}

	validator := &ConfigValidator{
		params: params,
	}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid argon2 config: %v", err)
	}

	// Determine Argon2 variant based on p.Name
	switch p.Name {
	case constants.ARGON2, constants.ARGON2id:
		// Default to Argon2id if generic Argon2 or Argon2id is requested
		// The argon2.IDKey function handles both Argon2id and Argon2i based on the version.
		// For simplicity, we'll use IDKey for both, assuming the underlying library handles the distinction.
	case constants.ARGON2i:
		// For Argon2i, we'd ideally use argon2.Key if there was a separate function.
		// Since argon2.IDKey is the primary, we proceed with it.
	default:
		return nil, fmt.Errorf("unsupported argon2 type: %s", p.Name)
	}

	return &Argon2{
		params: params,
		config: config,
		codec:  codecPkg.NewCodec(p),
	}, nil
}

// Hash implements the hash method
func (c *Argon2) Hash(password string) (string, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, string(salt))
}

// HashWithSalt implements the hash with salt method
func (c *Argon2) HashWithSalt(password, salt string) (string, error) {
	hash := argon2.IDKey(
		[]byte(password),
		[]byte(salt),
		c.params.TimeCost,
		c.params.MemoryCost,
		c.params.Threads,
		c.params.KeyLength,
	)

	return c.codec.Encode([]byte(salt), hash, c.params.String()), nil
}

// Verify implements the verify method
func (c *Argon2) Verify(parts *types.HashParts, password string) error {
	if !parts.Algorithm.Is(TypeArgon2) {
		return errors.ErrAlgorithmMismatch
	}
	// Parse parameters
	params, err := parseParams(parts.Params)
	if err != nil {
		return err
	}
	hash := argon2.IDKey(
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
