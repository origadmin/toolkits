/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package core

import "github.com/origadmin/toolkits/errors"

var (
	// ErrPasswordNotMatch error when password not match
	ErrPasswordNotMatch = errors.String("password not match")
	// ErrAlgorithmMismatch error when algorithm mismatch
	ErrAlgorithmMismatch = errors.String("algorithm mismatch")
)
