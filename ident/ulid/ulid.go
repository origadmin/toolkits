// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package ulid provides a ULID (Universally Unique Lexicographically Sortable Identifier) implementation.
package ulid

import (
	"github.com/oklog/ulid/v2"

	"github.com/origadmin/toolkits/ident"
)

// ULID is a struct that implements the ident.Identifier interface.
type ULID struct{}

// bitSize is the size of a ULID string in bits.
var bitSize = len(ulid.Make().String())

// init registers the ULID implementation with ident.
func init() {
	ident.Register(New())
}

// Name returns the name of the ULID implementation.
// It implements the ident.Identifier interface.
func (U ULID) Name() string {
	return "ulid"
}

// Gen generates a new ULID string.
// It implements the ident.Identifier interface.
func (U ULID) Gen() string {
	return ulid.Make().String()
}

// Validate checks if the given ID is a valid ULID.
// It implements the ident.Identifier interface.
func (U ULID) Validate(id string) bool {
	if len(id) != bitSize {
		return false
	}
	_, err := ulid.ParseStrict(id)
	return err == nil
}

// Size returns the size of a ULID string in bits.
// It implements the ident.Identifier interface.
func (U ULID) Size() int {
	return bitSize
}

// New creates a new ULID implementation.
func New() *ULID {
	return &ULID{}
}
