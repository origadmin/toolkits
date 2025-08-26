/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package shortid

import (
	"math/rand/v2"
	"strings"

	"github.com/teris-io/shortid"

	"github.com/origadmin/toolkits/identifier"
)

// Ensure the provider and generator implement the required interfaces at compile time.
var (
	_ identifier.GeneratorProvider      = (*provider)(nil)
	_ identifier.TypedGenerator[string] = (*stringGenerator)(nil)
)

// provider implements identifier.GeneratorProvider for shortid.
// It holds a configured instance of a shortid generator.
type provider struct {
	generator *shortid.Shortid
}

// Name returns the name of the identifier.
func (p *provider) Name() string {
	return "shortid"
}

// Size returns the size of the identifier in bits.
// For shortid, the length is variable, so we return 0.
func (p *provider) Size() int {
	return 0
}

// AsString returns a string-based generator for shortid.
func (p *provider) AsString() identifier.TypedGenerator[string] {
	return &stringGenerator{
		generator: p.generator,
	}
}

// AsNumber returns nil as shortid only generates strings.
func (p *provider) AsNumber() identifier.TypedGenerator[int64] {
	return nil
}

// stringGenerator implements identifier.TypedGenerator[string] for shortid.
type stringGenerator struct {
	generator *shortid.Shortid
}

// Name returns the name of the identifier.
func (g *stringGenerator) Name() string {
	return "shortid"
}

// Size returns the size of the identifier in bits.
// For shortid, the length is variable, so we return 0.
func (g *stringGenerator) Size() int {
	return 0
}

// Generate creates a new short, unique ID string.
func (g *stringGenerator) Generate() string {
	id, err := g.generator.Generate()
	if err != nil {
		// The underlying library is not expected to error with valid initialization,
		// but we return an empty string as a safeguard.
		return ""
	}
	return id
}

// Validate checks if the provided string is a plausible shortid.
// Note: This is a best-effort validation as the library does not expose a validator.
// It checks if the characters belong to the default alphabet.
func (g *stringGenerator) Validate(id string) bool {
	if id == "" {
		return false
	}
	return strings.ContainsOnly(id, shortid.DefaultABC)
}

// init registers the shortid provider with the global identifier registry.
func init() {
	// Create a default generator instance with a random seed and worker ID
	// to ensure uniqueness across different application instances.
	generator, err := shortid.New(
		uint8(rand.Uint32N(31)+1), // Worker ID from 1 to 31
		shortid.DefaultABC,
		rand.Uint64(),             // Random seed
	)
	if err != nil {
		// This panic is acceptable in an init() function if a core component fails to initialize.
		panic("identifier: failed to initialize shortid generator: " + err.Error())
	}

	// Register a provider instance containing the configured generator.
	identifier.Register(&provider{
		generator: generator,
	})
}
