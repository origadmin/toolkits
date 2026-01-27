/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package snowflake provides a unique ID generator based on the Twitter Snowflake algorithm.
package snowflake

import (
	"math/rand/v2"

	"github.com/bwmarrin/snowflake"

	"github.com/origadmin/toolkits/identifier"
)

// generator implements the identifier.Generator interface for Snowflake.
type generator struct {
	node *snowflake.Node
}

// New creates a new Snowflake generator with a random node ID.
// It panics if node creation fails.
func New() identifier.Generator {
	// Create a snowflake node with a random ID.
	nodeID := rand.Int64N(1024)
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		panic("snowflake: failed to initialize node: " + err.Error())
	}
	return &generator{node: node}
}

// Name returns the name of the generator.
func (g *generator) Name() string {
	return "snowflake"
}

// GenerateString generates a new Snowflake ID as a string.
func (g *generator) GenerateString() (string, error) {
	return g.node.Generate().String(), nil
}

// GenerateNumber generates a new Snowflake ID as an int64.
func (g *generator) GenerateNumber() (int64, error) {
	return g.node.Generate().Int64(), nil
}

// Validate checks if the provided id is a valid Snowflake ID.
// It supports both int64 and string representations.
func (g *generator) Validate(id any) bool {
	switch v := id.(type) {
	case int64:
		// A valid Snowflake ID should always be a positive number.
		return v > 0
	case string:
		_, err := snowflake.ParseString(v)
		return err == nil
	default:
		return false
	}
}

// init registers this Snowflake generator with the identifier registry.
// The registration in the parent 'identifier' package for the default generator
// runs first, so this registration will effectively be a no-op due to the duplicate name check.
func init() {
	tryRegister()
}

func tryRegister() {
	// The registry panics on duplicate registration. We can recover from it
	// as it's an expected condition in this specific setup.
	defer func() {
		_ = recover()
	}()
	identifier.Register(New())
}
