/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package crc

import (
	"fmt"
	"strings"

	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// CRC implements CRC32 and CRC64 hashing algorithms
type CRC struct {
	algType types.Type
	config  *types.Config
	stdHash stdhash.Hash
}

// Hash implements the hash method
func (c *CRC) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt is not applicable for simple checksum functions like CRC
func (c *CRC) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	h := c.stdHash.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		return nil, fmt.Errorf("crc: failed to write password: %w", err)
	}
	if len(salt) > 0 {
		_, err = h.Write(salt)
		if err != nil {
			return nil, fmt.Errorf("crc: failed to write salt: %w", err)
		}
	}
	return types.NewHashPartsFull(c.Type(), h.Sum(nil), salt, nil), nil
}

// Verify implements the verify method
func (c *CRC) Verify(parts *types.HashParts, password string) error {
	if parts.Hash == nil || len(parts.Hash) == 0 {
		return fmt.Errorf("crc: invalid hash parts: hash is nil or empty")
	}

	h := c.stdHash.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		return fmt.Errorf("crc: failed to write password for verification: %w", err)
	}
	if parts.Salt != nil && len(parts.Salt) > 0 {
		_, err = h.Write(parts.Salt)
		if err != nil {
			return fmt.Errorf("crc: failed to write salt for verification: %w", err)
		}
	}
	newHash := h.Sum(nil)

	if string(newHash) != string(parts.Hash) {
		return fmt.Errorf("crc: checksum does not match")
	}
	return nil
}

func (c *CRC) Type() types.Type {
	return c.algType
}

// NewCRC creates a new CRC crypto instance
func NewCRC(algType types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	var stdHashType stdhash.Hash
	switch strings.ToLower(algType.Name) {
	case "crc32":
		stdHashType = stdhash.CRC32
	case "crc64":
		stdHashType = stdhash.CRC64
	default:
		return nil, fmt.Errorf("crc: unsupported algorithm type: %s", algType.Name)
	}

	return &CRC{
		algType: algType,
		config:  config,
		stdHash: stdHashType,
	}, nil
}

// DefaultConfig returns the default configuration for CRC
func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 0, // CRC typically doesn't use salt, but allow for future extension
	}
}
