/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

import (
	"context"
)

type Policy struct {
	Subject string
	Object  string
	Action  string
	Domain  []string
	Extra   map[string]string
}

// PolicyManager is an interface that defines the methods for a policy manager
type PolicyManager interface {
	// AddPolicy adds a policy for a given subject, object, action, domain and extra
	AddPolicy(sec string, pt string, rule []string) error
	// RemovePolicy removes a policy for a given subject, object, action, domain and extra
	RemovePolicy(sec string, pt string, rule []string) error
	// GetPolicy returns the policy for a given subject, object, action, domain and extra
	GetPolicy(subject string, object string, action string, domain []string, extra map[string]string) Policy
	// SetPolicy sets the policy for a given subject, object, action, domain and extra
	SetPolicy(subject string, object string, action string, domain []string, extra map[string]string)
	// SetPolicies sets the policies for a given context
	SetPolicies(context.Context, map[string]Policy) error
}

// PolicyChecker is an interface that defines the methods for a policy checker
type PolicyChecker interface {
	// CheckPolicy checks if the policy for a given subject, object, action, domain and extra is allowed
	CheckPolicy(Policy) bool
	// CheckPolicyContext checks if the policy for a given subject, object, action
	CheckPolicyContext(context.Context, TokenType, Policy) bool
}

type Authorizer interface {
	SetPolicies(context.Context, map[string]any, map[string]any) error
	Authorized(context.Context, UserClaims) (bool, error)
}
