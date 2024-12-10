/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides hash functions for password encryption and comparison.
package hash

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/errors"
)

func TestGenerateScryptPassword(t *testing.T) {
	password := "password"
	salt := "salt"

	hashedPassword, err := GenerateScryptPassword(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestCompareScryptHashAndPassword(t *testing.T) {
	password := "password"
	salt := "salt"

	hashedPassword, err := GenerateScryptPassword(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = CompareScryptHashAndPassword(hashedPassword, password, salt)
	assert.NoError(t, err)
}

func TestGenerateArgon2Password(t *testing.T) {
	password := "password"
	salt := "salt"

	hashedPassword, err := GenerateArgon2Password(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)
}

func TestCompareArgon2HashAndPassword(t *testing.T) {
	password := "password"
	salt := "salt"

	hashedPassword, err := GenerateArgon2Password(password, salt)
	assert.NoError(t, err)
	assert.NotEmpty(t, hashedPassword)

	err = CompareArgon2HashAndPassword(hashedPassword, password, salt)
	assert.NoError(t, err)
}

func TestGenerate(t *testing.T) {
	// Test case 1: Password and salt are valid
	password := "myPassword123"
	salt := "mySalt123"

	expectedHash := "generatedHash123"

	// Mocking the Generate function of gac package
	gac.Generate = func(password, salt string) (string, error) {
		return expectedHash, nil
	}

	hash, err := Generate(password, salt)

	assert.NoError(t, err)
	assert.Equal(t, expectedHash, hash)

	// Test case 2: Generation fails
	expectedError := "generation failed"

	// Mocking the Generate function of gac package
	gac.Generate = func(password, salt string) (string, error) {
		return "", errors.New(expectedError)
	}

	hash, err = Generate(password, salt)

	assert.Error(t, err)
	assert.EqualError(t, err, expectedError)
	assert.Empty(t, hash)
}

func TestCompare(t *testing.T) {
	var err error
	hashedPassword := ""
	password := "password123"
	salt := "salt123"

	UseMD5()

	hashedPassword, err = Generate(password, salt)
	t.Logf("MD5: %s", hashedPassword)
	err = Compare(hashedPassword, password, salt)
	assert.NoError(t, err, "Expected no error on successful comparison")

	// Test case 2: failed comparison
	hashedPassword = "hashedPassword456"
	err = Compare(hashedPassword, password, salt)
	assert.Error(t, err, "Expected error on failed comparison")

	UseSHA1()

	hashedPassword, err = Generate(password, salt)
	t.Logf("SHA1: %s", hashedPassword)
	err = Compare(hashedPassword, password, salt)
	assert.NoError(t, err, "Expected no error on successful comparison")

	// Test case 2: failed comparison
	hashedPassword = "hashedPassword456"
	err = Compare(hashedPassword, password, salt)
	assert.Error(t, err, "Expected error on failed comparison")

	UseScrypt()

	hashedPassword, err = Generate(password, salt)
	t.Logf("Scrypt: %s", hashedPassword)
	err = Compare(hashedPassword, password, salt)
	assert.NoError(t, err, "Expected no error on successful comparison")

	// Test case 2: failed comparison
	hashedPassword = "hashedPassword456"
	err = Compare(hashedPassword, password, salt)
	assert.Error(t, err, "Expected error on failed comparison")

	UseBcrypt()

	hashedPassword, err = Generate(password, salt)
	t.Logf("Bcrypt: %s", hashedPassword)
	err = Compare(hashedPassword, password, salt)
	assert.NoError(t, err, "Expected no error on successful comparison")

	// Test case 2: failed comparison
	hashedPassword = "hashedPassword456"
	err = Compare(hashedPassword, password, salt)
	assert.Error(t, err, "Expected error on failed comparison")

	UseArgon2()

	hashedPassword, err = Generate(password, salt)
	t.Logf("Argon2: %s", hashedPassword)
	err = Compare(hashedPassword, password, salt)
	assert.NoError(t, err, "Expected no error on successful comparison")

	// Test case 2: failed comparison
	hashedPassword = "hashedPassword456"
	err = Compare(hashedPassword, password, salt)
	assert.Error(t, err, "Expected error on failed comparison")

	UseSHA256()

	hashedPassword, err = Generate(password, salt)
	t.Logf("SHA256: %s", hashedPassword)
	err = Compare(hashedPassword, password, salt)
	assert.NoError(t, err, "Expected no error on successful comparison")

	// Test case 2: failed comparison
	hashedPassword = "hashedPassword456"
	err = Compare(hashedPassword, password, salt)
	assert.Error(t, err, "Expected error on failed comparison")
}
