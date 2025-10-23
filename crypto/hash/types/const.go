/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

// ============================================================================ //
//                                 CONSTANTS
// ============================================================================ //

// This file defines constants related to hash algorithms and configurations.
// It includes general configuration constants, base algorithm names,
// composite algorithm identifiers, and default values.

// ---------------------------------------------------------------------------- //
//                       General Configuration Constants
// ---------------------------------------------------------------------------- //

const (
	// ENV is the environment variable name for hash type.
	ENV = "ORIGADMIN_HASH_TYPE"
	// DefaultType is the default hash type.
	DefaultType = "argon2"
	// DefaultVersion is the default hash version.
	DefaultVersion = "v1"
	// DefaultSaltLength is the default salt length.
	DefaultSaltLength = 16
	// DefaultTimeCost is the default time cost for Argon2.
	DefaultTimeCost = 3
	// DefaultMemoryCost is the default memory cost for Argon2.
	DefaultMemoryCost = 64 * 1024 // 64MB
	// DefaultThreads is the default number of threads for Argon2.
	DefaultThreads = 4
	// DefaultCost is the default cost for bcrypt.
	DefaultCost = 10

	// ParamSeparator is the separator for parameters in a hash string.
	ParamSeparator = ","
	// ParamValueSeparator is the separator for parameter key-value pairs.
	ParamValueSeparator = ":"
	// CodecSeparator is the separator used in the encoded hash string.
	CodecSeparator = "$"
)

// ---------------------------------------------------------------------------- //
//                            Base Algorithm Names
// ---------------------------------------------------------------------------- //

const (
	// UNKNOWN represents an unknown algorithm.
	UNKNOWN = "unknown"

	// MD5 is the MD5 hash algorithm.
	MD5 = "md5"
	// SHA1 is the SHA-1 hash algorithm.
	SHA1 = "sha1"
	// SHA224 is the SHA-224 hash algorithm.
	SHA224 = "sha224"
	// SHA256 is the SHA-256 hash algorithm.
	SHA256 = "sha256"
	// SHA384 is the SHA-384 hash algorithm.
	SHA384 = "sha384"
	// SHA512 is the SHA-512 hash algorithm.
	SHA512 = "sha512"

	// SHA3 is the SHA-3 hash algorithm.
	SHA3 = "sha3"
	// SHA3_224 is the SHA-3-224 hash algorithm.
	SHA3_224 = "sha3-224"
	// SHA3_256 is the SHA-3-256 hash algorithm.
	SHA3_256 = "sha3-256"
	// SHA3_384 is the SHA-3-384 hash algorithm.
	SHA3_384 = "sha3-384"
	// SHA3_512 is the SHA-3-512 hash algorithm.
	SHA3_512 = "sha3-512"
	// SHA3_512_224 is the SHA-3-512/224 hash algorithm.
	SHA3_512_224 = "sha3-512-224"
	// SHA3_512_256 is the SHA-3-512/256 hash algorithm.
	SHA3_512_256 = "sha3-512-256"

	// SHA512_224 is the truncated SHA512/224 hash algorithm.
	SHA512_224 = "sha512/224"
	// SHA512_256 is the truncated SHA512/256 hash algorithm.
	SHA512_256 = "sha512/256"

	// BLAKE2b is the BLAKE2b hash algorithm.
	BLAKE2b = "blake2b"
	// BLAKE2s is the BLAKE2s hash algorithm.
	BLAKE2s = "blake2s"
	// BLAKE2b_256 is the BLAKE2b-256 hash algorithm.
	BLAKE2b_256 = "blake2b-256"
	// BLAKE2b_384 is the BLAKE2b-384 hash algorithm.
	BLAKE2b_384 = "blake2b-384"
	// BLAKE2b_512 is the BLAKE2b-512 hash algorithm.
	BLAKE2b_512 = "blake2b-512"
	// BLAKE2s_128 is the BLAKE2s-128 hash algorithm.
	BLAKE2s_128 = "blake2s-128"
	// BLAKE2s_256 is the BLAKE2s-256 hash algorithm.
	BLAKE2s_256 = "blake2s-256"
	// DefaultBLAKE2b is the default BLAKE2b hash algorithm.
	DefaultBLAKE2b = BLAKE2b_512
	// DefaultBLAKE2s is the default BLAKE2s hash algorithm.
	DefaultBLAKE2s = BLAKE2s_256

	// ARGON2 is the Argon2 password hashing algorithm.
	ARGON2 = "argon2"
	// ARGON2i is the Argon2i password hashing algorithm.
	ARGON2i = "argon2i"
	// ARGON2id is the Argon2id password hashing algorithm.
	ARGON2id = "argon2id"
	// BCRYPT is the Bcrypt password hashing algorithm.
	BCRYPT = "bcrypt"
	// SCRYPT is the Scrypt password hashing algorithm.
	SCRYPT = "scrypt"

	// HMAC is the HMAC message authentication code algorithm.
	HMAC = "hmac"
	// PBKDF2 is the PBKDF2 key derivation function.
	PBKDF2 = "pbkdf2"

	// RIPEMD is the RIPEMD hash algorithm family.
	RIPEMD = "ripemd"
	// RIPEMD160 is the RIPEMD-160 hash algorithm.
	RIPEMD160 = "ripemd-160"
	// CRC32 is the CRC32 checksum algorithm.
	CRC32 = "crc32"
	// CRC32_ISO is the CRC32-ISO checksum algorithm.
	CRC32_ISO = "crc32-iso"
	// CRC32_CAST is the CRC32-CAST checksum algorithm.
	CRC32_CAST = "crc32-cast"
	// CRC32_KOOP is the CRC32-KOOP checksum algorithm.
	CRC32_KOOP = "crc32-koop"
	// CRC64 is the CRC64 checksum algorithm.
	CRC64 = "crc64"
	// CRC64_ISO is the CRC64-ISO checksum algorithm.
	CRC64_ISO = "crc64-iso"
	// CRC64_ECMA is the CRC64-ECMA checksum algorithm.
	CRC64_ECMA = "crc64-ecma"
)

