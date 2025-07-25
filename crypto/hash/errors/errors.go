/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package errors

import "github.com/origadmin/toolkits/errors"

var (
	// ErrPasswordNotMatch error when password not match
	ErrPasswordNotMatch = errors.String("password not match")
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
)
