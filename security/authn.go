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
	// AuthenticateToken returns a nil error and the AuthClaims info (if available).
	AuthenticateToken(context.Context, string) (Claims, error)
	// AuthenticateTokenContext returns a nil error and the AuthClaims info (if available).
	// if the subject is authenticated or a non-nil error with an appropriate error cause otherwise.
	AuthenticateTokenContext(context.Context, TokenType) (Claims, error)
	// Authenticate validates if a token is valid.
	Authenticate(context.Context, string) (bool, error)
	// AuthenticateContext validates if a token is valid.
	AuthenticateContext(context.Context, TokenType) (bool, error)
	// CreateToken inject user claims into token string.
	CreateToken(context.Context, Claims) (string, error)
	// CreateTokenContext inject user claims into context.
	CreateTokenContext(context.Context, TokenType, Claims) (context.Context, error)
	// DestroyToken invalidate a token by removing it from the token store.
	DestroyToken(context.Context, string) error
	// DestroyTokenContext invalidate a token by removing it from the token store.
	DestroyTokenContext(context.Context, TokenType) error
	// Close Cleans up the authenticator.
	Close()
}
