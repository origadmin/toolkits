/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
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
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type algorithm struct {
	algSpec       types.Spec // Added to store the algorithm type
	creator       interfaces.AlgorithmCreator
	defaultConfig interfaces.AlgorithmConfig
	resolver      interfaces.SpecResolver // New field: Resolver for this algorithm type
}

// defaultSpecResolver is a pass-through resolver for algorithms that don't need special resolution.
var defaultSpecResolver interfaces.SpecResolver = interfaces.AlgorithmResolver(func(algSpec types.Spec) (types.Spec, error) {
	algSpec.Name = algSpec.String()
	algSpec.Underlying = ""
	return algSpec, nil
})

func wrapCreator(oldCreator func(*types.Config) (interfaces.Cryptographic, error)) interfaces.AlgorithmCreator {
	return func(_ types.Spec, cfg *types.Config) (interfaces.Cryptographic, error) {
		return oldCreator(cfg)
	}
}

var (
	// algorithmMap stores all supported hash algorithmMap, keyed by their main name (string)
	algorithmMap = map[string]algorithm{
		types.ARGON2: {
			algSpec:       types.New(types.ARGON2),
			creator:       argon2.NewArgon2,
			defaultConfig: argon2.DefaultConfig,
			resolver:      interfaces.AlgorithmResolver(argon2.ResolveSpec), // 使用argon2自定义解析器
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
		types.BCRYPT: {
			algSpec:       types.New(types.BCRYPT),
			creator:       wrapCreator(bcrypt.NewBcrypt),
			defaultConfig: bcrypt.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.BLAKE2b: {
			algSpec:       types.New(types.BLAKE2b),
			creator:       blake2.NewBlake2,
			defaultConfig: blake2.DefaultConfig,
			resolver:      interfaces.AlgorithmResolver(blake2.ResolveSpec),
		},
		types.BLAKE2s: {
			algSpec:       types.New(types.BLAKE2s),
			creator:       blake2.NewBlake2,
			defaultConfig: blake2.DefaultConfig,
			resolver:      interfaces.AlgorithmResolver(blake2.ResolveSpec),
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
		},
		types.SHA224: {
			algSpec:       types.New(types.SHA224),
			creator:       wrapCreator(sha.NewSha224),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.SHA256: {
			algSpec:       types.New(types.SHA256),
			creator:       wrapCreator(sha.NewSha256),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.SHA512: {
			algSpec:       types.New(types.SHA512),
			creator:       wrapCreator(sha.NewSha512),
			defaultConfig: sha.DefaultConfig,
			resolver:      defaultSpecResolver,
		},
		types.SHA3: {
			algSpec:       types.New(types.SHA3),
			creator:       sha.NewSHA,
			defaultConfig: sha.DefaultConfig,
			resolver:      interfaces.AlgorithmResolver(sha.ResolveSpec),
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
			resolver:      interfaces.AlgorithmResolver(hmac.ResolveSpec), // 使用hmac自定义解析器
		},
		types.PBKDF2: { // PBKDF2 creator will now handle the Underlying type
			algSpec:       types.New(types.PBKDF2),
			creator:       pbkdf2.NewPBKDF2,
			defaultConfig: pbkdf2.DefaultConfig,
			resolver:      interfaces.AlgorithmResolver(pbkdf2.ResolveSpec), // 使用pbkdf2自定义解析器
		},
		types.RIPEMD: {
			algSpec:       types.New(types.RIPEMD),
			creator:       wrapCreator(ripemd160.NewRIPEMD160),
			defaultConfig: ripemd160.DefaultConfig,
			resolver:      interfaces.AlgorithmResolver(ripemd160.ResolveSpec), // 使用ripemd160自定义解析器
		},
		types.CRC32: {
			algSpec:       types.New(types.CRC32),
			creator:       crc.NewCRC,
			defaultConfig: crc.DefaultConfig,
			resolver:      interfaces.AlgorithmResolver(crc.ResolveSpec), // 使用crc自定义解析器
		},
		types.CRC64: {
			algSpec:       types.New(types.CRC64),
			creator:       crc.NewCRC,
			defaultConfig: crc.DefaultConfig,
			resolver:      interfaces.AlgorithmResolver(crc.ResolveSpec), // 使用crc自定义解析器
		},
	}
)
