/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package md5

import (
	"crypto/subtle"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/internal/validator"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

var md5AlgType = types.NewType(types.MD5)

type ConfigValidator struct {
}

func (v ConfigValidator) String() string {
	return ""
}

func (v ConfigValidator) ToMap() map[string]string {
	return map[string]string{}
}

func (v ConfigValidator) FromMap(params map[string]string) error {
	return nil
}

func (v ConfigValidator) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return fmt.Errorf("salt length must be at least 8 bytes")
	}
	return nil
}

// MD5 implements the MD5 hashing algorithm
type MD5 struct {
	config   *types.Config
	hashHash stdhash.Hash
}

func (c *MD5) Type() types.Type {
	return md5AlgType
}

// NewMD5 creates a new MD5 crypto instance
func NewMD5(config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	v := validator.WithParams(&ConfigValidator{})
	if err := v.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid md5 config: %v", err)
	}

	hashHash, err := types.TypeHash(md5AlgType.Name)
	if err != nil {
		return nil, err
	}

	return &MD5{
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
func (c *MD5) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *MD5) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	hash := c.hashHash.New()
	hash.Write([]byte(password))
	if len(salt) > 0 {
		hash.Write(salt)
	}
	hashBytes := hash.Sum(nil)
	return types.NewHashPartsWithHashSalt(c.Type(), hashBytes[:], salt), nil
}

// Verify implements the verify method
func (c *MD5) Verify(parts *types.HashParts, password string) error {
	// parts.Algorithm is already of type types.Type, so no need to parse it again.
	// We can directly use parts.Algorithm.Name for comparison.
	if parts.Algorithm.Name != types.MD5 {
		return errors.ErrAlgorithmMismatch
	}

	hash := c.hashHash.New()
	hash.Write([]byte(password))
	if len(parts.Salt) > 0 {
		hash.Write(parts.Salt)
	}
	hashBytes := hash.Sum(nil)
	if subtle.ConstantTimeCompare(hashBytes[:], parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}
	return nil
}
