// Copyright (c) 2024 GodCong. All rights reserved.

// Package hash provides the hash functions
package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// MD5 hash
//
//nolint:gosec
func MD5(b []byte) []byte {
	hashBytes := md5.Sum(b)
	return hashBytes[:]
}

// MD5String md5 hash
func MD5String(s string) string {
	return hex.EncodeToString(MD5([]byte(s)))
}

// SHA1 sha1 hash
//
//nolint:gosec
func SHA1(b []byte) []byte {
	hashBytes := sha1.Sum(b)
	return hashBytes[:]
}

// SHA1String sha1 hash
//
//nolint:gosec
func SHA1String(s string) string {
	return hex.EncodeToString(SHA1([]byte(s)))
}

// SHA256 sha256 hash
//
//nolint:gosec
func SHA256(b []byte) []byte {
	hashBytes := sha256.Sum256(b)
	return hashBytes[:]
}

// SHA256String sha256 hash
//
//nolint:gosec
func SHA256String(s string) string {
	return hex.EncodeToString(SHA256([]byte(s)))
}
