/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package blake2

import (
	"crypto/subtle"
	"fmt"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// Blake2 implements the BLAKE2 hashing algorithm
type Blake2 struct {
	config   *types.Config
	codec    interfaces.Codec
	hashHash stdhash.Hash
}

func (c *Blake2) Type() string {
	return c.codec.Type().String()
}

type ConfigValidator struct{}

func (v ConfigValidator) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return errors.ErrSaltLengthTooShort
	}
	return nil
}

// NewBlake2Crypto creates a new BLAKE2 crypto instance
func NewBlake2Crypto(hashType types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	validator := &ConfigValidator{}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid blake2 config: %v", err)
	}

	hashHash, err := stdhash.ParseHash(hashType.String())
	if err != nil {
		return nil, err
	}
	//var stdhashType stdhash.Hash
	//switch hashType {
	//case types.TypeBlake2b:
	//	stdhashType, err = stdhash.ParseHash("blake2b-512") // Use blake2b-512 as default
	//	if err != nil {
	//		return nil, fmt.Errorf("failed to parse blake2b hash: %v", err)
	//	}
	//case types.TypeBlake2s:
	//	stdhashType, err = stdhash.ParseHash("blake2s-256") // Use blake2s-256 as default
	//	if err != nil {
	//		return nil, fmt.Errorf("failed to parse blake2s hash: %v", err)
	//	}
	//default:
	//	return nil, fmt.Errorf("unsupported BLAKE2 hash type: %s", hashType)
	//}

	return &Blake2{
		config:   config,
		codec:    codecPkg.NewCodec(hashType),
		hashHash: hashHash,
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
	h := c.hashHash.New()
	h.Write([]byte(password))
	h.Write([]byte(salt))
	hash := h.Sum(nil)

	return c.codec.Encode([]byte(salt), hash), nil
}

// Verify implements the verify method
func (c *Blake2) Verify(parts *types.HashParts, password string) error {
	if !parts.Algorithm.Is(c.codec.Type()) {
		return errors.ErrAlgorithmMismatch
	}

	hash, err := stdhash.ParseHash(parts.Algorithm.String())
	if err != nil {
		return err
	}

	h := hash.New()
	h.Write([]byte(password))
	h.Write(parts.Salt)
	newHash := h.Sum(nil)

	if subtle.ConstantTimeCompare(newHash, parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}

	return nil
}

// NewBlake2bCrypto creates a new BLAKE2b crypto instance
func NewBlake2bCrypto(config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2Crypto(types.TypeBlake2b, config)
}

// NewBlake2sCrypto creates a new BLAKE2s crypto instance
func NewBlake2sCrypto(config *types.Config) (interfaces.Cryptographic, error) {
	return NewBlake2Crypto(types.TypeBlake2s, config)
}
