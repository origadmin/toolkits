/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package uuid_test // Use black-box testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	"github.com/origadmin/toolkits/identifier/uuid"
	// Blank import to trigger the uuid provider registration
	_ "github.com/origadmin/toolkits/identifier/uuid"
)

// TestRegistryProviders ensures all registered UUID providers work as expected via the central registry.
func TestRegistryProviders(t *testing.T) {
	// A list of all registered UUID providers to test.
	providerNames := []string{"uuid", "uuid-v1", "uuid-v4", "uuid-v6", "uuid-v7"}

	for _, name := range providerNames {
		// Run tests in a subtest to get clear output for each provider.
		t.Run(fmt.Sprintf("Provider_%s", name), func(t *testing.T) {
			// 1. Check provider availability
			generator := identifier.Get[string](name)
			if !assert.NotNil(t, generator, "Expected to get a non-nil string generator for '%s'", name) {
				t.FailNow()
			}

			// 2. Check that number generator is not supported
			numGenerator := identifier.Get[int64](name)
			assert.Nil(t, numGenerator, "Expected to get a nil number generator for '%s'", name)

			// 3. Generate and validate an ID
			id := generator.Generate()
			assert.NotEmpty(t, id, "Generated ID should not be empty")
			isValid := generator.Validate(id)
			assert.True(t, isValid, "A freshly generated ID should be valid")

			// 4. Check properties
			assert.Equal(t, name, generator.Name(), "Generator name should match")
			assert.Equal(t, 128, generator.Size(), "Generator size should be 128 bits")
		})
	}
}

// TestConvenienceConstructors tests that the package-level New...() functions work correctly.
func TestConvenienceConstructors(t *testing.T) {
	tests := []struct {
		name       string
		constructor func() identifier.Generator[string]
	}{
		{"uuid", uuid.New},
		{"uuid-v1", uuid.NewV1},
		{"uuid-v4", uuid.NewV4},
		{"uuid-v6", uuid.NewV6},
		{"uuid-v7", uuid.NewV7},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			generator := tt.constructor()
			assert.NotNil(t, generator)
			id := generator.Generate()
			assert.NotEmpty(t, id)
			assert.True(t, generator.Validate(id))
			assert.Equal(t, tt.name, generator.Name())
		})
	}
}

// TestUUIDValidationFailure specifically tests the validation logic with a bad input.
func TestUUIDValidationFailure(t *testing.T) {
	generator := uuid.New()
	if !assert.NotNil(t, generator) {
		t.FailNow()
	}
	isInvalid := generator.Validate("not-a-valid-uuid")
	assert.False(t, isInvalid, "An invalid string should not be considered a valid UUID")
}
