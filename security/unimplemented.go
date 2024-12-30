/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

import (
	"time"
)

// UnimplementedClaims is a struct that implements the Claims interface
type UnimplementedClaims struct {
}

func (u UnimplementedClaims) GetJWTID() string {
	return ""
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

type UnimplementedUserClaims struct {
}

func (u UnimplementedUserClaims) IsRoot() bool {
	return false
}

func (u UnimplementedUserClaims) GetSubject() string {
	return ""
}

func (u UnimplementedUserClaims) GetObject() string {
	return ""
}

func (u UnimplementedUserClaims) GetAction() string {
	return ""
}

func (u UnimplementedUserClaims) GetDomain() string {
	return ""
}

func (u UnimplementedUserClaims) GetClaims() Claims {
	return &UnimplementedClaims{}
}

func (u UnimplementedUserClaims) GetExtra() map[string]string {
	return nil
}
