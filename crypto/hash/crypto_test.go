/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides hash functions for password encryption and comparison.
package hash

import (
	"testing"

	"github.com/origadmin/toolkits/crypto/hash/constants"
)

const (
	origin = "OrigAdmin@123456"
)

var (
	slatKey = []byte("key")
)

func TestGeneratePassword(t *testing.T) {
	crypto, err := NewCrypto(constants.MD5)
	if err != nil {
		t.Error("NewCrypto Failed: ", err.Error())
		return
	}

	hashPwd, err := crypto.Hash(origin)
	if err != nil {
		t.Error("Hash Failed: ", err.Error())
		return
	}

	if err := crypto.Verify(hashPwd, origin); err != nil {
		t.Error("Verify Failed: ", err.Error())
	}
}

func TestMD5(t *testing.T) {
	crypto, err := NewCrypto(constants.MD5)
	if err != nil {
		t.Error("NewCrypto Failed: ", err.Error())
		return
	}

	hashPwd, err := crypto.HashWithSalt(origin, slatKey)
	if err != nil {
		t.Error("HashWithSalt Failed: ", err.Error())
		return
	}

	if err := crypto.Verify(hashPwd, origin); err != nil {
		t.Error("Verify Failed: ", err.Error())
	}
}

func TestSHA1(t *testing.T) {
	crypto, err := NewCrypto(constants.SHA1)
	if err != nil {
		t.Error("NewCrypto Failed: ", err.Error())
		return
	}

	hashPwd, err := crypto.HashWithSalt(origin, slatKey)
	if err != nil {
		t.Error("HashWithSalt Failed: ", err.Error())
		return
	}

	if err := crypto.Verify(hashPwd, origin); err != nil {
		t.Error("Verify Failed: ", err.Error())
	}
}

func TestSHA256(t *testing.T) {
	crypto, err := NewCrypto(constants.SHA256)
	if err != nil {
		t.Error("NewCrypto Failed: ", err.Error())
		return
	}

	hashPwd, err := crypto.HashWithSalt(origin, slatKey)
	if err != nil {
		t.Error("HashWithSalt Failed: ", err.Error())
		return
	}

	if err := crypto.Verify(hashPwd, origin); err != nil {
		t.Error("Verify Failed: ", err.Error())
	}
}
