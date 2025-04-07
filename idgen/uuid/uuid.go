/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package uuid provides a UUID (Universally Unique Identifier) implementation.
package uuid

import (
	"github.com/google/uuid"

	"github.com/origadmin/toolkits/idgen"
)

// UUID represents a UUID generator
type UUID struct{}

var (
	bitSize = len(uuid.New().String())
)

// init registers UUID as an identifier
func init() {
	idgen.RegisterStringIdentifier(New())
}

// Name returns the name of the UUID generator
func (U UUID) Name() string {
	return "uuid"
}

// Gen generates a new UUID
func (U UUID) String() string {
	return uuid.Must(uuid.NewRandom()).String()
}

// ValidateString checks if the provided ID is a valid UUID
func (U UUID) ValidateString(id string) bool {
	if len(id) != bitSize {
		return false
	}
	_, err := uuid.Parse(id)
	return err == nil
}

// Size returns the size of the UUID in bits
func (U UUID) Size() int {
	return bitSize
}

type Setting struct {
}

// New creates a new UUID instance
func New(_ ...Setting) *UUID {
	return &UUID{}
}
