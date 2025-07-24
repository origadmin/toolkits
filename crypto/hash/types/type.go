/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

import (
	"strings"
)

// Type represents the hash algorithm type
type Type string

const (
	// TypeMD5 md5 type
	TypeMD5 Type = "md5"
	// TypeSha1 sha1 type
	TypeSha1 Type = "sha1"
	// TypeSha224 sha224 type
	TypeSha224 Type = "sha224"
	// TypeSha256 sha256 type
	TypeSha256 Type = "sha256"
	// TypeSha384 sha384 type
	TypeSha384 Type = "sha384"
	// TypeSha512 sha512 type
	TypeSha512 Type = "sha512"
	// TypeSha3224 sha3-224 type
	TypeSha3224 Type = "sha3-224"
	// TypeSha3256 sha3-256 type
	TypeSha3256 Type = "sha3-256"
	// TypeSha3384 sha3-384 type
	TypeSha3384 Type = "sha3-384"
	// TypeSha3512 sha3-512 type
	TypeSha3512 Type = "sha3-512"
	// TypeSha3512224 sha3-512-224 type
	TypeSha3512224 Type = "sha3-512-224"
	// TypeSha3512256 sha3-512-256 type
	TypeSha3512256 Type = "sha3-512-256"
	// TypeArgon2 argon2 type
	TypeArgon2 Type = "argon2"
	// TypeScrypt scrypt type
	TypeScrypt Type = "scrypt"
	// TypeBcrypt bcrypt type
	TypeBcrypt Type = "bcrypt"
	// TypeHMAC hmac type
	TypeHMAC Type = "hmac"
	// TypeHMAC256 hmac sha256 type
	TypeHMAC256 Type = "hmac256"
	// TypeHMAC512 hmac sha512 type
	TypeHMAC512 Type = "hmac512"
	// TypePBKDF2 pbkdf2 type
	TypePBKDF2 Type = "pbkdf2"
	// TypePBKDF2SHA256 pbkdf2 sha256 type
	TypePBKDF2SHA256 Type = "pbkdf2-sha256"
	// TypePBKDF2SHA512 pbkdf2 sha512 type
	TypePBKDF2SHA512 Type = "pbkdf2-sha512"
	// TypePBKDF2SHA384 pbkdf2 sha384 type
	TypePBKDF2SHA384 Type = "pbkdf2-sha384"
	// TypePBKDF2SHA3256 pbkdf2 sha3-256 type
	TypePBKDF2SHA3256 Type = "pbkdf2-sha3-256"
	// TypePBKDF2SHA3224 pbkdf2 sha3-224 type
	TypePBKDF2SHA3224 Type = "pbkdf2-sha3-224"
	// TypePBKDF2SHA3384 pbkdf2 sha3-384 type
	TypePBKDF2SHA3384 Type = "pbkdf2-sha3-384"
	// TypePBKDF2SHA3512224 pbkdf2 sha3-512-224 type
	TypePBKDF2SHA3512224 Type = "pbkdf2-sha3-512-224"
	// TypePBKDF2SHA3512256 pbkdf2 sha3-512-256 type
	TypePBKDF2SHA3512256 Type = "pbkdf2-sha3-512-256"
	// TypePBKDF2SHA3512384 pbkdf2 sha3-512-384 type
	TypePBKDF2SHA3512384 Type = "pbkdf2-sha3-512-384"
	// TypePBKDF2SHA3512512 pbkdf2 sha3-512-512 type
	TypePBKDF2SHA3512512 Type = "pbkdf2-sha3-512-512"
	// TypeBlake2b blake2b type
	TypeBlake2b Type = "blake2b"
	// TypeBlake2s blake2s type
	TypeBlake2s Type = "blake2s"
	// TypeUnknown unknown type
	TypeUnknown Type = "unknown"
)

// String returns the string representation of the type
func (t Type) String() string {
	return string(t)
}

func (t Type) Is(t2 Type) bool {
	return TypeIs(t, t2)
}

// ParseType parses a string into a Type
func ParseType(s string) Type {
	return Type(s)
}

func TypeIs(t Type, t2 Type) bool {
	if t == t2 {
		return true
	}
	if t == TypeUnknown || t2 == TypeUnknown {
		return false
	}
	return len(t) >= len(t2) && strings.HasPrefix(t.String(), t2.String())
}
