/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"github.com/origadmin/toolkits/crypto/hash/algorithms/argon2"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/hmac"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/md5"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/pbkdf2"
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
	// algorithms stores all supported hash algorithms, keyed by their main name (string)
	algorithms = map[string]algorithm{
		constants.ARGON2: {
			creator:       argon2.NewArgon2,
			defaultConfig: argon2.DefaultConfig,
		},
		constants.BCRYPT: {
			creator:       wrapCreator(bcrypt.NewBcrypt),
			defaultConfig: bcrypt.DefaultConfig,
		},
		//		constants.BLAKE2b_256: {
		//	creator:       wrapCreator(blake2.NewBlake2b),
		//	defaultConfig: blake2.DefaultConfig,
		//},
		//constants.BLAKE2s_256: {
		//	creator:       wrapCreator(blake2.NewBlake2s),
		//	defaultConfig: blake2.DefaultConfig,
		//},
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
	}
)
