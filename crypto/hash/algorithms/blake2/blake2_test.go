/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package blake2

import (
	"testing"

	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestNewBlake2(t *testing.T) {
	tests := []struct {
		name     string
		hashType types.Type
		config   *types.Config
		wantErr  bool
	}{
		{
			name:     "BLAKE2b Default Config",
			hashType: blake2b256Type,
			config:   types.DefaultConfig(),
			wantErr:  false,
		},
		{
			name:     "BLAKE2s Default Config",
			hashType: blake2s256Type,
			config:   types.DefaultConfig(),
			wantErr:  false,
		},
		{
			name:     "BLAKE2b Custom SaltLength",
			hashType: blake2b256Type,
			config:   &types.Config{SaltLength: 32},
			wantErr:  false,
		},
		{
			name:     "BLAKE2b Invalid SaltLength",
			hashType: blake2b256Type,
			config:   &types.Config{SaltLength: 4}, // Less than 8
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewBlake2(tt.hashType, tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBlake2() error = %v, wantErr %v", err, tt.wantErr)
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
			hashType: blake2b256Type,
			password: "testpassword",
			wantErr:  false,
		},
		{
			name:     "BLAKE2s Hash",
			hashType: blake2s256Type,
			password: "testpassword",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto, err := NewBlake2(tt.hashType, types.DefaultConfig())
			if err != nil {
				t.Fatalf("Failed to create crypto: %v", err)
			}
			hash, err := crypto.Hash(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("Hash() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && hash.Algorithm == "" {
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
		hashType types.Type
		password string
		salt     []byte
		wantErr  bool
	}{
		{
			name:     "BLAKE2b HashWithSalt",
			hashType: blake2b256Type,
			password: "testpassword",
			salt:     []byte("somesalt"),
			wantErr:  false,
		},
		{
			name:     "BLAKE2s HashWithSalt",
			hashType: blake2s256Type,
			password: "testpassword",
			salt:     []byte("somesalt"),
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto, err := NewBlake2(tt.hashType, types.DefaultConfig())
			if err != nil {
				t.Fatalf("Failed to create crypto: %v", err)
			}
			hash, err := crypto.HashWithSalt(tt.password, tt.salt)
			if (err != nil) != tt.wantErr {
				t.Errorf("HashWithSalt() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && hash.Algorithm == "" {
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
		salt     []byte
		wantErr  bool
	}{
		{
			name:     "BLAKE2b Verify Correct",
			hashType: blake2b256Type,
			password: "testpassword",
			salt:     []byte("somesalt"),
			wantErr:  false,
		},
		{
			name:     "BLAKE2s Verify Correct",
			hashType: blake2s256Type,
			password: "testpassword",
			salt:     []byte("somesalt"),
			wantErr:  false,
		},
		{
			name:     "BLAKE2b Verify Wrong Password",
			hashType: blake2b256Type,
			password: "testpassword",
			salt:     []byte("somesalt"),
			wantErr:  true, // Expect error for wrong password
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			crypto, err := NewBlake2(tt.hashType, types.DefaultConfig())
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
