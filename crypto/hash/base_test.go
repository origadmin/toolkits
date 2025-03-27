/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides hash functions for password encryption and comparison.
package hash

import (
	"testing"

	"github.com/origadmin/toolkits/crypto/hash/types"
)

const (
	origin  = "OrigAdmin@123456"
	slatKey = "key"
)

func TestGeneratePassword(t *testing.T) {
	crypto, err := NewCrypto(types.WithAlgorithm(types.TypeBcrypt))
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
	crypto, err := NewCrypto(types.WithAlgorithm(types.TypeMD5))
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
	crypto, err := NewCrypto(types.WithAlgorithm(types.TypeSHA1))
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
	crypto, err := NewCrypto(types.WithAlgorithm(types.TypeSHA256))
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

func TestHMAC256(t *testing.T) {
	crypto, err := NewCrypto(types.WithAlgorithm(types.TypeHMAC256))
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
