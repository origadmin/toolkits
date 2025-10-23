/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package md5

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestNewMD5(t *testing.T) {
	tests := []struct {
		name    string
		config  *types.Config
		wantErr bool
	}{
		{
			name:    "Default config",
			config:  DefaultConfig(),
			wantErr: false,
		},
		{
			name:    "Custom SaltLength",
			config:  &types.Config{SaltLength: 32},
			wantErr: false,
		},
		{
			name:    "Invalid SaltLength",
			config:  &types.Config{SaltLength: 4}, // Less than 8
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewMD5(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMD5() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.NotNil(t, c)
				assert.Equal(t, types.MD5, c.Type().Name)
			}
		})
	}
}

func TestMD5_HashAndVerify(t *testing.T) {
	password := "testpassword"
	salt := []byte("testsalt12345678") // Must be at least 8 bytes for MD5

	tests := []struct {
		name    string
		algType types.Type
	}{
		{name: "MD5", algType: types.NewType(types.MD5)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			md5Alg, err := NewMD5(DefaultConfig())
			assert.NoError(t, err)
			assert.NotNil(t, md5Alg)

			// Test Hash with salt
			hashedParts, err := md5Alg.HashWithSalt(password, salt)
			assert.NoError(t, err)
			assert.NotNil(t, hashedParts)
			assert.Equal(t, tt.algType, hashedParts.Algorithm)
			assert.Equal(t, salt, hashedParts.Salt)

			// Test Verify with correct password and salt
			err = md5Alg.Verify(hashedParts, password)
			assert.NoError(t, err)

			// Test Verify with incorrect password
			err = md5Alg.Verify(hashedParts, "wrongpassword")
			assert.Error(t, err)
			assert.EqualError(t, err, "password does not match")

			// Test Hash without salt (SaltLength 0) - should return error
			cfg := DefaultConfig()
			cfg.SaltLength = 0 // This should cause an error due to ConfigValidator
			md5NoSalt, err := NewMD5(cfg)
			assert.Error(t, err)
			assert.Nil(t, md5NoSalt)
		})
	}
}
