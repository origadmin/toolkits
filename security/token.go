// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package security is a toolkit for security check and authorization
package security

import (
	"encoding/json"
)

// Token is an interface for getting token information
type Token interface {
	json.Marshaler
	// TokenType returns the token type
	TokenType() TokenType
	// Domain returns the domain
	Domain() string
	// Subject returns the subject
	Subject() string
	// ExpiresAt returns the expiration time
	ExpiresAt() int64
	// AccessToken returns the access token
	AccessToken() string
	// Others returns the others
	Others() map[string]any
}
