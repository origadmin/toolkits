/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"fmt"

	"github.com/goexts/generic/settings"

	"github.com/origadmin/toolkits/crypto/hash/algorithms/argon2"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/dummy"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/hmac256"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/md5"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/scrypt"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/sha1"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/sha256"
	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type algorithm struct {
	creator       func(*types.Config) (interfaces.Cryptographic, error)
	defaultConfig *types.Config
}

var (
	// algorithms stores all supported hash algorithms
	algorithms = map[types.Type]algorithm{
		types.TypeArgon2: {
			creator:       argon2.NewArgon2Crypto,
			defaultConfig: argon2.DefaultConfig(),
		},
		types.TypeBcrypt: {
			creator: bcrypt.NewBcryptCrypto,
			defaultConfig: &types.Config{
				TimeCost:   10,
				SaltLength: 16,
			},
		},
		types.TypeHMAC256: {
			creator: hmac256.NewHMAC256Crypto,
			defaultConfig: &types.Config{
				KeyLength: 32,
			},
		},
		types.TypeMD5: {
			creator: md5.NewMD5Crypto,
			defaultConfig: &types.Config{
				SaltLength: 16,
			},
		},
		types.TypeScrypt: {
			creator: scrypt.NewScryptCrypto,
			defaultConfig: &types.Config{
				TimeCost:   3,
				MemoryCost: 64 * 1024,
				Threads:    2,
				SaltLength: 16,
				KeyLength:  32,
			},
		},
		types.TypeSha1: {
			creator: sha1.NewSHA1Crypto,
			defaultConfig: &types.Config{
				SaltLength: 16,
			},
		},
		types.TypeSha256: {
			creator: sha256.NewSHA256Crypto,
			defaultConfig: &types.Config{
				SaltLength: 16,
			},
		},
		// Unimplemented cryptos use dummy implementation
		types.TypeCustom: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypeSha512: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypeSha384: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypeSha3256: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypeHMAC512: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypePBKDF2: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypePBKDF2SHA256: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypePBKDF2SHA512: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypePBKDF2SHA384: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypePBKDF2SHA3256: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypePBKDF2SHA3224: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypePBKDF2SHA3384: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypePBKDF2SHA3512224: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypePBKDF2SHA3512256: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypePBKDF2SHA3512384: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
		types.TypePBKDF2SHA3512512: {
			creator:       dummy.NewDummyCrypto,
			defaultConfig: &types.Config{},
		},
	}
)

type crypto struct {
	algorithm types.Type
	codec     interfaces.Codec
	crypto    interfaces.Cryptographic
	cryptos   map[types.Type]interfaces.Cryptographic
}

func (c crypto) Hash(password string) (string, error) {
	return c.crypto.Hash(password)
}

func (c crypto) HashWithSalt(password, salt string) (string, error) {
	return c.crypto.HashWithSalt(password, salt)
}

func (c crypto) Verify(hashed, password string) error {
	// Decode the hash value
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}

	// Get algorithm instance from cache or create new one
	crypto, exists := c.cryptos[parts.Algorithm]
	if !exists {
		algorithm, exists := algorithms[parts.Algorithm]
		if !exists {
			return fmt.Errorf("unsupported algorithm: %s", parts.Algorithm)
		}

		// Create cryptographic instance and cache it
		var err error
		crypto, err = algorithm.creator(algorithm.defaultConfig)
		if err != nil {
			return err
		}
		c.cryptos[parts.Algorithm] = crypto
	}

	return crypto.Verify(hashed, password)
}

// NewCrypto creates a new cryptographic instance
func NewCrypto(alg types.Type, opts ...types.ConfigOption) (interfaces.Cryptographic, error) {
	// Get algorithm creator and default config
	algorithm, exists := algorithms[alg]
	if !exists {
		return nil, fmt.Errorf("unsupported algorithm: %s", alg)
	}
	// Apply default config if not set
	cfg := settings.Apply(algorithm.defaultConfig, opts)
	cryptographic, err := algorithm.creator(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create cryptographic instance: %v", err)
	}
	// Create cryptographic instance
	return &crypto{
		algorithm: alg,
		crypto:    cryptographic,
		codec:     core.NewCodec(alg),
		cryptos:   make(map[types.Type]interfaces.Cryptographic),
	}, nil
}

// RegisterAlgorithm registers a new hash algorithm
func RegisterAlgorithm(t types.Type, creator func(*types.Config) (interfaces.Cryptographic, error), defaultConfig *types.Config) {
	algorithms[t] = struct {
		creator       func(*types.Config) (interfaces.Cryptographic, error)
		defaultConfig *types.Config
	}{
		creator:       creator,
		defaultConfig: defaultConfig,
	}
}

// Verify verifies a password
func Verify(hashed, password string) error {
	return defaultCrypto.Verify(hashed, password)
}

func Generate(password string) (string, error) {
	return defaultCrypto.Hash(password)
}

func GenerateWithSalt(password, salt string) (string, error) {
	return defaultCrypto.HashWithSalt(password, salt)
}