// ---------------------------------------------------------------------------- //
//                       Composite Algorithm Identifiers
// ---------------------------------------------------------------------------- //

const (
	// HMAC_SHA1 is the HMAC-SHA1 composite algorithm.
	HMAC_SHA1 = HMAC + "-" + SHA1
	// HMAC_SHA256 is the HMAC-SHA256 composite algorithm.
	HMAC_SHA256 = HMAC + "-" + SHA256
	// HMAC_SHA384 is the HMAC-SHA384 composite algorithm.
	HMAC_SHA384 = HMAC + "-" + SHA384
	// HMAC_SHA512 is the HMAC-SHA512 composite algorithm.
	HMAC_SHA512 = HMAC + "-" + SHA512
	// HMAC_SHA3_224 is the HMAC-SHA3-224 composite algorithm.
	HMAC_SHA3_224 = HMAC + "-" + SHA3_224
	// HMAC_SHA3_256 is the HMAC-SHA3-256 composite algorithm.
	HMAC_SHA3_256 = HMAC + "-" + SHA3_256
	// HMAC_SHA3_384 is the HMAC-SHA3-384 composite algorithm.
	HMAC_SHA3_384 = HMAC + "-" + SHA3_384
	// HMAC_SHA3_512 is the HMAC-SHA3-512 composite algorithm.
	HMAC_SHA3_512 = HMAC + "-" + SHA3_512
	// DefaultHMAC is the default HMAC composite algorithm.
	DefaultHMAC = HMAC_SHA256
	// HMAC_PREFIX is the prefix for HMAC composite algorithms.
	HMAC_PREFIX = HMAC + "-"

	// PBKDF2_SHA1 is the PBKDF2-SHA1 composite algorithm.
	PBKDF2_SHA1 = PBKDF2 + "-" + SHA1
	// PBKDF2_SHA256 is the PBKDF2-SHA256 composite algorithm.
	PBKDF2_SHA256 = PBKDF2 + "-" + SHA256
	// PBKDF2_SHA384 is the PBKDF2-SHA384 composite algorithm.
	PBKDF2_SHA384 = PBKDF2 + "-" + SHA384
	// PBKDF2_SHA512 is the PBKDF2-SHA512 composite algorithm.
	PBKDF2_SHA512 = PBKDF2 + "-" + SHA512
	// PBKDF2_SHA3_224 is the PBKDF2-SHA3-224 composite algorithm.
	PBKDF2_SHA3_224 = PBKDF2 + "-" + SHA3_224
	// PBKDF2_SHA3_256 is the PBKDF2-SHA3-256 composite algorithm.
	PBKDF2_SHA3_256 = PBKDF2 + "-" + SHA3_256
	// PBKDF2_SHA3_384 is the PBKDF2-SHA3-384 composite algorithm.
	PBKDF2_SHA3_384 = PBKDF2 + "-" + SHA3_384
	// PBKDF2_SHA3_512 is the PBKDF2-SHA3-512 composite algorithm.
	PBKDF2_SHA3_512 = PBKDF2 + "-" + SHA3_512
	// DefaultPBKDF2 is the default PBKDF2 composite algorithm.
	DefaultPBKDF2 = PBKDF2_SHA256
	// PBKDF2_PREFIX is the prefix for PBKDF2 composite algorithms.
	PBKDF2_PREFIX = PBKDF2 + "-"
)
