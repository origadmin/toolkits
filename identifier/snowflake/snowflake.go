/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package snowflake provides a unique ID generation based on Twitter Snowflake algorithm.
package snowflake

import (
	"fmt"
	"math/rand/v2"

	"github.com/bwmarrin/snowflake"

	"github.com/origadmin/toolkits/identifier"
)

// Config holds the configuration for creating a new Snowflake node.
type Config struct {
	// Node is the unique node ID for this generator. It must be between 0 and 1023.
	Node int64
}

// Ensure the provider and generators implement the required interfaces at compile time.
var (
	_ identifier.Provider      = (*provider)(nil)
	_ identifier.Generator[int64]  = (*numberGenerator)(nil)
	_ identifier.Generator[string] = (*stringGenerator)(nil)
)

// provider implements identifier.Provider for Snowflake.
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
func (p *provider) AsString() identifier.Generator[string] {
	return &stringGenerator{node: p.node}
}

// AsNumber returns a number-based generator for Snowflake.
func (p *provider) AsNumber() identifier.Generator[int64] {
	return &numberGenerator{node: p.node}
}

// --- Number Generator ---

// numberGenerator implements identifier.Generator[int64] for Snowflake.
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
func (g *numberGenerator) Validate(id int64) bool {
	return id > 0
}

// --- String Generator ---

// stringGenerator implements identifier.Generator[string] for Snowflake.
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

// --- Advanced Usage ---

// New creates a new, local, configured Snowflake provider.
// This instance is NOT managed by the global identifier registry.
func New(cfg Config) (identifier.Provider, error) {
	if cfg.Node < 0 || cfg.Node > 1023 {
		return nil, fmt.Errorf("snowflake node ID %d is out of range (0-1023)", cfg.Node)
	}
	node, err := snowflake.NewNode(cfg.Node)
	if err != nil {
		return nil, fmt.Errorf("snowflake: failed to create new node: %w", err)
	}
	return &provider{node: node}, nil
}

// --- Default Global Instance ---

// init registers the default Snowflake provider with the global identifier registry.
// This provider uses a random node ID.
func init() {
	// Create a snowflake node with a random ID for the default global instance.
	nodeID := rand.Int64N(1024)
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		panic("identifier: failed to initialize default snowflake node: " + err.Error())
	}

	// Register a provider instance containing the default node.
	identifier.Register(&provider{
		node: node,
	})
}
