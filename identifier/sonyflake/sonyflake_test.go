/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package sonyflake_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	// Blank import to trigger the sonyflake generator registration
	_ "github.com/origadmin/toolkits/identifier/sonyflake"
)

func TestSonyflakeGenerator(t *testing.T) {
	// Get the generator from the registry
	gen := identifier.Get("sonyflake")
	if !assert.NotNil(t, gen, "Generator 'sonyflake' should be registered") {
		t.FailNow()
	}

	t.Run("Properties", func(t *testing.T) {
		assert.Equal(t, "sonyflake", gen.Name(), "Name() should be 'sonyflake'")
	})

	t.Run("GenerateString", func(t *testing.T) {
		id, err := gen.GenerateString()
		assert.NoError(t, err, "GenerateString() should not return an error")
		assert.NotEmpty(t, id, "Generated string ID should not be empty")
	})

	t.Run("GenerateNumber", func(t *testing.T) {
		id, err := gen.GenerateNumber()
		assert.NoError(t, err, "GenerateNumber() should not return an error")
		assert.NotZero(t, id, "Generated number ID should not be zero")
	})

	t.Run("Validate", func(t *testing.T) {
		strID, _ := gen.GenerateString()
		numID, _ := gen.GenerateNumber()

		// Test with a valid string ID
		assert.True(t, gen.Validate(strID), "Validate() should return true for a valid string ID")

		// Test with a valid int64 ID
		assert.True(t, gen.Validate(numID), "Validate() should return true for a valid int64 ID")

		// Test with an invalid string ID
		assert.False(t, gen.Validate("invalid-sonyflake"), "Validate() should return false for an invalid string")

		// Test with an invalid int64 ID (e.g., 0 or negative)
		assert.False(t, gen.Validate(int64(0)), "Validate() should return false for 0")
		assert.False(t, gen.Validate(int64(-1)), "Validate() should return false for a negative int64")

		// Test with an incorrect type
		assert.False(t, gen.Validate(123.45), "Validate() should return false for a float")

		// Test with nil
		assert.False(t, gen.Validate(nil), "Validate() should return false for nil")
	})
}
