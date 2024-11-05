// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package hash provides the hash functions
package hash

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 hash
func MD5(b []byte) []byte {
	hashBytes := md5.Sum(b)
	return hashBytes[:]
}

// MD5String md5 hash
func MD5String(s string) string {
	return hex.EncodeToString(MD5([]byte(s)))
}

// SHA1 sha1 hash
func SHA1(b []byte) []byte {
	hashBytes := sha1.Sum(b)
	return hashBytes[:]
}

// SHA1String sha1 hash
func SHA1String(s string) string {
	return hex.EncodeToString(SHA1([]byte(s)))
}

// SHA256 sha256 hash
func SHA256(b []byte) []byte {
	hashBytes := sha256.Sum256(b)
	return hashBytes[:]
}

// SHA256String sha256 hash
func SHA256String(s string) string {
	return hex.EncodeToString(SHA256([]byte(s)))
}

// HMAC256 hmac256 hash
func HMAC256(b []byte, key []byte) []byte {
	hashBytes := hmac.New(sha256.New, key)
	hashBytes.Write(b)
	return hashBytes.Sum(nil)
}

// HMAC256String hmac256 hash
func HMAC256String(s string, key string) string {
	return hex.EncodeToString(HMAC256([]byte(s), []byte(key)))
}
