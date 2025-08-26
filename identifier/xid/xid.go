/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package xid provides a unique identifier generator based on xid.
package xid

import (
	"github.com/rs/xid"

	"github.com/origadmin/toolkits/identifier"
)

// Ensure the provider and generator implement the required interfaces at compile time.
var (
	_ identifier.GeneratorProvider      = (*provider)(nil)
	_ identifier.TypedGenerator[string] = (*stringGenerator)(nil)
)

// provider implements identifier.GeneratorProvider for XID.
// It's a stateless singleton that vends the actual generator.
type provider struct{}

// Name returns the name of the identifier.
func (p *provider) Name() string {
	return "xid"
}

// Size returns the size of the identifier in bits.
// An XID is 12 bytes, so 96 bits.
func (p *provider) Size() int {
	return 96
}

// AsString returns a string-based generator for XID.
func (p *provider) AsString() identifier.TypedGenerator[string] {
	return &stringGenerator{}
}

// AsNumber returns nil as XID does not have a standard integer representation.
func (p *provider) AsNumber() identifier.TypedGenerator[int64] {
	return nil
}

// stringGenerator implements identifier.TypedGenerator[string] for XID.
// This is the actual workhorse for generating and validating IDs.
type stringGenerator struct{}

// Name returns the name of the identifier.
func (g *stringGenerator) Name() string {
	return "xid"
}

// Size returns the size of the identifier in bits.
func (g *stringGenerator) Size() int {
	return 96
}

// Generate creates a new XID and returns it as a string.
func (g *stringGenerator) Generate() string {
	return xid.New().String()
}

// Validate checks if the provided string is a valid XID.
func (g *stringGenerator) Validate(id string) bool {
	_, err := xid.FromString(id)
	return err == nil
}

// init registers the XID provider with the global identifier registry.
func init() {
	// Register a singleton instance of our provider.
	identifier.Register(&provider{})
}
