/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package identifier provides a unified interface for generating and validating unique identifiers.
package identifier

// Identifier defines the basic interface for all identifier types.
type Identifier interface {
	Name() string // Name returns the name of the identifier.
	Size() int    // Size returns the size of the identifier in bits.
}

// TypedGenerator defines a unified, generic interface for generating identifiers of a specific type.
// This is the interface that consuming code will typically interact with after initialization.
type TypedGenerator[T ~int64 | ~string] interface {
	Identifier
	// Generate creates a new identifier of type T.
	Generate() T
	// Validate checks if the provided value is a valid identifier of type T.
	Validate(T) bool
}

// GeneratorProvider is an interface for an algorithm that can provide
// generators for different types. A single provider can vend either a string
// or a number generator, or both. This is the object you initialize via `New()`.
type GeneratorProvider interface {
    Identifier
    // AsString returns a string-based generator. Returns nil if not supported.
    AsString() TypedGenerator[string]
    // AsNumber returns a number-based generator. Returns nil if not supported.
    AsNumber() TypedGenerator[int64]
}
