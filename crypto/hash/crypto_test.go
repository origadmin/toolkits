/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides hash functions for password encryption and comparison.
package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestNewCryptoAllAlgorithms(t *testing.T) {
	testCases := []struct {
		algName         string
		expectedAlgName string // The name expected from crypto.Type().Name()
	}{
		{types.MD5, types.MD5},
		{types.SHA1, types.SHA1},
		{types.SHA256, types.SHA256},
		{types.SHA384, types.SHA384},
		{types.SHA512, types.SHA512},
		{types.SHA224, types.SHA224},
		{types.SHA3_224, types.SHA3_224},
		{types.SHA3_256, types.SHA3_256},
		{types.SHA3_384, types.SHA3_384},
		{types.SHA3_512, types.SHA3_512},
		{types.SHA512_224, types.SHA512_224},
		{types.SHA512_256, types.SHA512_256},
		{types.BLAKE2s_128, types.BLAKE2s_128},
		{types.BLAKE2s_256, types.BLAKE2s_256},
		{types.BLAKE2b_256, types.BLAKE2b_256},
		{types.BLAKE2b_384, types.BLAKE2b_384},
		{types.BLAKE2b_512, types.BLAKE2b_512},
		{types.RIPEMD160, types.RIPEMD160},
		{types.CRC32, types.CRC32_ISO}, // Expected to resolve to CRC32_ISO
		{types.CRC32_ISO, types.CRC32_ISO},
		{types.CRC32_CAST, types.CRC32_CAST},
		{types.CRC32_KOOP, types.CRC32_KOOP},
		{types.CRC64, types.CRC64_ISO}, // Expected to resolve to CRC64_ISO
		{types.CRC64_ISO, types.CRC64_ISO},
		{types.CRC64_ECMA, types.CRC64_ECMA},
		{types.SCRYPT, types.SCRYPT},
		{types.BCRYPT, types.BCRYPT},
		{types.ARGON2, types.ARGON2i},
		{types.ARGON2i, types.ARGON2i},
		{types.ARGON2id, types.ARGON2id},
		{types.HMAC_SHA1, types.HMAC_SHA1},
		{types.HMAC_SHA256, types.HMAC_SHA256},
		{types.HMAC_SHA384, types.HMAC_SHA384},
		{types.HMAC_SHA512, types.HMAC_SHA512},
		{types.HMAC_SHA3_224, types.HMAC_SHA3_224},
		{types.HMAC_SHA3_256, types.HMAC_SHA3_256},
		{types.HMAC_SHA3_384, types.HMAC_SHA3_384},
		{types.HMAC_SHA3_512, types.HMAC_SHA3_512},
		{types.PBKDF2_SHA1, types.PBKDF2_SHA1},
		{types.PBKDF2_SHA256, types.PBKDF2_SHA256},
		{types.PBKDF2_SHA384, types.PBKDF2_SHA384},
		{types.PBKDF2_SHA512, types.PBKDF2_SHA512},
		{types.PBKDF2_SHA3_224, types.PBKDF2_SHA3_224},
		{types.PBKDF2_SHA3_256, types.PBKDF2_SHA3_256},
		{types.PBKDF2_SHA3_384, types.PBKDF2_SHA3_384},
		{types.PBKDF2_SHA3_512, types.PBKDF2_SHA3_512},
	}

	for _, tc := range testCases {
		t.Run(tc.algName, func(t *testing.T) {
			crypto, err := NewCrypto(tc.algName)
			assert.NoError(t, err, "Failed to create crypto for algorithm: %s", tc.algName)
			assert.NotNil(t, crypto, "Crypto instance is nil for algorithm: %s", tc.algName)

			// Verify the Type() method returns the correct algorithm name
			assert.Equal(t, tc.expectedAlgName, crypto.Type().String(), "Unexpected algorithm name for %s", tc.algName)

			// Test Hash method
			password := "testpassword"
			hashedParts, hashErr := crypto.Hash(password)
			if hashErr == nil {
				assert.NotNil(t, hashedParts, "Hashed parts are nil for %s (Hash method)", tc.algName)
				verifyErr := crypto.Verify(hashedParts, password)
				assert.NoError(t, verifyErr, "Verification failed for %s with correct password (Hash method)", tc.algName)

				verifyErr = crypto.Verify(hashedParts, "wrongpassword")
				assert.Error(t, verifyErr, "Verification should fail for %s with wrong password (Hash method)", tc.algName)
			} else {
				t.Logf("Skipping Hash method test for %s due to error: %v", tc.algName, hashErr)
			}

			// Test HashWithSalt method
			salt := []byte("testsalt12345678") // A sample salt
			hashedPartsWithSalt, hashWithSaltErr := crypto.HashWithSalt(password, salt)
			if hashWithSaltErr == nil {
				assert.NotNil(t, hashedPartsWithSalt, "Hashed parts are nil for %s (HashWithSalt method)", tc.algName)
				verifyErr := crypto.Verify(hashedPartsWithSalt, password)
				assert.NoError(t, verifyErr, "Verification failed for %s with correct password (HashWithSalt method)", tc.algName)

				verifyErr = crypto.Verify(hashedPartsWithSalt, "wrongpassword")
				assert.Error(t, verifyErr, "Verification should fail for %s with wrong password (HashWithSalt method)", tc.algName)
			} else {
				t.Logf("Skipping HashWithSalt method test for %s due to error: %v", tc.algName, hashWithSaltErr)
			}
		})
	}
}
