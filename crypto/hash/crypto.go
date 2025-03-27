/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/algorithms/argon2"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/bcrypt"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/dummy"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/hmac256"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/md5"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/scrypt"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/sha1"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/sha256"
	"github.com/origadmin/toolkits/crypto/hash/base"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

var (
	// algorithms 存储所有支持的哈希算法
	algorithms = map[types.Type]func(*types.Config) (interfaces.Cryptographic, error){
		types.TypeArgon2:  argon2.NewArgon2Crypto,
		types.TypeBcrypt:  bcrypt.NewBcryptCrypto,
		types.TypeHMAC256: hmac256.NewHMAC256Crypto,
		types.TypeMD5:     md5.NewMD5Crypto,
		types.TypeScrypt:  scrypt.NewScryptCrypto,
		types.TypeSHA1:    sha1.NewSHA1Crypto,
		types.TypeSHA256:  sha256.NewSHA256Crypto,
		// 未实现的算法使用dummy实现
		types.TypeCustom:           dummy.NewDummyCrypto,
		types.TypeSHA512:           dummy.NewDummyCrypto,
		types.TypeSHA384:           dummy.NewDummyCrypto,
		types.TypeSHA3256:          dummy.NewDummyCrypto,
		types.TypeHMAC512:          dummy.NewDummyCrypto,
		types.TypePBKDF2:           dummy.NewDummyCrypto,
		types.TypePBKDF2SHA256:     dummy.NewDummyCrypto,
		types.TypePBKDF2SHA512:     dummy.NewDummyCrypto,
		types.TypePBKDF2SHA384:     dummy.NewDummyCrypto,
		types.TypePBKDF2SHA3256:    dummy.NewDummyCrypto,
		types.TypePBKDF2SHA3224:    dummy.NewDummyCrypto,
		types.TypePBKDF2SHA3384:    dummy.NewDummyCrypto,
		types.TypePBKDF2SHA3512224: dummy.NewDummyCrypto,
		types.TypePBKDF2SHA3512256: dummy.NewDummyCrypto,
		types.TypePBKDF2SHA3512384: dummy.NewDummyCrypto,
		types.TypePBKDF2SHA3512512: dummy.NewDummyCrypto,
	}
)

// NewCrypto 创建一个新的加密实例
func NewCrypto(opts ...types.ConfigOption) (interfaces.Cryptographic, error) {
	// 创建默认配置
	cfg := &types.Config{
		Algorithm:  types.TypeArgon2,
		TimeCost:   3,
		MemoryCost: 64 * 1024, // 64MB
		Threads:    2,
		SaltLength: 16,
	}

	// 应用配置选项
	for _, opt := range opts {
		opt(cfg)
	}

	// 获取算法创建器
	creator, exists := algorithms[cfg.Algorithm]
	if !exists {
		return nil, fmt.Errorf("unsupported algorithm: %s", cfg.Algorithm)
	}

	// 创建加密实例
	return creator(cfg)
}

// RegisterAlgorithm 注册新的哈希算法
func RegisterAlgorithm(t types.Type, creator func(*types.Config) (interfaces.Cryptographic, error)) {
	algorithms[t] = creator
}

// DecodeHash 解码哈希值
func DecodeHash(hashed string) (*types.HashParts, error) {
	// 使用通用编解码器解码密码
	codec := base.GetCodec(types.TypeCustom) // 使用通用编解码器
	return codec.Decode(hashed)
}

// Verify 验证密码
func Verify(hashed, password string) error {
	// 解码哈希值
	parts, err := DecodeHash(hashed)
	if err != nil {
		return err
	}

	// 获取算法创建器
	creator, exists := algorithms[parts.Algorithm]
	if !exists {
		return fmt.Errorf("unsupported algorithm: %s", parts.Algorithm)
	}

	// 创建加密实例并验证
	crypto, err := creator(&types.Config{
		Algorithm: parts.Algorithm,
	})
	if err != nil {
		return err
	}

	return crypto.Verify(hashed, password)
}
