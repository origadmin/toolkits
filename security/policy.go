/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

import (
	"context"
)

// PolicyParser is an interface that defines the methods for a user claims parser
// It takes a context and a Claims object as input and returns a Policy object and an error
type PolicyParser func(ctx context.Context, claims Claims) (Policy, error)

// Policy is an interface that defines the methods for a policy
// It provides methods for getting the subject, object, action, domain, roles, and permissions of the policy
type Policy interface {
	// GetSubject returns the subject of the casbin policy
	GetSubject() string
	// GetObject returns the object of the casbin policy
	GetObject() string
	// GetAction returns the action of the casbin policy
	GetAction() string
	// GetDomain returns the domain of the casbin policy
	GetDomain() string
	// GetRoles returns a list of roles for the user
	GetRoles() []string
	// GetPermissions returns a list of permissions for the user
	GetPermissions() []string
}

// RegisteredPolicy represents a registered policy in your system.
// It contains information about the subject, object, action, domain, roles, and permissions.
type RegisteredPolicy struct {
	Subject    string   `json:"subject"`    // The subject of the policy.
	Object     string   `json:"object"`     // The object of the policy.
	Action     string   `json:"action"`     // The action of the policy.
	Domain     string   `json:"domain"`     // The domain of the policy.
	Roles      []string `json:"roles"`      // The roles associated with the policy.
	Permission []string `json:"permission"` // The permissions associated with the policy.
}

// GetSubject returns the subject of the policy.
func (r RegisteredPolicy) GetSubject() string {
	return r.Subject
}

// GetObject returns the object of the policy.
func (r RegisteredPolicy) GetObject() string {
	return r.Object
}

// GetAction returns the action of the policy.
func (r RegisteredPolicy) GetAction() string {
	return r.Action
}

// GetDomain returns the domain of the policy.
func (r RegisteredPolicy) GetDomain() string {
	return r.Domain
}

// GetRoles returns the roles associated with the policy.
func (r RegisteredPolicy) GetRoles() []string {
	return r.Roles
}

// GetPermissions returns the permissions associated with the policy.
func (r RegisteredPolicy) GetPermissions() []string {
	return r.Permission
}

// Policy is an interface that RegisteredPolicy satisfies.
var _ Policy = (*RegisteredPolicy)(nil)
