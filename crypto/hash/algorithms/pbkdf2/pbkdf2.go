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
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/validator"
	"github.com/origadmin/toolkits/crypto/rand"
)

var specPBKDF2_SHA256 = types.Spec{
	Name:       types.PBKDF2,
	Underlying: types.SHA256, // Default underlying hash for PBKDF2 is SHA-256
}

var unsafeHashes = map[string]bool{
	types.CRC32:      true,
	types.CRC32_ISO:  true,
	types.CRC32_CAST: true,
	types.CRC32_KOOP: true,
	types.CRC64_ISO:  true,
	types.CRC64_ECMA: true,
	"fnv32":          true,
	"fnv32a":         true,
	"fnv64":          true,
	"fnv64a":         true,
	"fnv128":         true,
	"fnv128a":        true,
	"adler32":        true,
	"maphash":        true,
	// 添加其他不安全的哈希算法
}

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
	return types.NewHashParts(c.Spec()).WithHashSalt(hashBytes, salt).WithParams(c.params.ToMap()), nil
}

// Verify implements the verify method
func (c *PBKDF2) Verify(parts *types.HashParts, password string) error {
	// parts.Spec is already of type types.Spec, so no need to parse it again.
	// We can directly use parts.Spec for comparison.
	if parts.Spec.Name != types.PBKDF2 {
		return errors.ErrInvalidAlgorithm
	}

	// Parse parameters
	params, err := FromMap(parts.Params)
	if err != nil {
		return err
	}

	prf, err := getPRF(parts.Spec) // Pass parts.Spec directly
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
		hmacHash, err := types.Hash(hmacHashName)
		if err != nil {
			return nil, err
		}
		// Explicitly check for unsuitable hash types for HMAC
		// MAPHASH, ADLER32, CRC32, FNV are not cryptographically secure and should not be used with HMAC
		switch hmacHash {
		case stdhash.MAPHASH, stdhash.ADLER32, stdhash.CRC32, stdhash.CRC32_ISO, stdhash.CRC32_CAST, stdhash.CRC32_KOOP,
			stdhash.CRC64_ISO, stdhash.CRC64_ECMA, stdhash.FNV32, stdhash.FNV32A, stdhash.FNV64, stdhash.FNV64A,
			stdhash.FNV128, stdhash.FNV128A:
			return nil, errors.ErrUnsupportedHashForHMAC
		default:
		}

		// PBKDF2 uses an internal key for HMAC, so we pass a dummy key here.
		// The actual key is derived internally by the pbkdf2.Key function.
		return func() hash.Hash { return hmac.New(hmacHash.New, []byte{}) }, nil
	} else {
		if unsafeHashes[algSpec.Underlying] {
			return nil, fmt.Errorf("unsupported hash algorithm for PBKDF2: %s", algSpec.Underlying)
		}

		hashHash, err := stdhash.ParseHash(algSpec.Underlying)
		if err != nil {
			return nil, err
		}
		return hashHash.New, nil
	}
}

// NewPBKDF2 creates a new PBKDF2 crypto instance
func NewPBKDF2(algSpec types.Spec, config *types.Config) (scheme.Scheme, error) {
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
