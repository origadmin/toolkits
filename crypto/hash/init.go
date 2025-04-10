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
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type (
	AlgorithmCreator func(*types.Config) (interfaces.Cryptographic, error)
	AlgorithmConfig  func() *types.Config
)

type algorithm struct {
	creator       AlgorithmCreator
	defaultConfig AlgorithmConfig
}

var (
	// algorithms stores all supported hash algorithms
	algorithms = map[types.Type]algorithm{
		types.TypeArgon2: {
			creator:       argon2.NewArgon2Crypto,
			defaultConfig: argon2.DefaultConfig,
		},
		types.TypeBcrypt: {
			creator:       bcrypt.NewBcryptCrypto,
			defaultConfig: bcrypt.DefaultConfig,
		},
		types.TypeHMAC: {
			creator:       hmac.NewDefaultHMACCrypto,
			defaultConfig: hmac.DefaultConfig,
		},
		//types.TypeHMAC256: {
		//	creator:       hmac.NewHMAC256Crypto,
		//	defaultConfig: hmac.DefaultConfig,
		//},
		//types.TypeHMAC512: {
		//	creator:       hmac.NewHMAC512Crypto,
		//	defaultConfig: hmac.DefaultConfig,
		//},
		types.TypeMD5: {
			creator:       md5.NewMD5Crypto,
			defaultConfig: types.DefaultConfig,
		},
		types.TypeScrypt: {
			creator:       scrypt.NewScryptCrypto,
			defaultConfig: scrypt.DefaultConfig,
		},
		types.TypeSha1: {
			creator:       sha.NewSha1Crypto,
			defaultConfig: sha.DefaultConfig,
		},
		types.TypeSha224: {
			creator:       sha.NewSha224Crypto,
			defaultConfig: sha.DefaultConfig,
		},
		types.TypeSha256: {
			creator:       sha.NewSha256Crypto,
			defaultConfig: sha.DefaultConfig,
		},
		types.TypeSha512: {
			creator:       sha.NewSha512Crypto,
			defaultConfig: sha.DefaultConfig,
		},
		types.TypeSha3224: {
			creator:       sha.NewSha3224,
			defaultConfig: sha.DefaultConfig,
		},
		types.TypeSha3256: {
			creator:       sha.NewSha3256,
			defaultConfig: sha.DefaultConfig,
		},
		types.TypeSha384: {
			creator:       sha.NewSha384Crypto,
			defaultConfig: sha.DefaultConfig,
		},
		types.TypeSha3512: {
			creator:       sha.NewSha3512,
			defaultConfig: sha.DefaultConfig,
		},
		types.TypeSha3512224: {
			creator:       sha.NewSha3512224,
			defaultConfig: sha.DefaultConfig,
		},
		types.TypeSha3512256: {
			creator:       sha.NewSha3512256,
			defaultConfig: sha.DefaultConfig,
		},
		types.TypePBKDF2: {
			creator:       pbkdf2.NewPBKDF2Crypto,
			defaultConfig: pbkdf2.DefaultConfig,
		},
	}
)
