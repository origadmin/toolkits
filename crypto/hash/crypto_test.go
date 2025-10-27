/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides hash functions for password encryption and comparison.
package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/origadmin/toolkits/crypto/hash/algorithms/blake2"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestNewCryptoAllAlgorithms(t *testing.T) {
	// 基础算法测试用例
	baseAlgorithms := []struct {
		algName         string
		expectedAlgName string
		options         []Option
	}{{
		algName:         types.MD5,
		expectedAlgName: types.MD5,
		options:         nil,
	}, {
		algName:         types.SHA1,
		expectedAlgName: types.SHA1,
		options:         nil,
	}, {
		algName:         types.SHA224,
		expectedAlgName: types.SHA224,
		options:         nil,
	}, {
		algName:         types.SHA256,
		expectedAlgName: types.SHA256,
		options:         nil,
	}, {
		algName:         types.SHA384,
		expectedAlgName: types.SHA384,
		options:         nil,
	}, {
		algName:         types.SHA512,
		expectedAlgName: types.SHA512,
		options:         nil,
	}, {
		algName:         types.SHA3,
		expectedAlgName: types.SHA3_256,
		options:         nil,
	}, {
		algName:         types.SHA3_224,
		expectedAlgName: types.SHA3_224,
		options:         nil,
	}, {
		algName:         types.SHA3_256,
		expectedAlgName: types.SHA3_256,
		options:         nil,
	}, {
		algName:         types.SHA3_384,
		expectedAlgName: types.SHA3_384,
		options:         nil,
	}, {
		algName:         types.SHA3_512,
		expectedAlgName: types.SHA3_512,
		options:         nil,
	}, {
		algName:         types.SHA512_224,
		expectedAlgName: types.SHA512_224,
		options:         nil,
	}, {
		algName:         types.SHA512_256,
		expectedAlgName: types.SHA512_256,
		options:         nil,
	}, {
		algName:         types.BLAKE2b,
		expectedAlgName: types.DefaultBLAKE2b,
		options:         nil,
	}, {
		algName:         types.BLAKE2s,
		expectedAlgName: types.DefaultBLAKE2s,
		options:         nil,
	}, {
		algName:         types.BLAKE2b_256,
		expectedAlgName: types.BLAKE2b_256,
		options:         nil,
	}, {
		algName:         types.BLAKE2b_384,
		expectedAlgName: types.BLAKE2b_384,
		options:         nil,
	}, {
		algName:         types.BLAKE2b_512,
		expectedAlgName: types.BLAKE2b_512,
		options:         nil,
	}, {
		algName:         types.BLAKE2s_128,
		expectedAlgName: types.BLAKE2s_128,
		options:         []Option{blake2.WithKey([]byte("default-16-byte-key!!"))},
	}, {
		algName:         types.BLAKE2s_256,
		expectedAlgName: types.BLAKE2s_256,
		options:         []Option{blake2.WithKey([]byte("default-32-byte-key-for-blake2s-256!!"))},
	}, {
		algName:         types.ARGON2,
		expectedAlgName: types.ARGON2i,
		options:         nil,
	}, {
		algName:         types.ARGON2i,
		expectedAlgName: types.ARGON2i,
		options:         nil,
	}, {
		algName:         types.ARGON2id,
		expectedAlgName: types.ARGON2id,
		options:         nil,
	}, {
		algName:         types.BCRYPT,
		expectedAlgName: types.BCRYPT,
		options:         nil,
	}, {
		algName:         types.SCRYPT,
		expectedAlgName: types.SCRYPT,
		options:         nil,
	}, {
		algName:         types.RIPEMD,
		expectedAlgName: types.RIPEMD160,
		options:         nil,
	}, {
		algName:         types.RIPEMD160,
		expectedAlgName: types.RIPEMD160,
		options:         nil,
	}, {
		algName:         types.CRC32,
		expectedAlgName: types.CRC32_ISO,
		options:         nil,
	}, {
		algName:         types.CRC32_ISO,
		expectedAlgName: types.CRC32_ISO,
		options:         nil,
	}, {
		algName:         types.CRC32_CAST,
		expectedAlgName: types.CRC32_CAST,
		options:         nil,
	}, {
		algName:         types.CRC32_KOOP,
		expectedAlgName: types.CRC32_KOOP,
		options:         nil,
	}, {
		algName:         types.CRC64,
		expectedAlgName: types.CRC64_ISO,
		options:         nil,
	}, {
		algName:         types.CRC64_ISO,
		expectedAlgName: types.CRC64_ISO,
		options:         nil,
	}, {
		algName:         types.CRC64_ECMA,
		expectedAlgName: types.CRC64_ECMA,
		options:         nil,
	}}

	// 复合算法测试用例
	compositeAlgorithms := []struct {
		algName         string
		expectedAlgName string
		options         []Option
	}{{
		algName:         types.HMAC,
		expectedAlgName: types.DefaultHMAC,
		options:         nil,
	}, {
		algName:         types.HMAC_SHA1,
		expectedAlgName: types.HMAC_SHA1,
		options:         nil,
	}, {
		algName:         types.HMAC_SHA256,
		expectedAlgName: types.HMAC_SHA256,
		options:         nil,
	}, {
		algName:         types.HMAC_SHA384,
		expectedAlgName: types.HMAC_SHA384,
		options:         nil,
	}, {
		algName:         types.HMAC_SHA512,
		expectedAlgName: types.HMAC_SHA512,
		options:         nil,
	}, {
		algName:         types.HMAC_SHA3_224,
		expectedAlgName: types.HMAC_SHA3_224,
		options:         nil,
	}, {
		algName:         types.HMAC_SHA3_256,
		expectedAlgName: types.HMAC_SHA3_256,
		options:         nil,
	}, {
		algName:         types.HMAC_SHA3_384,
		expectedAlgName: types.HMAC_SHA3_384,
		options:         nil,
	}, {
		algName:         types.HMAC_SHA3_512,
		expectedAlgName: types.HMAC_SHA3_512,
		options:         nil,
	}, {
		algName:         types.PBKDF2,
		expectedAlgName: types.DefaultPBKDF2,
		options:         nil,
	}, {
		algName:         types.PBKDF2_SHA1,
		expectedAlgName: types.PBKDF2_SHA1,
		options:         nil,
	}, {
		algName:         types.PBKDF2_SHA256,
		expectedAlgName: types.PBKDF2_SHA256,
	}, {
		algName:         types.PBKDF2_SHA384,
		expectedAlgName: types.PBKDF2_SHA384,
	}, {
		algName:         types.PBKDF2_SHA512,
		expectedAlgName: types.PBKDF2_SHA512,
	}, {
		algName:         types.PBKDF2_SHA3_224,
		expectedAlgName: types.PBKDF2_SHA3_224,
	}, {
		algName:         types.PBKDF2_SHA3_256,
		expectedAlgName: types.PBKDF2_SHA3_256,
	}, {
		algName:         types.PBKDF2_SHA3_384,
		expectedAlgName: types.PBKDF2_SHA3_384,
	}, {
		algName:         types.PBKDF2_SHA3_512,
		expectedAlgName: types.PBKDF2_SHA3_512,
	}, {
		algName:         "sha-256",
		expectedAlgName: types.SHA256,
	}, {
		algName:         "sha-512",
		expectedAlgName: types.SHA512,
	},
	}

	// 合并所有测试用例
	allAlgorithms := append(baseAlgorithms, compositeAlgorithms...)

	// 测试所有算法
	for _, tc := range allAlgorithms {
		t.Run(tc.algName, func(t *testing.T) {
			// 测试创建算法实例
			crypto, err := NewCrypto(tc.algName, tc.options...)
			require.NoError(t, err, "Failed to Create crypto for algorithm: %s", tc.algName)
			require.NotNil(t, crypto, "Crypto instance is nil for algorithm: %s", tc.algName)

			// 测试算法类型是否正确
			assert.Equal(t, tc.expectedAlgName, crypto.Spec().String(), "Unexpected algorithm name for %s", tc.algName)

			// 测试 Hash 方法
			password := "testpassword"
			hashed, err := crypto.Hash(password)
			if err == nil {
				assert.NotEmpty(t, hashed, "Hashed string is empty for %s (Hash method)", tc.algName)

				// 测试 Verify 方法 - 正确的密码
				verifyErr := crypto.Verify(hashed, password)
				assert.NoError(t, verifyErr, "Verification failed for %s with correct password (Hash method)", tc.algName)

				// 测试 Verify 方法 - 错误的密码
				verifyErr = crypto.Verify(hashed, "wrongpassword")
				assert.Error(t, verifyErr, "Verification should fail for %s with wrong password (Hash method)", tc.algName)
			} else {
				t.Logf("Skipping Hash method test for %s due to error: %v", tc.algName, err)
			}

			// 测试 HashWithSalt 方法
			salt := []byte("testsalt12345678") // 示例盐值
			hashedWithSalt, err := crypto.HashWithSalt(password, salt)
			if err == nil {
				assert.NotEmpty(t, hashedWithSalt, "Hashed string is empty for %s (HashWithSalt method)", tc.algName)

				// 测试 Verify 方法 - 正确的密码
				verifyErr := crypto.Verify(hashedWithSalt, password)
				assert.NoError(t, verifyErr, "Verification failed for %s with correct password (HashWithSalt method)", tc.algName)

				// 测试 Verify 方法 - 错误的密码
				verifyErr = crypto.Verify(hashedWithSalt, "wrongpassword")
				assert.Error(t, verifyErr, "Verification should fail for %s with wrong password (HashWithSalt method)", tc.algName)
			} else {
				t.Logf("Skipping HashWithSalt method test for %s due to error: %v", tc.algName, err)
			}
		})
	}
}

