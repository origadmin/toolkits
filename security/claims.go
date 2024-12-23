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
type UserClaimsParser func(ctx context.Context, claims Claims) (UserClaims, error)

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

type rootCtxKey struct{}

func WithRootContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, rootCtxKey{}, true)
}

func ContextIsRoot(ctx context.Context) bool {
	if _, ok := ctx.Value(rootCtxKey{}).(bool); ok {
		return true
	}
	return false
}
