/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides hash functions for password encryption and comparison.
package hash

import (
	"fmt"
	"strings"
	"testing"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
)

const (
	origin  = "OrigAdmin@123456"
	slatKey = "key"
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

func TestHMAC(t *testing.T) {
	global, err := NewCrypto("hmac")
	if err != nil {
		t.Error("NewCrypto Failed: ", err.Error())
		return
	}
	for i := stdhash.Hash(1); i < stdhash.MAPHASH; i++ {
		t.Logf("test hash:%s starting", i.String())
		crypto, err := NewCrypto(fmt.Sprintf("hmac-%s", i.String()))
		if err != nil {
			t.Error("NewCrypto Failed: ", err.Error())
			return
		}
		crypto = CachedCrypto(crypto)
		t.Run("HashWithSalt", func(t *testing.T) {
			hashPwd, err := crypto.HashWithSalt(origin, slatKey)
			if err != nil {
				t.Error("HashWithSalt Failed: ", err.Error())
				return
			}
			t.Logf("hashPwd:%s", hashPwd)
			if err := global.Verify(hashPwd, origin); err != nil {
				t.Error("Verify Failed: ", err.Error())
			}
		})

		t.Run("Hash", func(t *testing.T) {
			hashPwd, err := crypto.Hash(origin)
			if err != nil {
				t.Error("Hash Failed: ", err.Error())
				return
			}
			t.Logf("hashPwd:%s", hashPwd)
			if err := global.Verify(hashPwd, origin); err != nil {
				t.Error("Verify Failed: ", err.Error())
			}
		})
	}
}

func TestUnsupportedAlgorithms(t *testing.T) {
	tests := []struct {
		name        string
		algorithm   string
		expectedErr string
	}{
		{"MapHash", "maphash", "cannot compare hash with maphash"},
	}

	for _, tt := range tests {
		crypto, err := NewCrypto(fmt.Sprintf("hmac-%s", tt.algorithm))
		if err != nil {
			t.Error("NewCrypto Failed: ", err.Error())
			return
		}
		crypto = CachedCrypto(crypto)
		t.Run("HashWithSalt", func(t *testing.T) {
			hashPwd, err := crypto.HashWithSalt(origin, slatKey)
			if err != nil {
				t.Error("HashWithSalt Failed: ", err.Error())
				return
			}
			t.Logf("hashPwd:%s", hashPwd)
			if err := crypto.Verify(hashPwd, origin); err == nil || !strings.Contains(err.Error(), tt.expectedErr) {
				t.Errorf("expected error: %s, but got %v", tt.expectedErr, err)
			}
		})
	}
}
