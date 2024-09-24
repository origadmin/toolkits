// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package uuid provides a UUID (Universally Unique Identifier) implementation.
package uuid

import (
	"github.com/google/uuid"

	"github.com/origadmin/toolkits/ident"
)

// UUID represents a UUID generator
type UUID struct{}

var (
	bitSize = len(uuid.New().String())
)

// init registers UUID as an identifier
func init() {
	ident.Register(New())
}

// Name returns the name of the UUID generator
func (U UUID) Name() string {
	return "uuid"
}

// Gen generates a new UUID
func (U UUID) Gen() string {
	return uuid.Must(uuid.NewRandom()).String()
}

// Validate checks if the provided ID is a valid UUID
func (U UUID) Validate(id string) bool {
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

// New creates a new UUID instance
func New() *UUID {
	return &UUID{}
}
