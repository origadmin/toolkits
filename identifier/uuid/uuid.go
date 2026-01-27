/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package uuid provides generators for various UUID versions.
package uuid

import (
	"errors"

	"github.com/google/uuid"

	"github.com/origadmin/toolkits/identifier"
)

// generator implements the identifier.Generator interface for UUID.
type generator struct {
	name      string
	generator func() (string, error)
}

// New creates a new UUID generator for a specific version.
func newGenerator(name string, genFunc func() (string, error)) identifier.Generator {
	return &generator{name: name, generator: genFunc}
}

// Name returns the name of the generator (e.g., "uuid-v4").
func (g *generator) Name() string {
	return g.name
}

// GenerateString generates a new UUID string.
func (g *generator) GenerateString() (string, error) {
	return g.generator()
}

// GenerateNumber is not supported by UUID and returns an error.
func (g *generator) GenerateNumber() (int64, error) {
	return 0, errors.New("uuid: does not support number generation")
}

// Validate checks if the provided id is a valid UUID string.
func (g *generator) Validate(id any) bool {
	idStr, ok := id.(string)
	if !ok {
		return false
	}
	_, err := uuid.Parse(idStr)
	return err == nil
}

// init registers generators for different UUID versions.
func init() {
	// Register UUIDv4 as the default "uuid" and also as "uuid-v4".
	// UUIDv4 is the most common, random-based UUID.
	v4gen := newGenerator("uuid-v4", func() (string, error) {
		return uuid.New().String(), nil
	})
	identifier.Register(v4gen)

	// Register a "uuid" alias pointing to the v4 generator.
	aliasGen := newGenerator("uuid", v4gen.GenerateString)
	identifier.Register(aliasGen)

	// Register UUIDv7, the new time-sortable standard.
	identifier.Register(newGenerator("uuid-v7", func() (string, error) {
		id, err := uuid.NewV7()
		return id.String(), err
	}))

	// Register UUIDv6, another time-sortable version.
	identifier.Register(newGenerator("uuid-v6", func() (string, error) {
		id, err := uuid.NewV6()
		return id.String(), err
	}))

	// Register UUIDv1, the classic time-based version.
	identifier.Register(newGenerator("uuid-v1", func() (string, error) {
		id, err := uuid.NewUUID()
		return id.String(), err
	}))
}
