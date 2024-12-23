/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

import (
	"context"
)

// PolicyManager is an interface that defines the methods for a policy manager
type PolicyManager interface {
	// AddPolicy adds a policy for a given subject, object, action, domain
	AddPolicy(sec string, pt string, rule []string) error
	// RemovePolicy removes a policy for a given subject, object, action, domain
	RemovePolicy(sec string, pt string, rule []string) error
	// SetPolicies sets the policies for a given context
	SetPolicies(context.Context, map[string]any) error
}

// Authorizer is an interface that defines the methods for an authorizer.
// It is used to manage policies and check authorization.
type Authorizer interface {
	// SetPolicies sets the policies for a given context.
	// It takes a context, a map of policies, and a map of roles as input.
	// It returns an error if the policies cannot be set.
	SetPolicies(ctx context.Context, policies map[string]any, roles map[string]any) error

	// Authorized checks if a user is authorized to perform an action.
	// It takes a context and a UserClaims object as input.
	// It returns a boolean indicating whether the user is authorized and an error if the check fails.
	Authorized(ctx context.Context, claims UserClaims) (bool, error)
}