func TestNewCryptoWithOptions(t *testing.T) {
	// 测试带有选项的算法创建
	testCases := []struct {
		algName string
		options []Option
	}{{
		algName: types.ARGON2,
		options: []Option{
			WithSaltLength(32),
			// Argon2的具体参数应该在其算法实现中处理
		},
	}, {
		algName: types.BCRYPT,
		options: []Option{
			WithSaltLength(16),
			// Bcrypt的具体参数应该在其算法实现中处理
		},
	}, {
		algName: types.SHA256,
		options: []Option{
			WithSaltLength(24),
		},
	}, {
		algName: types.HMAC_SHA256,
		options: []Option{
			WithSaltLength(32),
		},
	}, {
		algName: types.PBKDF2_SHA256,
		options: []Option{
			WithSaltLength(24),
			// PBKDF2的具体参数应该在其算法实现中处理
		},
	}}

	for _, tc := range testCases {
		t.Run(tc.algName+"_with_options", func(t *testing.T) {
			crypto, err := NewCrypto(tc.algName, tc.options...)
			require.NoError(t, err, "Failed to Create crypto with options for algorithm: %s", tc.algName)
			require.NotNil(t, crypto, "Crypto instance is nil with options for algorithm: %s", tc.algName)

			// 测试带选项的哈希和验证功能
			password := "testpassword"
			hashed, err := crypto.Hash(password)
			assert.NoError(t, err, "Hash failed with options for %s", tc.algName)
			assert.NotEmpty(t, hashed, "Hashed string is empty with options for %s", tc.algName)

			verifyErr := crypto.Verify(hashed, password)
			assert.NoError(t, verifyErr, "Verification failed with options for %s with correct password", tc.algName)

			verifyErr = crypto.Verify(hashed, "wrongpassword")
			assert.Error(t, verifyErr, "Verification should fail with options for %s with wrong password", tc.algName)
		})
	}
}

func TestNewCryptoInvalidAlgorithm(t *testing.T) {
	// 测试无效的算法名称
	invalidAlgorithms := []string{
		"invalid_algorithm",
		"hmac-invalid",
		"pbkdf2-invalid",
		"",
		"unknown",
	}

	for _, algName := range invalidAlgorithms {
		t.Run(algName, func(t *testing.T) {
			crypto, err := NewCrypto(algName)
			assert.Error(t, err, "Expected error for invalid algorithm: %s", algName)
			assert.Nil(t, crypto, "Expected nil crypto for invalid algorithm: %s", algName)
		})
	}
}

func TestAlgorithmMap(t *testing.T) {
	// 测试 AlgorithmMap 函数返回的算法映射表
	algMap := AlgorithmMap()
	assert.NotEmpty(t, algMap, "Spec map is empty")

	// 检查一些关键算法是否存在
	expectedAlgs := []string{
		types.ARGON2,
		types.BCRYPT,
		types.SHA256,
		types.HMAC,
		types.PBKDF2,
	}

	for _, algName := range expectedAlgs {
		_, exists := algMap[algName]
		assert.True(t, exists, "Expected algorithm not found in map: %s", algName)
	}
}
