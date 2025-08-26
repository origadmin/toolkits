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
	_ identifier.GeneratorProvider      = (*provider)(nil)
	_ identifier.TypedGenerator[string] = (*stringGenerator)(nil)
)

// provider implements identifier.GeneratorProvider for a specific UUID version.
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
func (p *provider) AsString() identifier.TypedGenerator[string] {
	return &stringGenerator{
		name:      p.name,
		generator: p.generator,
	}
}

// AsNumber returns nil as UUIDs do not have a standard integer representation.
func (p *provider) AsNumber() identifier.TypedGenerator[int64] {
	return nil
}

// stringGenerator implements identifier.TypedGenerator[string].
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

// init registers multiple providers for different UUID versions.
func init() {
	// Register UUIDv4 as the default "uuid" and also as "uuid-v4".
	// UUIDv4 is the most common, random-based UUID.
	v4Provider := &provider{
		name:      "uuid",
		generator: func() string { return uuid.New().String() },
	}
	identifier.Register(v4Provider)
	identifier.Register(&provider{
		name:      "uuid-v4",
		generator: v4Provider.generator, // Reuse the same generator function
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
