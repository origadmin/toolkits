/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package ripemd160

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestNewRIPEMD160(t *testing.T) {
	tests := []struct {
		name            string
		algSpec         types.Spec
		config          *types.Config
		expectedAlgName string
		wantErr         bool
	}{
		{
			name:            "Default config",
			algSpec:         types.New(types.RIPEMD160),
			config:          DefaultConfig(),
			expectedAlgName: types.RIPEMD160,
			wantErr:         false,
		},
		{
			name:            "With underlying type (should be ignored)",
			algSpec:         types.New(types.RIPEMD160, "sha256"),
			config:          DefaultConfig(),
			expectedAlgName: types.RIPEMD160,
			wantErr:         false,
		},
		{
			name:            "Invalid SaltLength",
			algSpec:         types.New(types.RIPEMD160),
			config:          &types.Config{SaltLength: 4}, // Less than 8
			expectedAlgName: "",
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewRIPEMD160(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewRIPEMD160() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.NotNil(t, c)
				assert.Equal(t, tt.expectedAlgName, c.Spec().Name)
				assert.Empty(t, c.Spec().Underlying) // Underlying should always be empty after ResolveSpec
			}
		})
	}
}

func TestRIPEMD160_HashAndVerify(t *testing.T) {
	password := "testpassword"
	salt := []byte("testsalt12345678") // Must be at least 8 bytes for RIPEMD160

	tests := []struct {
		name    string
		algSpec types.Spec
	}{
		{name: "RIPEMD160", algSpec: types.New(types.RIPEMD160)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ripemd160Alg, err := NewRIPEMD160(DefaultConfig())
			assert.NoError(t, err)
			assert.NotNil(t, ripemd160Alg)

			// Test Hash with salt
			hashedParts, err := ripemd160Alg.HashWithSalt(password, salt)
			assert.NoError(t, err)
			assert.NotNil(t, hashedParts)
			assert.Equal(t, tt.algSpec, hashedParts.Algorithm)
			assert.Equal(t, salt, hashedParts.Salt)

			// Test Verify with correct password and salt
			err = ripemd160Alg.Verify(hashedParts, password)
			assert.NoError(t, err)

			// Test Verify with incorrect password
			err = ripemd160Alg.Verify(hashedParts, "wrongpassword")
			assert.Error(t, err)
			assert.EqualError(t, err, "password does not match")

		})
	}
}
