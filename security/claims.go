/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security is a toolkit for security check and authorization
package security

import (
	"context"
	"time"
)

// UserClaimsParser is an interface that defines the methods for a user claims parser
type UserClaimsParser interface {
	Parse(ctx context.Context, id string) (UserClaims, error)
}

// UserClaims is an interface that defines the methods for a casbin policy
type UserClaims interface {
	// GetSubject returns the subject of the casbin policy
	GetSubject() string
	// GetObject returns the object of the casbin policy
	GetObject() string
	// GetAction returns the action of the casbin policy
	GetAction() string
	// GetDomain returns the domain of the casbin policy
	GetDomain() string
	// GetClaims returns the claims of the casbin policy
	GetClaims() Claims
	// GetExtra returns the extra information of the casbin policy
	GetExtra() map[string]string
}

// Claims is an interface that defines the methods that a security claims object should have
type Claims interface {
	// GetSubject returns the subject of the security
	GetSubject() string
	// GetIssuer returns the issuer of the security
	GetIssuer() string
	// GetAudience returns the audience of the security
	GetAudience() []string
	// GetExpiration returns the expiration time of the security
	GetExpiration() time.Time
	// GetNotBefore returns the time before which the security cannot be accepted
	GetNotBefore() time.Time
	// GetIssuedAt returns the time at which the security was issued
	GetIssuedAt() time.Time
	// GetJWTID returns the unique identifier for the security
	GetJWTID() string
	// GetScopes returns the scopes associated with the security
	GetScopes() map[string]bool
	// GetExtra returns any additional data associated with the security
	GetExtra() map[string]string
}

// UnimplementedClaims is a struct that implements the Claims interface
type UnimplementedClaims struct {
}

// GetSubject returns an empty string
func (u UnimplementedClaims) GetSubject() string {
	return ""
}

// GetIssuer returns an empty string
func (u UnimplementedClaims) GetIssuer() string {
	return ""
}

// GetAudience returns an empty slice
func (u UnimplementedClaims) GetAudience() []string {
	return []string{}
}

// GetExpiration returns the current time
func (u UnimplementedClaims) GetExpiration() time.Time {
	return time.Now()
}

// GetNotBefore returns the current time
func (u UnimplementedClaims) GetNotBefore() time.Time {
	return time.Now()
}

// GetIssuedAt returns the current time
func (u UnimplementedClaims) GetIssuedAt() time.Time {
	return time.Now()
}

// GetJwtID returns an empty string
func (u UnimplementedClaims) GetJwtID() string {
	return ""
}

// GetScopes returns an empty map
func (u UnimplementedClaims) GetScopes() map[string]bool {
	return make(map[string]bool)
}

// GetExtra returns an empty map
func (u UnimplementedClaims) GetExtra() map[string]string {
	return make(map[string]string)
}
