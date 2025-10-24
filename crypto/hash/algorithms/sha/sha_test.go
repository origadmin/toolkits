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

func TestAllSHASpecs(t *testing.T) {
	testCases := []struct {
		name    string
		algSpec types.Spec
	}{
		{name: "SHA1", algSpec: sha1AlgSpec},
		{name: "SHA224", algSpec: sha224AlgSpec},
		{name: "SHA256", algSpec: sha256AlgSpec},
		{name: "SHA384", algSpec: sha384AlgSpec},
		{name: "SHA512", algSpec: sha512AlgSpec},
		{name: "SHA512_224", algSpec: sha512_224AlgSpec},
		{name: "SHA512_256", algSpec: sha512_256AlgSpec},
		{name: "SHA3_224", algSpec: sha3_224AlgSpec},
		{name: "SHA3_256", algSpec: sha3_256AlgSpec},
		{name: "SHA3_384", algSpec: sha3_384AlgSpec},
		{name: "SHA3_512", algSpec: sha3_512AlgSpec},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a new SHA instance
			crypto, err := NewSHA(tc.algSpec, nil)
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
		expected types.Spec
	}{
		{"NewSha1", NewSha1, sha1AlgSpec},
		{"NewSha224", NewSha224, sha224AlgSpec},
		{"NewSha256", NewSha256, sha256AlgSpec},
		{"NewSha384", NewSha384, sha384AlgSpec},
		{"NewSha512", NewSha512, sha512AlgSpec},
		{"NewSha3224", NewSha3224, sha3_224AlgSpec},
		{"NewSha3256", NewSha3256, sha3_256AlgSpec},
		{"NewSha3384", NewSha3384, sha3_384AlgSpec},
		{"NewSha3512", NewSha3512, sha3_512AlgSpec},
		{"NewSha3512224", NewSha3512224, sha512_224AlgSpec},
		{"NewSha3512256", NewSha3512256, sha512_256AlgSpec},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			crypto, err := tc.newFunc(nil)
			assert.NoError(t, err)
			assert.NotNil(t, crypto)
			assert.Equal(t, tc.expected, crypto.Spec())
		})
	}
}
