/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package crc

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestNewCRC(t *testing.T) {
	tests := []struct {
		name            string
		algSpec         types.Spec
		expectedAlgName string
		expectedStdHash stdhash.Hash
		expectedErr     bool
	}{
		{name: "CRC32", algSpec: types.New(types.CRC32_ISO), expectedAlgName: "crc32-iso", expectedStdHash: stdhash.CRC32_ISO, expectedErr: false},
		{name: "CRC32-ISO", algSpec: types.New(types.CRC32_ISO), expectedAlgName: "crc32-iso", expectedStdHash: stdhash.CRC32_ISO, expectedErr: false},
		{name: "CRC32-CAST", algSpec: types.New("crc32-cast"), expectedAlgName: "crc32-cast", expectedStdHash: stdhash.CRC32_CAST, expectedErr: false},
		{name: "CRC32-KOOP", algSpec: types.New("crc32-koop"), expectedAlgName: "crc32-koop", expectedStdHash: stdhash.CRC32_KOOP, expectedErr: false},
		{name: "CRC64", algSpec: types.New(types.CRC64_ISO), expectedAlgName: "crc64-iso", expectedStdHash: stdhash.CRC64_ISO, expectedErr: false}, // Default to ISO
		{name: "CRC64-ISO", algSpec: types.New(types.CRC64_ISO), expectedAlgName: "crc64-iso", expectedStdHash: stdhash.CRC64_ISO, expectedErr: false},
		{name: "CRC64-ECMA", algSpec: types.New("crc64-ecma"), expectedAlgName: "crc64-ecma", expectedStdHash: stdhash.CRC64_ECMA, expectedErr: false},
		{name: "Unsupported", algSpec: types.New("unsupported"), expectedAlgName: "", expectedStdHash: 0, expectedErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crc, err := NewCRC(tt.algSpec, nil)
			if tt.expectedErr {
				assert.Error(t, err)
				assert.Nil(t, crc)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, crc)
				assert.Equal(t, tt.expectedAlgName, crc.Spec().Name)
				assert.Equal(t, tt.expectedStdHash, crc.(*CRC).hashHash)
			}
		})
	}
}

func TestCRCHashAndVerify(t *testing.T) {
	password := "testpassword"
	salt := []byte("testsalt")

	tests := []struct {
		name              string
		algSpec           types.Spec
		expectedAlgorithm string
	}{
		{name: "CRC32", algSpec: types.New(types.CRC32_ISO), expectedAlgorithm: "crc32-iso"},
		{name: "CRC64", algSpec: types.New(types.CRC64_ISO), expectedAlgorithm: "crc64-iso"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crc, err := NewCRC(tt.algSpec, DefaultConfig())
			assert.NoError(t, err)
			assert.NotNil(t, crc)

			// Test Hash with salt
			hashedParts, err := crc.HashWithSalt(password, salt)
			assert.NoError(t, err)
			assert.NotNil(t, hashedParts)
			assert.Equal(t, tt.expectedAlgorithm, hashedParts.Spec.String())
			assert.Equal(t, salt, hashedParts.Salt)

			// Test Verify with correct password and salt
			err = crc.Verify(hashedParts, password)
			assert.NoError(t, err)

			// Test Verify with incorrect password
			err = crc.Verify(hashedParts, "wrongpassword")
			assert.Error(t, err)
			assert.EqualError(t, err, "password does not match")

			// Test Hash without salt (SaltLength 0)
			cfg := DefaultConfig()
			cfg.SaltLength = 0
			crcNoSalt, err := NewCRC(tt.algSpec, cfg)
			assert.NoError(t, err)
			assert.NotNil(t, crcNoSalt)

			hashedPartsNoSalt, err := crcNoSalt.Hash(password)
			assert.NoError(t, err)
			assert.NotNil(t, hashedPartsNoSalt)
			assert.Empty(t, hashedPartsNoSalt.Salt)

			// Test Verify without salt
			err = crcNoSalt.Verify(hashedPartsNoSalt, password)
			assert.NoError(t, err)

			// Test Verify with incorrect password without salt
			err = crcNoSalt.Verify(hashedPartsNoSalt, "wrongpassword")
			assert.Error(t, err)
			assert.EqualError(t, err, "password does not match")
		})
	}
}

func TestCRCHashWithSaltEmpty(t *testing.T) {
	password := "testpassword"
	algSpec := types.New(types.CRC32_ISO)

	crc, err := NewCRC(algSpec, DefaultConfig())
	assert.NoError(t, err)

	// Test HashWithSalt with empty salt
	hashedParts, err := crc.HashWithSalt(password, []byte{})
	assert.NoError(t, err)
	assert.NotNil(t, hashedParts)
	assert.Empty(t, hashedParts.Salt)

	// Verify with empty salt
	err = crc.Verify(hashedParts, password)
	assert.NoError(t, err)

	err = crc.Verify(hashedParts, "wrongpassword")
	assert.Error(t, err)
	assert.EqualError(t, err, "password does not match")
}
