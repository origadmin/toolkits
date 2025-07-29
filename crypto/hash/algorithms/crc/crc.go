/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package crc

import (
	"crypto/subtle"

	"github.com/goexts/generic"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// CRC implements CRC32 and CRC64 hashing algorithms
type CRC struct {
	algType  types.Type
	config   *types.Config
	hashHash stdhash.Hash
}

// Hash implements the hash method
func (c *CRC) Hash(password string) (*types.HashParts, error) {
	var salt []byte
	var err error
	if c.config.SaltLength > 0 {
		salt, err = rand.RandomBytes(c.config.SaltLength)
		if err != nil {
			return nil, err
		}
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *CRC) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	h := c.hashHash.New()
	_, _ = h.Write([]byte(password)) // Error is always nil for standard hash.Hash.Write
	if len(salt) > 0 {
		_, _ = h.Write(salt) // Error is always nil for standard hash.Hash.Write
	}
	return types.NewHashPartsFull(c.Type(), h.Sum(nil), salt, nil), nil
}

// Verify implements the verify method
func (c *CRC) Verify(parts *types.HashParts, password string) error {
	hashHash, err := types.TypeHash(parts.Algorithm)
	if err != nil {
		return err
	}
	h := hashHash.New()
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

func (c *CRC) Type() types.Type {
	return c.algType
}

// NewCRC creates a new CRC crypto instance
func NewCRC(algType types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}

	// No validator needed here, as CRC doesn't have complex params beyond SaltLength
	// which is handled by the main hash package's config validation.

	// The algType.Underlying will contain the specific stdhash name after ResolveType
	// We need to call ResolveType here to get the correct underlying hash.
	algType = generic.Must(ResolveType(algType))

	hashHash, err := stdhash.ParseHash(algType.Name)
	if err != nil {
		return nil, err
	}

	return &CRC{
		algType:  algType,
		config:   config,
		hashHash: hashHash,
	}, nil
}

// DefaultConfig returns the default configuration for CRC
func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 0, // CRC typically doesn't use salt, but allow for future extension
	}
}
