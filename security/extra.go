/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

// ExtraData is an interface that defines methods for handling extra data associated with the security claims
type ExtraData interface {
	// GetClaims returns the Claims object associated with the extra data,if Claims exists
	GetClaims() (Claims, bool)
	// GetPolicy returns the Policy object associated with the extra data,if Policy exists
	GetPolicy() (Policy, bool)
	// GetExtra returns the extra data as a map of strings
	GetExtra() map[string]string
	// Get returns the value associated with the given key
	Get(key string) (string, bool)
	// Set sets the value associated with the given key
	Set(key string, value string)
}

type extraData struct {
	Claims
	Policy
	Extra map[string]string
}

func (e extraData) GetClaims() (Claims, bool) {
	return e.Claims, e.Claims != nil
}

func (e extraData) GetPolicy() (Policy, bool) {
	return e.Policy, e.Policy != nil
}

func (e extraData) GetExtra() map[string]string {
	return e.Extra
}

func (e extraData) Get(key string) (string, bool) {
	return e.Extra[key], true
}

func (e extraData) Set(key string, value string) {
	e.Extra[key] = value
}

// ExtraDataObject retrieves the ExtraData object from a Policy if it implements the ExtraData interface
func ExtraDataObject(extra any) (ExtraData, bool) {
	if ex, ok := extra.(ExtraData); ok {
		return ex, true
	}
	return nil, false
}

func ClaimsWithExtra(claims Claims, extra map[string]string) Claims {
	return &extraData{
		Claims: claims,
		Extra:  extra,
	}
}

func PolicyWithExtra(policy Policy, extra map[string]string) Policy {
	return &extraData{
		Policy: policy,
		Extra:  extra,
	}
}

var _ ExtraData = (*extraData)(nil)
