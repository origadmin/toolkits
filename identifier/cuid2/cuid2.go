/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package cuid2 provides a secure, collision-resistant, and scalable unique ID generator.
package cuid2

import (
	"errors"

	"github.com/nrednav/cuid2"

	"github.com/origadmin/toolkits/identifier"
)

// generator implements the identifier.Generator interface for CUID2.
type generator struct{}

// New creates a new CUID2 generator.
func New() identifier.Generator {
	return &generator{}
}

// Name returns the name of the generator.
func (g *generator) Name() string {
	return "cuid2"
}

// GenerateString generates a new CUID2 identifier as a string.
func (g *generator) GenerateString() (string, error) {
	return cuid2.Generate(), nil
}

// GenerateNumber is not supported by CUID2 and returns an error.
func (g *generator) GenerateNumber() (int64, error) {
	return 0, errors.New("cuid2: does not support number generation")
}

// Validate checks if the provided id is a valid CUID2 string.
func (g *generator) Validate(id any) bool {
	idStr, ok := id.(string)
	if !ok {
		return false
	}
	return cuid2.IsCuid(idStr)
}

// init registers the CUID2 generator.
func init() {
	identifier.Register(New())
}
