/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package ulid_test // Use black-box testing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	"github.com/origadmin/toolkits/identifier/ulid"
	// Blank import to trigger the ulid provider registration
	_ "github.com/origadmin/toolkits/identifier/ulid"
)

// TestGeneratorCreation ensures the ulid provider is registered correctly.
func TestGeneratorCreation(t *testing.T) {
	// Test getting from the global registry
	registryGenerator := identifier.Get[string]("ulid")
	assert.NotNil(t, registryGenerator, "Expected to get a non-nil string generator for 'ulid'")

	// Test getting via the convenience function
	convenienceGenerator := ulid.New()
	assert.NotNil(t, convenienceGenerator, "Expected to get a non-nil generator from ulid.New()")

	// Number generator should NOT be available.
	numGenerator := identifier.Get[int64]("ulid")
	assert.Nil(t, numGenerator, "Expected to get a nil number generator for 'ulid' as it is not supported")
}

// TestGenerateAndValidate tests the generation and validation of a ULID.
func TestGenerateAndValidate(t *testing.T) {
	generator := ulid.New()
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	// 1. Generate a new ID
	id := generator.Generate()
	assert.NotEmpty(t, id, "Generated ID should not be empty")
	assert.Len(t, id, 26, "A standard ULID should have a length of 26 characters")

	// 2. Validate the new ID
	isValid := generator.Validate(id)
	assert.True(t, isValid, "A freshly generated ID should be valid")

	// 3. Validate a known invalid ID
	isInvalid := generator.Validate("not_a_valid_ulid_string")
	assert.False(t, isInvalid, "An invalid string should not be considered a valid ULID")
}

// TestGeneratorProperties checks the metadata of the ulid generator.
func TestGeneratorProperties(t *testing.T) {
	generator := ulid.New()
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	assert.Equal(t, "ulid", generator.Name(), "Generator name should be 'ulid'")
	assert.Equal(t, 128, generator.Size(), "Generator size should be 128 bits")
}
