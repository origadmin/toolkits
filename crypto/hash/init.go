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
	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type (
	AlgorithmConfig func() *types.Config
	// AlgorithmCreator now accepts the full types.Type struct
	AlgorithmCreator func(algType types.Type, cfg *types.Config) (interfaces.Cryptographic, error)
)

type algorithm struct {
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
		constants.ARGON2: {
			creator:       argon2.NewArgon2,
			defaultConfig: argon2.DefaultConfig,
		},
		constants.ARGON2i: {
			creator:       wrapCreator(argon2.NewArgon2i),
			defaultConfig: argon2.DefaultConfig,
		},
		constants.ARGON2id: {
			creator:       wrapCreator(argon2.NewArgon2id),
			defaultConfig: argon2.DefaultConfig,
		},
		constants.BCRYPT: {
			creator:       wrapCreator(bcrypt.NewBcrypt),
			defaultConfig: bcrypt.DefaultConfig,
		},
		constants.BLAKE2b: {
			creator:       blake2.NewBlake2,
			defaultConfig: blake2.DefaultConfig,
		},
		constants.BLAKE2s: {
			creator:       blake2.NewBlake2,
			defaultConfig: blake2.DefaultConfig,
		},
		constants.MD5: {
			creator:       wrapCreator(md5.NewMD5),
			defaultConfig: md5.DefaultConfig,
		},
		constants.SCRYPT: {
			creator:       wrapCreator(scrypt.NewScrypt),
			defaultConfig: scrypt.DefaultConfig,
		},
		constants.SHA1: {
			creator:       wrapCreator(sha.NewSha1),
			defaultConfig: sha.DefaultConfig,
		},
		constants.SHA224: {
			creator:       wrapCreator(sha.NewSha224),
			defaultConfig: sha.DefaultConfig,
		},
		constants.SHA256: {
			creator:       wrapCreator(sha.NewSha256),
			defaultConfig: sha.DefaultConfig,
		},
		constants.SHA512: {
			creator:       wrapCreator(sha.NewSha512),
			defaultConfig: sha.DefaultConfig,
		},
		constants.SHA3_224: {
			creator:       wrapCreator(sha.NewSha3224),
			defaultConfig: sha.DefaultConfig,
		},
		constants.SHA3_256: {
			creator:       wrapCreator(sha.NewSha3256),
			defaultConfig: sha.DefaultConfig,
		},
		constants.SHA3_384: {
			creator:       wrapCreator(sha.NewSha3384),
			defaultConfig: sha.DefaultConfig,
		},
		constants.SHA384: {
			creator:       wrapCreator(sha.NewSha384),
			defaultConfig: sha.DefaultConfig,
		},
		constants.SHA3_512: {
			creator:       wrapCreator(sha.NewSha3512),
			defaultConfig: sha.DefaultConfig,
		},
		constants.SHA3_512_224: {
			creator:       wrapCreator(sha.NewSha3512224),
			defaultConfig: sha.DefaultConfig,
		},
		constants.SHA3_512_256: {
			creator:       wrapCreator(sha.NewSha3512256),
			defaultConfig: sha.DefaultConfig,
		},
		constants.HMAC: { // HMAC creator will now handle the Underlying type
			creator:       hmac.NewHMAC,
			defaultConfig: hmac.DefaultConfig,
		},
		constants.PBKDF2: { // PBKDF2 creator will now handle the Underlying type
			creator:       pbkdf2.NewPBKDF2,
			defaultConfig: pbkdf2.DefaultConfig,
		},
		constants.RIPEMD: {
			creator:       wrapCreator(ripemd160.NewRIPEMD160),
			defaultConfig: ripemd160.DefaultConfig,
		},
		constants.CRC32: {
			creator:       crc.NewCRC,
			defaultConfig: crc.DefaultConfig,
		},
		constants.CRC64: {
			creator:       crc.NewCRC,
			defaultConfig: crc.DefaultConfig,
		},
	}
)
