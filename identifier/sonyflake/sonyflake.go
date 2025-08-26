/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package sonyflake provides a unique ID generation based on Sony's snowflake variant.
package sonyflake // Corrected package name from 'shortid'

import (
	"github.com/sony/sonyflake"

	"github.com/origadmin/toolkits/identifier"
)

// Ensure the provider and generator implement the required interfaces at compile time.
var (
	_ identifier.GeneratorProvider      = (*provider)(nil)
	_ identifier.TypedGenerator[int64] = (*numberGenerator)(nil)
)

// provider implements identifier.GeneratorProvider for Sonyflake.
// It holds a configured, stateful sonyflake instance.
type provider struct {
	sf *sonyflake.Sonyflake
}

// Name returns the name of the identifier.
func (p *provider) Name() string {
	return "sonyflake"
}

// Size returns the size of the identifier in bits.
// A Sonyflake ID is 63 bits.
func (p *provider) Size() int {
	return 63
}

// AsString returns nil as Sonyflake does not have a standard string representation.
func (p *provider) AsString() identifier.TypedGenerator[string] {
	return nil
}

// AsNumber returns a number-based generator for Sonyflake.
func (p *provider) AsNumber() identifier.TypedGenerator[int64] {
	return &numberGenerator{sf: p.sf}
}

// numberGenerator implements identifier.TypedGenerator[int64] for Sonyflake.
type numberGenerator struct {
	sf *sonyflake.Sonyflake
}

// Name returns the name of the identifier.
func (g *numberGenerator) Name() string {
	return "sonyflake"
}

// Size returns the size of the identifier in bits.
func (g *numberGenerator) Size() int {
	return 63
}

// Generate creates a new Sonyflake ID and returns it as an int64.
func (g *numberGenerator) Generate() int64 {
	id, err := g.sf.NextID()
	if err != nil {
		// This can happen if the system clock goes backwards.
		// Returning 0 is a reasonable, if not ideal, fallback.
		return 0
	}
	return int64(id)
}

// Validate checks if the provided int64 is a plausible Sonyflake ID.
func (g *numberGenerator) Validate(id int64) bool {
	// Any positive number could be a valid ID.
	return id > 0
}

// init registers the Sonyflake provider with the global identifier registry.
func init() {
	// Create a default sonyflake instance.
	// An empty Settings struct will cause the library to use the private IP
	// as the machine ID, which is a sensible default.
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		// This can happen if the machine ID function fails.
		panic("identifier: failed to initialize sonyflake")
	}

	// Register a provider instance containing the configured generator.
	identifier.Register(&provider{sf: sf})
}
