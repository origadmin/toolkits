/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package hash

import (
	"testing"
)

const (
	origin  = "OrigAdmin@123456"
	slatKey = "key"
)

func TestGeneratePassword(t *testing.T) {
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
	hashVal := "06f684620c2e8f7caf9bb5a4fcba2ff2"
	if v := MD5String(origin); v != hashVal {
		t.Error("Failed to generate MD5 hash: ", v)
	}
}

func TestSHA1(t *testing.T) {
	hashVal := "6d94221e0f42005e332ff9b468614ebe798786c5"
	if v := SHA1String(origin); v != hashVal {
		t.Error("Failed to generate MD5 hash: ", v)
	}
}

func TestSHA256(t *testing.T) {
	hashVal := "7784a4de3b48eb1b5b562097ba46230e778a51147129f77c60a1294b411ead13"
	if v := SHA256String(origin); v != hashVal {
		t.Error("Failed to generate MD5 hash: ", v)
	}
}

func TestHMAC256(t *testing.T) {
	hashVal := "47f03a422b85f8bc524f283e78b70c9f026db157de8c21cf2330238cfb54cd56"
	if v := HMAC256String(origin, slatKey); v != hashVal {
		t.Error("Failed to generate HMAC256 hash: ", v)
	}
}
