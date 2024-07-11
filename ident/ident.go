// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package ident provides the helpers functions.
package ident

// Identifier is the interface of ident.
type Identifier interface {
	Name() string
	Gen() string
	Validate(id string) bool
	Size() int
}
