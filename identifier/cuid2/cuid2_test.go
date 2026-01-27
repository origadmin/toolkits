/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package cuid2_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	// Blank import to trigger the cuid2 generator registration
	_ "github.com/origadmin/toolkits/identifier/cuid2"
)

func TestCUID2Generator(t *testing.T) {
	// Get the generator from the registry
	gen := identifier.Get("cuid2")
	if !assert.NotNil(t, gen, "Generator 'cuid2' should be registered") {
		t.FailNow()
	}

	t.Run("Properties", func(t *testing.T) {
		assert.Equal(t, "cuid2", gen.Name(), "Name() should be 'cuid2'")
	})

	t.Run("GenerateString", func(t *testing.T) {
		id, err := gen.GenerateString()
		assert.NoError(t, err, "GenerateString() should not return an error")
		assert.NotEmpty(t, id, "Generated string ID should not be empty")
	})

	t.Run("GenerateNumber", func(t *testing.T) {
		// CUID2 does not support number generation
		_, err := gen.GenerateNumber()
		assert.Error(t, err, "GenerateNumber() should return an error for cuid2")
	})

	t.Run("Validate", func(t *testing.T) {
		id, _ := gen.GenerateString()

		// Test with a valid string ID
		assert.True(t, gen.Validate(id), "Validate() should return true for a valid string ID")

		// Test with an invalid string ID
		assert.False(t, gen.Validate("invalid-cuid2"), "Validate() should return false for an invalid string")

		// Test with an incorrect type (int64)
		assert.False(t, gen.Validate(int64(12345)), "Validate() should return false for an int64")

		// Test with nil
		assert.False(t, gen.Validate(nil), "Validate() should return false for nil")
	})
}
