/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package hmac

import (
	"crypto/hmac"
	"crypto/subtle"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/internal/validator"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// HMAC implements the HMAC hashing algorithm
type HMAC struct {
	p        types.Type
	config   *types.Config
	hashHash stdhash.Hash
}

var hmacType = types.Type{Name: constants.HMAC, Underlying: constants.SHA256}

func (c *HMAC) Type() types.Type {
	return c.p
}

type ConfigValidator struct {
	SaltLength int
}

func (v ConfigValidator) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return fmt.Errorf("salt length must be at least 8 bytes")
	}
	return nil
}

// NewHMAC creates a new HMAC crypto instance
func NewHMAC(p types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	v := validator.WithParams(&Params{})
	if err := v.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid hmac config: %v", err)
	}
	p, err := ResolveType(p)
	if err != nil {
		return nil, err
	}
	hashHash, err := types.TypeHash(p.Underlying)
	if err != nil {
		return nil, err
	}

	return &HMAC{
		config:   config,
		hashHash: hashHash,
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}

// Hash implements the hash method
func (c *HMAC) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *HMAC) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	h := hmac.New(c.hashHash.New, salt)
	h.Write([]byte(password))
	hash := h.Sum(nil)
	return types.NewHashPartsWithHashSalt(c.Type(), hash, salt), nil
}

// Verify implements the verify method
func (c *HMAC) Verify(parts *types.HashParts, password string) error {
	algorithm, err := types.ParseType(parts.Algorithm)
	if err != nil {
		return err
	}
	algorithm, err = ResolveType(algorithm)
	if err != nil {
		return err
	}
	if constants.HMAC != algorithm.Name {
		return errors.ErrAlgorithmMismatch
	}

	hashHash, err := types.TypeHash(algorithm.Underlying)
	if err != nil {
		return err
	}
	// Explicitly check for unsuitable hash types for HMAC
	// MAPHASH, ADLER32, CRC32, FNV are not cryptographically secure and should not be used with HMAC
	switch hashHash {
	case stdhash.MAPHASH, stdhash.ADLER32, stdhash.CRC32, stdhash.CRC32_ISO, stdhash.CRC32_CAST, stdhash.CRC32_KOOP,
		stdhash.CRC64_ISO, stdhash.CRC64_ECMA, stdhash.FNV32, stdhash.FNV32A, stdhash.FNV64, stdhash.FNV64A,
		stdhash.FNV128, stdhash.FNV128A:
		return errors.ErrUnsupportedHashForHMAC
	default:
	}
	h := hmac.New(hashHash.New, parts.Salt)
	h.Write([]byte(password))
	newHash := h.Sum(nil)
	if subtle.ConstantTimeCompare(newHash, parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}

	return nil
}
