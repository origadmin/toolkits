/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package snowflake_test // Use black-box testing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	// Blank import to trigger the snowflake provider registration
	_ "github.com/origadmin/toolkits/identifier/snowflake"
)

// TestSnowflakeProvider ensures the snowflake provider is registered correctly
// and provides generators for both string and int64.
func TestSnowflakeProvider(t *testing.T) {
	// Both string and number generators should be available.
	strGenerator := identifier.New[string]("snowflake")
	assert.NotNil(t, strGenerator, "Expected to get a non-nil string generator for 'snowflake'")

	numGenerator := identifier.New[int64]("snowflake")
	assert.NotNil(t, numGenerator, "Expected to get a non-nil number generator for 'snowflake'")
}

// TestGenerateAndValidateNumber tests the number-based snowflake generator.
func TestGenerateAndValidateNumber(t *testing.T) {
	generator := identifier.New[int64]("snowflake")
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
	generator := identifier.New[string]("snowflake")
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
	strGenerator := identifier.New[string]("snowflake")
	numGenerator := identifier.New[int64]("snowflake")
	if !assert.NotNil(t, strGenerator) || !assert.NotNil(t, numGenerator) {
		t.FailNow()
	}

	assert.Equal(t, "snowflake", strGenerator.Name())
	assert.Equal(t, 64, strGenerator.Size())

	assert.Equal(t, "snowflake", numGenerator.Name())
	assert.Equal(t, 64, numGenerator.Size())
}
