/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package nanoid_test // Use black-box testing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	// Blank import to trigger the nanoid provider registration
	_ "github.com/origadmin/toolkits/identifier/nanoid"
)

// TestNanoidProvider ensures the nanoid provider is registered correctly.
func TestNanoidProvider(t *testing.T) {
	// String generator should be available.
	strGenerator := identifier.New[string]("nanoid")
	assert.NotNil(t, strGenerator, "Expected to get a non-nil string generator for 'nanoid'")

	// Number generator should NOT be available.
	numGenerator := identifier.New[int64]("nanoid")
	assert.Nil(t, numGenerator, "Expected to get a nil number generator for 'nanoid' as it is not supported")
}

// TestGenerateAndValidate tests the generation and validation of a NanoID.
func TestGenerateAndValidate(t *testing.T) {
	generator := identifier.New[string]("nanoid")
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	// 1. Generate a new ID
	id := generator.Generate()
	assert.NotEmpty(t, id, "Generated ID should not be empty")
	assert.Len(t, id, 21, "A standard NanoID should have a length of 21 characters")

	// 2. Validate the new ID
	isValid := generator.Validate(id)
	assert.True(t, isValid, "A freshly generated ID should be valid")

	// 3. Validate a known invalid ID (wrong length and contains invalid char '%')
	isInvalid := generator.Validate("invalid-nanoid-%")
	assert.False(t, isInvalid, "An invalid string should not be considered a valid NanoID")
}

// TestGeneratorProperties checks the metadata of the nanoid generator.
func TestGeneratorProperties(t *testing.T) {
	generator := identifier.New[string]("nanoid")
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	assert.Equal(t, "nanoid", generator.Name(), "Generator name should be 'nanoid'")
	// Size is variable for nanoid, so we expect 0.
	assert.Equal(t, 0, generator.Size(), "Generator size should be 0 for variable-length IDs")
}
