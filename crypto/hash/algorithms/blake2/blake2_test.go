/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package blake2

import (
	"testing"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

var codec = codecPkg.NewCodec(types.TypeBlake2b) // Use Blake2b for testing codec

func TestNewBlake2Crypto(t *testing.T) {
	tests := []struct {
		name     string
		hashType types.Type
		config   *types.Config
		wantErr  bool
	}{
		{
			name:     "BLAKE2b Default Config",
			hashType: types.TypeBlake2b,
			config:   DefaultConfig(),
			wantErr:  false,
		},
		{
			name:     "BLAKE2s Default Config",
			hashType: types.TypeBlake2s,
			config:   DefaultConfig(),
			wantErr:  false,
		},
		{
			name:     "BLAKE2b Custom SaltLength",
			hashType: types.TypeBlake2b,
			config:   &types.Config{SaltLength: 32},
			wantErr:  false,
		},
		{
			name:     "BLAKE2b Invalid SaltLength",
			hashType: types.TypeBlake2b,
			config:   &types.Config{SaltLength: 4}, // Less than 8
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewBlake2Crypto(tt.hashType, tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBlake2Crypto() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCrypto_Hash(t *testing.T) {
	tests := []struct {
		name     string
		hashType types.Type
		password string
		wantErr  bool
	}{
		{
			name:     "BLAKE2b Hash",
			hashType: types.TypeBlake2b,
			password: "testpassword",
			wantErr:  false,
		},
		{
			name:     "BLAKE2s Hash",
			hashType: types.TypeBlake2s,
			password: "testpassword",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto, err := NewBlake2Crypto(tt.hashType, DefaultConfig())
			if err != nil {
				t.Fatalf("Failed to create crypto: %v", err)
			}
			hash, err := crypto.Hash(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && hash == "" {
				t.Error("Hash() returned empty string")
			}
		})
	}
}

func TestCrypto_HashWithSalt(t *testing.T) {
	tests := []struct {
		name     string
		hashType types.Type
		password string
		salt     string
		wantErr  bool
	}{
		{
			name:     "BLAKE2b HashWithSalt",
			hashType: types.TypeBlake2b,
			password: "testpassword",
			salt:     "somesalt",
			wantErr:  false,
		},
		{
			name:     "BLAKE2s HashWithSalt",
			hashType: types.TypeBlake2s,
			password: "testpassword",
			salt:     "somesalt",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto, err := NewBlake2Crypto(tt.hashType, DefaultConfig())
			if err != nil {
				t.Fatalf("Failed to create crypto: %v", err)
			}
			hash, err := crypto.HashWithSalt(tt.password, tt.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashWithSalt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && hash == "" {
				t.Error("HashWithSalt() returned empty string")
			}
		})
	}
}

func TestCrypto_Verify(t *testing.T) {
	tests := []struct {
		name     string
		hashType types.Type
		password string
		salt     string
		wantErr  bool
	}{
		{
			name:     "BLAKE2b Verify Correct",
			hashType: types.TypeBlake2b,
			password: "testpassword",
			salt:     "somesalt",
			wantErr:  false,
		},
		{
			name:     "BLAKE2s Verify Correct",
			hashType: types.TypeBlake2s,
			password: "testpassword",
			salt:     "somesalt",
			wantErr:  false,
		},
		{
			name:     "BLAKE2b Verify Wrong Password",
			hashType: types.TypeBlake2b,
			password: "testpassword",
			salt:     "somesalt",
			wantErr:  true, // Expect error for wrong password
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto, err := NewBlake2Crypto(tt.hashType, DefaultConfig())
			if err != nil {
				t.Fatalf("Failed to create crypto: %v", err)
			}

			hashed, err := crypto.HashWithSalt(tt.password, tt.salt)
			if err != nil {
				t.Fatalf("HashWithSalt() error = %v", err)
			}

			parts, err := codec.Decode(hashed)
			if err != nil {
				t.Fatalf("codec.Decode() error = %v", err)
			}

			// For wrong password test, change the password before verification
			verifyPassword := tt.password
			if tt.wantErr {
				verifyPassword = "wrongpassword"
			}

			err = crypto.Verify(parts, verifyPassword)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
