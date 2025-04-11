/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package ulid provides a ULID (Universally Unique Lexicographically Sortable Identifier) implementation.
package ulid

import (
	"github.com/oklog/ulid/v2"

	"github.com/origadmin/toolkits/identifier"
)

// ULID is a struct that implements the ident.Identifier interface.
type ULID struct{}

func (u ULID) Generate() string {
	return u.GenerateString()
}

func (u ULID) Validate(id string) bool {
	return u.ValidateString(id)
}

// Name returns the name of the ULID implementation.
func (u ULID) Name() string {
	return "ulid"
}

// GenerateString generates a new ULID string.
func (u ULID) GenerateString() string {
	return ulid.Make().String()
}

// ValidateString checks if the given ID is a valid ULID.
func (u ULID) ValidateString(id string) bool {
	if len(id) != bitSize {
		return false
	}
	_, err := ulid.ParseStrict(id)
	return err == nil
}

// bitSize is the size of a ULID string in bits.
var bitSize = len(ulid.Make().String())

// Size returns the size of a ULID string in bits.
func (u ULID) Size() int {
	return bitSize
}

// init registers the ULID implementation with ident.
func init() {
	identifier.RegisterStringIdentifier(New())
}

type Options struct {
}

// New creates a new ULID implementation.
func New(_ ...Options) *ULID {
	return &ULID{}
}

var _ identifier.TypedIdentifier[string] = &ULID{}
