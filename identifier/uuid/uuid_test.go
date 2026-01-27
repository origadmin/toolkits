/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package uuid_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	// Blank import to trigger the uuid generator registration
	_ "github.com/origadmin/toolkits/identifier/uuid"
)

func TestUUIDGenerators(t *testing.T) {
	// A list of all registered UUID generator names to test.
	generatorNames := []string{"uuid", "uuid-v1", "uuid-v4", "uuid-v6", "uuid-v7"}

	for _, name := range generatorNames {
		t.Run(fmt.Sprintf("Generator_%s", name), func(t *testing.T) {
			gen := identifier.Get(name)
			if !assert.NotNil(t, gen, "Generator '%s' should be registered", name) {
				t.FailNow()
			}

			t.Run("Properties", func(t *testing.T) {
				assert.Equal(t, name, gen.Name(), "Name() should be '%s'", name)
			})

			t.Run("GenerateString", func(t *testing.T) {
				id, err := gen.GenerateString()
				assert.NoError(t, err, "GenerateString() should not return an error for '%s'", name)
				assert.NotEmpty(t, id, "Generated string ID should not be empty for '%s'", name)
			})

			t.Run("GenerateNumber", func(t *testing.T) {
				// UUIDs do not support number generation
				_, err := gen.GenerateNumber()
				assert.Error(t, err, "GenerateNumber() should return an error for '%s'", name)
			})

			t.Run("Validate", func(t *testing.T) {
				id, _ := gen.GenerateString()

				// Test with a valid string ID
				assert.True(t, gen.Validate(id), "Validate() should return true for a valid string ID for '%s'", name)

				// Test with an invalid string ID
				assert.False(t, gen.Validate("invalid-uuid"), "Validate() should return false for an invalid string for '%s'", name)

				// Test with an incorrect type (int64)
				assert.False(t, gen.Validate(int64(12345)), "Validate() should return false for an int64 for '%s'", name)

				// Test with nil
				assert.False(t, gen.Validate(nil), "Validate() should return false for nil for '%s'", name)
			})
		})
	}
}
