/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides the hash functions
package hash

import (
	"encoding/hex"
	"os"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/scrypt"

	"github.com/origadmin/toolkits/errors"
)

const (
	// ENV environment variable name
	ENV = "ORIG_ADMIN_HASH_TYPE"
	// ErrPasswordNotMatch error when password not match
	ErrPasswordNotMatch = errors.String("password not match")
)

// GenerateFunc generate password hash function
type GenerateFunc func(password string, salt string) (string, error)

// CompareFunc compare password hash function
type CompareFunc func(hashpass, password, salt string) error

// GenerateAndCompare hash type Generate and Compare functions
type GenerateAndCompare struct {
	Generate GenerateFunc
	Compare  CompareFunc
}

var (
	hashes = map[Type]GenerateAndCompare{
		TypeMD5:     {Generate: GenerateMD5Password, Compare: CompareMD5HashAndPassword},
		TypeSHA1:    {Generate: GenerateSHA1Password, Compare: CompareSHA1HashAndPassword},
		TypeSHA256:  {Generate: GenerateSHA256Password, Compare: CompareSHA256HashAndPassword},
		TypeArgon2:  {Generate: GenerateArgon2Password, Compare: CompareArgon2HashAndPassword},
		TypeScrypt:  {Generate: GenerateScryptPassword, Compare: CompareScryptHashAndPassword},
		TypeBcrypt:  {Generate: GenerateBcryptPassword, Compare: CompareBcryptHashAndPassword},
		TypeHMAC256: {Generate: GenerateHMAC256Password, Compare: CompareHMAC256HashAndPassword},
	}

	// defaultGAC default generate and compare function
	defaultGAC GenerateAndCompare
)

func init() {
	defaultGAC = GenerateAndCompare{
		Generate: GenerateMD5Password,
		Compare:  CompareMD5HashAndPassword,
	}

	env := os.Getenv(ENV)
	if env == "" {
		return
	}

	if v, ok := hashes[Type(env)]; ok {
		defaultGAC = v
	}
}

// UseCrypto updates the default generate and compare functions based on the hash type provided.
//
// t Type - the hash type to update the functions with.
func UseCrypto(t Type) {
	if v, ok := hashes[t]; ok {
		defaultGAC = v
	}
}

// UseMD5 sets the global hash compare function to the one corresponding to the TypeMD5 hash type.
func UseMD5() {
	defaultGAC = hashes[TypeMD5]
}

// UseSHA1 sets the global hash compare function to the one corresponding to the TypeSHA1 hash type.
func UseSHA1() {
	defaultGAC = hashes[TypeSHA1]
}

// UseSHA256 sets the global hash compare function to the one corresponding to the TypeSHA256 hash type.
func UseSHA256() {
	defaultGAC = hashes[TypeSHA256]
}

// UseBcrypt sets the global hash compare function to the one corresponding to the TypeBcrypt hash type.
func UseBcrypt() {
	defaultGAC = hashes[TypeBcrypt]
}

// UseScrypt sets the global hash compare function to the one corresponding to the TypeScrypt hash type.
func UseScrypt() {
	defaultGAC = hashes[TypeScrypt]
}

// UseArgon2 sets the global hash compare function to the one corresponding to the TypeArgon2 hash type.
func UseArgon2() {
	defaultGAC = hashes[TypeArgon2]
}

// Generate generates a password hash using the global hash compare function.
//
// It takes in two parameters:
// - password: a string representing the password to be hashed.
// - salt: a string representing the salt value to be used in the hashing process.
//
// It returns a string representing the generated password hash and an error if the generation fails.
func Generate(password string, salt string) (string, error) {
	return defaultGAC.Generate(password, salt)
}

// Compare compares the given hashed password, password, and salt with the global
// hash compare function. It returns an error if the comparison fails.
//
// Parameters:
// - hashpass: the hashed password to compare.
// - password: the plaintext password to compare.
// - salt: the salt to use for the comparison.
//
// Returns:
// - error: an error if the comparison fails, otherwise nil.
func Compare(hashpass, password, salt string) error {
	return defaultGAC.Compare(hashpass, password, salt)
}

// GenerateSHA1Password generates a SHA1 password hash using the provided password and salt.
//
// Parameters:
// - password: the plaintext password to be hashed.
// - salt: the salt value to be used for hashing.
//
// Returns:
// - string: the SHA1 hashed password.
// - error: nil if the hash generation is successful, otherwise an error.
func GenerateSHA1Password(password string, salt string) (string, error) {
	return SHA1String(salt + password + salt), nil
}

// CompareSHA1HashAndPassword compares a given SHA1 hashed password with a plaintext password
// using a provided salt value. It returns an error if the comparison fails.
//
// Parameters:
// - hashpass: the hashed password to compare.
// - password: the plaintext password to compare.
// - salt: the salt to use for the comparison.
//
// Returns:
// - error: an error if the comparison fails, otherwise nil.
func CompareSHA1HashAndPassword(hashpass string, password string, salt string) error {
	pass := SHA1String(salt + password + salt)
	if hashpass != pass {
		return ErrPasswordNotMatch
	}
	return nil
}

// GenerateSHA256Password generates an SHA256 password hash using the provided password and salt.
//
// Parameters:
// - password: the plaintext password to be hashed.
// - salt: the salt value to be used for hashing.
//
// Returns:
// - string: the MD5 hashed password.
// - error: nil if the hash generation is successful, otherwise an error.
func GenerateSHA256Password(password string, salt string) (string, error) {
	return SHA256String(salt + password + salt), nil
}

