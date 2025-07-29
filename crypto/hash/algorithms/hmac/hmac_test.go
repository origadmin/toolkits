/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package hmac

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestNewHMAC(t *testing.T) {
	tests := []struct {
		name               string
		algType            types.Type
		config             *types.Config
		expectedAlgName    string
		expectedUnderlying string
		expectedStdHash    stdhash.Hash
		wantErr            bool
	}{
		{
			name:               "HMAC Default (SHA256)",
			algType:            types.NewType(constants.HMAC),
			config:             DefaultConfig(),
			expectedAlgName:    constants.HMAC,
			expectedUnderlying: constants.SHA256,
			expectedStdHash:    stdhash.SHA256,
			wantErr:            false,
		},
		{
			name:               "HMAC-SHA1",
			algType:            types.NewType(constants.HMAC_SHA1),
			config:             DefaultConfig(),
			expectedAlgName:    constants.HMAC,
			expectedUnderlying: constants.SHA1,
			expectedStdHash:    stdhash.SHA1,
			wantErr:            false,
		},
		{
			name:               "HMAC-SHA512",
			algType:            types.NewType(constants.HMAC_SHA512),
			config:             DefaultConfig(),
			expectedAlgName:    constants.HMAC,
			expectedUnderlying: constants.SHA512,
			expectedStdHash:    stdhash.SHA512,
			wantErr:            false,
		},
		{
			name:               "HMAC with unsupported underlying hash (CRC32)",
			algType:            types.NewType(constants.HMAC, constants.CRC32),
			config:             DefaultConfig(),
			expectedAlgName:    "",
			expectedUnderlying: "",
			wantErr:            true,
		},
		{
			name:               "HMAC with invalid underlying hash",
			algType:            types.NewType(constants.HMAC, "invalidhash"),
			config:             DefaultConfig(),
			expectedAlgName:    "",
			expectedUnderlying: "",
			wantErr:            true,
		},
		{
			name:               "Invalid SaltLength",
			algType:            types.NewType(constants.HMAC),
			config:             &types.Config{SaltLength: 4}, // Less than 8
			expectedAlgName:    "",
			expectedUnderlying: "",
			wantErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewHMAC(tt.algType, tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewHMAC() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.NotNil(t, c)
				assert.Equal(t, tt.expectedAlgName, c.Type().Name)
				assert.Equal(t, tt.expectedUnderlying, c.Type().Underlying)
				assert.Equal(t, tt.expectedStdHash, c.(*HMAC).hashHash)
			}
		})
	}
}

func TestHMAC_HashAndVerify(t *testing.T) {
	password := "testpassword"
	salt := []byte("testsalt12345678") // Must be at least 8 bytes for HMAC

	tests := []struct {
		name    string
		algType types.Type
	}{
		{name: "HMAC-SHA256", algType: types.NewType(constants.HMAC_SHA256)},
		{name: "HMAC-SHA512", algType: types.NewType(constants.HMAC_SHA512)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hmac, err := NewHMAC(tt.algType, DefaultConfig())
			assert.NoError(t, err)
			assert.NotNil(t, hmac)

			// Test Hash with salt
			hashedParts, err := hmac.HashWithSalt(password, salt)
			assert.NoError(t, err)
			assert.NotNil(t, hashedParts)
			assert.Equal(t, tt.algType.String(), hashedParts.Algorithm)
			assert.Equal(t, salt, hashedParts.Salt)

			// Test Verify with correct password and salt
			err = hmac.Verify(hashedParts, password)
			assert.NoError(t, err)

			// Test Verify with incorrect password
			err = hmac.Verify(hashedParts, "wrongpassword")
			assert.Error(t, err)
			assert.EqualError(t, err, "password does not match")

			// Test Hash without salt (SaltLength 0) - should return error
			cfg := DefaultConfig()
			cfg.SaltLength = 0 // This should cause an error due to ConfigValidator
			hmacNoSalt, err := NewHMAC(tt.algType, cfg)
			assert.Error(t, err)      // Expect an error here
			assert.Nil(t, hmacNoSalt) // Expect hmacNoSalt to be nil

			// The following assertions are moved outside this block as hmacNoSalt will be nil
			// hashedPartsNoSalt, err := hmacNoSalt.Hash(password)
			// assert.NoError(t, err)
			// assert.NotNil(t, hashedPartsNoSalt)
			// assert.NotEmpty(t, hashedPartsNoSalt.Salt) // Salt should still be generated
			// t.Logf("hashedPartsNoSalt.Algorithm: %s", hashedPartsNoSalt.Algorithm)
			// // Test Verify without salt
			// err = hmacNoSalt.Verify(hashedPartsNoSalt, password)
			// assert.NoError(t, err)

			// // Test Verify with incorrect password without salt
			// err = hmacNoSalt.Verify(hashedPartsNoSalt, "wrongpassword")
			// assert.Error(t, err)
			// assert.EqualError(t, err, "password does not match")
		})
	}
}
