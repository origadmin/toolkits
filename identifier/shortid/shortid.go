/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package shortid provides a short, URL-friendly, unique ID generator.
package shortid

import (
	"errors"
	"math/rand/v2"
	"strings"

	"github.com/teris-io/shortid"

	"github.com/origadmin/toolkits/identifier"
)

// generator implements the identifier.Generator interface for ShortID.
type generator struct {
	gen      *shortid.Shortid
	alphabet string
}

// New creates a new ShortID generator with default settings.
func New() (identifier.Generator, error) {
	// Create a default generator instance with a random seed and worker ID
	// to ensure uniqueness across different application instances.
	alphabet := shortid.DefaultABC
	gen, err := shortid.New(
		uint8(rand.Uint32N(31)+1), // Worker ID from 1 to 31
		alphabet,
		rand.Uint64(),             // Random seed
	)
	if err != nil {
		return nil, err
	}
	return &generator{gen: gen, alphabet: alphabet}, nil
}

// Name returns the name of the generator.
func (g *generator) Name() string {
	return "shortid"
}

// GenerateString generates a new ShortID identifier as a string.
func (g *generator) GenerateString() (string, error) {
	return g.gen.Generate()
}

// GenerateNumber is not supported by ShortID and returns an error.
func (g *generator) GenerateNumber() (int64, error) {
	return 0, errors.New("shortid: does not support number generation")
}

// Validate checks if the provided id is a valid ShortID string.
func (g *generator) Validate(id any) bool {
	idStr, ok := id.(string)
	if !ok {
		return false
	}
	return strings.Trim(idStr, g.alphabet) == ""
}

// init registers the ShortID generator.
// It panics on initialization failure.
func init() {
	gen, err := New()
	if err != nil {
		panic("shortid: failed to initialize generator: " + err.Error())
	}
	identifier.Register(gen)
}
