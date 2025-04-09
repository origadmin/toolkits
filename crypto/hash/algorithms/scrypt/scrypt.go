/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package scrypt

import (
	"crypto/subtle"
	"fmt"
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
	params *Params
	config *types.Config
	codec  interfaces.Codec
}

func (c *Scrypt) Type() string {
	return types.TypeScrypt.String()
}

type ConfigValidator struct {
	params *Params
}

func (v ConfigValidator) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return fmt.Errorf("salt length must be at least 8 bytes")
	}
	// N must be > 1 and a power of 2
	if v.params.N <= 1 || v.params.N&(v.params.N-1) != 0 {
		return fmt.Errorf("N must be > 1 and a power of 2")
	}

	return nil
}

// NewScryptCrypto creates a new Scrypt crypto instance
func NewScryptCrypto(config *types.Config) (interfaces.Cryptographic, error) {
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
		params: params,
		config: config,
		codec:  core.NewCodec(types.TypeScrypt),
	}, nil
}

func DefaultParams() *Params {
	return &Params{
		N:      16384,
		R:      8,
		P:      1,
		KeyLen: 32,
	}
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength:  16,
		ParamConfig: DefaultParams().String(),
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
	salt, err := utils.GenerateSaltString(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *Scrypt) HashWithSalt(password, salt string) (string, error) {
	hash, err := scrypt.Key([]byte(password), []byte(salt), c.params.N, c.params.R, c.params.P, c.params.KeyLen)
	if err != nil {
		return "", err
	}
	return c.codec.Encode([]byte(salt), hash, c.params.String()), nil
}

// Verify implements the verify method
func (c *Scrypt) Verify(parts *types.HashParts, password string) error {
	if parts.Algorithm != types.TypeScrypt {
		return core.ErrAlgorithmMismatch
	}
	// Parse parameters
	params, err := parseParams(parts.Params)
	if err != nil {
		return err
	}
	hash, err := scrypt.Key([]byte(password), []byte(parts.Salt), params.N, params.R, params.P, params.KeyLen)
	if err != nil {
		return err
	}
	if subtle.ConstantTimeCompare(hash, parts.Hash) != 1 {
		return core.ErrPasswordNotMatch
	}

	return nil
}
