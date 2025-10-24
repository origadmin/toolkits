/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package ripemd160

import (
	"crypto/subtle"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// RIPEMD160 implements the RIPEMD-160 hashing algorithm
type RIPEMD160 struct {
	config   *types.Config
	hashHash stdhash.Hash
}

var ripemd160Spec = types.Spec{Name: types.RIPEMD160}

// Hash implements the hash method
func (r *RIPEMD160) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(r.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return r.HashWithSalt(password, salt)
}

// HashWithSalt is not applicable for simple hash functions like RIPEMD-160
func (r *RIPEMD160) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	h := r.hashHash.New()
	h.Write([]byte(password))
	if len(salt) > 0 {
		h.Write(salt)
	}
	return types.NewHashPartsFull(r.Spec(), h.Sum(nil), salt, nil), nil
}

// Verify implements the verify method
func (r *RIPEMD160) Verify(parts *types.HashParts, password string) error {
	h := r.hashHash.New()
	h.Write([]byte(password))
	if len(parts.Salt) > 0 {
		h.Write(parts.Salt)
	}
	newHash := h.Sum(nil)

	if subtle.ConstantTimeCompare(newHash, parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}
	return nil
}

func (r *RIPEMD160) Spec() types.Spec {
	return ripemd160Spec
}

// NewRIPEMD160 creates a new RIPEMD-160 crypto instance
func NewRIPEMD160(config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	if config.SaltLength < 8 {
		return nil, errors.ErrSaltLengthTooShort
	}

	hashHash, err := stdhash.ParseHash(ripemd160Spec.Name)
	if err != nil {
		return nil, err
	}

	return &RIPEMD160{
		config:   config,
		hashHash: hashHash,
	}, nil
}

// DefaultConfig returns the default configuration for RIPEMD-160
func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}
