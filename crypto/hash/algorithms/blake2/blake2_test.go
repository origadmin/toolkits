/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package blake2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestNewBlake2(t *testing.T) {
	tests := []struct {
		name            string
		algType         types.Type
		config          *types.Config
		expectedAlgName string
		wantErr         bool
	}{
		{
			name:            "BLAKE2b Default Type",
			algType:         types.NewType(types.BLAKE2b),
			config:          types.DefaultConfig(),
			expectedAlgName: types.BLAKE2b_512,
			wantErr:         false,
		},
		{
			name:            "BLAKE2s Default Type",
			algType:         types.NewType(types.BLAKE2s),
			config:          types.DefaultConfig(),
			expectedAlgName: types.BLAKE2s_256,
			wantErr:         false,
		},
		{
			name:            "BLAKE2b_256 Explicit",
			algType:         types.NewType(types.BLAKE2b_256),
			config:          types.DefaultConfig(),
			expectedAlgName: types.BLAKE2b_256,
			wantErr:         false,
		},
		{
			name:            "BLAKE2s_128 Explicit",
			algType:         types.NewType(types.BLAKE2s_128),
			config:          types.DefaultConfig(),
			expectedAlgName: types.BLAKE2s_128,
			wantErr:         false,
		},
		{
			name:            "BLAKE2b Custom SaltLength",
			algType:         types.NewType(types.BLAKE2b_256),
			config:          &types.Config{SaltLength: 32},
			expectedAlgName: types.BLAKE2b_256,
			wantErr:         false,
		},
		{
			name:            "BLAKE2b Invalid SaltLength",
			algType:         types.NewType(types.BLAKE2b_256),
			config:          &types.Config{SaltLength: 4},
			expectedAlgName: "",
			wantErr:         true,
		},
		{
			name:            "Unsupported Type",
			algType:         types.NewType("unsupported"),
			config:          types.DefaultConfig(),
			expectedAlgName: "",
			wantErr:         true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, err := NewBlake2(tt.algType, tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBlake2() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				assert.NotNil(t, c)
				assert.Equal(t, tt.expectedAlgName, c.Type().Name)
			}
		})
	}
}

func TestCrypto_Hash(t *testing.T) {
	tests := []struct {
		name     string
		algType  types.Type
		password string
		wantErr  bool
	}{
		{
			name:     "BLAKE2b Hash",
			algType:  blake2b256Type,
			password: "testpassword",
			wantErr:  false,
		},
		{
			name:     "BLAKE2s Hash",
			algType:  blake2s256Type,
			password: "testpassword",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto, err := NewBlake2(tt.algType, types.DefaultConfig())
			if err != nil {
				t.Fatalf("Failed to create crypto: %v", err)
			}
			hash, err := crypto.Hash(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && hash.Algorithm.Name == "" {
				t.Error("Hash() returned empty string")
			}
			if err := crypto.Verify(hash, tt.password); err != nil {
				t.Errorf("Verify() error = %v", err)
			}
		})
	}
}

func TestCrypto_HashWithSalt(t *testing.T) {
	tests := []struct {
		name     string
		algType  types.Type
		password string
		salt     []byte
		wantErr  bool
	}{
		{
			name:     "BLAKE2b HashWithSalt",
			algType:  blake2b256Type,
			password: "testpassword",
			salt:     []byte("somesalt"),
			wantErr:  false,
		},
		{
			name:     "BLAKE2s HashWithSalt",
			algType:  blake2s256Type,
			password: "testpassword",
			salt:     []byte("somesalt"),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto, err := NewBlake2(tt.algType, types.DefaultConfig())
			if err != nil {
				t.Fatalf("Failed to create crypto: %v", err)
			}
			hash, err := crypto.HashWithSalt(tt.password, tt.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashWithSalt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && hash.Algorithm.Name == "" {
				t.Error("HashWithSalt() returned empty string")
			}
		})
	}
}

func TestCrypto_Verify(t *testing.T) {
	tests := []struct {
		name     string
		algType  types.Type
		password string
		salt     []byte
		wantErr  bool
	}{
		{
			name:     "BLAKE2b Verify Correct",
			algType:  blake2b256Type,
			password: "testpassword",
			salt:     []byte("somesalt"),
			wantErr:  false,
		},
		{
			name:     "BLAKE2s Verify Correct",
			algType:  blake2s256Type,
			password: "testpassword",
			salt:     []byte("somesalt"),
			wantErr:  false,
		},
		{
			name:     "BLAKE2b Verify Wrong Password",
			algType:  blake2b256Type,
			password: "testpassword",
			salt:     []byte("somesalt"),
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto, err := NewBlake2(tt.algType, types.DefaultConfig())
			if err != nil {
				t.Fatalf("Failed to create crypto: %v", err)
			}

			hashed, err := crypto.HashWithSalt(tt.password, tt.salt)
			if err != nil {
				t.Fatalf("HashWithSalt() error = %v", err)
			}

			// For wrong password test, change the password before verification
			verifyPassword := tt.password
			if tt.wantErr {
				verifyPassword = "wrongpassword"
			}

			err = crypto.Verify(hashed, verifyPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
