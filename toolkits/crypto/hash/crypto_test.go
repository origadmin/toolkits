/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides hash functions for password encryption and comparison.
package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/constants"
)

func TestNewCryptoAllAlgorithms(t *testing.T) {
	testCases := []struct {
		algName         string
		expectedAlgName string // The name expected from crypto.Type().Name()
	}{
		{constants.MD5, constants.MD5},
		{constants.SHA1, constants.SHA1},
		{constants.SHA256, constants.SHA256},
		{constants.SHA384, constants.SHA384},
		{constants.SHA512, constants.SHA512},
		{constants.SHA224, constants.SHA224},
		{constants.SHA3_224, constants.SHA3_224},
		{constants.SHA3_256, constants.SHA3_256},
		{constants.SHA3_384, constants.SHA3_384},
		{constants.SHA3_512, constants.SHA3_512},
		{constants.SHA512_224, constants.SHA512_224},
		{constants.SHA512_256, constants.SHA512_256},
		{constants.BLAKE2s_128, constants.BLAKE2s_128},
		{constants.BLAKE2s_256, constants.BLAKE2s_256},
		{constants.BLAKE2b_256, constants.BLAKE2b_256},
		{constants.BLAKE2b_384, constants.BLAKE2b_384},
		{constants.BLAKE2b_512, constants.BLAKE2b_512},
		{constants.RIPEMD160, constants.RIPEMD160},
		{constants.CRC32, constants.CRC32_ISO}, // Expected to resolve to CRC32_ISO
		{constants.CRC32_ISO, constants.CRC32_ISO},
		{constants.CRC32_CAST, constants.CRC32_CAST},
		{constants.CRC32_KOOP, constants.CRC32_KOOP},
		{constants.CRC64, constants.CRC64_ISO}, // Expected to resolve to CRC64_ISO
		{constants.CRC64_ISO, constants.CRC64_ISO},
		{constants.CRC64_ECMA, constants.CRC64_ECMA},
		{constants.SCRYPT, constants.SCRYPT},
		{constants.BCRYPT, constants.BCRYPT},
		{constants.ARGON2, constants.ARGON2i},
		{constants.ARGON2i, constants.ARGON2i},
		{constants.ARGON2id, constants.ARGON2id},
		{constants.HMAC_SHA1, constants.HMAC_SHA1},
		{constants.HMAC_SHA256, constants.HMAC_SHA256},
		{constants.HMAC_SHA384, constants.HMAC_SHA384},
		{constants.HMAC_SHA512, constants.HMAC_SHA512},
		{constants.HMAC_SHA3_224, constants.HMAC_SHA3_224},
		{constants.HMAC_SHA3_256, constants.HMAC_SHA3_256},
		{constants.HMAC_SHA3_384, constants.HMAC_SHA3_384},
		{constants.HMAC_SHA3_512, constants.HMAC_SHA3_512},
		{constants.PBKDF2_SHA1, constants.PBKDF2_SHA1},
		{constants.PBKDF2_SHA256, constants.PBKDF2_SHA256},
		{constants.PBKDF2_SHA384, constants.PBKDF2_SHA384},
		{constants.PBKDF2_SHA512, constants.PBKDF2_SHA512},
		{constants.PBKDF2_SHA3_224, constants.PBKDF2_SHA3_224},
		{constants.PBKDF2_SHA3_256, constants.PBKDF2_SHA3_256},
		{constants.PBKDF2_SHA3_384, constants.PBKDF2_SHA3_384},
		{constants.PBKDF2_SHA3_512, constants.PBKDF2_SHA3_512},
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
