/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package shortid

import (
	"github.com/segmentio/ksuid"

	"github.com/origadmin/toolkits/identifier"
)

var (
	bitSize = len(ksuid.New().String()) // bitSize is used to store the length of generated ID.
)

// init registers the Snowflake generator with the ident package and initializes bitSize.
func init() {
	s := New()
	bitSize = len(s.GenerateString())
	identifier.RegisterStringIdentifier(s)
}

type KSUID struct {
	generator ksuid.KSUID
}

func (s KSUID) ValidateString(id string) bool {
	if len(id) != bitSize {
		return false
	}
	_, err := ksuid.Parse(id)
	return err == nil
}

func (s KSUID) GenerateString() string {
	return s.generator.String()
}

// Name returns the name of the generator.
func (s KSUID) Name() string {
	return "ksuid"
}

// Generate generates a new KSUID ID as a string.
func (s KSUID) Generate() string {
	return s.GenerateString()
}

// Validate checks if the provided ID is a valid KSUID ID.
func (s KSUID) Validate(id string) bool {
	return s.ValidateString(id)
}

// Size returns the bit size of the generated KSUID ID.
func (s KSUID) Size() int {
	return bitSize
}

type Options struct {
}

// New creates a new KSUID generator with a unique node.
func New(_ ...Options) *KSUID {
	generator, err := ksuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return &KSUID{
		generator: generator,
	}
}

var _ identifier.TypedIdentifier[string] = &KSUID{}
