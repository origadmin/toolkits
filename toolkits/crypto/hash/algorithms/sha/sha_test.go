/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package sha

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestAllSHATypes(t *testing.T) {
	testCases := []struct {
		name    string
		algType types.Type
	}{
		{name: "SHA1", algType: sha1AlgType},
		{name: "SHA224", algType: sha224AlgType},
		{name: "SHA256", algType: sha256AlgType},
		{name: "SHA384", algType: sha384AlgType},
		{name: "SHA512", algType: sha512AlgType},
		{name: "SHA512_224", algType: sha512_224AlgType},
		{name: "SHA512_256", algType: sha512_256AlgType},
		{name: "SHA3_224", algType: sha3_224AlgType},
		{name: "SHA3_256", algType: sha3_256AlgType},
		{name: "SHA3_384", algType: sha3_384AlgType},
		{name: "SHA3_512", algType: sha3_512AlgType},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new SHA instance
			crypto, err := NewSHA(tc.algType, nil)
			assert.NoError(t, err)
			assert.NotNil(t, crypto)

			// Test hashing and verification
			password := "mysecretpassword"
			hashedParts, err := crypto.Hash(password)
			assert.NoError(t, err)
			assert.NotNil(t, hashedParts)

			// Verify with correct password
			err = crypto.Verify(hashedParts, password)
			assert.NoError(t, err)

			// Verify with incorrect password
			err = crypto.Verify(hashedParts, "wrongpassword")
			assert.Error(t, err)
		})
	}
}

func TestNewShaFunctions(t *testing.T) {
	testCases := []struct {
		name     string
		newFunc  func(config *types.Config) (interfaces.Cryptographic, error)
		expected types.Type
	}{
		{"NewSha1", NewSha1, sha1AlgType},
		{"NewSha224", NewSha224, sha224AlgType},
		{"NewSha256", NewSha256, sha256AlgType},
		{"NewSha384", NewSha384, sha384AlgType},
		{"NewSha512", NewSha512, sha512AlgType},
		{"NewSha3224", NewSha3224, sha3_224AlgType},
		{"NewSha3256", NewSha3256, sha3_256AlgType},
		{"NewSha3384", NewSha3384, sha3_384AlgType},
		{"NewSha3512", NewSha3512, sha3_512AlgType},
		{"NewSha3512224", NewSha3512224, sha512_224AlgType},
		{"NewSha3512256", NewSha3512256, sha512_256AlgType},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			crypto, err := tc.newFunc(nil)
			assert.NoError(t, err)
			assert.NotNil(t, crypto)
			assert.Equal(t, tc.expected, crypto.Type())
		})
	}
}
