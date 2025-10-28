/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"log/slog"
	"os"

	"github.com/origadmin/toolkits/crypto/hash/algorithms/argon2"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/blake2"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/crc"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/hmac"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/md5"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/pbkdf2"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/ripemd160"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/scrypt"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/sha"
	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// --- Legacy Algorithm Registration ---

type algorithm struct {
	algSpec       types.Spec
	creator       scheme.AlgorithmCreator
	defaultConfig scheme.AlgorithmConfig
	resolver      scheme.AlgorithmResolver
	alias         []string
}

var defaultSpecResolver = scheme.AlgorithmResolver(func(algSpec types.Spec) (types.Spec, error) {
	algSpec.Name = algSpec.String()
	algSpec.Underlying = ""
	return algSpec, nil
})

func wrapCreator(oldCreator func(*types.Config) (scheme.Scheme, error)) scheme.AlgorithmCreator {
	return func(_ types.Spec, cfg *types.Config) (scheme.Scheme, error) {
		return oldCreator(cfg)
	}
}

func init() {
	algorithmMap := map[string]algorithm{
		types.ARGON2: {
			algSpec:       types.New(types.ARGON2),
			creator:       argon2.NewArgon2,
			defaultConfig: argon2.DefaultConfig,
			resolver:      scheme.AlgorithmResolver(argon2.ResolveSpec),
			alias:         []string{types.ARGON2},
		},
		types.ARGON2i: {
			algSpec:       types.New(types.ARGON2i),
			creator:       wrapCreator(argon2.NewArgon2i),
			defaultConfig: argon2.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.ARGON2id: {
			algSpec:       types.New(types.ARGON2id),
			creator:       wrapCreator(argon2.NewArgon2id),
			defaultConfig: argon2.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.BLAKE2b: {
			algSpec:       types.New(types.BLAKE2b),
			creator:       blake2.NewBlake2,
			defaultConfig: blake2.DefaultConfig,
			resolver:      scheme.AlgorithmResolver(blake2.ResolveSpec),
		},
		types.BLAKE2s: {
			algSpec:       types.New(types.BLAKE2s),
			creator:       blake2.NewBlake2,
			defaultConfig: blake2.DefaultConfig,
			resolver:      scheme.AlgorithmResolver(blake2.ResolveSpec),
		},
		types.BCRYPT: {
			algSpec:       types.New(types.BCRYPT),
			creator:       wrapCreator(bcrypt.NewBcrypt),
			defaultConfig: bcrypt.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.MD5: {
			algSpec:       types.New(types.MD5),
			creator:       wrapCreator(md5.NewMD5),
			defaultConfig: md5.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.SCRYPT: {
			algSpec:       types.New(types.SCRYPT),
			creator:       wrapCreator(scrypt.NewScrypt),
			defaultConfig: scrypt.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.SHA1: {
			algSpec:       types.New(types.SHA1),
			creator:       wrapCreator(sha.NewSha1),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
			alias:         []string{"sha-1"},
		},
		types.SHA224: {
			algSpec:       types.New(types.SHA224),
			creator:       wrapCreator(sha.NewSha224),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
			alias:         []string{"sha-224"},
		},
		types.SHA256: {
			algSpec:       types.New(types.SHA256),
			creator:       wrapCreator(sha.NewSha256),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
			alias:         []string{"sha-256"},
		},
		types.SHA512: {
			algSpec:       types.New(types.SHA512),
			creator:       wrapCreator(sha.NewSha512),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
			alias:         []string{"sha-512"},
		},
		types.SHA3: {
			algSpec:       types.New(types.SHA3),
			creator:       sha.NewSHA,
			defaultConfig: sha.DefaultConfig,
			resolver:      scheme.AlgorithmResolver(sha.ResolveSpec),
		},
		types.SHA3_224: {
			algSpec:       types.New(types.SHA3_224),
			creator:       wrapCreator(sha.NewSha3224),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.SHA3_256: {
			algSpec:       types.New(types.SHA3_256),
			creator:       wrapCreator(sha.NewSha3256),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.SHA3_384: {
			algSpec:       types.New(types.SHA3_384),
			creator:       wrapCreator(sha.NewSha3384),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.SHA3_512: {
			algSpec:       types.New(types.SHA3_512),
			creator:       wrapCreator(sha.NewSha3512),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.SHA384: {
			algSpec:       types.New(types.SHA384),
			creator:       wrapCreator(sha.NewSha384),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
			alias:         []string{"sha-384"},
		},
		types.SHA512_224: {
			algSpec:       types.New(types.SHA512_224),
			creator:       wrapCreator(sha.NewSha3512224),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.SHA512_256: {
			algSpec:       types.New(types.SHA512_256),
			creator:       wrapCreator(sha.NewSha3512256),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.HMAC: { // HMAC creator will now handle the Underlying type
			algSpec:       types.New(types.HMAC),
			creator:       hmac.NewHMAC,
			defaultConfig: hmac.DefaultConfig,
			resolver:      scheme.AlgorithmResolver(hmac.ResolveSpec),
		},
		types.PBKDF2: { // PBKDF2 creator will now handle the Underlying type
			algSpec:       types.New(types.PBKDF2),
			creator:       pbkdf2.NewPBKDF2,
			defaultConfig: pbkdf2.DefaultConfig,
			resolver:      scheme.AlgorithmResolver(pbkdf2.ResolveSpec),
		},
		types.RIPEMD: {
			algSpec:       types.New(types.RIPEMD),
			creator:       wrapCreator(ripemd160.NewRIPEMD160),
			defaultConfig: ripemd160.DefaultConfig,
			resolver:      scheme.AlgorithmResolver(ripemd160.ResolveSpec),
		},
		types.CRC32: {
			algSpec:       types.New(types.CRC32),
			creator:       crc.NewCRC,
			defaultConfig: crc.DefaultConfig,
			resolver:      scheme.AlgorithmResolver(crc.ResolveSpec),
		},
		types.CRC64: {
			algSpec:       types.New(types.CRC64),
			creator:       crc.NewCRC,
			defaultConfig: crc.DefaultConfig,
			resolver:      scheme.AlgorithmResolver(crc.ResolveSpec),
		},
	}

	// --- Legacy Algorithm Registration (Moved from init.go) ---
	// Iterate over the old algorithmMap and register each one using the new factory system.
	for _, alg := range algorithmMap {
		// Create an adapter for each old algorithm entry.
		adapter := &legacyFactoryAdapter{
			creator:       alg.creator,
			defaultConfig: alg.defaultConfig,
			resolver:      alg.resolver,
		}

		// Use the new Register function.
		// We use the algSpec from the map as the canonical spec.
		// For now, no aliases are explicitly passed here, as the old system didn't define them this way.
		// Aliases will be added when individual algorithm packages are migrated.
		Register(adapter, alg.algSpec, alg.alias...)
	}

	// --- Global Crypto Instance Initialization (Moved from hash.go) ---
	algStr := os.Getenv(types.ENV)
	if algStr == "" {
		algStr = types.DefaultSpec
	}

	// Try to Create an encryption instance with the defined algorithm type
	crypto, err := NewCrypto(algStr)
	globalCryptoMutex.Lock()
	defer globalCryptoMutex.Unlock()
	if err != nil {
		slog.Error("hash: failed to initialize active crypto", "type", algStr, "error", err)
		globalCrypto = &uninitializedCrypto{}
	} else {
		globalCrypto = crypto
	}

}
