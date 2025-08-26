/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package snowflake provides a unique ID generation based on Twitter Snowflake algorithm.
package snowflake

import (
	"math/rand/v2"

	"github.com/bwmarrin/snowflake"

	"github.com/origadmin/toolkits/identifier"
)

// Ensure the provider and generators implement the required interfaces at compile time.
var (
	_ identifier.GeneratorProvider      = (*provider)(nil)
	_ identifier.TypedGenerator[int64]  = (*numberGenerator)(nil)
	_ identifier.TypedGenerator[string] = (*stringGenerator)(nil)
)

// provider implements identifier.GeneratorProvider for Snowflake.
// It holds a configured, stateful snowflake node.
type provider struct {
	node *snowflake.Node
}

// Name returns the name of the identifier.
func (p *provider) Name() string {
	return "snowflake"
}

// Size returns the size of the identifier in bits.
func (p *provider) Size() int {
	return 64
}

// AsString returns a string-based generator for Snowflake.
func (p *provider) AsString() identifier.TypedGenerator[string] {
	return &stringGenerator{node: p.node}
}

// AsNumber returns a number-based generator for Snowflake.
func (p *provider) AsNumber() identifier.TypedGenerator[int64] {
	return &numberGenerator{node: p.node}
}

// --- Number Generator ---

// numberGenerator implements identifier.TypedGenerator[int64] for Snowflake.
type numberGenerator struct {
	node *snowflake.Node
}

// Name returns the name of the identifier.
func (g *numberGenerator) Name() string {
	return "snowflake"
}

// Size returns the size of the identifier in bits.
func (g *numberGenerator) Size() int {
	return 64
}

// Generate creates a new Snowflake ID and returns it as an int64.
func (g *numberGenerator) Generate() int64 {
	return g.node.Generate().Int64()
}

// Validate checks if the provided int64 is a plausible Snowflake ID.
// Note: This is a basic check, as any positive int64 could be valid.
func (g *numberGenerator) Validate(id int64) bool {
	return id > 0
}

// --- String Generator ---

// stringGenerator implements identifier.TypedGenerator[string] for Snowflake.
type stringGenerator struct {
	node *snowflake.Node
}

// Name returns the name of the identifier.
func (g *stringGenerator) Name() string {
	return "snowflake"
}

// Size returns the size of the identifier in bits.
func (g *stringGenerator) Size() int {
	return 64
}

// Generate creates a new Snowflake ID and returns it as a string.
func (g *stringGenerator) Generate() string {
	return g.node.Generate().String()
}

// Validate checks if the provided string is a valid Snowflake ID.
func (g *stringGenerator) Validate(id string) bool {
	_, err := snowflake.ParseString(id)
	return err == nil
}

// init registers the Snowflake provider with the global identifier registry.
func init() {
	// Create a snowflake node with a random ID between 0 and 1023.
	// This helps prevent collisions when multiple application instances are running.
	nodeID := rand.Int64N(1024)
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		// This panic is acceptable in an init() function if a core component fails to initialize.
		panic("identifier: failed to initialize snowflake node: " + err.Error())
	}

	// Register a provider instance containing the configured node.
	identifier.Register(&provider{
		node: node,
	})
}
