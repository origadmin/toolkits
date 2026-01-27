/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package ulid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	// Blank import to trigger the ulid generator registration
	_ "github.com/origadmin/toolkits/identifier/ulid"
)

func TestULIDGenerator(t *testing.T) {
	// Get the generator from the registry
	gen := identifier.Get("ulid")
	if !assert.NotNil(t, gen, "Generator 'ulid' should be registered") {
		t.FailNow()
	}

	t.Run("Properties", func(t *testing.T) {
		assert.Equal(t, "ulid", gen.Name(), "Name() should be 'ulid'")
	})

	t.Run("GenerateString", func(t *testing.T) {
		id, err := gen.GenerateString()
		assert.NoError(t, err, "GenerateString() should not return an error")
		assert.NotEmpty(t, id, "Generated string ID should not be empty")
		assert.Len(t, id, 26, "A standard ULID should have a length of 26 characters")
	})

	t.Run("GenerateNumber", func(t *testing.T) {
		// ULID's GenerateNumber returns the timestamp part, which is a valid use case.
		num, err := gen.GenerateNumber()
		assert.NoError(t, err, "GenerateNumber() should not return an error")
		assert.NotZero(t, num, "Generated number (timestamp) should not be zero")
	})

	t.Run("Validate", func(t *testing.T) {
		id, _ := gen.GenerateString()

		// Test with a valid string ID
		assert.True(t, gen.Validate(id), "Validate() should return true for a valid string ID")

		// Test with an invalid string ID
		assert.False(t, gen.Validate("invalid-ulid"), "Validate() should return false for an invalid string")

		// Test with an incorrect type (int64) - ulid validation is string-based
		assert.False(t, gen.Validate(int64(12345)), "Validate() should return false for an int64")

		// Test with nil
		assert.False(t, gen.Validate(nil), "Validate() should return false for nil")
	})
}
