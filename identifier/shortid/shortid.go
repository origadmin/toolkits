/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package shortid

import (
	"fmt"
	"math/rand/v2"

	"github.com/teris-io/shortid"

	"github.com/origadmin/toolkits/identifier"
)

const (
	// minAlphabetLen is the minimum number of characters an alphabet must have.
	// This value is based on the internal requirements of the teris-io/shortid library.
	minAlphabetLen = 64
)

// Config holds the configuration for creating a new Shortid generator.
type Config struct {
	// Worker is the worker ID for the generator. It must be between 0 and 31.
	Worker uint8
	// Alphabet is the custom character set to use for generation.
	// It must contain at least 64 unique characters.
	Alphabet string
	// Seed is the initial seed for the random number generator.
	Seed uint64
}

// Ensure the provider and generator implement the required interfaces at compile time.
var (
	_ identifier.GeneratorProvider      = (*provider)(nil)
	_ identifier.TypedGenerator[string] = (*stringGenerator)(nil)
)

// provider implements identifier.GeneratorProvider for shortid.
// It holds a configured instance of a shortid generator.
type provider struct {
	generator *shortid.Shortid
	alphabet  string
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
		alphabet:  p.alphabet,
	}
}

// AsNumber returns nil as shortid only generates strings.
func (p *provider) AsNumber() identifier.TypedGenerator[int64] {
	return nil
}

// stringGenerator implements identifier.TypedGenerator[string] for shortid.
type stringGenerator struct {
	generator *shortid.Shortid
	alphabet  string
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
		return ""
	}
	return id
}

// Validate checks if the provided string is a plausible shortid.
// It checks if the characters belong to the alphabet used by the generator.
func (g *stringGenerator) Validate(id string) bool {
	if id == "" {
		return false
	}
	// Compatible replacement for strings.ContainsOnly
	for _, r := range id {
		present := false
		for _, a := range g.alphabet {
			if r == a {
				present = true
				break
			}
		}
		if !present {
			return false
		}
	}
	return true
}

// --- Advanced Usage ---

// NewGenerator creates a new, local, configured Shortid generator.
// This instance is NOT managed by the global identifier registry.
func NewGenerator(cfg Config) (identifier.TypedGenerator[string], error) {
	// The library panics on invalid alphabet, so we check it first.
	if len(cfg.Alphabet) < minAlphabetLen {
		return nil, fmt.Errorf("shortid: alphabet must contain at least %d unique characters", minAlphabetLen)
	}

	gen, err := shortid.New(cfg.Worker, cfg.Alphabet, cfg.Seed)
	if err != nil {
		return nil, fmt.Errorf("shortid: failed to initialize with the given settings: %w", err)
	}
	return &stringGenerator{generator: gen, alphabet: cfg.Alphabet}, nil
}

// --- Default Global Instance ---

// init registers the shortid provider with the global identifier registry.
// This provider uses a random seed and worker ID.
func init() {
	// Create a default generator instance with a random seed and worker ID
	// to ensure uniqueness across different application instances.
	generator, err := shortid.New(
		uint8(rand.Uint32N(31)+1), // Worker ID from 1 to 31
		shortid.DefaultABC,
		rand.Uint64(),             // Random seed
	)
	if err != nil {
		panic("identifier: failed to initialize default shortid generator: " + err.Error())
	}

	// Register a provider instance containing the configured generator.
	identifier.Register(&provider{
		generator: generator,
		alphabet:  shortid.DefaultABC,
	})
}
