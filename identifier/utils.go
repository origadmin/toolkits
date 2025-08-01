/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package identifier implements the functions, types, and interfaces for the module.
package identifier

import (
	"math/rand/v2"
	"sync/atomic"
)

var (
	registry = NewRegistry()
)

func init() {
	registry.SetDefaultString(NewStringGenerate(16))
	registry.SetDefaultNumber(NewNumberGenerate(rand.Int64()))
}

// GenerateString generates a string identifier using the specified generator name.
func GenerateString(name string) string {
	if gen := registry.GetString(name); gen != nil {
		return gen.GenerateString()
	}
	return registry.DefaultString().GenerateString()
}

// GenerateNumber generates a number identifier using the specified generator name.
func GenerateNumber(name string) int64 {
	if gen := registry.GetNumber(name); gen != nil {
		return gen.GenerateNumber()
	}
	return registry.DefaultNumber().GenerateNumber()
}

// Generate generates an identifier of type T (int64 or string) using the specified generator name.
func Generate[T int64 | string](name string) T {
	var t T
	switch any(t).(type) {
	case int64:
		return any(GenerateNumber(name)).(T)
	case string:
		return any(GenerateString(name)).(T)
	}
	return t
}

// Validate validates an identifier of type T (int64 or string) using the specified generator name.
func Validate[T int64 | string](name string, id T) bool {
	switch v := any(id).(type) {
	case int64:
		if gen := registry.GetNumber(name); gen != nil {
			return gen.ValidateNumber(v)
		}
	case string:
		if gen := registry.GetString(name); gen != nil {
			return gen.ValidateString(v)
		}
	}
	return false
}

// StringGenerate is a simple string identifier generator.
type StringGenerate struct {
	size int
}

// NewStringGenerate creates a new StringGenerate instance with the specified size.
func NewStringGenerate(size int) *StringGenerate {
	return &StringGenerate{size: size}
}

// Name returns the name of the string generator.
func (s *StringGenerate) Name() string {
	return "random_string"
}

// Size returns the size of the generated string identifier.
func (s *StringGenerate) Size() int {
	return s.size
}

// GenerateString generates a new string identifier.
func (s *StringGenerate) GenerateString() string {
	// Implementation for generating a string identifier.
	return "generated_string"
}

// ValidateString validates a string identifier.
func (s *StringGenerate) ValidateString(id string) bool {
	// Implementation for validating a string identifier.
	return len(id) == s.size
}

// StartNumber is the start number of the randNumber.
var StartNumber = int64(rand.Int32()) << 32

// NumberGenerate is a simple number identifier generator.
type NumberGenerate struct {
	seed int64
}

// NewNumberGenerate creates a new NumberGenerate instance with the specified seed.
func NewNumberGenerate(seed int64) *NumberGenerate {
	return &NumberGenerate{seed: seed}
}

// Name returns the name of the number generator.
func (n *NumberGenerate) Name() string {
	return "random_number"
}

// Size returns the size of the generated number identifier.
func (n *NumberGenerate) Size() int {
	return 64 // Assuming the generated number is 64-bit.
}

// GenerateNumber generates a new number identifier.
func (n *NumberGenerate) GenerateNumber() int64 {
	// Implementation for generating a number identifier.
	return atomic.AddInt64(&n.seed, 1)
}

// ValidateNumber validates a number identifier.
func (n *NumberGenerate) ValidateNumber(id int64) bool {
	// Implementation for validating a number identifier.
	return id != 0
}
