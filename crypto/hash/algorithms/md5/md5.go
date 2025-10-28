/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package md5

import (
	"crypto/subtle"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/validator"
	"github.com/origadmin/toolkits/crypto/rand"
)

var specMD5 = types.Spec{Name: types.MD5}

// MD5 implements the MD5 hashing algorithm
type MD5 struct {
	config   *types.Config
	hashHash stdhash.Hash
}

func (c *MD5) Spec() types.Spec {
	return specMD5
}

// NewMD5 creates a new MD5 crypto instance
func NewMD5(config *types.Config) (scheme.Scheme, error) {
	if config == nil {
		config = DefaultConfig()
	}
	cfg, err := validator.ValidateConfig(config)
	if err != nil {
		return nil, fmt.Errorf("invalid md5 config: %v", err)
	}

	hashHash, err := types.Hash(specMD5.Name)
	if err != nil {
		return nil, err
	}

	return &MD5{
		config:   cfg,
		hashHash: hashHash,
	}, nil
}

// Hash implements the hash method
func (c *MD5) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *MD5) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	h := c.hashHash.New()
	h.Write([]byte(password))
	if len(salt) > 0 {
		h.Write(salt)
	}
	hashBytes := h.Sum(nil)
	return types.NewHashParts(c.Spec(), hashBytes, salt, nil), nil
}

// Verify implements the verify method
func (c *MD5) Verify(parts *types.HashParts, password string) error {
	// parts.Spec is already of type types.Spec, so no need to parse it again.
	// We can directly use parts.Spec.Name for comparison.
	if parts.Spec.Name != types.MD5 {
		return errors.ErrAlgorithmMismatch
	}

	h := c.hashHash.New()
	h.Write([]byte(password))
	if len(parts.Salt) > 0 {
		h.Write(parts.Salt)
	}
	hashBytes := h.Sum(nil)
	if subtle.ConstantTimeCompare(hashBytes[:], parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}
	return nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: types.DefaultSaltLength,
	}
}
