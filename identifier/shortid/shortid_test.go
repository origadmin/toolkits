/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package shortid_test // Use black-box testing

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	"github.com/origadmin/toolkits/identifier/shortid"
	// Blank import to trigger the shortid provider registration
	_ "github.com/origadmin/toolkits/identifier/shortid"
)

// TestGeneratorCreation ensures the generator can be retrieved correctly.
func TestGeneratorCreation(t *testing.T) {
	// Test getting from the global registry
	registryGenerator := identifier.Get[string]("shortid")
	assert.NotNil(t, registryGenerator, "Expected to get a non-nil string generator for 'shortid'")

	// Test getting via the convenience function
	convenienceGenerator := shortid.New()
	assert.NotNil(t, convenienceGenerator, "Expected to get a non-nil generator from shortid.New()")

	// Number generator should NOT be available.
	numGenerator := identifier.Get[int64]("shortid")
	assert.Nil(t, numGenerator, "Expected to get a nil number generator for 'shortid' as it is not supported")
}

// TestGenerateAndValidate tests the generation and validation of a shortid.
func TestGenerateAndValidate(t *testing.T) {
	generator := shortid.New()
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	// 1. Generate a new ID
	id := generator.Generate()
	assert.NotEmpty(t, id, "Generated ID should not be empty")

	// 2. Validate the new ID
	isValid := generator.Validate(id)
	assert.True(t, isValid, "A freshly generated ID should be valid")

	// 3. Validate a known invalid ID (contains a character not in the default alphabet)
	isInvalid := generator.Validate("invalid-id-with-percent-sign-%")
	assert.False(t, isInvalid, "An invalid string should not be considered a valid shortid")
}

// TestGeneratorProperties checks the metadata of the shortid generator.
func TestGeneratorProperties(t *testing.T) {
	generator := shortid.New()
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	assert.Equal(t, "shortid", generator.Name(), "Generator name should be 'shortid'")
	// Size is variable for shortid, so we expect 0.
	assert.Equal(t, 0, generator.Size(), "Generator size should be 0 for variable-length IDs")
}

// TestCustomGenerator tests creating a shortid generator with custom settings.
func TestCustomGenerator(t *testing.T) {
	t.Run("SuccessfulCreation", func(t *testing.T) {
		cfg := shortid.Config{
			Worker:   1,
			Alphabet: "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@", // 64 chars
			Seed:     12345,
		}
		generator, err := shortid.NewGenerator(cfg)
		assert.NoError(t, err)
		assert.NotNil(t, generator)
		id := generator.Generate()
		assert.NotEmpty(t, id)
		assert.True(t, generator.Validate(id))
	})

	t.Run("FailedCreation", func(t *testing.T) {
		// Alphabet is too short, which should cause an error.
		cfg := shortid.Config{
			Alphabet: "abc",
		}
		_, err := shortid.NewGenerator(cfg)
		assert.Error(t, err)
	})
}
