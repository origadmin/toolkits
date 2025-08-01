/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides hash functions for password encryption and comparison.
package hash

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/crypto/hash/algorithms/scrypt"
	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func TestGenerateScryptPassword(t *testing.T) {
	password := "password"
	salt := "salt"

	crypto, err := NewCrypto(constants.SCRYPT)
	assert.NoError(t, err)

	hashedPassword, err := crypto.HashWithSalt(password, []byte(salt))
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestCompareScryptHashAndPassword(t *testing.T) {
	password := "password"
	salt := "salt"

	crypto, err := NewCrypto(constants.SCRYPT)
	assert.NoError(t, err)

	hashedPassword, err := crypto.HashWithSalt(password, []byte(salt))
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)
}

func TestGenerateArgon2Password(t *testing.T) {
	password := "password"
	salt := "salt"

	crypto, err := NewCrypto(constants.ARGON2)
	assert.NoError(t, err)

	hashedPassword, err := crypto.HashWithSalt(password, []byte(salt))
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = UseCrypto(constants.ARGON2, types.WithSaltLength(16))
	assert.NoError(t, err)
}

func TestCompareArgon2HashAndPassword(t *testing.T) {
	password := "password"
	salt := "salt"

	crypto, err := NewCrypto(constants.ARGON2)
	assert.NoError(t, err)

	hashedPassword, err := crypto.HashWithSalt(password, []byte(salt))
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)
}

func TestGenerate(t *testing.T) {
	password := "myPassword123"
	salt := "mySalt123"

	crypto, err := NewCrypto(constants.ARGON2)
	assert.NoError(t, err)

	hash, err := crypto.HashWithSalt(password, []byte(salt))
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestCompare(t *testing.T) {
	password := "password123"
	salt := "salt123"

	// Test MD5
	crypto, err := NewCrypto(constants.MD5)
	assert.NoError(t, err)
	hashedPassword, err := crypto.HashWithSalt(password, []byte(salt))
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test SHA1
	crypto, err = NewCrypto(constants.SHA1)
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, []byte(salt))
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test Scrypt
	crypto, err = NewCrypto(constants.SCRYPT)
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, []byte(salt))
	crypto2, err := NewCrypto(constants.SCRYPT, types.WithSaltLength(16))
	assert.NoError(t, err)
	err = crypto2.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test Bcrypt
	crypto, err = NewCrypto(constants.BCRYPT)
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, []byte(salt))
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test Argon2
	crypto, err = NewCrypto(constants.ARGON2)
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, []byte(salt))
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test SHA256
	crypto, err = NewCrypto(constants.SHA256)
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, []byte(salt))
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)
}

