/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package shortid_test // Use black-box testing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	// Blank import to trigger the shortid provider registration
	_ "github.com/origadmin/toolkits/identifier/shortid"
)

// TestShortidProvider ensures the shortid provider is registered correctly.
func TestShortidProvider(t *testing.T) {
	// String generator should be available.
	strGenerator := identifier.New[string]("shortid")
	assert.NotNil(t, strGenerator, "Expected to get a non-nil string generator for 'shortid'")

	// Number generator should NOT be available.
	numGenerator := identifier.New[int64]("shortid")
	assert.Nil(t, numGenerator, "Expected to get a nil number generator for 'shortid' as it is not supported")
}

// TestGenerateAndValidate tests the generation and validation of a shortid.
func TestGenerateAndValidate(t *testing.T) {
	generator := identifier.New[string]("shortid")
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	// 1. Generate a new ID
	id := generator.Generate()
	assert.NotEmpty(t, id, "Generated ID should not be empty")

	// 2. Validate the new ID
	isValid := generator.Validate(id)
	assert.True(t, isValid, "A freshly generated ID should be valid")

	// 3. Validate a known invalid ID (contains a character not in the default alphabet)
	isInvalid := generator.Validate("invalid-id-with-percent-sign-%")
	assert.False(t, isInvalid, "An invalid string should not be considered a valid shortid")
}

// TestGeneratorProperties checks the metadata of the shortid generator.
func TestGeneratorProperties(t *testing.T) {
	generator := identifier.New[string]("shortid")
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	assert.Equal(t, "shortid", generator.Name(), "Generator name should be 'shortid'")
	// Size is variable for shortid, so we expect 0.
	assert.Equal(t, 0, generator.Size(), "Generator size should be 0 for variable-length IDs")
}
