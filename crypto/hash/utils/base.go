/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package utils

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 calculates MD5 hash
func MD5(b []byte) []byte {
	hashBytes := md5.Sum(b)
	return hashBytes[:]
}

// MD5String calculates MD5 hash and returns hex string
func MD5String(s string) string {
	return hex.EncodeToString(MD5([]byte(s)))
}

// SHA1 calculates SHA1 hash
func SHA1(b []byte) []byte {
	hashBytes := sha1.Sum(b)
	return hashBytes[:]
}

// SHA1String calculates SHA1 hash and returns hex string
func SHA1String(s string) string {
	return hex.EncodeToString(SHA1([]byte(s)))
}

// SHA256 calculates SHA256 hash
func SHA256(b []byte) []byte {
	hashBytes := sha256.Sum256(b)
	return hashBytes[:]
}

// SHA256String calculates SHA256 hash and returns hex string
func SHA256String(s string) string {
	return hex.EncodeToString(SHA256([]byte(s)))
}

// HMAC256 calculates HMAC-SHA256 hash
func HMAC256(b []byte, key []byte) []byte {
	hashBytes := hmac.New(sha256.New, key)
	hashBytes.Write(b)
	return hashBytes.Sum(nil)
}

// HMAC256String calculates HMAC-SHA256 hash and returns hex string
func HMAC256String(s string, key string) string {
	return hex.EncodeToString(HMAC256([]byte(s), []byte(key)))
}
