/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package nanoid provides a NanoID generator.
package nanoid

import (
	"fmt"
	"regexp"

	"github.com/jaevor/go-nanoid"
	"github.com/origadmin/toolkits/identifier"
)

// Ensure the provider and generator implement the required interfaces at compile time.
var (
	_ identifier.GeneratorProvider      = (*provider)(nil)
	_ identifier.TypedGenerator[string] = (*stringGenerator)(nil)
)

// validationRegex is used to perform a basic validation of a standard nanoid string (21 chars, URL-friendly).
var validationRegex = regexp.MustCompile("^[a-zA-Z0-9_-]{21}$")

// provider implements identifier.GeneratorProvider for NanoID.
type provider struct {
	generator func() (string, error)
}

// Name returns the name of the identifier.
func (p *provider) Name() string {
	return "nanoid"
}

// Size returns the size of the identifier in bits.
// For nanoid, the length is variable, so we return 0. The default is 21 characters.
func (p *provider) Size() int {
	return 0
}

// AsString returns a string-based generator for NanoID.
func (p *provider) AsString() identifier.TypedGenerator[string] {
	return &stringGenerator{generator: p.generator}
}

// AsNumber returns nil as NanoID is a string-based identifier.
func (p *provider) AsNumber() identifier.TypedGenerator[int64] {
	return nil
}

// stringGenerator implements identifier.TypedGenerator[string] for NanoID.
type stringGenerator struct {
	generator func() (string, error)
}

// Name returns the name of the identifier.
func (g *stringGenerator) Name() string {
	return "nanoid"
}

// Size returns 0 as the length can be variable.
func (g *stringGenerator) Size() int {
	return 0
}

// Generate creates a new, URL-friendly, unique string ID with a default length of 21.
func (g *stringGenerator) Generate() string {
	// This library is fast and uses crypto/rand by default.
	// We panic on error for consistency with the generator pattern.
	id, err := g.generator()
	if err != nil {
		panic(fmt.Sprintf("nanoid: failed to generate id: %v", err))
	}
	return id
}

// Validate checks if the provided string is a plausible, standard NanoID.
// This checks for the default length of 21 and the URL-friendly character set.
func (g *stringGenerator) Validate(id string) bool {
	return validationRegex.MatchString(id)
}

// --- Convenience Constructor ---

// New creates a new, default NanoID generator.
// This is a convenience function for direct use of the nanoid package,
// and it returns the globally registered default generator.
func New() identifier.TypedGenerator[string] {
	// This relies on the init() function having registered the provider.
	return identifier.New[string]("nanoid")
}

// init registers the NanoID provider with the global identifier registry.
func init() {
	// Create a new generator function with the standard alphabet and a length of 21.
	genFunc, err := nanoid.Standard(21)
	if err != nil {
		panic(fmt.Sprintf("nanoid: failed to initialize generator: %v", err))
	}

	// Wrap the generator to match the provider's expected signature.
	wrappedGen := func() (string, error) {
		return genFunc(), nil
	}

	identifier.Register(&provider{generator: wrappedGen})
}
