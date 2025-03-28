/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package argon2

import (
	"crypto/subtle"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// Params represents parameters for Argon2 algorithm
type Params struct {
	TimeCost   uint32
	MemoryCost uint32
	Threads    uint8
	KeyLength  uint32
}

// parseParams parses Argon2 parameters from string
func parseParams(params string) (*Params, error) {
	result := &Params{}

	// Handle empty string case
	if params == "" {
		return result, nil
	}

	kv := make(map[string]string)
	for _, param := range strings.Split(params, ",") {
		parts := strings.Split(param, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid argon2 param format: %s", param)
		}
		kv[parts[0]] = parts[1]
	}

	// Parse time cost
	if v, ok := kv["t"]; ok {
		timeCost, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid time cost: %v", err)
		}
		result.TimeCost = uint32(timeCost)
	}

	// Parse memory cost
	if v, ok := kv["m"]; ok {
		memoryCost, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid memory cost: %v", err)
		}
		result.MemoryCost = uint32(memoryCost)
	}

	// Parse threads
	if v, ok := kv["p"]; ok {
		threads, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return nil, fmt.Errorf("invalid threads: %v", err)
		}
		result.Threads = uint8(threads)
	}

	// Parse key length
	if v, ok := kv["k"]; ok {
		keyLength, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return nil, fmt.Errorf("invalid key length: %v", err)
		}
		result.KeyLength = uint32(keyLength)
	}

	return result, nil
}

// String returns the string representation of parameters
func (p *Params) String() string {
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
	params *Params
	config *types.Config
	codec  interfaces.Codec
}

// ConfigValidator implements the config validator for Argon2
type ConfigValidator struct{}

// Validate validates the Argon2 configuration
func (v *ConfigValidator) Validate(config *types.Config) error {
	if config.TimeCost < 1 {
		return fmt.Errorf("invalid time cost: %d", config.TimeCost)
	}
	if config.MemoryCost < 1 {
		return fmt.Errorf("invalid memory cost: %d", config.MemoryCost)
	}
	if config.Threads < 1 {
		return fmt.Errorf("invalid threads: %d", config.Threads)
	}
	if config.SaltLength < 8 {
		return core.ErrSaltLengthTooShort
	}
	if config.KeyLength < 4 || config.KeyLength > 1024 {
		return fmt.Errorf("invalid key length: %d, must be between 4 and 1024", config.KeyLength)
	}
	return nil
}

// DefaultConfig returns the default configuration for Argon2
func DefaultConfig() *types.Config {
	return &types.Config{
		TimeCost:   3,         // Default time cost
		MemoryCost: 64 * 1024, // Default memory cost (64MB)
		Threads:    4,         // Default threads
		SaltLength: 16,        // Default salt length
		KeyLength:  32,        // Default key length
	}
}

// NewArgon2Crypto creates a new Argon2 crypto instance
func NewArgon2Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	// Use default config if provided config is nil
	if config == nil {
		config = DefaultConfig()
	}
	validator := &ConfigValidator{}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid argon2 config: %v", err)
	}
	params := &Params{
		TimeCost:   config.TimeCost,
		MemoryCost: config.MemoryCost,
		Threads:    config.Threads,
		KeyLength:  config.KeyLength,
	}
	return &Argon2{
		params: params,
		config: config,
		codec:  core.NewCodec(types.TypeArgon2),
	}, nil
}

// Hash implements the hash method
func (c *Argon2) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, salt)
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
func (c *Argon2) Verify(hashed, password string) error {
	slog.Info("Verify", "hashed", hashed, "password", password)
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}
	slog.Info("Verify", "parts", parts)
	if parts.Algorithm != types.TypeArgon2 {
		return fmt.Errorf("algorithm mismatch")
	}
	slog.Info("Verify", "parts.Params", parts.Params)
	// Parse parameters
	params, err := parseParams(parts.Params)
	if err != nil {
		return err
	}
	slog.Info("Verify", "params", params)
	hash := argon2.IDKey(
		[]byte(password),
		parts.Salt,
		params.TimeCost,
		params.MemoryCost,
		params.Threads,
		params.KeyLength,
	)

	if subtle.ConstantTimeCompare(hash, parts.Hash) != 1 {
		return core.ErrPasswordNotMatch
	}

	return nil
}
