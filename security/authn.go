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
	// CreateIdentityClaims creates a new identity claims.It should be used when a new user is created.
	CreateIdentityClaims(context.Context, string) (Claims, error)
	// CreateIdentityClaimsContext creates a new identity.It should be used when a new user is created.
	CreateIdentityClaimsContext(context.Context, TokenType, string) (context.Context, error)
	// Authenticate returns a nil error and the AuthClaims info (if available).
	Authenticate(context.Context, string) (Claims, error)
	// AuthenticateContext returns a nil error and the AuthClaims info (if available).
	// if the subject is authenticated or a non-nil error with an appropriate error cause otherwise.
	AuthenticateContext(context.Context, TokenType) (Claims, error)
	// Verify validates if a token is valid.
	Verify(context.Context, string) (bool, error)
	// VerifyContext validates if a token is valid.
	VerifyContext(context.Context, TokenType) (bool, error)
	// CreateToken inject user claims into token string. bool true is for refresh token
	CreateToken(context.Context, Claims, bool) (string, error)
	// CreateTokenContext inject user claims into context.
	CreateTokenContext(context.Context, TokenType, Claims) (context.Context, error)
	// DestroyToken invalidate a token by removing it from the token store.
	DestroyToken(context.Context, string) error
	// DestroyTokenContext invalidate a token by removing it from the token store.
	DestroyTokenContext(context.Context, TokenType) error
	// Close Cleans up the authenticator.
	Close(context.Context) error
}
