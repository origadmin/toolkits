/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package pbkdf2

import (
	"crypto/subtle"
	"fmt"
	"hash"
	"strconv"
	"strings"

	"golang.org/x/crypto/pbkdf2"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// PBKDF2 implements the PBKDF2 hashing algorithm
type PBKDF2 struct {
	params *Params
	config *types.Config
	codec  interfaces.Codec
}

func (c *PBKDF2) Type() string {
	return types.TypePBKDF2.String()
}

type ConfigValidator struct {
	params *Params
}

func (v ConfigValidator) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return fmt.Errorf("salt length must be at least 8 bytes")
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
	_, err := core.ParseHash(v.params.HashType)
	if err != nil {
		return err
	}
	return nil
}

// NewPBKDF2Crypto creates a new PBKDF2 crypto instance
func NewPBKDF2Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}

	if config.ParamConfig == "" {
		config.ParamConfig = DefaultParams().String()
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
		codec:  core.NewCodec(types.TypePBKDF2),
	}, nil
}

func DefaultParams() *Params {
	return &Params{
		Iterations: 10000,
		KeyLength:  32,
		HashType:   core.SHA256.String(),
	}
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength:  16,
		ParamConfig: DefaultParams().String(),
	}
}

// Hash implements the hash method
func (c *PBKDF2) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, string(salt))
}

func (c *PBKDF2) hashFromName(name string) (func() hash.Hash, error) {
	parseHash, err := core.ParseHash(name)
	if err != nil {
		return nil, err
	}
	return parseHash.New, nil
}

// HashWithSalt implements the hash with salt method
func (c *PBKDF2) HashWithSalt(password, salt string) (string, error) {
	hashHash, err := c.hashFromName(c.params.HashType)
	if err != nil {
		return "", err
	}
	newHash := pbkdf2.Key([]byte(password), []byte(salt), c.params.Iterations, int(c.params.KeyLength), hashHash)
	return c.codec.Encode([]byte(salt), newHash, c.params.String()), nil
}

// Verify implements the verify method
func (c *PBKDF2) Verify(parts *types.HashParts, password string) error {
	if !parts.Algorithm.Is(types.TypePBKDF2) {
		return core.ErrAlgorithmMismatch
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
		return core.ErrPasswordNotMatch
	}
	return nil
}

// Params represents parameters for PBKDF2 algorithm
type Params struct {
	Iterations int
	KeyLength  uint32
	HashType   string
}

// String returns the string representation of parameters
func (p *Params) String() string {
	var parts []string
	if p.Iterations > 0 {
		parts = append(parts, fmt.Sprintf("i:%d", p.Iterations))
	}
	if p.KeyLength > 0 {
		parts = append(parts, fmt.Sprintf("k:%d", p.KeyLength))
	}
	_, err := core.ParseHash(p.HashType)
	if err == nil {
		parts = append(parts, fmt.Sprintf("h:%s", p.HashType))
	}
	return strings.Join(parts, ",")
}

// parseParams parses PBKDF2 parameters from string
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
			return nil, fmt.Errorf("invalid pbkdf2 param format: %s", param)
		}
		kv[parts[0]] = parts[1]
	}

	// Parse iterations
	if v, ok := kv["i"]; ok {
		iterations, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid iterations: %v", err)
		}
		result.Iterations = iterations
	}

	// Parse key length
	if v, ok := kv["k"]; ok {
		keyLength, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid key length: %v", err)
		}
		result.KeyLength = uint32(keyLength)
	}

	// Parse hash type
	if v, ok := kv["h"]; ok {
		_, err := core.ParseHash(v)
		if err == nil {
			result.HashType = v
		}
	}

	return result, nil
}