func TestAllHashTypes(t *testing.T) {
	tests := []struct {
		name     string
		algType  string
		password string
		salt     string
		options  []types.Option // Special configuration options
		crypto   Crypto
	}{
		{
			name:     "MD5",
			algType:  constants.MD5,
			password: "password123",
			salt:     "salt123",
		},
		{
			name:     "SHA1",
			algType:  constants.SHA1,
			password: "securePass!",
			salt:     "pepper456",
		},
		{
			name:     "ScryptDefault",
			algType:  constants.SCRYPT,
			password: "scryptPass",
			salt:     "scryptSalt",
		},
		{
			name:     "ScryptCustom",
			algType:  constants.SCRYPT,
			password: "scryptPass2",
			salt:     "scryptSalt2",
			options: []types.Option{types.WithSaltLength(16), types.WithParams(&scrypt.Params{
				N:      2,
				R:      1,
				P:      1,
				KeyLen: 16,
			})},
		},
		{
			name:     "Bcrypt",
			algType:  constants.BCRYPT,
			password: "bcryptPassword",
			salt:     "bcryptSalt",
		},
		{
			name:     "Argon2",
			algType:  constants.ARGON2,
			password: "argon2Password",
			salt:     "argon2Salt",
		},
		{
			name:     "SHA256",
			algType:  constants.SHA256,
			password: "sha256Password",
			salt:     "sha256Salt",
		},
	}

	var hashes []string
	var hashes2 []string
	for i, tt := range tests {
		generator, err := NewCrypto(tt.algType, tt.options...)
		assert.NoError(t, err)
		assert.NotNil(t, generator)
		tests[i].crypto = generator
		hashed, err := generator.HashWithSalt(tt.password, []byte(tt.salt))
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
			t.Logf("Verify %s hash", tt.algType)

			err := tt.crypto.Verify(hashes[i], tt.password)
			assert.NoError(t, err, "Failed to verify %s hash", tt.algType)

			err = tt.crypto.Verify(hashes[i], tt.password+"_invalid")
			if tt.algType == constants.MD5 {
				assert.NoError(t, err, "MD5 should not return an error for wrong password")
			} else {
				assert.Error(t, err, "Should fail with wrong password for %s", tt.algType)
			}

			err2 := tt.crypto.Verify(hashes2[i], tt.password)
			assert.NoError(t, err2, "Failed to verify %s hash", tt.algType)

			err2 = tt.crypto.Verify(hashes2[i], tt.password+"_invalid")
			if tt.algType == constants.MD5 {
				assert.NoError(t, err, "MD5 should not return an error for wrong password")
			} else {
				assert.Error(t, err2, "Should fail with wrong password for %s", tt.algType)
			}
		})
	}

	argonCrypto, _ := NewCrypto(constants.ARGON2)
	hashed, err := argonCrypto.Hash("passwordWithoutSalt")
	assert.NoError(t, err)
	for i, tt := range tests {
		t.Run(tt.name+"_VerifyWithCrypto", func(t *testing.T) {
			err = tt.crypto.Verify(hashed, "passwordWithoutSalt")
			assert.NoError(t, err, "Should verify hash without explicit salt")
		})
		t.Run(tt.name+"_VerifyWithCommon", func(t *testing.T) {
			t.Logf("hash details: %s", hashes[i])
			err := Verify(hashes[i], tt.password)
			assert.NoError(t, err, "Should verify hash without explicit salt")

			err2 := Verify(hashes[i], tt.password+"_invalid")
			if tt.algType == constants.MD5 {
				assert.NoError(t, err2, "MD5 should not return an error for wrong password")
			} else {
				assert.Error(t, err2, "Should fail with wrong password")
			}

			err3 := Verify(hashes2[i], tt.password)
			assert.NoError(t, err3, "Should verify hash without explicit salt")

			err4 := Verify(hashes2[i], tt.password+"_invalid")
			if tt.algType == constants.MD5 {
				assert.NoError(t, err4, "MD5 should not return an error for wrong password")
			} else {
				assert.Error(t, err4, "Should fail with wrong password")
			}

		})
	}
}

func TestUseCrypto(t *testing.T) {
	err := UseCrypto(constants.BCRYPT)
	assert.NoError(t, err)
	assert.Equal(t, constants.BCRYPT, activeCrypto.Type().Name)

	err = UseCrypto(constants.SCRYPT)
	assert.NoError(t, err)
	assert.Equal(t, constants.SCRYPT, activeCrypto.Type().Name)
}

func TestGenerateAndVerify(t *testing.T) {
	password := "testPassword"
	err := UseCrypto(constants.BCRYPT)
	assert.NoError(t, err)

	hashed, err := Generate(password)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)

	err = Verify(hashed, password)
	assert.NoError(t, err)

	err = Verify(hashed, "wrongPassword")
	assert.Error(t, err)
}

func TestGenerateWithSaltAndVerify(t *testing.T) {
	password := "testPassword"
	salt := []byte("testSalt")
	err := UseCrypto(constants.SCRYPT)
	assert.NoError(t, err)

	hashed, err := GenerateWithSalt(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashed)

	err = Verify(hashed, password)
	assert.NoError(t, err)

	err = Verify(hashed, "wrongPassword")
	assert.Error(t, err)
}

func TestInvalidAlgorithm(t *testing.T) {
	_, err := NewCrypto("invalid_alg")
	assert.Error(t, err)
	assert.Equal(t, errors.New("unsupported algorithm: invalid_alg"), err)
}

func TestUninitializedCrypto(t *testing.T) {
	activeCrypto = &uninitializedCrypto{}
	password := "testPassword"
	hashed, err := Generate(password)
	assert.Error(t, err)
	assert.Empty(t, hashed)
	assert.Equal(t, ErrHashModuleNotInitialized, err)

	_, err = GenerateWithSalt(password, []byte("salt"))
	assert.Error(t, err)
	assert.Equal(t, ErrHashModuleNotInitialized, err)

	err = Verify("hashed", password)
	assert.Error(t, err)
	assert.Equal(t, ErrHashModuleNotInitialized, err)
}
