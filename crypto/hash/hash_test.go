// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package hash provides the hash functions
package hash

import (
	"testing"

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
