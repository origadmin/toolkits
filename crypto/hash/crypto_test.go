/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides hash functions for password encryption and comparison.
package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/origadmin/toolkits/crypto/hash/algorithms/argon2"
	"github.com/origadmin/toolkits/crypto/hash/algorithms/blake2"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// runCryptoTest is a helper function to run a standard set of tests on a Crypto instance.
func runCryptoTest(t *testing.T, c Crypto, algName string) {
	require.NotNil(t, c, "Crypto instance is nil for algorithm: %s", algName)

	password := "testpassword"

	// Test Hash method
	hashed, err := c.Hash(password)
	if err == nil {
		assert.NotEmpty(t, hashed, "Hashed string is empty for %s (Hash method)", algName)
		t.Logf("Hashed string for %s (Hash method): %s", algName, hashed)

		// Test Verify method - correct password
		verifyErr := c.Verify(hashed, password)
		assert.NoError(t, verifyErr, "Verification failed for %s with correct password (Hash method)", algName)

		// Test Verify method - incorrect password
		verifyErr = c.Verify(hashed, "wrongpassword")
		assert.Error(t, verifyErr, "Verification should fail for %s with wrong password (Hash method)", algName)
	} else {
		t.Logf("Skipping Hash method test for %s due to error: %v", algName, err)
	}

	// Test HashWithSalt method
	salt := []byte("testsalt12345678") // Example salt
	hashedWithSalt, err := c.HashWithSalt(password, salt)
	if err == nil {
		assert.NotEmpty(t, hashedWithSalt, "Hashed string is empty for %s (HashWithSalt method)", algName)

		// Test Verify method - correct password
		verifyErr := c.Verify(hashedWithSalt, password)
		assert.NoError(t, verifyErr, "Verification failed for %s with correct password (HashWithSalt method)", algName)

		// Test Verify method - incorrect password
		verifyErr = c.Verify(hashedWithSalt, "wrongpassword")
		assert.Error(t, verifyErr, "Verification should fail for %s with wrong password (HashWithSalt method)", algName)
	} else {
		t.Logf("Skipping HashWithSalt method test for %s due to error: %v", algName, err)
	}
}

func TestBaseAlgorithms(t *testing.T) {
	testCases := []struct {
		algName         string
		expectedAlgName string
		options         []Option
	}{
		{algName: types.MD5, expectedAlgName: types.MD5},
		{algName: types.SHA1, expectedAlgName: types.SHA1},
		{algName: types.SHA224, expectedAlgName: types.SHA224},
		{algName: types.SHA256, expectedAlgName: types.SHA256},
		{algName: types.SHA384, expectedAlgName: types.SHA384},
		{algName: types.SHA512, expectedAlgName: types.SHA512},
		{algName: types.SHA3, expectedAlgName: types.SHA3_256},
		{algName: types.SHA3_224, expectedAlgName: types.SHA3_224},
		{algName: types.SHA3_256, expectedAlgName: types.SHA3_256},
		{algName: types.SHA3_384, expectedAlgName: types.SHA3_384},
		{algName: types.SHA3_512, expectedAlgName: types.SHA3_512},
		{algName: types.SHA512_224, expectedAlgName: types.SHA512_224},
		{algName: types.SHA512_256, expectedAlgName: types.SHA512_256},
		{algName: types.BLAKE2b, expectedAlgName: types.DefaultBLAKE2b},
		{algName: types.BLAKE2s, expectedAlgName: types.DefaultBLAKE2s},
		{algName: types.BLAKE2b_256, expectedAlgName: types.BLAKE2b_256},
		{algName: types.BLAKE2b_384, expectedAlgName: types.BLAKE2b_384},
		{algName: types.BLAKE2b_512, expectedAlgName: types.BLAKE2b_512},
		{algName: types.BLAKE2s_128, expectedAlgName: types.BLAKE2s_128, options: []Option{blake2.WithKey([]byte("default-16-byte-key!!"))}},
		{algName: types.BLAKE2s_256, expectedAlgName: types.BLAKE2s_256, options: []Option{blake2.WithKey([]byte("default-32-byte-key-for-blake2s-256!!"))}},
		{algName: types.ARGON2, expectedAlgName: types.ARGON2i, options: []Option{argon2.WithParams(argon2.DefaultParams())}},
		{algName: types.ARGON2i, expectedAlgName: types.ARGON2i, options: []Option{argon2.WithParams(argon2.DefaultParams())}},
		{algName: types.ARGON2id, expectedAlgName: types.ARGON2id, options: []Option{argon2.WithParams(argon2.DefaultParams())}},
		{algName: types.BCRYPT, expectedAlgName: types.BCRYPT},
		{algName: types.SCRYPT, expectedAlgName: types.SCRYPT},
		{algName: types.RIPEMD, expectedAlgName: types.RIPEMD160},
		{algName: types.RIPEMD160, expectedAlgName: types.RIPEMD160},
		{algName: types.CRC32, expectedAlgName: types.CRC32_ISO},
		{algName: types.CRC32_ISO, expectedAlgName: types.CRC32_ISO},
		{algName: types.CRC32_CAST, expectedAlgName: types.CRC32_CAST},
		{algName: types.CRC32_KOOP, expectedAlgName: types.CRC32_KOOP},
		{algName: types.CRC64, expectedAlgName: types.CRC64_ISO},
		{algName: types.CRC64_ISO, expectedAlgName: types.CRC64_ISO},
		{algName: types.CRC64_ECMA, expectedAlgName: types.CRC64_ECMA},
	}

	for _, tc := range testCases {
		t.Run(tc.algName, func(t *testing.T) {
			c, err := NewCrypto(tc.algName, tc.options...)
			require.NoError(t, err, "Failed to create crypto for algorithm: %s", tc.algName)
			assert.Equal(t, tc.expectedAlgName, c.Spec().String(), "Unexpected algorithm name for %s", tc.algName)
			runCryptoTest(t, c, tc.algName)
		})
	}
}

