/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package uuid_test // Use black-box testing

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	// Blank import to trigger the uuid provider registration
	_ "github.com/origadmin/toolkits/identifier/uuid"
)

// TestUUIDProviders ensures all registered UUID providers work as expected.
func TestUUIDProviders(t *testing.T) {
	// A list of all registered UUID providers to test.
	providerNames := []string{"uuid", "uuid-v1", "uuid-v4", "uuid-v6", "uuid-v7"}

	for _, name := range providerNames {
		// Run tests in a subtest to get clear output for each provider.
		t.Run(fmt.Sprintf("Provider_%s", name), func(t *testing.T) {
			// 1. Check provider availability
			generator := identifier.New[string](name)
			if !assert.NotNil(t, generator, "Expected to get a non-nil string generator for '%s'", name) {
				t.FailNow()
			}

			// 2. Check that number generator is not supported
			numGenerator := identifier.New[int64](name)
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

// TestUUIDValidationFailure specifically tests the validation logic with a bad input.
func TestUUIDValidationFailure(t *testing.T) {
	generator := identifier.New[string]("uuid")
	if !assert.NotNil(t, generator) {
		t.FailNow()
	}
	isInvalid := generator.Validate("not-a-valid-uuid")
	assert.False(t, isInvalid, "An invalid string should not be considered a valid UUID")
}
