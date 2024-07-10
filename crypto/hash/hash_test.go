// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package hash provides the hash functions
package hash

import (
	"testing"

	"github.com/origadmin/toolkits/errors"
	"github.com/stretchr/testify/assert"
)

func TestGeneratePassword(t *testing.T) {
	origin := "OrigAdmin@123456"
	hashPwd, err := GenerateBcryptPassword(origin, "")
	if err != nil {
		t.Error("GenerateBcryptPassword Failed: ", err.Error())
	}
	// t.Log("test password: ", hashPwd, ",length: ", len(hashPwd))

	if err := CompareBcryptHashAndPassword(hashPwd, origin, ""); err != nil {
		t.Error("Unmatched password: ", err.Error())
	}
}

func TestMD5(t *testing.T) {
	origin := "OrigAdmin@123456"
	hashVal := "06f684620c2e8f7caf9bb5a4fcba2ff2"
	if v := MD5String(origin); v != hashVal {
		t.Error("Failed to generate MD5 hash: ", v)
	}
}

func TestSHA1(t *testing.T) {
	origin := "OrigAdmin@123456"
	hashVal := "6d94221e0f42005e332ff9b468614ebe798786c5"
	if v := SHA1String(origin); v != hashVal {
		t.Error("Failed to generate MD5 hash: ", v)
	}
}

func TestSHA256(t *testing.T) {
	origin := "OrigAdmin@123456"
	hashVal := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"
	if v := SHA256String(origin); v != hashVal {
		t.Error("Failed to generate MD5 hash: ", v)
	}
}

func TestHMAC256(t *testing.T) {
	origin := "OrigAdmin@123456"
	key := "key"
	hashVal := "47f03a422b85f8bc524f283e78b70c9f026db157de8c21cf2330238cfb54cd56"
	if v := HMAC256String(origin, key); v != hashVal {
		t.Error("Failed to generate HMAC256 hash: ", v)
	}
}

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

	// Mocking the Generate function of defaultGAC package
	defaultGAC.Generate = func(password, salt string) (string, error) {
		return expectedHash, nil
	}

	hash, err := Generate(password, salt)

	assert.NoError(t, err)
	assert.Equal(t, expectedHash, hash)

	// Test case 2: Generation fails
	expectedError := "generation failed"

	// Mocking the Generate function of defaultGAC package
	defaultGAC.Generate = func(password, salt string) (string, error) {
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
