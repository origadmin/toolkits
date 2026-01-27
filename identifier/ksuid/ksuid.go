/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package ksuid provides a K-Sortable Unique Identifier (KSUID) generator.
package ksuid

import (
	"errors"

	"github.com/segmentio/ksuid"

	"github.com/origadmin/toolkits/identifier"
)

// generator implements the identifier.Generator interface for KSUID.
type generator struct{}

// New creates a new KSUID generator.
func New() identifier.Generator {
	return &generator{}
}

// Name returns the name of the generator.
func (g *generator) Name() string {
	return "ksuid"
}

// GenerateString generates a new KSUID identifier as a string.
// It can panic if the system's entropy source is unavailable.
func (g *generator) GenerateString() (string, error) {
	return ksuid.New().String(), nil
}

// GenerateNumber returns the Unix timestamp part of a new KSUID.
// This is not a unique number but provides the time component.
func (g *generator) GenerateNumber() (int64, error) {
	// ksuid.New() can panic, so we need to recover from it to return an error.
	var id ksuid.KSUID
	var err error
	func() {
		defer func() {
			if r := recover(); r != nil {
				err = errors.New("ksuid: failed to generate id, entropy source might be unavailable")
			}
		}()
		id = ksuid.New()
	}()
	if err != nil {
		return 0, err
	}
	return id.Time().Unix(), nil
}

// Validate checks if the provided id is a valid KSUID string.
func (g *generator) Validate(id any) bool {
	idStr, ok := id.(string)
	if !ok {
		return false
	}
	_, err := ksuid.Parse(idStr)
	return err == nil
}

// init registers the KSUID generator.
func init() {
	identifier.Register(New())
}
