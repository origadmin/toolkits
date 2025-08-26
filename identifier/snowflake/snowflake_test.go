/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package snowflake_test // Use black-box testing

import (
	"testing"

	"github.com/bwmarrin/snowflake"
	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	sf "github.com/origadmin/toolkits/identifier/snowflake"
	// Blank import to trigger the snowflake provider registration
	_ "github.com/origadmin/toolkits/identifier/snowflake"
)

// TestDefaultGenerator ensures the default snowflake provider is registered correctly
// and provides generators for both string and int64.
func TestDefaultGenerator(t *testing.T) {
	// Both string and number generators should be available from the registry.
	strGenerator := identifier.Get[string]("snowflake")
	assert.NotNil(t, strGenerator, "Expected to get a non-nil string generator for 'snowflake'")

	numGenerator := identifier.Get[int64]("snowflake")
	assert.NotNil(t, numGenerator, "Expected to get a non-nil number generator for 'snowflake'")
}

// TestGenerateAndValidateNumber tests the number-based snowflake generator.
func TestGenerateAndValidateNumber(t *testing.T) {
	generator := identifier.Get[int64]("snowflake")
	if !assert.NotNil(t, generator, "Number generator should not be nil") {
		t.FailNow()
	}

	id := generator.Generate()
	assert.NotEqual(t, int64(0), id, "Generated number ID should not be zero")

	isValid := generator.Validate(id)
	assert.True(t, isValid, "A freshly generated number ID should be valid")

	isInvalid := generator.Validate(0)
	assert.False(t, isInvalid, "ID '0' should be considered invalid")
}

// TestGenerateAndValidateString tests the string-based snowflake generator.
func TestGenerateAndValidateString(t *testing.T) {
	generator := identifier.Get[string]("snowflake")
	if !assert.NotNil(t, generator, "String generator should not be nil") {
		t.FailNow()
	}

	id := generator.Generate()
	assert.NotEmpty(t, id, "Generated string ID should not be empty")

	isValid := generator.Validate(id)
	assert.True(t, isValid, "A freshly generated string ID should be valid")

	isInvalid := generator.Validate("not-a-snowflake-id")
	assert.False(t, isInvalid, "An invalid string should not be considered a valid snowflake ID")
}

// TestGeneratorProperties checks the metadata of the snowflake generators.
func TestGeneratorProperties(t *testing.T) {
	strGenerator := identifier.Get[string]("snowflake")
	numGenerator := identifier.Get[int64]("snowflake")
	if !assert.NotNil(t, strGenerator) || !assert.NotNil(t, numGenerator) {
		t.FailNow()
	}

	assert.Equal(t, "snowflake", strGenerator.Name())
	assert.Equal(t, 64, strGenerator.Size())

	assert.Equal(t, "snowflake", numGenerator.Name())
	assert.Equal(t, 64, numGenerator.Size())
}

// TestCustomNodeGenerator tests creating a snowflake generator with a specific node ID.
func TestCustomNodeGenerator(t *testing.T) {
	t.Run("ValidNodeID", func(t *testing.T) {
		const nodeID int64 = 478
		cfg := sf.Config{Node: nodeID}

		// Create a new provider with the custom config
		provider, err := sf.New(cfg)
		assert.NoError(t, err)
		assert.NotNil(t, provider)

		// Get a number generator and verify the node ID
		numGenerator := provider.AsNumber()
		assert.NotNil(t, numGenerator)
		generatedID := numGenerator.Generate()
		parsedID := snowflake.ID(generatedID)
		assert.Equal(t, nodeID, parsedID.Node())

		// Get a string generator and verify the node ID
		strGenerator := provider.AsString()
		assert.NotNil(t, strGenerator)
		generatedStrID := strGenerator.Generate()
		parsedStrID, err := snowflake.ParseString(generatedStrID)
		assert.NoError(t, err)
		assert.Equal(t, nodeID, parsedStrID.Node())
	})

	t.Run("InvalidNodeID", func(t *testing.T) {
		// Node ID is out of the valid range (0-1023)
		cfg := sf.Config{Node: 2000}
		_, err := sf.New(cfg)
		assert.Error(t, err, "Expected an error for an out-of-range node ID")
	})
}