func TestCompositeAlgorithms(t *testing.T) {
	testCases := []struct {
		algName         string
		expectedAlgName string
		options         []Option
	}{
		{algName: types.HMAC, expectedAlgName: types.DefaultHMAC},
		{algName: types.HMAC_SHA1, expectedAlgName: types.HMAC_SHA1},
		{algName: types.HMAC_SHA256, expectedAlgName: types.HMAC_SHA256},
		{algName: types.HMAC_SHA384, expectedAlgName: types.HMAC_SHA384},
		{algName: types.HMAC_SHA512, expectedAlgName: types.HMAC_SHA512},
		{algName: types.HMAC_SHA3_224, expectedAlgName: types.HMAC_SHA3_224},
		{algName: types.HMAC_SHA3_256, expectedAlgName: types.HMAC_SHA3_256},
		{algName: types.HMAC_SHA3_384, expectedAlgName: types.HMAC_SHA3_384},
		{algName: types.HMAC_SHA3_512, expectedAlgName: types.HMAC_SHA3_512},
		{algName: types.PBKDF2, expectedAlgName: types.DefaultPBKDF2},
		{algName: types.PBKDF2_SHA1, expectedAlgName: types.PBKDF2_SHA1},
		{algName: types.PBKDF2_SHA256, expectedAlgName: types.PBKDF2_SHA256},
		{algName: types.PBKDF2_SHA384, expectedAlgName: types.PBKDF2_SHA384},
		{algName: types.PBKDF2_SHA512, expectedAlgName: types.PBKDF2_SHA512},
		{algName: types.PBKDF2_SHA3_224, expectedAlgName: types.PBKDF2_SHA3_224},
		{algName: types.PBKDF2_SHA3_256, expectedAlgName: types.PBKDF2_SHA3_256},
		{algName: types.PBKDF2_SHA3_384, expectedAlgName: types.PBKDF2_SHA3_384},
		{algName: types.PBKDF2_SHA3_512, expectedAlgName: types.PBKDF2_SHA3_512},
		// Test with aliases
		{algName: "sha-256", expectedAlgName: types.SHA256},
		{algName: "sha-512", expectedAlgName: types.SHA512},
	}

	for _, tc := range testCases {
		t.Run(tc.algName, func(t *testing.T) {
			c, err := NewCrypto(tc.algName, tc.options...)
			require.NoError(t, err, "Failed to create crypto for algorithm: %s", tc.algName)
			assert.Equal(t, tc.expectedAlgName, c.Spec().String(), "Unexpected algorithm name for %s", tc.algName)
			runCryptoTest(t, c, tc.algName)
		})
	}
}

func TestNewCryptoWithOptions(t *testing.T) {
	testCases := []struct {
		algName string
		options []Option
	}{
		{algName: types.ARGON2, options: []Option{WithSaltLength(32)}},
		{algName: types.BCRYPT, options: []Option{WithSaltLength(16)}},
		{algName: types.SHA256, options: []Option{WithSaltLength(24)}},
		{algName: types.HMAC_SHA256, options: []Option{WithSaltLength(32)}},
		{algName: types.PBKDF2_SHA256, options: []Option{WithSaltLength(24)}},
	}

	for _, tc := range testCases {
		t.Run(tc.algName+"_with_options", func(t *testing.T) {
			crypto, err := NewCrypto(tc.algName, tc.options...)
			require.NoError(t, err, "Failed to create crypto with options for algorithm: %s", tc.algName)
			runCryptoTest(t, crypto, tc.algName)
		})
	}
}

func TestNewCryptoInvalidAlgorithm(t *testing.T) {
	invalidAlgorithms := []string{
		"invalid_algorithm",
		"hmac-invalid",
		"pbkdf2-invalid",
		"",
		"unknown",
	}

	for _, algName := range invalidAlgorithms {
		t.Run(algName, func(t *testing.T) {
			_, err := NewCrypto(algName)
			assert.Error(t, err, "Expected error for invalid algorithm: %s", algName)
		})
	}
}

func TestAvailableAlgorithms(t *testing.T) {
	availableAlgs := AvailableAlgorithms()
	assert.NotEmpty(t, availableAlgs, "AvailableAlgorithms list is empty")

	// Convert to a map for efficient lookup
	availableAlgsMap := make(map[string]bool)
	for _, alg := range availableAlgs {
		availableAlgsMap[alg] = true
	}

	// Check that some key algorithms and aliases are present
	expectedAlgs := []string{
		types.ARGON2,
		types.BCRYPT,
		types.SHA256,
		types.HMAC,
		types.PBKDF2,
		"sha-256", // alias
	}

	for _, algName := range expectedAlgs {
		assert.True(t, availableAlgsMap[algName], "Expected algorithm not found in available list: %s", algName)
	}
}
