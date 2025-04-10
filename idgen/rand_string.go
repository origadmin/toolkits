/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package idgen implements the functions, types, and interfaces for the module.
package idgen

import (
	"crypto/rand"
)

func init() {
	RegisterStringIdentifier(NewString())
}

type StringGenerator struct {
	size int
}

// Name returns the name of the object identifier.
func (obj *StringGenerator) Name() string {
	return "string"
}

// String method generates an identifier.
func (obj *StringGenerator) String() string {
	id := make([]byte, obj.size)
	_, err := rand.Read(id)
	if err != nil {
		return ""
	}
	return string(id)
}

// ValidateString method checks if the given identifier is valid.
func (obj *StringGenerator) ValidateString(id string) bool {
	return true
}

// Size method returns the size of the identifier.
func (obj *StringGenerator) Size() int {
	return obj.size
}

// NewString method creates a new instance of the StringGenerator type.
func NewString() StringIdentifier {
	return &StringGenerator{
		size: 16,
	}
}
