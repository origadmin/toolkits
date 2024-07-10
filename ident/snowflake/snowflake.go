package snowflake

import (
	"fmt"

	"github.com/bwmarrin/snowflake"

	"github.com/origadmin/toolkits/ident"
)

// Generator represents a Snowflake-based ID Generator.
type Generator struct {
	node *snowflake.Node // Node is the snowflake node for generating IDs.
	size int             // size indicates the length of the generated ID string.
}

// DefaultNodeNumber is the default node number used to initialize the ID Generator.
var DefaultNodeNumber = int64(1) // Number range from 0 to 1023.

// snowflakeIdentify is the global instance of the Snowflake ID Generator.
var snowflakeIdentify *Generator

// init initializes the global Snowflake ID Generator when the package is first loaded.
func init() {
	// Creates a new Snowflake Node with a node number of 1.
	var err error
	snowflakeIdentify, err = New(DefaultNodeNumber)
	if err != nil {
		// If initialization fails, panic with an error message.
		panic(fmt.Sprintf("snowflake: failed to initialize snowflake: %v", err))
	}
}

// Name returns the name of the Generator, which is "snowflake".
func (i Generator) Name() string {
	return "snowflake"
}

// Gen generates and returns a new Snowflake ID as a string.
func (i Generator) Gen() string {
	return i.node.Generate().String()
}

// Validate checks if the provided ID string is a valid Snowflake ID.
func (i Generator) Validate(id string) bool {
	_, err := snowflake.ParseString(id)
	return err == nil
}

// Size returns the size of the generated Snowflake ID string.
func (i Generator) Size() int {
	return i.size
}

// New creates a new Snowflake Generator with the specified node number.
// It returns a pointer to the Generator and an error if any.
func New(nodeNumber int64) (*Generator, error) {
	node, err := snowflake.NewNode(nodeNumber)
	if err != nil {
		return nil, err
	}
	return &Generator{
		node: node,
		size: len(node.Generate().String()),
	}, nil
}

// Default returns the global Snowflake ID Generator instance.
func Default() ident.Identifier {
	return snowflakeIdentify
}
