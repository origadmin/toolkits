// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package ident provides the helpers functions.
package ident

import (
	"sync"
)

// Identifier is the interface of ident.
type Identifier interface {
	Name() string
	Gen() string
	Validate(id string) bool
	Size() int
}

var (
	defaultIdentifier = NewNumber(StartNumber)
	once              sync.Once
)

// Register sets the defaultIdentifier ident.
func Register(identifier Identifier) {
	once.Do(func() {
		defaultIdentifier = identifier
	})
}

// Default method returns the default defaultGenerator ident.
func Default() Identifier {
	return defaultIdentifier
}

// GenID The function "GenID" generates a new unique identifier and returns it as a string.
func GenID() string {
	return defaultIdentifier.Gen()
}

// GenSize The function "GenSize" returns the size of the generated identifier
func GenSize() int {
	return defaultIdentifier.Size()
}

// Validate The function "Validate" checks whether the given identifier is valid or not.
func Validate(id string) bool {
	return defaultIdentifier.Validate(id)
}

var (
	_ = GenID
	_ = GenSize
	_ = Validate
)
