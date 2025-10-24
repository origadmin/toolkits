/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package pbkdf2

import (
	"crypto/hmac"
	"crypto/subtle"
	"fmt"
	"hash"
	"strings"

	"golang.org/x/crypto/pbkdf2"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/validator"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// PBKDF2 implements the PBKDF2 hashing algorithm
type PBKDF2 struct {
	algSpec types.Spec
	params  *Params
	config  *types.Config
	prf     func() hash.Hash
}

// Hash implements the hash method
func (c *PBKDF2) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *PBKDF2) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	hashBytes := pbkdf2.Key([]byte(password), salt, c.params.Iterations, int(c.params.KeyLength), c.prf)
	return types.NewHashPartsFull(c.Spec(), hashBytes, salt, c.params.ToMap()), nil
}

// Verify implements the verify method
func (c *PBKDF2) Verify(parts *types.HashParts, password string) error {
	// parts.Algorithm is already of type types.Spec, so no need to parse it again.
	// We can directly use parts.Algorithm for comparison.
	if parts.Algorithm.Name != types.PBKDF2 {
		return errors.ErrInvalidAlgorithm
	}

	// Parse parameters
	params, err := FromMap(parts.Params)
	if err != nil {
		return err
	}

	prf, err := getPRF(parts.Algorithm) // Pass parts.Algorithm directly
	if err != nil {
		return err
	}

	newHash := pbkdf2.Key([]byte(password), parts.Salt, params.Iterations, int(params.KeyLength), prf)
	if subtle.ConstantTimeCompare(newHash, parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}
	return nil
}

func (c *PBKDF2) Spec() types.Spec {
	return c.algSpec
}

// getPRF determines the Pseudo-Random Function (PRF) based on the algorithm type's underlying hash.
// It supports both direct hash functions and HMAC-based PRFs.
func getPRF(algSpec types.Spec) (func() hash.Hash, error) {
	if strings.HasPrefix(algSpec.Underlying, types.HMAC_PREFIX) {
		// Extract the underlying hash for HMAC
		hmacHashName := strings.TrimPrefix(algSpec.Underlying, types.HMAC_PREFIX)
		hmacHash, err := stdhash.ParseHash(hmacHashName)
		if err != nil {
			return nil, err
		}
		// PBKDF2 uses an internal key for HMAC, so we pass a dummy key here.
		// The actual key is derived internally by the pbkdf2.Key function.
		return func() hash.Hash { return hmac.New(hmacHash.New, []byte{}) }, nil
	} else {
		hashHash, err := stdhash.ParseHash(algSpec.Underlying)
		if err != nil {
			return nil, err
		}
		return hashHash.New, nil
	}
}

// NewPBKDF2 creates a new PBKDF2 crypto instance
func NewPBKDF2(algSpec types.Spec, config *types.Config) (interfaces.Cryptographic, error) {
	// Ensure algorithm-specific default config is applied when caller passes nil.
	if config == nil {
		config = DefaultConfig()
	}

	v, err := validator.ValidateParams(config, DefaultParams())
	if err != nil {
		return nil, fmt.Errorf("invalid pbkdf2 config: %v", err)
	}

	prf, err := getPRF(algSpec)
	if err != nil {
		return nil, err
	}

	return &PBKDF2{
		algSpec: algSpec,
		params:  v.Params,
		config:  v.Config,
		prf:     prf,
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}
