package hash

// Type hash type
type Type string

const (
	// TypeMD5 md5 type
	TypeMD5 Type = "md5"
	// TypeSHA1 sha1 type
	TypeSHA1 Type = "sha1"
	// TypeSHA256 sha256 type
	TypeSHA256 Type = "sha256"
	// TypeArgon2 argon2 type
	TypeArgon2 Type = "argon2"
	// TypeScrypt scrypt type
	TypeScrypt Type = "scrypt"
	// TypeBcrypt bcrypt type
	TypeBcrypt Type = "bcrypt"
	// TypeHMAC256 hmac sha256 type
	TypeHMAC256 Type = "hmac256"
)
