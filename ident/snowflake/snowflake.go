// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package snowflake provides a unique ID generation based on Twitter Snowflake algorithm.
package snowflake

import (
	"math/rand/v2"

	"github.com/bwmarrin/snowflake"

	"github.com/origadmin/toolkits/ident"
)

// nodeNumber range from 0 to 1023 is used for generating unique node ID.
var (
	nodeNumber = rand.Int63n(1023)
	bitSize    = 0 // bitSize is used to store the length of generated ID.
)

// Snowflake represents a Snowflake generator with a unique node.
type Snowflake struct {
	node *snowflake.Node
}

// init registers the Snowflake generator with the ident package and initializes bitSize.
func init() {
	s := New()
	bitSize = len(s.Gen())
	ident.Register(s)
}

// Name returns the name of the generator.
func (s Snowflake) Name() string {
	return "snowflake"
}

// Gen generates a new Snowflake ID as a string.
func (s Snowflake) Gen() string {
	return s.node.Generate().String()
}

// Validate checks if the provided ID is a valid Snowflake ID.
func (s Snowflake) Validate(id string) bool {
	_, err := snowflake.ParseString(id)
	return err != nil
}

// Size returns the bit size of the generated Snowflake ID.
func (s Snowflake) Size() int {
	return bitSize
}

// New creates a new Snowflake generator with a unique node.
func New() *Snowflake {
	node, err := snowflake.NewNode(nodeNumber)
	if err != nil {
		panic(err)
	}
	return &Snowflake{
		node: node,
	}
}
