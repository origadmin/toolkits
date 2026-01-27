/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package sonyflake provides a unique ID generator based on Sony's snowflake variant.
package sonyflake

import (
	"strconv"

	"github.com/sony/sonyflake"

	"github.com/origadmin/toolkits/identifier"
)

// generator implements the identifier.Generator interface for Sonyflake.
type generator struct {
	sf *sonyflake.Sonyflake
}

// New creates a new Sonyflake generator with default settings.
// It panics if the generator cannot be initialized.
func New() identifier.Generator {
	sf := sonyflake.NewSonyflake(sonyflake.Settings{})
	if sf == nil {
		panic("sonyflake: failed to initialize generator")
	}
	return &generator{sf: sf}
}

// Name returns the name of the generator.
func (g *generator) Name() string {
	return "sonyflake"
}

// GenerateString generates a new Sonyflake ID and returns it as a string.
func (g *generator) GenerateString() (string, error) {
	id, err := g.sf.NextID()
	if err != nil {
		return "", err
	}
	return strconv.FormatUint(id, 10), nil
}

// GenerateNumber generates a new Sonyflake ID as an int64.
func (g *generator) GenerateNumber() (int64, error) {
	id, err := g.sf.NextID()
	if err != nil {
		return 0, err
	}
	return int64(id), nil
}

// Validate checks if the provided id is a valid Sonyflake ID.
// A valid ID must be a positive number.
func (g *generator) Validate(id any) bool {
	switch v := id.(type) {
	case int64:
		return v > 0
	case uint64:
		return v > 0
	case string:
		parsed, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			return false
		}
		// A valid ID must be greater than 0.
		return parsed > 0
	default:
		return false
	}
}

// init registers the Sonyflake generator.
func init() {
	identifier.Register(New())
}
