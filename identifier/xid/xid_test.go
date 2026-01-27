/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package xid_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	// Blank import to trigger the xid generator registration
	_ "github.com/origadmin/toolkits/identifier/xid"
)

func TestXIDGenerator(t *testing.T) {
	// Get the generator from the registry.
	// Note: The default generator is 'snowflake', but 'xid' should still be registered
	// by its own package's init() function.
	gen := identifier.Get("xid")
	if !assert.NotNil(t, gen, "Generator 'xid' should be registered") {
		t.FailNow()
	}

	t.Run("Properties", func(t *testing.T) {
		assert.Equal(t, "xid", gen.Name(), "Name() should be 'xid'")
	})

	t.Run("GenerateString", func(t *testing.T) {
		id, err := gen.GenerateString()
		assert.NoError(t, err, "GenerateString() should not return an error")
		assert.NotEmpty(t, id, "Generated string ID should not be empty")
	})

	t.Run("GenerateNumber", func(t *testing.T) {
		// XID's GenerateNumber returns the timestamp part.
		num, err := gen.GenerateNumber()
		assert.NoError(t, err, "GenerateNumber() should not return an error")
		assert.NotZero(t, num, "Generated number (timestamp) should not be zero")
	})

	t.Run("Validate", func(t *testing.T) {
		id, _ := gen.GenerateString()

		// Test with a valid string ID
		assert.True(t, gen.Validate(id), "Validate() should return true for a valid string ID")

		// Test with an invalid string ID
		assert.False(t, gen.Validate("invalid-xid"), "Validate() should return false for an invalid string")

		// Test with an incorrect type (int64) - xid validation is string-based
		assert.False(t, gen.Validate(int64(12345)), "Validate() should return false for an int64")

		// Test with nil
		assert.False(t, gen.Validate(nil), "Validate() should return false for nil")
	})
}
