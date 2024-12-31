/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security is a toolkit for security check and authorization
package security

// Claims is an interface that defines the methods that a security claims object should have
// It provides methods for getting the subject, issuer, audience, expiration, not before, issued at, JWT ID, and scopes of the claims
type Claims interface {
	// GetSubject returns the subject of the security
	GetSubject() string
	// GetIssuer returns the issuer of the security
	GetIssuer() string
	// GetAudience returns the audience of the security
	GetAudience() []string
	// GetExpiration returns the expiration time of the security
	GetExpiration() int64
	// GetNotBefore returns the time before which the security cannot be accepted
	GetNotBefore() int64
	// GetIssuedAt returns the time at which the security was issued
	GetIssuedAt() int64
	// GetID returns the unique identifier for the security
	GetID() string
	// GetScopes returns the scopes associated with the security
	GetScopes() map[string]bool
}

// RegisteredClaims is a struct that implements the Claims interface
// It provides fields for the subject, issuer, audience, expiration, not before, issued at, JWT ID, and scopes of the claims
// json example:
//
//	{
//	  "sub": "test_subject",
//	  "iss": "test_issuer",
//	  "aud": [
//	    "test_audience1",
//	    "test_audience2"
//	  ],
//	  "exp": 1735647621,
//	  "nbf": 1735644021,
//	  "iat": 1735644021,
//	  "jti": "test_jti",
//	  "scopes": {
//	    "scope1": true,
//	    "scope2": false
//	  }
//	}
type RegisteredClaims struct {
	ID         string          `json:"jti,omitempty"`
	Subject    string          `json:"sub,omitempty"`
	Issuer     string          `json:"iss,omitempty"`
	Audience   []string        `json:"aud,omitempty"`
	Expiration int64           `json:"exp,omitempty"`
	NotBefore  int64           `json:"nbf,omitempty"`
	IssuedAt   int64           `json:"iat,omitempty"`
	Scopes     map[string]bool `json:"scopes,omitempty"`
}

// GetSubject returns the subject of the claims
func (r RegisteredClaims) GetSubject() string {
	return r.Subject
}

// GetIssuer returns the issuer of the claims
func (r RegisteredClaims) GetIssuer() string {
	return r.Issuer
}

// GetAudience returns the audience of the claims
func (r RegisteredClaims) GetAudience() []string {
	return r.Audience
}

// GetExpiration returns the expiration time of the claims
func (r RegisteredClaims) GetExpiration() int64 {
	return r.Expiration
}

// GetNotBefore returns the time before which the claims cannot be accepted
func (r RegisteredClaims) GetNotBefore() int64 {
	return r.NotBefore
}

// GetIssuedAt returns the time at which the claims were issued
func (r RegisteredClaims) GetIssuedAt() int64 {
	return r.IssuedAt
}

// GetID returns the unique identifier for the claims
func (r RegisteredClaims) GetID() string {
	return r.ID
}

// GetScopes returns the scopes associated with the RegisteredClaims.
func (r RegisteredClaims) GetScopes() map[string]bool {
	// Return the scopes map directly, as it is already a map[string]bool.
	return r.Scopes
}

// Assert that RegisteredClaims implements the Claims interface.
var _ Claims = (*RegisteredClaims)(nil)
