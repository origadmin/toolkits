/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package ksuid provides a KSUID (K-Sortable Unique Identifier) implementation.
package ksuid

import (
	"github.com/segmentio/ksuid"

	"github.com/origadmin/toolkits/identifier"
)

// Ensure the provider and generator implement the required interfaces at compile time.
var (
	_ identifier.GeneratorProvider      = (*provider)(nil)
	_ identifier.TypedGenerator[string] = (*stringGenerator)(nil)
)

// provider implements identifier.GeneratorProvider for KSUID.
// It's a stateless singleton that vends the actual generator.
type provider struct{}

// Name returns the name of the identifier.
func (p *provider) Name() string {
	return "ksuid"
}

// Size returns the size of the identifier in bits.
func (p *provider) Size() int {
	// A KSUID is 160 bits (20 bytes).
	return 160
}

// AsString returns a string-based generator for KSUID.
func (p *provider) AsString() identifier.TypedGenerator[string] {
	return &stringGenerator{}
}

// AsNumber returns nil as KSUID does not have a standard integer representation.
func (p *provider) AsNumber() identifier.TypedGenerator[int64] {
	return nil
}

// stringGenerator implements identifier.TypedGenerator[string] for KSUID.
// This is the actual workhorse for generating and validating IDs.
type stringGenerator struct{}

// Name returns the name of the identifier.
func (g *stringGenerator) Name() string {
	return "ksuid"
}

// Size returns the size of the identifier in bits.
func (g *stringGenerator) Size() int {
	return 160
}

// Generate creates a new KSUID and returns it as a string.
func (g *stringGenerator) Generate() string {
	// ksuid.New() panics if it fails to read from the system's entropy source.
	return ksuid.New().String()
}

// Validate checks if the provided string is a valid KSUID.
func (g *stringGenerator) Validate(id string) bool {
	_, err := ksuid.Parse(id)
	return err == nil
}

// init registers the KSUID provider with the global identifier registry.
func init() {
	// Register a singleton instance of our provider.
	identifier.Register(&provider{})
}
