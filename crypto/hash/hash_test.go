/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides hash functions for password encryption and comparison.
package hash

import (
	"testing"

	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/stretchr/testify/assert"
)

func TestGenerateScryptPassword(t *testing.T) {
	password := "password"
	salt := "salt"

	crypto, err := NewCrypto(types.WithAlgorithm(types.TypeScrypt))
	assert.NoError(t, err)

	hashedPassword, err := crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestCompareScryptHashAndPassword(t *testing.T) {
	password := "password"
	salt := "salt"

	crypto, err := NewCrypto(types.WithAlgorithm(types.TypeScrypt))
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

	crypto, err := NewCrypto(types.WithAlgorithm(types.TypeArgon2))
	assert.NoError(t, err)

	hashedPassword, err := crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestCompareArgon2HashAndPassword(t *testing.T) {
	password := "password"
	salt := "salt"

	crypto, err := NewCrypto(types.WithAlgorithm(types.TypeArgon2))
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

	crypto, err := NewCrypto(types.WithAlgorithm(types.TypeArgon2))
	assert.NoError(t, err)

	hash, err := crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
}

func TestCompare(t *testing.T) {
	password := "password123"
	salt := "salt123"

	// Test MD5
	crypto, err := NewCrypto(types.WithAlgorithm(types.TypeMD5))
	assert.NoError(t, err)
	hashedPassword, err := crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test SHA1
	crypto, err = NewCrypto(types.WithAlgorithm(types.TypeSHA1))
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test Scrypt
	crypto, err = NewCrypto(types.WithAlgorithm(types.TypeScrypt))
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test Bcrypt
	crypto, err = NewCrypto(types.WithAlgorithm(types.TypeBcrypt))
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test Argon2
	crypto, err = NewCrypto(types.WithAlgorithm(types.TypeArgon2))
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)

	// Test SHA256
	crypto, err = NewCrypto(types.WithAlgorithm(types.TypeSHA256))
	assert.NoError(t, err)
	hashedPassword, err = crypto.HashWithSalt(password, salt)
	assert.NoError(t, err)
	err = crypto.Verify(hashedPassword, password)
	assert.NoError(t, err)
}
