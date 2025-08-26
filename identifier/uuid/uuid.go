/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package uuid provides implementations for various UUID versions.
package uuid

import (
	"github.com/google/uuid"

	"github.com/origadmin/toolkits/identifier"
)

// Ensure the provider and generator implement the required interfaces at compile time.
var (
	_ identifier.Provider          = (*provider)(nil)
	_ identifier.Generator[string] = (*stringGenerator)(nil)
)

// provider implements identifier.Provider for a specific UUID version.
// It holds the name and the function to generate a new UUID string.
type provider struct {
	name      string
	generator func() string
}

// Name returns the name of the identifier (e.g., "uuid", "uuid-v7").
func (p *provider) Name() string {
	return p.name
}

// Size returns the size of the identifier in bits, which is always 128 for UUIDs.
func (p *provider) Size() int {
	return 128
}

// AsString returns a string-based generator for this UUID version.
func (p *provider) AsString() identifier.Generator[string] {
	return &stringGenerator{
		name:      p.name,
		generator: p.generator,
	}
}

// AsNumber returns nil as UUIDs do not have a standard integer representation.
func (p *provider) AsNumber() identifier.Generator[int64] {
	return nil
}

// stringGenerator implements identifier.Generator[string].
// This is the actual workhorse that generates and validates UUIDs.
type stringGenerator struct {
	name      string
	generator func() string
}

// Name returns the name of the identifier.
func (g *stringGenerator) Name() string {
	return g.name
}

// Size returns the size of the identifier in bits.
func (g *stringGenerator) Size() int {
	return 128
}

// Generate creates a new UUID string using the configured generator function.
func (g *stringGenerator) Generate() string {
	return g.generator()
}

// Validate checks if the provided string is a valid UUID.
func (g *stringGenerator) Validate(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// --- Convenience Constructors ---

// New creates a new, default UUID (v4) generator.
func New() identifier.Generator[string] {
	return identifier.Get[string]("uuid")
}

// NewV1 creates a new UUIDv1 (time-based) generator.
func NewV1() identifier.Generator[string] {
	return identifier.Get[string]("uuid-v1")
}

// NewV4 creates a new UUIDv4 (random) generator.
func NewV4() identifier.Generator[string] {
	return identifier.Get[string]("uuid-v4")
}

// NewV6 creates a new UUIDv6 (time-sortable) generator.
func NewV6() identifier.Generator[string] {
	return identifier.Get[string]("uuid-v6")
}

// NewV7 creates a new UUIDv7 (time-sortable) generator.
func NewV7() identifier.Generator[string] {
	return identifier.Get[string]("uuid-v7")
}

// --- Default Global Instance ---

// init registers multiple providers for different UUID versions.
func init() {
	// Register UUIDv4 as the default "uuid" and also as "uuid-v4".
	// UUIDv4 is the most common, random-based UUID.
	v4provider := &provider{
		name:      "uuid",
		generator: func() string { return uuid.New().String() },
	}
	identifier.Register(v4provider)
	identifier.Register(&provider{
		name:      "uuid-v4",
		generator: v4provider.generator, // Reuse the same generator function
	})

	// Register UUIDv7, the new time-sortable standard.
	identifier.Register(&provider{
		name:      "uuid-v7",
		generator: func() string { return uuid.Must(uuid.NewV7()).String() },
	})

	// Register UUIDv6, another time-sortable version.
	identifier.Register(&provider{
		name:      "uuid-v6",
		generator: func() string { return uuid.Must(uuid.NewV6()).String() },
	})

	// Register UUIDv1, the classic time-based version.
	identifier.Register(&provider{
		name:      "uuid-v1",
		generator: func() string { return uuid.Must(uuid.NewUUID()).String() },
	})
}
