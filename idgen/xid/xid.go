/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package xid provides a unique identifier generator based on xid.
package xid

import (
	"github.com/rs/xid"

	"github.com/origadmin/toolkits/idgen"
)

// XID represents a unique identifier generator.
type XID struct{}

var (
	// bitSize represents the size of the xid in bits.
	bitSize = len(xid.New().String())
)

// init registers the XID identifier with the ident package.
func init() {
	idgen.RegisterStringIdentifier(New())
}

// Name returns the name of the identifier.
func (x XID) Name() string {
	return "xid"
}

// String generates a new unique identifier.
func (x XID) String() string {
	return xid.New().String()
}

// ValidateString checks if the provided id is a valid xid.
func (x XID) ValidateString(id string) bool {
	if len(id) != bitSize {
		return false
	}
	_, err := xid.FromString(id)
	return err == nil
}

// Size returns the size of the xid in bits.
func (x XID) Size() int {
	return bitSize
}

type Setting struct {
}

// New creates a new instance of the XID identifier.
func New(_ ...Setting) *XID {
	return &XID{}
}
