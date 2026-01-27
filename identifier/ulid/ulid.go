/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package ulid provides a Universally Unique Lexicographically Sortable Identifier (ULID) generator.
package ulid

import (
	"crypto/rand"
	"time"

	"github.com/oklog/ulid/v2"

	"github.com/origadmin/toolkits/identifier"
)

// generator implements the identifier.Generator interface for ULID.
type generator struct {
	entropy *ulid.MonotonicEntropy
}

// New creates a new ULID generator.
func New() identifier.Generator {
	// Create a new monotonic entropy source seeded with a cryptographically secure random reader.
	entropy := ulid.Monotonic(rand.Reader, 0)
	return &generator{entropy: entropy}
}

// Name returns the name of the generator.
func (g *generator) Name() string {
	return "ulid"
}

// GenerateString generates a new ULID identifier as a string.
func (g *generator) GenerateString() (string, error) {
	id, err := ulid.New(ulid.Timestamp(time.Now()), g.entropy)
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// GenerateNumber returns the Unix timestamp part of a new ULID in milliseconds.
// This is not a unique number but provides the time component.
func (g *generator) GenerateNumber() (int64, error) {
	id, err := ulid.New(ulid.Timestamp(time.Now()), g.entropy)
	if err != nil {
		return 0, err
	}
	return int64(id.Time()), nil
}

// Validate checks if the provided id is a valid ULID string.
func (g *generator) Validate(id any) bool {
	idStr, ok := id.(string)
	if !ok {
		return false
	}
	_, err := ulid.ParseStrict(idStr)
	return err == nil
}

// init registers the ULID generator.
func init() {
	// This can panic if the entropy source fails, which is a critical failure.
	defer func() {
		if r := recover(); r != nil {
			panic("ulid: failed to initialize generator, entropy source might be unavailable")
		}
	}()
	identifier.Register(New())
}
