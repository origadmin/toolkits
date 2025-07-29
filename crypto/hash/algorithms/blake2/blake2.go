/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package blake2

import (
	"crypto/subtle"
	"fmt"
	"hash"

	"github.com/goexts/generic"
	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/validator"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

type hashFunc func(key []byte) (hash.Hash, error)

var (
	blake2b512Type = types.Type{Name: constants.BLAKE2b_512}
	blake2b384Type = types.Type{Name: constants.BLAKE2b_384}
	blake2b256Type = types.Type{Name: constants.BLAKE2b_256}
	blake2s256Type = types.Type{Name: constants.BLAKE2s_256}
	blake2s128Type = types.Type{Name: constants.BLAKE2s_128}
)

var hashFuncs = map[string]hashFunc{
	constants.BLAKE2b_512: blake2b.New512,
	constants.BLAKE2b_384: blake2b.New384,
	constants.BLAKE2b_256: blake2b.New256,
	constants.BLAKE2s_256: blake2s.New256,
	constants.BLAKE2s_128: blake2s.New128,
}

// Blake2 implements the BLAKE2 hashing algorithm
type Blake2 struct {
	p          types.Type
	params     *Params
	config     *types.Config
	hashFunc   func(key []byte) (hash.Hash, error)
	outputSize int
}

func (c *Blake2) Type() types.Type {
	return c.p
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
	h.Write(salt)
	hashBytes := h.Sum(nil)
	return types.NewHashPartsFull(c.p, hashBytes, salt, c.params.ToMap()), nil
}

// Verify implements the verify method
func (c *Blake2) Verify(parts *types.HashParts, password string) error {
	hashFunc, ok := hashFuncs[parts.Algorithm]
	if !ok {
		return fmt.Errorf("unsupported blake2 type for keyed hash: %s", parts.Algorithm)
	}
	// Recreate the hash function based on the stored parameters
	h, err := hashFunc(c.params.Key)
	if err != nil {
		return err
	}
	h.Write([]byte(password))
	h.Write(parts.Salt)
	hashBytes := h.Sum(nil)

	if subtle.ConstantTimeCompare(hashBytes, parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}

	return nil
}

func NewBlake2(p types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	// Use default config if provided config is nil
	if config == nil {
		config = types.DefaultConfig()
	}
	if config.ParamConfig == "" {
		config.ParamConfig = DefaultParams().String()
	}
	v := validator.WithParams(&Params{})
	if err := v.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid blake2 config: %v", err)
	}
	p = generic.Must(ResolveType(p))
	hashFunc, ok := hashFuncs[p.Name]
	if !ok {
		return nil, fmt.Errorf("unsupported blake2 type for keyed hash: %s", p.Name)
	}
	return &Blake2{
		p:        p,
		params:   v.Params(),
		config:   config,
		hashFunc: hashFunc,
	}, nil
}

// NewBlake2b256 creates a new BLAKE2b crypto instance
func NewBlake2b256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2(blake2b256Type, config)
}

// NewBlake2b384 creates a new BLAKE2b crypto instance
func NewBlake2b384(config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2(blake2b384Type, config)
}

// NewBlake2b512 creates a new BLAKE2b crypto instance
func NewBlake2b512(config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2(blake2b512Type, config)
}

// NewBlake2s128 creates a new BLAKE2s crypto instance
func NewBlake2s128(config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2(blake2s128Type, config)
}

// NewBlake2s256 creates a new BLAKE2s crypto instance
func NewBlake2s256(config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2(blake2s256Type, config)
}
