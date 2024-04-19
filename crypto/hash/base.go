// Copyright (c) 2024 GodCong. All rights reserved.

// Package hash provides the hash functions
package hash

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
)

// MD5 hash
//
//nolint:gosec
func MD5(b []byte) string {
	hasher := md5.New()
	_, _ = hasher.Write(b)
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
	// return fmt.Sprintf("%x", hasher.Sum(nil))
}

// MD5String md5 hash
func MD5String(s string) string {
	return MD5([]byte(s))
}

// SHA1 sha1 hash
//
//nolint:gosec
func SHA1(b []byte) string {
	hasher := sha1.New()
	_, _ = hasher.Write(b)
	hashBytes := hasher.Sum(nil)
	return hex.EncodeToString(hashBytes)
	// return fmt.Sprintf("%x", hasher.Sum(nil))
}

// SHA1String sha1 hash
func SHA1String(s string) string {
	return SHA1([]byte(s))
}
