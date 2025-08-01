/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package snowflake

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
)

func TestSnowflakeGen(t *testing.T) {
	s := New()
	id := s.GenerateString()
	assert.NotEmpty(t, id)
}

func TestSnowflakeValidateValidID(t *testing.T) {
	s := New()
	id := s.GenerateString()
	valid := s.ValidateString(id)
	assert.True(t, valid)
}

func TestSnowflakeValidateInvalidID(t *testing.T) {
	s := New()
	invalidID := "invalidID"
	valid := s.ValidateString(invalidID)
	assert.False(t, valid)
}

func TestSnowflakeSize(t *testing.T) {
	s := New()
	size := s.Size()
	assert.Equal(t, bitSize, size)
}

func TestRegister(t *testing.T) {
	identifier.SetDefaultString(New())
	// Check that the default identifier is set
	defaultIdentifier := identifier.DefaultString()
	if defaultIdentifier == nil {
		t.Fatal("Expected default identifier to be set, but it was not")
	}

	// Check that the default identifier is of type Snowflake
	if defaultIdentifier.Name() != "snowflake" {
		t.Errorf("Expected default identifier to be 'snowflake', but got '%s'", defaultIdentifier.Name())
	}
}

func TestGenID(t *testing.T) {
	identifier.SetDefaultString(New())
	// Generate an ID
	generatedID := identifier.DefaultString().GenerateString()

	// Check that the generated ID is valid
	if !identifier.DefaultString().ValidateString(generatedID) {
		t.Errorf("Generated ID is not valid")
	}
}

func TestGenSize(t *testing.T) {
	identifier.SetDefaultString(New())
	// Check that the size of the generated ID is correct
	if identifier.DefaultString().Size() != bitSize {
		t.Errorf("Expected size of generated ID to be %d, but it was %d", bitSize, identifier.DefaultString().Size())
	}
}

func TestValidate(t *testing.T) {
	// Create a new identifier
	generator := New()

	// Generate an ID
	generatedID := generator.GenerateString()

	// Check that the generated ID is valid
	if !generator.ValidateString(generatedID) {
		t.Errorf("Generated ID is not valid")
	}

	// Check that an invalid ID is not valid
	if generator.ValidateString("invalid") {
		t.Errorf("Invalid ID is valid")
	}
}
