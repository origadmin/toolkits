/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package snowflake provides a unique ID generation based on Twitter Snowflake algorithm.
package snowflake

import (
	"cmp"
	"math/rand/v2"

	"github.com/bwmarrin/snowflake"

	"github.com/origadmin/toolkits/idgen"
)

// nodeNumber range from 0 to 1023 is used for generating unique generator ID.
var (
	bitSize = 0 // bitSize is used to store the length of generated ID.
)

// Snowflake represents a Snowflake generator with a unique generator.
type Snowflake struct {
	generator *snowflake.Node
}

// init registers the Snowflake generator with the ident package and initializes bitSize.
func init() {
	s := New()
	bitSize = len(s.Gen())
	idgen.Register(s)
}

// Name returns the name of the generator.
func (s Snowflake) Name() string {
	return "snowflake"
}

// Gen generates a new Snowflake ID as a string.
func (s Snowflake) Gen() string {
	return s.generator.Generate().String()
}

// Validate checks if the provided ID is a valid Snowflake ID.
func (s Snowflake) Validate(id string) bool {
	if len(id) != bitSize {
		return false
	}
	_, err := snowflake.ParseString(id)
	return err == nil
}

// Size returns the bit size of the generated Snowflake ID.
func (s Snowflake) Size() int {
	return bitSize
}

type Setting struct {
	Node int64
}

// New creates a new Snowflake generator with a unique generator.
func New(ss ...*Setting) *Snowflake {
	ss = append(ss, &Setting{})
	o := cmp.Or(ss...)
	if o.Node < 0 || o.Node > 1023 {
		o.Node = rand.Int64N(1023)
	}
	generator, err := snowflake.NewNode(o.Node)
	if err != nil {
		panic(err)
	}
	return &Snowflake{
		generator: generator,
	}
}