// CompareSHA256HashAndPassword compares an SHA256 hashed password with a plaintext password using a provided salt value.
//
// Parameters:
// - hashpass: the hashed password to compare.
// - password: the plaintext password to compare.
// - salt: the salt value to use for the comparison.
//
// Returns:
// - error: an error if the comparison fails, otherwise nil.
func CompareSHA256HashAndPassword(hashpass string, password string, salt string) error {
	pass := SHA256String(salt + password + salt)
	if hashpass != pass {
		return ErrPasswordNotMatch
	}
	return nil
}

// GenerateHMAC256Password generates an HMAC256 password hash using the provided password and salt.
//
// Parameters:
// - password: the plaintext password to be hashed.
// - salt: the salt value to be used for hashing.
//
// Returns:
// - string: the MD5 hashed password.
// - error: nil if the hash generation is successful, otherwise an error.
func GenerateHMAC256Password(password string, salt string) (string, error) {
	return HMAC256String(password, salt), nil
}

// CompareHMAC256HashAndPassword compares an HMAC256 hashed password with a plaintext password using a provided salt value.
//
// Parameters:
// - hashpass: the hashed password to compare.
// - password: the plaintext password to compare.
// - salt: the salt value to use for the comparison.
//
// Returns:
// - error: an error if the comparison fails, otherwise nil.
func CompareHMAC256HashAndPassword(hashpass string, password string, salt string) error {
	pass := HMAC256String(password, salt)
	if hashpass != pass {
		return ErrPasswordNotMatch
	}
	return nil
}

// GenerateMD5Password generates an MD5 password hash using the provided password and salt.
//
// Parameters:
// - password: the plaintext password to be hashed.
// - salt: the salt value to be used for hashing.
//
// Returns:
// - string: the MD5 hashed password.
// - error: nil if the hash generation is successful, otherwise an error.
func GenerateMD5Password(password string, salt string) (string, error) {
	return MD5String(salt + password + salt), nil
}

// CompareMD5HashAndPassword compares an MD5 hashed password with a plaintext password using a provided salt value.
//
// Parameters:
// - hashpass: the hashed password to compare.
// - password: the plaintext password to compare.
// - salt: the salt value to use for the comparison.
//
// Returns:
// - error: an error if the comparison fails, otherwise nil.
func CompareMD5HashAndPassword(hashpass, password, salt string) error {
	pass := MD5String(salt + password + salt)
	if hashpass != pass {
		return ErrPasswordNotMatch
	}
	return nil
}

// GenerateBcryptPassword generates a bcrypt hash for the given password.
//
// password: the password to hash.
// _: a placeholder for unused variable.
// (string, error): the generated bcrypt hash and any error encountered.
func GenerateBcryptPassword(password string, _ string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// CompareBcryptHashAndPassword compares a bcrypt hashed password with a plaintext password.
//
// Parameters:
// - hashpass: The hashed password to compare.
// - password: The plaintext password to compare.
// - _: Unused parameter.
//
// Returns:
// - error: An error if the hashed password does not match the plaintext password.
func CompareBcryptHashAndPassword(hashpass, password, _ string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashpass), []byte(password))
}

// GenerateScryptPassword generates a password hash using the scrypt algorithm.
//
// Parameters:
// - password: the password to hash (string).
// - salt: the salt value to use for hashing (string).
//
// Returns:
// - string: the generated password hash (hex-encoded).
// - error: an error if the hashing process fails.
func GenerateScryptPassword(password, salt string) (string, error) {
	var rb []byte
	rb, err := scrypt.Key([]byte(password), []byte(salt), 1<<15, 8, 1, 32)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(rb), nil
}

// CompareScryptHashAndPassword compares a scrypt hashed password with a plaintext password using a provided salt value.
//
// Parameters:
// - hashpass: the hashed password to compare.
// - password: the plaintext password to compare.
// - salt: the salt value to use for the comparison.
//
// Returns:
// - error: an error if the comparison fails, otherwise nil.
func CompareScryptHashAndPassword(hashpass, password string, salt string) error {
	pass, err := GenerateScryptPassword(password, salt)
	if err != nil {
		return err
	}
	if hashpass != pass {
		return ErrPasswordNotMatch
	}
	return nil
}

// GenerateArgon2Password generates an Argon2 hashed password based on the provided password and salt.
//
// Parameters:
// - password: the password to hash.
// - salt: the salt value to use for hashing.
// Return type(s): a string containing the hashed password and an error.
func GenerateArgon2Password(password, salt string) (string, error) {
	pass := argon2.Key([]byte(password), []byte(salt), 3, 32*1024, 4, 32)
	return hex.EncodeToString(pass), nil
}

// CompareArgon2HashAndPassword compares an Argon2 hashed password with a plaintext password using a provided salt value.
//
// Parameters:
// - hashpass: the hashed password to compare.
// - password: the plaintext password to compare.
// - salt: the salt value to use for the comparison.
// Returns:
// - error: an error if the comparison fails, otherwise nil.
func CompareArgon2HashAndPassword(hashpass, password string, salt string) error {
	pass, _ := GenerateArgon2Password(password, salt)
	if hashpass != pass {
		return ErrPasswordNotMatch
	}
	return nil
}
