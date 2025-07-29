package constants

// UNKNOWN Algorithm
const (
	UNKNOWN = "unknown"
)

// Standard Hashes
const (
	MD5    = "md5"
	SHA1   = "sha1"
	SHA256 = "sha256"
	SHA384 = "sha384"
	SHA512 = "sha512"
	SHA224 = "sha224"
)

// SHA-3 Hashes
const (
	SHA3_224     = "sha3-224"
	SHA3_256     = "sha3-256"
	SHA3_384     = "sha3-384"
	SHA3_512     = "sha3-512"
	SHA3_512_224 = "sha3-512-224"
	SHA3_512_256 = "sha3-512-256"
)

// SHA512 Hashes
const (
	SHA512_224 = "sha512-224"
	SHA512_256 = "sha512-256"
)

// BLAKE2 Hashes
const (
	BLAKE2s_128     = "blake2s-128"
	BLAKE2s_256     = "blake2s-256"
	BLAKE2b_256     = "blake2b-256"
	BLAKE2b_384     = "blake2b-384"
	BLAKE2b_512     = "blake2b-512"
	DEFAULT_BLAKE2b = BLAKE2b_256
	DEFAULT_BLAKE2s = BLAKE2s_256
)

// Base Algorithm Names (for composite algorithms)
const (
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
)

// Composite Algorithm Identifiers (HMAC)
const (
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
)

// Composite Algorithm Identifiers (PBKDF2)
const (
	PBKDF2_SHA1     = PBKDF2 + "-" + SHA1
	PBKDF2_SHA256   = PBKDF2 + "-" + SHA256
	PBKDF2_SHA384   = PBKDF2 + "-" + SHA384
	PBKDF2_SHA512   = PBKDF2 + "-" + SHA512
	PBKDF2_SHA3_224 = PBKDF2 + "-" + SHA3_224
	PBKDF2_SHA3_256 = PBKDF2 + "-" + SHA3_256
	PBKDF2_SHA3_384 = PBKDF2 + "-" + SHA3_384
	PBKDF2_SHA3_512 = PBKDF2 + "-" + SHA3_512
	DEFAULT_PBKDF2  = PBKDF2_SHA256
)
