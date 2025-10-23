/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

import (
	"fmt"
	"strings"

	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
)

// Constants for hash algorithms and configurations.
const (
	// General Constants
	ENV                 = "ORIGADMIN_HASH_TYPE"
	DefaultType         = "argon2"
	DefaultVersion      = "v1"
	DefaultSaltLength   = 16
	DefaultTimeCost     = 3
	DefaultMemoryCost   = 64 * 1024 // 64MB
	DefaultThreads      = 4
	DefaultCost         = 10
	ParamSeparator      = ","
	ParamValueSeparator = ":"
	CodecSeparator      = "$"

	// Algorithm Names
	UNKNOWN = "unknown"

	// Standard Hashes
	MD5    = "md5"
	SHA1   = "sha1"
	SHA256 = "sha256"
	SHA384 = "sha384"
	SHA512 = "sha512"
	SHA224 = "sha224"

	// SHA-3 Hashes
	SHA3         = "sha3"
	SHA3_224     = "sha3-224"
	SHA3_256     = "sha3-256"
	SHA3_384     = "sha3-384"
	SHA3_512     = "sha3-512"
	SHA3_512_224 = "sha3-512-224"
	SHA3_512_256 = "sha3-512-256"

	// SHA512 Hashes (distinct from SHA-3 variants)
	SHA512_224 = "sha512/224" // Renamed to avoid conflict with SHA3_512_224
	SHA512_256 = "sha512/256" // Renamed to avoid conflict with SHA3_512_256

	// BLAKE2 Hashes
	BLAKE2s_128     = "blake2s-128"
	BLAKE2s_256     = "blake2s-256"
	BLAKE2b_256     = "blake2b-256"
	BLAKE2b_384     = "blake2b-384"
	BLAKE2b_512     = "blake2b-512"
	DEFAULT_BLAKE2b = BLAKE2b_256
	DEFAULT_BLAKE2s = BLAKE2s_256

	// Base Algorithm Names (for composite algorithms)
	HMAC       = "hmac"
	PBKDF2     = "pbkdf2"
	SCRYPT     = "scrypt"
	BCRYPT     = "bcrypt"
	ARGON2     = "argon2"
	ARGON2i    = "argon2i"
	ARGON2id   = "argon2id"
	BLAKE2b    = "blake2b"
	BLAKE2s    = "blake2s"
	RIPEMD     = "ripemd"
	RIPEMD160  = "ripemd-160"
	CRC32      = "crc32"
	CRC32_ISO  = "crc32-iso"
	CRC32_CAST = "crc32-cast"
	CRC32_KOOP = "crc32-koop"
	CRC64      = "crc64"
	CRC64_ISO  = "crc64-iso"
	CRC64_ECMA = "crc64-ecma"

	// Composite Algorithm Identifiers (HMAC)
	HMAC_SHA1     = HMAC + "-" + SHA1
	HMAC_SHA256   = HMAC + "-" + SHA256
	HMAC_SHA384   = HMAC + "-" + SHA384
	HMAC_SHA512   = HMAC + "-" + SHA512
	HMAC_SHA3_224 = HMAC + "-" + SHA3_224
	HMAC_SHA3_256 = HMAC + "-" + SHA3_256
	HMAC_SHA3_384 = HMAC + "-" + SHA3_384
	HMAC_SHA3_512 = HMAC + "-" + SHA3_512
	DEFAULT_HMAC  = HMAC_SHA256
	HMAC_PREFIX   = HMAC + "-"

	// Composite Algorithm Identifiers (PBKDF2)
	PBKDF2_SHA1     = PBKDF2 + "-" + SHA1
	PBKDF2_SHA256   = PBKDF2 + "-" + SHA256
	PBKDF2_SHA384   = PBKDF2 + "-" + SHA384
	PBKDF2_SHA512   = PBKDF2 + "-" + SHA512
	PBKDF2_SHA3_224 = PBKDF2 + "-" + SHA3_224
	PBKDF2_SHA3_256 = PBKDF2 + "-" + SHA3_256
	PBKDF2_SHA3_384 = PBKDF2 + "-" + SHA3_384
	PBKDF2_SHA3_512 = PBKDF2 + "-" + SHA3_512
	DEFAULT_PBKDF2  = PBKDF2_SHA256
	PBKDF2_PREFIX   = PBKDF2 + "-"
)

// AlgorithmResolver defines an interface for resolving and normalizing algorithm types.
type AlgorithmResolver interface {
	ResolveType(t Type) (Type, error)
}

// Type represents a structured hash algorithm definition.
// It separates the main algorithm from its underlying hash function,
// allowing for clear and extensible handling of composite algorithms
// like HMAC and PBKDF2.
type Type struct {
	// Name is the main algorithm's name, e.g., "hmac", "pbkdf2", "sha256".
	// This is the key field for logical dispatch.
	Name string

	// Underlying is the full string representation of the underlying hash algorithm.
	// For simple hashes, this field is empty.
	// For composite hashes, it specifies the hash function to be used,
	// e.g., "sha256" for "hmac-sha256", or "sha3-512" for "pbkdf2-sha3-512".
	Underlying string
}

// String returns the string representation of the type
func (t Type) String() string {
	if t.Underlying != "" {
		return t.Name + "-" + t.Underlying
	}
	return t.Name
}

// Is compares two Type instances for equality.
func (t Type) Is(t2 Type) bool {
	return t.Name == t2.Name && t.Underlying == t2.Underlying
}

// ParseType parses an algorithm string into its structured Type.
// It handles common aliases and composite algorithm formats.
func ParseType(algorithm string) (Type, error) {
	algorithm = strings.ToLower(algorithm)

	parts := strings.SplitN(algorithm, "-", 2)
	var t Type
	if len(parts) == 2 {
		// This is a composite algorithm like "hmac-sha256" or "pbkdf2-sha512"
		t = Type{Name: parts[0], Underlying: parts[1]}
	} else {
		// This is a simple algorithm like "sha256"
		t = Type{Name: algorithm}
	}

	// If no specific resolver is registered, return the parsed type as is.
	return t, nil
}

// NewType creates a new Type instance with the specified name and underlying hash.
func NewType(name string, underlying ...string) Type {
	t := Type{Name: name}
	if len(underlying) > 0 {
		t.Underlying = underlying[0]
	}
	return t
}

// TypeHash is a helper function that might need to be refactored
// depending on how stdhash.ParseHash is updated to handle the new Type struct.
// For now, it assumes subAlg is a simple string.
func TypeHash(subAlg string) (stdhash.Hash, error) {
	h, err := stdhash.ParseHash(subAlg)
	if err != nil {
		return 0, fmt.Errorf("unsupported hash type: %s", subAlg)
	}
	return h, nil
}
