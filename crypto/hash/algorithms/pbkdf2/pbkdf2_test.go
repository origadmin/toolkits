/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package pbkdf2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestNewPBKDF2(t *testing.T) {
	tests := []struct {
		name               string
		algSpec            types.Spec
		config             *types.Config
		expectedUnderlying string
		expectedErr        bool
	}{
		{
			name:               "PBKDF2 Default (SHA256)",
			algSpec:            types.New(types.PBKDF2),
			config:             DefaultConfig(),
			expectedUnderlying: types.SHA256,
			expectedErr:        false,
		},
		{
			name:               "PBKDF2-SHA1",
			algSpec:            types.New(types.PBKDF2, types.SHA1),
			config:             DefaultConfig(),
			expectedUnderlying: types.SHA1,
			expectedErr:        false,
		},
		{
			name:               "PBKDF2-HMAC-SHA256",
			algSpec:            types.New(types.PBKDF2, types.HMAC_SHA256),
			config:             DefaultConfig(),
			expectedUnderlying: types.HMAC_SHA256,
			expectedErr:        false,
		},
		{
			name:               "PBKDF2 with unsupported underlying hash (CRC32)",
			algSpec:            types.New(types.PBKDF2, types.CRC32),
			config:             DefaultConfig(),
			expectedUnderlying: "",
			expectedErr:        true,
		},
		{
			name:               "PBKDF2 with invalid underlying hash",
			algSpec:            types.New(types.PBKDF2, "invalidhash"),
			config:             DefaultConfig(),
			expectedUnderlying: "",
			expectedErr:        true,
		},
		{
			name:               "Invalid SaltLength",
			algSpec:            types.New(types.PBKDF2),
			config:             &types.Config{SaltLength: 4}, // Less than 8
			expectedUnderlying: "",
			expectedErr:        true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewPBKDF2(tt.algSpec, tt.config)
			if (err != nil) != tt.expectedErr {
				t.Errorf("NewPBKDF2() error = %v, expectedErr %v", err, tt.expectedErr)
			}
			if !tt.expectedErr {
				assert.NotNil(t, c)
				assert.Equal(t, types.PBKDF2, c.Spec().Name)
				assert.Equal(t, tt.expectedUnderlying, c.Spec().Underlying)
			}
		})
	}
}

func TestPBKDF2_HashAndVerify(t *testing.T) {
	password := "testpassword"
	salt := []byte("testsalt12345678") // Must be at least 8 bytes for PBKDF2

	tests := []struct {
		name                    string
		algSpec                 types.Spec
		expectedResolvedAlgSpec types.Spec
	}{
		{name: "PBKDF2-SHA256", algSpec: types.New(types.PBKDF2_SHA256), expectedResolvedAlgSpec: types.New(types.PBKDF2, types.SHA256)},
		{name: "PBKDF2-HMAC-SHA512", algSpec: types.New(types.PBKDF2, types.HMAC_SHA512), expectedResolvedAlgSpec: types.New(types.PBKDF2, types.HMAC_SHA512)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pbkdf2Alg, err := NewPBKDF2(tt.algSpec, DefaultConfig())
			assert.NoError(t, err)
			assert.NotNil(t, pbkdf2Alg)

			// Test Hash with salt
			hashedParts, err := pbkdf2Alg.HashWithSalt(password, salt)
			assert.NoError(t, err)
			assert.NotNil(t, hashedParts)
			assert.Equal(t, tt.expectedResolvedAlgSpec.String(), hashedParts.Spec)
			parsedAlgSpec := hashedParts.Spec
			assert.Equal(t, tt.expectedResolvedAlgSpec.Name, parsedAlgSpec.Name)
			assert.Equal(t, tt.expectedResolvedAlgSpec.Underlying, parsedAlgSpec.Underlying)
			assert.Equal(t, salt, hashedParts.Salt)

			// Test Verify with correct password and salt
			err = pbkdf2Alg.Verify(hashedParts, password)
			assert.NoError(t, err)

			// Test Verify with incorrect password
			err = pbkdf2Alg.Verify(hashedParts, "wrongpassword")
			assert.Error(t, err)
			assert.EqualError(t, err, "password does not match")

			// Test Hash without salt (SaltLength 0) - should return error
			cfg := DefaultConfig()
			cfg.SaltLength = 0 // This should cause an error due to ConfigValidator
			pbkdf2NoSalt, err := NewPBKDF2(tt.algSpec, cfg)
			assert.Error(t, err)
			assert.Nil(t, pbkdf2NoSalt)
		})
	}
}
