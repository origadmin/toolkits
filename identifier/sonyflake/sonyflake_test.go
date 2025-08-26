/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package sonyflake_test // Corrected package declaration for black-box testing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	// This blank import is necessary to ensure the sonyflake's init() function is called,
	// which registers the provider.
	_ "github.com/origadmin/toolkits/identifier/sonyflake"
)

// TestSonyflakeProvider ensures the sonyflake provider is registered and can be retrieved
// via the public identifier.New API.
func TestSonyflakeProvider(t *testing.T) {
	// Attempt to get the number generator, which should succeed.
	numGenerator := identifier.New[int64]("sonyflake")
	assert.NotNil(t, numGenerator, "Expected to get a non-nil number generator for 'sonyflake'")

	// Ensure the string generator is not supported and returns nil.
	strGenerator := identifier.New[string]("sonyflake")
	assert.Nil(t, strGenerator, "Expected to get a nil string generator for 'sonyflake' as it is not supported")
}

// TestGenerateAndValidate tests the full lifecycle of generating and validating an ID.
func TestGenerateAndValidate(t *testing.T) {
	generator := identifier.New[int64]("sonyflake")
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	// 1. Generate a new ID
	id := generator.Generate()
	assert.NotEqual(t, int64(0), id, "Generated ID should not be zero")

	// 2. Validate the newly generated ID
	isValid := generator.Validate(id)
	assert.True(t, isValid, "A freshly generated ID should be considered valid")

	// 3. Validate a known invalid ID
	isInvalid := generator.Validate(0)
	assert.False(t, isInvalid, "ID '0' should be considered invalid")
}

// TestGeneratorProperties checks the metadata of the sonyflake generator.
func TestGeneratorProperties(t *testing.T) {
	generator := identifier.New[int64]("sonyflake")
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	assert.Equal(t, "sonyflake", generator.Name(), "Generator name should be 'sonyflake'")
	assert.Equal(t, 63, generator.Size(), "Generator size should be 63 bits")
}
