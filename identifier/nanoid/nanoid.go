/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package nanoid provides a NanoID generator.
package nanoid

import (
	"regexp"

	gonanoid "github.com/jaevor/go-nanoid"

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
// It's a stateless singleton that vends the actual generator.
type provider struct{}

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
	return &stringGenerator{}
}

// AsNumber returns nil as NanoID is a string-based identifier.
func (p *provider) AsNumber() identifier.TypedGenerator[int64] {
	return nil
}

// stringGenerator implements identifier.TypedGenerator[string] for NanoID.
type stringGenerator struct{}

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
	// Using a generator which is fast and uses crypto/rand.
	id, err := gonanoid.New()
	if err != nil {
		// This is highly unlikely to fail. Return an empty string as a fallback.
		return ""
	}
	return id
}

// Validate checks if the provided string is a plausible, standard NanoID.
// This checks for the default length of 21 and the URL-friendly character set.
func (g *stringGenerator) Validate(id string) bool {
	return validationRegex.MatchString(id)
}

// init registers the NanoID provider with the global identifier registry.
func init() {
	// Register a singleton instance of our provider.
	identifier.Register(&provider{})
}
