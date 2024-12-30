/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

import (
	"context"
)

// Authenticator interface
type Authenticator interface {
	// CreateIdentityClaims creates a new identity claims. bool true is for refresh token
	CreateIdentityClaims(context.Context, string, bool) (Claims, error)
	// Authenticate returns a nil error and the AuthClaims info (if available).
	Authenticate(context.Context, string) (Claims, error)
	// Verify validates if a token is valid.
	Verify(context.Context, string) (bool, error)
	// CreateToken inject user claims into token string.
	CreateToken(context.Context, Claims) (string, error)
}
