/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package cuid2_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
	"github.com/origadmin/toolkits/identifier/cuid2"
	// Blank import to trigger the cuid2 provider registration
	_ "github.com/origadmin/toolkits/identifier/cuid2"
)

// TestDefaultProvider ensures the default cuid2 provider is registered correctly.
func TestDefaultProvider(t *testing.T) {
	generator := identifier.New[string]("cuid2")
	assert.NotNil(t, generator, "Expected to get a non-nil string generator for 'cuid2'")

	id := generator.Generate()
	assert.NotEmpty(t, id)
	assert.True(t, generator.Validate(id))
}

// TestCustomGenerator tests the creation of a generator with custom settings.
func TestCustomGenerator(t *testing.T) {
	t.Run("WithCustomConfig", func(t *testing.T) {
		// Test that providing a custom config returns an error, as it's not supported.
		cfg := cuid2.Config{
			Length:      32,
			Fingerprint: "my-app-worker-1",
		}
		generator, err := cuid2.NewGenerator(cfg)
		assert.Error(t, err)
		assert.Nil(t, generator)
	})

	t.Run("WithDefaultConfig", func(t *testing.T) {
		// Test that a default (empty) config works correctly.
		generator, err := cuid2.NewGenerator(cuid2.Config{})
		assert.NoError(t, err)
		assert.NotNil(t, generator)

		id := generator.Generate()
		assert.NotEmpty(t, id)
		assert.True(t, generator.Validate(id))
	})
}

// TestGeneratorProperties checks the metadata of the cuid2 generator.
func TestGeneratorProperties(t *testing.T) {
	generator := identifier.New[string]("cuid2")
	if !assert.NotNil(t, generator, "Generator should not be nil") {
		t.FailNow()
	}

	assert.Equal(t, "cuid2", generator.Name(), "Generator name should be 'cuid2'")
	assert.Equal(t, 0, generator.Size(), "Generator size should be 0 for variable-length IDs")
}

// TestValidationFailure checks that validation correctly identifies an invalid CUID2.
func TestValidationFailure(t *testing.T) {
	generator := identifier.New[string]("cuid2")
	if !assert.NotNil(t, generator) {
		t.FailNow()
	}
	assert.False(t, generator.Validate("this-is-not-a-cuid2"))
}
