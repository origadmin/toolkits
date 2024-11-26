/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package idgen provides the helpers functions.
package idgen

import (
	"sync"
)

// Identifier is the interface of randNumber.
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

// Register sets the defaultIdentifier randNumber.
func Register(identifier Identifier) {
	once.Do(func() {
		defaultIdentifier = identifier
	})
}

// Default method returns the default defaultGenerator randNumber.
func Default() Identifier {
	return defaultIdentifier
}

// GenID The function "GenID" generates a new unique identifier and returns it as a string.
func GenID() string {
	return defaultIdentifier.Gen()
}

// Size The function "Size" returns the size of the generated identifier
func Size() int {
	return defaultIdentifier.Size()
}

// Validate The function "Validate" checks whether the given identifier is valid or not.
func Validate(id string) bool {
	return defaultIdentifier.Validate(id)
}

var (
	_ = GenID
	_ = Size
	_ = Validate
)
