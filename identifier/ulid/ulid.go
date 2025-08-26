/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package ulid provides a ULID (Universally Unique Lexicographically Sortable Identifier) implementation.
package ulid

import (
	"github.com/oklog/ulid/v2"

	"github.com/origadmin/toolkits/identifier"
)

// Ensure the provider and generator implement the required interfaces at compile time.
var (
	_ identifier.GeneratorProvider      = (*provider)(nil)
	_ identifier.TypedGenerator[string] = (*stringGenerator)(nil)
)

// provider implements identifier.GeneratorProvider for ULID.
// It's a stateless singleton that vends the actual generator.
type provider struct{}

// Name returns the name of the identifier.
func (p *provider) Name() string {
	return "ulid"
}

// Size returns the size of the identifier in bits.
func (p *provider) Size() int {
	// A ULID is always 128 bits.
	return 128
}

// AsString returns a string-based generator for ULID.
func (p *provider) AsString() identifier.TypedGenerator[string] {
	return &stringGenerator{}
}

// AsNumber returns nil as ULID does not have a standard integer representation.
func (p *provider) AsNumber() identifier.TypedGenerator[int64] {
	return nil
}

// stringGenerator implements identifier.TypedGenerator[string] for ULID.
// This is the actual workhorse for generating and validating IDs.
type stringGenerator struct{}

// Name returns the name of the identifier.
func (g *stringGenerator) Name() string {
	return "ulid"
}

// Size returns the size of the identifier in bits.
func (g *stringGenerator) Size() int {
	return 128
}

// Generate creates a new ULID and returns it as a string.
func (g *stringGenerator) Generate() string {
	return ulid.Make().String()
}

// Validate checks if the provided string is a valid ULID.
func (g *stringGenerator) Validate(id string) bool {
	_, err := ulid.ParseStrict(id)
	return err == nil
}

// init registers the ULID provider with the global identifier registry.
func init() {
	// Register a singleton instance of our provider.
	identifier.Register(&provider{})
}
