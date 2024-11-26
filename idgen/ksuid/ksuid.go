/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package shortid

import (
	"github.com/segmentio/ksuid"

	"github.com/origadmin/toolkits/idgen"
)

var (
	bitSize = 27 // bitSize is used to store the length of generated ID.
)

// init registers the Snowflake generator with the ident package and initializes bitSize.
func init() {
	s := New()
	idgen.Register(s)
}

type KSUID struct {
	generator ksuid.KSUID
}

// Name returns the name of the generator.
func (s KSUID) Name() string {
	return "ksuid"
}

// Gen generates a new KSUID ID as a string.
func (s KSUID) Gen() string {
	return s.generator.String()
}

// Validate checks if the provided ID is a valid KSUID ID.
func (s KSUID) Validate(id string) bool {
	if len(id) != bitSize {
		return false
	}
	_, err := ksuid.Parse(id)
	return err == nil
}

// Size returns the bit size of the generated KSUID ID.
func (s KSUID) Size() int {
	return bitSize
}

type Setting struct {
}

// New creates a new KSUID generator with a unique node.
func New(_ ...Setting) *KSUID {

	generator, err := ksuid.NewRandom()
	if err != nil {
		panic(err)
	}
	return &KSUID{
		generator: generator,
	}
}
