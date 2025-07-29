/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package errors

import "github.com/origadmin/toolkits/errors"

var (
	// ErrPasswordNotMatch error when password does not match
	ErrPasswordNotMatch = errors.String("password does not match")
	// ErrAlgorithmMismatch error when algorithm mismatch
	ErrAlgorithmMismatch = errors.String("algorithm mismatch")
	// ErrInvalidHashFormat error when invalid hash format
	ErrInvalidHashFormat = errors.String("invalid hash format")
	// ErrSaltLengthTooShort error when salt length too short
	ErrSaltLengthTooShort = errors.String("salt length must be at least 8 bytes")
	// ErrCostOutOfRange error when cost out of range
	ErrCostOutOfRange = errors.String("cost must be between 4 and 31")
	// ErrHashModuleNotInitialized is returned when the hash module fails to initialize.
	ErrHashModuleNotInitialized = errors.String("hash module not initialized")
	// ErrUnsupportedHashForHMAC is returned when an unsupported hash type is used for HMAC.
	ErrUnsupportedHashForHMAC = errors.String("unsupported hash type for HMAC")
	// ErrInvalidAlgorithm is returned when an invalid algorithm is used.
	ErrInvalidAlgorithm = errors.String("invalid algorithm")
	// ErrResolverNotRegistered is returned when a resolver is not registered.
	ErrResolverNotRegistered = errors.String("resolver not registered")
	// ErrInvalidHash is returned when the provided hash string is invalid.
	ErrInvalidHash = errors.String("invalid hash string")
	// ErrInvalidHashParts is returned when HashParts or its critical fields are invalid.
	ErrInvalidHashParts = errors.String("invalid hash parts or missing hash/salt")
	// ErrUnsupportedHashForPBKDF2 is returned when an unsupported hash type is used for PBKDF2.
	ErrUnsupportedHashForPBKDF2 = errors.String("unsupported hash type for PBKDF2")
	// ErrKeyLengthTooShort is returned when the key length is too short.
	ErrKeyLengthTooShort = errors.String("key length must be at least 8 bytes")
)
