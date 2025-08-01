/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package xid provides a unique identifier generator based on xid.
package xid

import (
	"github.com/rs/xid"

	"github.com/origadmin/toolkits/identifier"
)

var (
	bitSize = len(xid.New().String()) // bitSize is used to store the length of generated ID.
)

func init() {
	identifier.RegisterStringIdentifier(New())
}

// XID represents a unique identifier generator.
type XID struct{}

func (x XID) Generate() string {
	return x.GenerateString()
}

func (x XID) Validate(id string) bool {
	return x.ValidateString(id)
}

// Name returns the name of the identifier.
func (x XID) Name() string {
	return "xid"
}

// GenerateString generates a new unique identifier.
func (x XID) GenerateString() string {
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

type Options struct {
}

// New creates a new instance of the XID identifier.
func New(_ ...Options) *XID {
	return &XID{}
}

var _ identifier.TypedIdentifier[string] = &XID{}
