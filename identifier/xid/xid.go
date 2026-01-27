/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package xid provides a globally unique identifier generator.
package xid

import (
	"github.com/rs/xid"

	"github.com/origadmin/toolkits/identifier"
)

// generator implements the identifier.Generator interface for XID.
type generator struct{}

// New creates a new XID generator.
// Note: This generator is distinct from the default generator in the parent 'identifier' package,
// but it uses the same underlying algorithm.
func New() identifier.Generator {
	return &generator{}
}

// Name returns the name of the generator.
func (g *generator) Name() string {
	// This name must match the one used by the internal default generator
	// to avoid registration conflicts if a user explicitly registers this one.
	// The registration logic prevents duplicates.
	return "xid"
}

// GenerateString generates a new XID identifier as a string.
func (g *generator) GenerateString() (string, error) {
	return xid.New().String(), nil
}

// GenerateNumber returns the Unix timestamp part of a new XID.
// This is not a unique number but can be useful for time-based sorting.
func (g *generator) GenerateNumber() (int64, error) {
	return xid.New().Time().Unix(), nil
}

// Validate checks if the provided id is a valid XID string.
func (g *generator) Validate(id any) bool {
	idStr, ok := id.(string)
	if !ok {
		return false
	}
	_, err := xid.FromString(idStr)
	return err == nil
}

// init registers this XID generator with the identifier registry.
// The registration in the parent 'identifier' package for the default generator
// runs first, so this registration will effectively be a no-op due to the duplicate name check.
// This ensures that if the parent package's logic changes, this package remains self-registering.
func init() {
	// This registration will likely be skipped because the default 'xid' is already registered.
	// We use a separate function to handle the panic gracefully for clarity.
	tryRegister()
}

func tryRegister() {
	// The registry panics on duplicate registration. We can recover from it
	// as it's an expected condition in this specific setup.
	defer func() {
		_ = recover()
	}()
	identifier.Register(New())
}
