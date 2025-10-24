/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package blake2

import (
	"crypto/subtle"
	"fmt"
	"hash"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/validator"
	"github.com/origadmin/toolkits/crypto/rand"
)

type hashFunc func(key []byte) (hash.Hash, error)

var (
	blake2b512Spec = types.Spec{Name: types.BLAKE2b_512}
	blake2b384Spec = types.Spec{Name: types.BLAKE2b_384}
	blake2b256Spec = types.Spec{Name: types.BLAKE2b_256}
	blake2s256Spec = types.Spec{Name: types.BLAKE2s_256}
	blake2s128Spec = types.Spec{Name: types.BLAKE2s_128}
)

var hashFuncs = map[string]hashFunc{
	types.BLAKE2b_512: blake2b.New512,
	types.BLAKE2b_384: blake2b.New384,
	types.BLAKE2b_256: blake2b.New256,
	types.BLAKE2s_256: blake2s.New256,
	types.BLAKE2s_128: blake2s.New128,
}

// Blake2 implements the BLAKE2 hashing algorithm
type Blake2 struct {
	algSpec    types.Spec
	params     *Params
	config     *types.Config
	hashFunc   func(key []byte) (hash.Hash, error)
	outputSize int
}

func (c *Blake2) Spec() types.Spec {
	return c.algSpec
}

// Hash implements the hash method
func (c *Blake2) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *Blake2) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	h, err := c.hashFunc(c.params.Key)
	if err != nil {
		return nil, err
	}
	h.Write([]byte(password))
	if len(salt) > 0 {
		h.Write(salt)
	}
	hashBytes := h.Sum(nil)
	return types.NewHashPartsFull(c.algSpec, hashBytes, salt, c.params.ToMap()), nil
}

// Verify implements the verify method
func (c *Blake2) Verify(parts *types.HashParts, password string) error {
	// parts.Algorithm is already of type types.Spec. Use its Name field as the map key.
	hashFunc, ok := hashFuncs[parts.Algorithm.Name]
	if !ok {
		return fmt.Errorf("unsupported blake2 type for keyed hash: %s", parts.Algorithm.String())
	}
	// Recreate the hash function based on the stored parameters
	h, err := hashFunc(c.params.Key)
	if err != nil {
		return err
	}
	h.Write([]byte(password))
	if len(parts.Salt) > 0 {
		h.Write(parts.Salt)
	}
	hashBytes := h.Sum(nil)

	if subtle.ConstantTimeCompare(hashBytes, parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}

	return nil
}

func NewBlake2(algSpec types.Spec, config *types.Config) (interfaces.Cryptographic, error) {
	// Ensure algorithm-specific default config is applied when caller passes nil.
	if config == nil {
		config = DefaultConfig()
	}

	v, err := validator.ValidateParams(config, DefaultParams())
	if err != nil {
		return nil, fmt.Errorf("invalid blake2 config: %v", err)
	}

	hashFunc, ok := hashFuncs[algSpec.Name]
	if !ok {
		return nil, fmt.Errorf("unsupported blake2 type for keyed hash: %s", algSpec.Name)
	}
	return &Blake2{
		algSpec:  algSpec,
		params:   v.Params,
		config:   v.Config,
		hashFunc: hashFunc,
	}, nil
}

// NewBlake2b256 creates a new BLAKE2b crypto instance
func NewBlake2b256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2(blake2b256Spec, config)
}

// NewBlake2b384 creates a new BLAKE2b crypto instance
func NewBlake2b384(config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2(blake2b384Spec, config)
}

// NewBlake2b512 creates a new BLAKE2b crypto instance
func NewBlake2b512(config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2(blake2b512Spec, config)
}

// NewBlake2s128 creates a new BLAKE2s crypto instance
func NewBlake2s128(config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2(blake2s128Spec, config)
}

// NewBlake2s256 creates a new BLAKE2s crypto instance
func NewBlake2s256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2(blake2s256Spec, config)
}

func DefaultConfig() *types.Config {
	return types.DefaultConfig()
}
