/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package sha

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestAllSHASpecs(t *testing.T) {
	testCases := []struct {
		name    string
		algSpec types.Spec
	}{
		{name: "SHA1", algSpec: specSHA1},
		{name: "SHA224", algSpec: specSHA224},
		{name: "SHA256", algSpec: specSHA256},
		{name: "SHA384", algSpec: specSHA384},
		{name: "SHA512", algSpec: specSHA512},
		{name: "SHA512_224", algSpec: specSHA512_224},
		{name: "SHA512_256", algSpec: specSHA512_256},
		{name: "SHA3_224", algSpec: specSHA3_224},
		{name: "SHA3_256", algSpec: specSHA3_256},
		{name: "SHA3_384", algSpec: specSHA3_384},
		{name: "SHA3_512", algSpec: specSHA3_512},
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
		newFunc  func(config *types.Config) (scheme.Scheme, error)
		expected types.Spec
	}{
		{"NewSha1", NewSha1, specSHA1},
		{"NewSha224", NewSha224, specSHA224},
		{"NewSha256", NewSha256, specSHA256},
		{"NewSha384", NewSha384, specSHA384},
		{"NewSha512", NewSha512, specSHA512},
		{"NewSha3224", NewSha3224, specSHA3_224},
		{"NewSha3256", NewSha3256, specSHA3_256},
		{"NewSha3384", NewSha3384, specSHA3_384},
		{"NewSha3512", NewSha3512, specSHA3_512},
		{"NewSha3512224", NewSha3512224, specSHA512_224},
		{"NewSha3512256", NewSha3512256, specSHA512_256},
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
