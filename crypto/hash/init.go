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

type (
	AlgorithmConfig func() *types.Config
	// AlgorithmCreator now accepts the full types.Type struct
	AlgorithmCreator func(algType types.Type, cfg *types.Config) (interfaces.Cryptographic, error)
)

type algorithm struct {
	algType       types.Type // Added to store the algorithm type
	creator       AlgorithmCreator
	defaultConfig AlgorithmConfig
}

func wrapCreator(oldCreator func(*types.Config) (interfaces.Cryptographic, error)) AlgorithmCreator {
	return func(_ types.Type, cfg *types.Config) (interfaces.Cryptographic, error) {
		return oldCreator(cfg)
	}
}

var (
	// algorithmMap stores all supported hash algorithmMap, keyed by their main name (string)
	algorithmMap = map[string]algorithm{
		types.ARGON2: {
			algType:       types.NewType(types.ARGON2),
			creator:       argon2.NewArgon2,
			defaultConfig: argon2.DefaultConfig,
		},
		types.ARGON2i: {
			algType:       types.NewType(types.ARGON2i),
			creator:       wrapCreator(argon2.NewArgon2i),
			defaultConfig: argon2.DefaultConfig,
		},
		types.ARGON2id: {
			algType:       types.NewType(types.ARGON2id),
			creator:       wrapCreator(argon2.NewArgon2id),
			defaultConfig: argon2.DefaultConfig,
		},
		types.BCRYPT: {
			algType:       types.NewType(types.BCRYPT),
			creator:       wrapCreator(bcrypt.NewBcrypt),
			defaultConfig: bcrypt.DefaultConfig,
		},
		types.BLAKE2b: {
			algType:       types.NewType(types.BLAKE2b),
			creator:       blake2.NewBlake2,
			defaultConfig: blake2.DefaultConfig,
		},
		types.BLAKE2s: {
			algType:       types.NewType(types.BLAKE2s),
			creator:       blake2.NewBlake2,
			defaultConfig: blake2.DefaultConfig,
		},
		types.MD5: {
			algType:       types.NewType(types.MD5),
			creator:       wrapCreator(md5.NewMD5),
			defaultConfig: md5.DefaultConfig,
		},
		types.SCRYPT: {
			algType:       types.NewType(types.SCRYPT),
			creator:       wrapCreator(scrypt.NewScrypt),
			defaultConfig: scrypt.DefaultConfig,
		},
		types.SHA1: {
			algType:       types.NewType(types.SHA1),
			creator:       wrapCreator(sha.NewSha1),
			defaultConfig: sha.DefaultConfig,
		},
		types.SHA224: {
			algType:       types.NewType(types.SHA224),
			creator:       wrapCreator(sha.NewSha224),
			defaultConfig: sha.DefaultConfig,
		},
		types.SHA256: {
			algType:       types.NewType(types.SHA256),
			creator:       wrapCreator(sha.NewSha256),
			defaultConfig: sha.DefaultConfig,
		},
		types.SHA512: {
			algType:       types.NewType(types.SHA512),
			creator:       wrapCreator(sha.NewSha512),
			defaultConfig: sha.DefaultConfig,
		},
		types.SHA3: {
			algType:       types.NewType(types.SHA3),
			creator:       sha.NewSHA,
			defaultConfig: sha.DefaultConfig,
		},
		//types.SHA3_224: {
		//	algType:       types.NewType(types.SHA3_224),
		//	creator:       wrapCreator(sha.NewSha3224),
		//	defaultConfig: sha.DefaultConfig,
		//},
		//types.SHA3_256: {
		//	algType:       types.NewType(types.SHA3_256),
		//	creator:       wrapCreator(sha.NewSha3256),
		//	defaultConfig: sha.DefaultConfig,
		//},
		//types.SHA3_384: {
		//	algType:       types.NewType(types.SHA3_384),
		//	creator:       wrapCreator(sha.NewSha3384),
		//	defaultConfig: sha.DefaultConfig,
		//},
		//types.SHA3_512: {
		//	algType:       types.NewType(types.SHA3_512),
		//	creator:       wrapCreator(sha.NewSha3512),
		//	defaultConfig: sha.DefaultConfig,
		//},
		types.SHA384: {
			algType:       types.NewType(types.SHA384),
			creator:       wrapCreator(sha.NewSha384),
			defaultConfig: sha.DefaultConfig,
		},
		types.SHA512_224: {
			algType:       types.NewType(types.SHA512_224),
			creator:       wrapCreator(sha.NewSha3512224),
			defaultConfig: sha.DefaultConfig,
		},
		types.SHA512_256: {
			algType:       types.NewType(types.SHA512_256),
			creator:       wrapCreator(sha.NewSha3512256),
			defaultConfig: sha.DefaultConfig,
		},
		types.HMAC: { // HMAC creator will now handle the Underlying type
			algType:       types.NewType(types.HMAC),
			creator:       hmac.NewHMAC,
			defaultConfig: hmac.DefaultConfig,
		},
		types.PBKDF2: { // PBKDF2 creator will now handle the Underlying type
			algType:       types.NewType(types.PBKDF2),
			creator:       pbkdf2.NewPBKDF2,
			defaultConfig: pbkdf2.DefaultConfig,
		},
		types.RIPEMD: {
			algType:       types.NewType(types.RIPEMD),
			creator:       wrapCreator(ripemd160.NewRIPEMD160),
			defaultConfig: ripemd160.DefaultConfig,
		},
		types.CRC32: {
			algType:       types.NewType(types.CRC32),
			creator:       crc.NewCRC,
			defaultConfig: crc.DefaultConfig,
		},
		types.CRC64: {
			algType:       types.NewType(types.CRC64),
			creator:       crc.NewCRC,
			defaultConfig: crc.DefaultConfig,
		},
	}
)
