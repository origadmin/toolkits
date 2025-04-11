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

// StringValidator defines the interface for validating string identifiers.
type StringValidator interface {
	ValidateString(string) bool // ValidateString checks if the provided string is a valid identifier.
}

// StringGenerator defines the interface for generating string identifiers.
type StringGenerator interface {
	GenerateString() string // GenerateString generates a new string identifier.
}

// NumberValidator defines the interface for validating number identifiers.
type NumberValidator interface {
	ValidateNumber(int64) bool // ValidateNumber checks if the provided number is a valid identifier.
}

// NumberGenerator defines the interface for generating number identifiers.
type NumberGenerator interface {
	GenerateNumber() int64 // GenerateNumber generates a new number identifier.
}

// StringIdentifier combines the Identifier, StringValidator, and StringGenerator interfaces.
type StringIdentifier interface {
	Identifier
	StringValidator
	StringGenerator
}

// NumberIdentifier combines the Identifier, NumberValidator, and NumberGenerator interfaces.
type NumberIdentifier interface {
	Identifier
	NumberValidator
	NumberGenerator
}

// MultiTypeIdentifier combines the StringIdentifier and NumberIdentifier interfaces.
type MultiTypeIdentifier interface {
	StringIdentifier
	NumberIdentifier
}

type TypedIdentifier[T ~int64 | ~string] interface {
	Identifier
	Generate() T
	Validate(T) bool
}

// SetDefaultIdentifier sets the default identifier generator.
func SetDefaultIdentifier(gen Identifier) {
	switch v := gen.(type) {
	case StringIdentifier:
		registry.SetDefaultString(v)
	case NumberIdentifier:
		registry.SetDefaultNumber(v)
	}
}
