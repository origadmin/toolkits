/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package cuid2 provides a secure, collision-resistant, and scalable unique ID generator.
package cuid2

import (
	"fmt"

	"github.com/nrednav/cuid2"

	"github.com/origadmin/toolkits/identifier"
)

// Config holds the configuration for creating a new CUID2 generator.
type Config struct {
	// Length is the desired length of the CUID2 string.
	// If not set, it defaults to the library's default length.
	Length int
	// Fingerprint is a custom string used to further prevent collisions
	// in a distributed system. It should be unique to the machine or process.
	Fingerprint string
}

// Ensure the provider and generator implement the required interfaces at compile time.
var (
	_ identifier.Provider      = (*provider)(nil)
	_ identifier.Generator[string] = (*stringGenerator)(nil)
)

// provider implements identifier.Provider for CUID2.
// It holds a configured generator function.
type provider struct {
	generator func() string
}

// Name returns the name of the identifier.
func (p *provider) Name() string {
	return "cuid2"
}

// Size returns 0 as the length of a CUID2 is variable.
func (p *provider) Size() int {
	return 0
}

// AsString returns a string-based generator for CUID2.
func (p *provider) AsString() identifier.Generator[string] {
	return &stringGenerator{generator: p.generator}
}

// AsNumber returns nil as CUID2 only generates strings.
func (p *provider) AsNumber() identifier.Generator[int64] {
	return nil
}

// stringGenerator implements identifier.Generator[string] for CUID2.
type stringGenerator struct {
	generator func() string
}

// Name returns the name of the identifier.
func (g *stringGenerator) Name() string {
	return "cuid2"
}

// Size returns 0 as the length is variable.
func (g *stringGenerator) Size() int {
	return 0
}

// Generate creates a new, secure, unique ID string.
func (g *stringGenerator) Generate() string {
	return g.generator()
}

// Validate checks if the provided string is a valid CUID2.
func (g *stringGenerator) Validate(id string) bool {
	return cuid2.IsCuid(id)
}

// --- Convenience Constructor ---

// New creates a new, default CUID2 generator.
// This is a convenience function for direct use of the cuid2 package,
// and it returns the globally registered default generator.
func New() identifier.Generator[string] {
	// This relies on the init() function having registered the provider.
	return identifier.Get[string]("cuid2")
}

// --- Advanced Usage ---

// NewGenerator creates a new, local, configured CUID2 generator.
// This instance is NOT managed by the global identifier registry.
// NOTE: The underlying 'github.com/nrednav/cuid2' library does not support
// custom length or fingerprint. This function will return an error if
// configuration is provided.
func NewGenerator(cfg Config) (identifier.Generator[string], error) {
	if cfg.Length > 0 || cfg.Fingerprint != "" {
		return nil, fmt.Errorf("cuid2: custom length and fingerprint are not supported by this generator")
	}

	// The library provides a simple Generate function without configuration.
	return &stringGenerator{generator: cuid2.Generate}, nil
}

// --- Default Global Instance ---

// init registers the default CUID2 provider with the global identifier registry.
// This provider uses default, secure settings.
func init() {
	// The default generator uses a secure random source and a default fingerprint.
	identifier.Register(&provider{
		generator: cuid2.Generate,
	})
}
