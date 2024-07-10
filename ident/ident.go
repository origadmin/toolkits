// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package ident provides the helpers functions.
package ident

// Identifier is the interface of ident.
type Identifier interface {
	Name() string
	Gen() string
	Validate(id string) bool
	Size() int
}

var defaultGenerator Identifier

func init() {
	defaultGenerator = newNumber()
}

// Use sets the defaultGenerator ident.
func Use(ident Identifier) {
	defaultGenerator = ident
}

// Default method returns the default defaultGenerator ident.
func Default() Identifier {
	return defaultGenerator
}

// GenID The function "GenID" generates a new unique identifier and returns it as a string.
func GenID() string {
	return defaultGenerator.Gen()
}

// GenSize The function "GenSize" returns the size of the generated identifier
func GenSize() int {
	return defaultGenerator.Size()
}

// Validate The function "Validate" checks whether the given identifier is valid or not.
func Validate(id string) bool {
	return defaultGenerator.Validate(id)
}

var (
	_ = Use
	_ = GenID
	_ = GenSize
	_ = Validate
)
