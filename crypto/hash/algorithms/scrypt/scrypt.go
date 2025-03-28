/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package scrypt

import (
	"crypto/subtle"
	"fmt"
	"log/slog"
	"strconv"
	"strings"

	"golang.org/x/crypto/scrypt"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// Scrypt implements the Scrypt hashing algorithm
type Scrypt struct {
	config *types.Config
	codec  interfaces.Codec
}

type ConfigValidator struct {
}

func (v ConfigValidator) Validate(config *types.Config) interface{} {
	if config.SaltLength < 8 {
		return fmt.Errorf("salt length must be at least 8 bytes")
	}
	// N must be > 1 and a power of 2
	if config.Scrypt.N <= 1 || config.Scrypt.N&(config.Scrypt.N-1) != 0 {
		return fmt.Errorf("N must be > 1 and a power of 2")
	}

	return nil
}

// NewScryptCrypto creates a new Scrypt crypto instance
func NewScryptCrypto(config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	validator := &ConfigValidator{}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid scrypt config: %v", err)
	}
	return &Scrypt{
		config: config,
		codec:  core.NewCodec(types.TypeScrypt),
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
		KeyLength:  32,
		Scrypt: types.ScryptConfig{
			N: 16384,
			R: 8,
			P: 1,
		},
	}
}

// Params represents parameters for Scrypt algorithm
type Params struct {
	N      int
	R      int
	P      int
	KeyLen int
}

// parseParams parses Scrypt parameters from string
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
			return nil, fmt.Errorf("invalid scrypt param format: %s", param)
		}
		kv[parts[0]] = parts[1]
	}

	// Parse N
	if v, ok := kv["n"]; ok {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid N: %v", err)
		}
		result.N = n
	}

	// Parse R
	if v, ok := kv["r"]; ok {
		r, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid R: %v", err)
		}
		result.R = r
	}

	// Parse P
	if v, ok := kv["p"]; ok {
		p, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid P: %v", err)
		}
		result.P = p
	}

	// Parse KeyLen
	if v, ok := kv["k"]; ok {
		keyLen, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid KeyLen: %v", err)
		}
		result.KeyLen = keyLen
	}

	return result, nil
}

// String returns the string representation of parameters
func (p *Params) String() string {
	var parts []string
	if p.N > 0 {
		parts = append(parts, fmt.Sprintf("n:%d", p.N))
	}
	if p.R > 0 {
		parts = append(parts, fmt.Sprintf("r:%d", p.R))
	}
	if p.P > 0 {
		parts = append(parts, fmt.Sprintf("p:%d", p.P))
	}
	if p.KeyLen > 0 {
		parts = append(parts, fmt.Sprintf("k:%d", p.KeyLen))
	}
	return strings.Join(parts, ",")
}

// Hash implements the hash method
func (c *Scrypt) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *Scrypt) HashWithSalt(password, salt string) (string, error) {
	params := &Params{
		N:      c.config.Scrypt.N,
		R:      c.config.Scrypt.R,
		P:      c.config.Scrypt.P,
		KeyLen: int(c.config.KeyLength),
	}
	hash, err := scrypt.Key([]byte(password), []byte(salt), params.N, params.R, params.P, params.KeyLen)
	if err != nil {
		return "", err
	}
	return c.codec.Encode([]byte(salt), hash, params.String()), nil
}

// Verify implements the verify method
func (c *Scrypt) Verify(hashed, password string) error {
	slog.Info("Verify", "hashed", hashed, "password", password)
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}
	slog.Info("Verify", "parts", parts)
	if parts.Algorithm != types.TypeScrypt {
		return core.ErrAlgorithmMismatch
	}
	// Parse parameters
	slog.Info("Verify", "params", parts.Params)
	params, err := parseParams(parts.Params)
	if err != nil {
		return err
	}
	slog.Info("Verify", "params", params)
	hash, err := scrypt.Key([]byte(password), []byte(parts.Salt), params.N, params.R, params.P, params.KeyLen)
	if err != nil {
		return err
	}
	slog.Info("Verify", "hash", hash)
	if subtle.ConstantTimeCompare(hash, parts.Hash) != 1 {
		return core.ErrPasswordNotMatch
	}

	return nil
}
