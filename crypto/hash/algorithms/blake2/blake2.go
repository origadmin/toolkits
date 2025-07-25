/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package blake2

import (
	"crypto/subtle"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"hash"
	"strings"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

const (
	MinKeyLength = 16
	MaxKeyLength = 64
)

type hashFunc func(key []byte) (hash.Hash, error)

var hashFuncs = map[string]hashFunc{
	constants.BLAKE2b_512: blake2b.New512,
	constants.BLAKE2b_384: blake2b.New384,
	constants.BLAKE2b_256: blake2b.New256,
	constants.BLAKE2s_256: blake2s.New256,
	constants.BLAKE2s_128: blake2s.New128,
}

// Blake2 implements the BLAKE2 hashing algorithm
type Blake2 struct {
	params     Params
	config     *types.Config
	codec      interfaces.Codec
	hashFunc   func(key []byte) (hash.Hash, error)
	outputSize int
}

func (c *Blake2) Type() types.Type {
	return c.codec.Type()
}

type Params struct {
	Key []byte
}

func (p Params) String() string {
	var parts []string
	if len(p.Key) > 0 {
		parts = append(parts, fmt.Sprintf("k:%s", base64.RawURLEncoding.EncodeToString(p.Key)))
	}
	return strings.Join(parts, ",")
}

func parseParams(params string) (result Params, err error) {
	if params == "" {
		return result, nil
	}

	// Try to parse as JSON first
	if err := json.Unmarshal([]byte(params), &result); err == nil {
		return result, nil
	}

	// Fallback to internal key-value pair parsing
	kv, err := codecPkg.ParseParams(params)
	if err != nil {
		return Params{}, err
	}
	if v, ok := kv["k"]; ok {
		key, err := base64.RawURLEncoding.DecodeString(v)
		if err != nil {
			return result, fmt.Errorf("invalid key: %w", err)
		}
		result.Key = key
	}
	return result, nil
}

type ConfigValidator struct {
	params Params
}

func (v ConfigValidator) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return errors.ErrSaltLengthTooShort
	}
	if len(v.params.Key) > 0 {
		if len(v.params.Key) < MinKeyLength || len(v.params.Key) > MaxKeyLength {
			return fmt.Errorf("blake2b key length must be between %d and %d bytes", MinKeyLength, MaxKeyLength)
		}
	}
	return nil
}

// NewBlake2 creates a new BLAKE2 crypto instance
func NewBlake2(p types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}

	params, err := parseParams(config.ParamConfig)
	if err != nil {
		return nil, fmt.Errorf("invalid blake2 param config: %v", err)
	}

	validator := &ConfigValidator{
		params: params,
	}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid blake2 config: %v", err)
	}
	switch p.Name {
	case constants.BLAKE2b:
		p.Name = constants.DEFAULT_BLAKE2b
	case constants.BLAKE2s:
		p.Name = constants.DEFAULT_BLAKE2s
	}

	h, ok := hashFuncs[p.Name]
	if !ok {
		return nil, fmt.Errorf("unsupported blake2 type for keyed hash: %s", p.Name)
	}

	return &Blake2{
		params:   params,
		config:   config,
		codec:    codecPkg.NewCodec(p),
		hashFunc: h,
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}

// Hash implements the hash method
func (c *Blake2) Hash(password string) (string, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, string(salt))
}

// HashWithSalt implements the hash with salt method
func (c *Blake2) HashWithSalt(password, salt string) (string, error) {
	h, err := c.hashFunc(c.params.Key)
	if err != nil {
		return "", err
	}
	h.Write([]byte(password))
	h.Write([]byte(salt))
	hashBytes := h.Sum(nil)
	return c.codec.Encode([]byte(salt), hashBytes, c.params.String()), nil
}

// Verify implements the verify method
func (c *Blake2) Verify(parts *types.HashParts, password string) error {
	if parts.Algorithm.Name != c.codec.Type().Name {
		return errors.ErrAlgorithmMismatch
	}
	hashFunc, ok := hashFuncs[parts.Algorithm.Name]
	if !ok {
		return fmt.Errorf("unsupported blake2 type for keyed hash: %s", parts.Algorithm.Name)
	}
	// Recreate the hash function based on the stored parameters
	h, err := hashFunc(c.params.Key)
	if err != nil {
		return err
	}
	h.Write([]byte(password))
	h.Write(parts.Salt)
	newHash := h.Sum(nil)

	if subtle.ConstantTimeCompare(newHash, parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}

	return nil
}

// NewBlake2b creates a new BLAKE2b crypto instance
func NewBlake2b(p types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2(p, config)
}

// NewBlake2s creates a new BLAKE2s crypto instance
func NewBlake2s(p types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2(p, config)
}
