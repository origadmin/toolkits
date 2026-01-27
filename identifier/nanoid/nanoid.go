/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package nanoid provides a fast, URL-friendly, unique string ID generator.
package nanoid

import (
	"errors"
	"regexp"

	gonanoid "github.com/jaevor/go-nanoid"
	"github.com/origadmin/toolkits/identifier"
)

// validationRegex is used to perform a basic validation of a standard nanoid string (21 chars, URL-friendly).
var validationRegex = regexp.MustCompile("^[a-zA-Z0-9_-]{21}$")

// generator implements the identifier.Generator interface for NanoID.
type generator struct {
	generate func() (string, error)
}

// New creates a new NanoID generator with a standard length of 21 characters.
func New() (identifier.Generator, error) {
	// Create a generator function with the standard alphabet and a length of 21.
	genFunc, err := gonanoid.Standard(21)
	if err != nil {
		return nil, err
	}
	return &generator{generate: func() (string, error) { return genFunc(), nil }}, nil
}

// Name returns the name of the generator.
func (g *generator) Name() string {
	return "nanoid"
}

// GenerateString generates a new NanoID identifier as a string.
func (g *generator) GenerateString() (string, error) {
	return g.generate()
}

// GenerateNumber is not supported by NanoID and returns an error.
func (g *generator) GenerateNumber() (int64, error) {
	return 0, errors.New("nanoid: does not support number generation")
}

// Validate checks if the provided id is a valid NanoID string.
func (g *generator) Validate(id any) bool {
	idStr, ok := id.(string)
	if !ok {
		return false
	}
	// This checks for the default length of 21 and the URL-friendly character set.
	return validationRegex.MatchString(idStr)
}

// init registers the NanoID generator.
// It panics on initialization failure, as the generator would be unusable.
func init() {
	gen, err := New()
	if err != nil {
		panic("nanoid: failed to initialize generator: " + err.Error())
	}
	identifier.Register(gen)
}
