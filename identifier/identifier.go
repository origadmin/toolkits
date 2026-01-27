/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package identifier provides a simple, thread-safe, and efficient way to generate unique identifiers.
// It offers straightforward functions to get unique strings or numbers without requiring any configuration.
package identifier

import (
	"fmt"
	"math/rand/v2"
	"sync"

	"github.com/bwmarrin/snowflake"
)

// --- Generator Interface ---

// Generator defines the interface for an identifier generation algorithm.
// This is intended for advanced users who want to implement or use custom generators.
type Generator interface {
	// Name returns the unique name of the generator algorithm.
	Name() string
	// GenerateString generates a new identifier as a string.
	GenerateString() (string, error)
	// GenerateNumber generates a new identifier as an int64.
	// Note: May not be supported by all generators.
	GenerateNumber() (int64, error)
	// Validate checks if a provided identifier is valid.
	// The id can be a string or int64, depending on the generator.
	Validate(id any) bool
}

// --- Default Generator (Internal Snowflake Implementation) ---

// defaultSnowflakeGenerator is an internal implementation of the Snowflake algorithm.
// It is used as the default generator for the package.
type defaultSnowflakeGenerator struct {
	node *snowflake.Node
}

func (g *defaultSnowflakeGenerator) Name() string {
	return "snowflake"
}

func (g *defaultSnowflakeGenerator) GenerateString() (string, error) {
	return g.node.Generate().String(), nil
}

func (g *defaultSnowflakeGenerator) GenerateNumber() (int64, error) {
	return g.node.Generate().Int64(), nil
}

// Validate checks if the provided id is a valid Snowflake ID.
// It supports both int64 and string representations.
func (g *defaultSnowflakeGenerator) Validate(id any) bool {
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

// --- Registry ---

var (
	mu         sync.RWMutex
	generators = make(map[string]Generator)
	defaultGen Generator
)

// init registers the default generator.
func init() {
	// Snowflake is chosen as the default because it natively generates int64 IDs
	// that are time-sortable and highly unique. A random node ID is used for
	// out-of-the-box functionality.
	nodeID := rand.Int64N(1024)
	node, err := snowflake.NewNode(nodeID)
	if err != nil {
		panic("identifier: failed to initialize default snowflake node: " + err.Error())
	}

	defaultGen = &defaultSnowflakeGenerator{node: node}
	Register(defaultGen)
}

// Register makes a generator available by name.
// If Register is called twice with the same name or if the generator is nil, it panics.
func Register(g Generator) {
	mu.Lock()
	defer mu.Unlock()
	if g == nil {
		panic("identifier: cannot register a nil generator")
	}
	name := g.Name()
	if _, dup := generators[name]; dup {
		panic(fmt.Sprintf("identifier: generator %q is already registered", name))
	}
	generators[name] = g
}

// Get returns the generator with the given name.
// It returns nil if no generator is registered with that name.
func Get(name string) Generator {
	mu.RLock()
	defer mu.RUnlock()
	return generators[name]
}

// --- Global Functions ---

// GenerateString generates a new identifier string using the default generator.
func GenerateString() (string, error) {
	return defaultGen.GenerateString()
}

// GenerateNumber generates a new identifier number using the default generator.
func GenerateNumber() (int64, error) {
	return defaultGen.GenerateNumber()
}

// Validate checks if the provided id is a valid identifier using the default generator.
func Validate(id any) bool {
	return defaultGen.Validate(id)
}
