/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package ksuid_test // Use black-box testing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	"github.com/origadmin/toolkits/identifier/ksuid"
	// Blank import to trigger the ksuid provider registration
	_ "github.com/origadmin/toolkits/identifier/ksuid"
)

// TestGeneratorCreation ensures the ksuid provider is registered correctly.
func TestGeneratorCreation(t *testing.T) {
	// Test getting from the global registry
	registryGenerator := identifier.Get[string]("ksuid")
	assert.NotNil(t, registryGenerator, "Expected to get a non-nil string generator for 'ksuid'")

	// Test getting via the convenience function
	convenienceGenerator := ksuid.New()
	assert.NotNil(t, convenienceGenerator, "Expected to get a non-nil generator from ksuid.New()")

	// Number generator should NOT be available.
	numGenerator := identifier.Get[int64]("ksuid")
	assert.Nil(t, numGenerator, "Expected to get a nil number generator for 'ksuid' as it is not supported")
}

// TestGenerateAndValidate tests the generation and validation of a KSUID.
func TestGenerateAndValidate(t *testing.T) {
	generator := identifier.Get[string]("ksuid")
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	// 1. Generate a new ID
	id := generator.Generate()
	assert.NotEmpty(t, id, "Generated ID should not be empty")
	assert.Len(t, id, 27, "A standard KSUID should have a length of 27 characters")

	// 2. Validate the new ID
	isValid := generator.Validate(id)
	assert.True(t, isValid, "A freshly generated ID should be valid")

	// 3. Validate a known invalid ID
	isInvalid := generator.Validate("not_a_valid_ksuid")
	assert.False(t, isInvalid, "An invalid string should not be considered a valid KSUID")
}

// TestGeneratorProperties checks the metadata of the ksuid generator.
func TestGeneratorProperties(t *testing.T) {
	generator := identifier.Get[string]("ksuid")
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	assert.Equal(t, "ksuid", generator.Name(), "Generator name should be 'ksuid'")
	assert.Equal(t, 160, generator.Size(), "Generator size should be 160 bits")
}
