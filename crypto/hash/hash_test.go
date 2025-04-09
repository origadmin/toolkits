/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides hash functions for password encryption and comparison.
package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestGenerateScryptPassword(t *testing.T) {
	password := "password"
	salt := "salt"

	crypto, err := NewCrypto(types.TypeScrypt)
	assert.NoError(t, err)

	hashedPassword, err := crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestCompareScryptHashAndPassword(t *testing.T) {
	password := "password"
	salt := "salt"

	crypto, err := NewCrypto(types.TypeScrypt)
	assert.NoError(t, err)

	hashedPassword, err := crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)
}

func TestGenerateArgon2Password(t *testing.T) {
	password := "password"
	salt := "salt"

	crypto, err := NewCrypto(types.TypeArgon2)
	assert.NoError(t, err)

	hashedPassword, err := crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = UseCrypto(types.TypeArgon2, types.WithSaltLength(16))
	assert.NoError(t, err)
}

func TestCompareArgon2HashAndPassword(t *testing.T) {
	password := "password"
	salt := "salt"

	crypto, err := NewCrypto(types.TypeArgon2)
	assert.NoError(t, err)

	hashedPassword, err := crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)
}

func TestGenerate(t *testing.T) {
	password := "myPassword123"
	salt := "mySalt123"

	crypto, err := NewCrypto(types.TypeArgon2)
	assert.NoError(t, err)

	hash, err := crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestCompare(t *testing.T) {
	password := "password123"
	salt := "salt123"

	// Test MD5
	crypto, err := NewCrypto(types.TypeMD5)
	assert.NoError(t, err)
	hashedPassword, err := crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test SHA1
	crypto, err = NewCrypto(types.TypeSha1)
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test Scrypt
	crypto, err = NewCrypto(types.TypeScrypt)
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, salt)
	crypto2, err := NewCrypto(types.TypeScrypt, types.WithSaltLength(16))
	assert.NoError(t, err)
	err = crypto2.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test Bcrypt
	crypto, err = NewCrypto(types.TypeBcrypt)
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test Argon2
	crypto, err = NewCrypto(types.TypeArgon2)
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test SHA256
	crypto, err = NewCrypto(types.TypeSha256)
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)
}

func TestAllHashTypes(t *testing.T) {
	tests := []struct {
		name     string
		hashType types.Type
		password string
		salt     string
		options  []types.Option // Special configuration options
		crypto   Crypto
	}{
		{
			name:     "MD5",
			hashType: types.TypeMD5,
			password: "password123",
			salt:     "salt123",
		},
		{
			name:     "SHA1",
			hashType: types.TypeSha1,
			password: "securePass!",
			salt:     "pepper456",
		},
		{
			name:     "Scrypt-Default",
			hashType: types.TypeScrypt,
			password: "scryptPass",
			salt:     "scryptSalt",
		},
		{
			name:     "Scrypt-Custom",
			hashType: types.TypeScrypt,
			password: "scryptPass2",
			salt:     "scryptSalt2",
			options:  []types.Option{types.WithSaltLength(16)},
		},
		{
			name:     "Bcrypt",
			hashType: types.TypeBcrypt,
			password: "bcryptPassword",
			salt:     "bcryptSalt",
		},
		{
			name:     "Argon2",
			hashType: types.TypeArgon2,
			password: "argon2Password",
			salt:     "argon2Salt",
		},
		{
			name:     "SHA256",
			hashType: types.TypeSha256,
			password: "sha256Password",
			salt:     "sha256Salt",
		},
	}

	var hashes []string
	var hashes2 []string
	for i, tt := range tests {
		generator, err := NewCrypto(tt.hashType, tt.options...)
		assert.NoError(t, err)
		assert.NotNil(t, generator)
		tests[i].crypto = generator
		hashed, err := generator.HashWithSalt(tt.password, tt.salt)
		assert.NoError(t, err)
		assert.NotEmpty(t, hashed)
		hashes = append(hashes, hashed)

		hashed2, err2 := generator.Hash(tt.password)
		assert.NoError(t, err2)
		assert.NotEmpty(t, hashed2)
		hashes2 = append(hashes2, hashed2)
	}

	for i, tt := range tests {
		t.Run(tt.name+"_Verify", func(t *testing.T) {
			t.Logf("Verify %s hash", tt.hashType)

			err := tt.crypto.Verify(hashes[i], tt.password)
			assert.NoError(t, err, "Failed to verify %s hash", tt.hashType)

			err = tt.crypto.Verify(hashes[i], tt.password+"_invalid")
			assert.Error(t, err, "Should fail with wrong password for %s", tt.hashType)

			err2 := tt.crypto.Verify(hashes2[i], tt.password)
			assert.NoError(t, err2, "Failed to verify %s hash", tt.hashType)

			err2 = tt.crypto.Verify(hashes2[i], tt.password+"_invalid")
			assert.Error(t, err2, "Should fail with wrong password for %s", tt.hashType)
		})
	}

	argonCrypto, _ := NewCrypto(types.TypeArgon2)
	hashed, err := argonCrypto.Hash("passwordWithoutSalt")
	assert.NoError(t, err)
	for i, tt := range tests {
		t.Run("HashWithoutSalt", func(t *testing.T) {
			err = tt.crypto.Verify(hashed, "passwordWithoutSalt")
			assert.NoError(t, err, "Should verify hash without explicit salt")
		})
		t.Run("HashCommon", func(t *testing.T) {
			t.Logf("hash details: %s", hashes[i])
			err := Verify(hashes[i], tt.password)
			assert.NoError(t, err, "Should verify hash without explicit salt")

			err2 := Verify(hashes[i], tt.password+"_invalid")
			assert.Error(t, err2, "Should fail with wrong password")

			err3 := Verify(hashes2[i], tt.password)
			assert.NoError(t, err3, "Should verify hash without explicit salt")

			err4 := Verify(hashes2[i], tt.password+"_invalid")
			assert.Error(t, err4, "Should fail with wrong password")

		})
	}
}
