/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package ulid

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
)

func TestULID(t *testing.T) {
	s := New()
	id := s.Generate()
	assert.NotEmpty(t, id)
}

func TestULIDValidateValidID(t *testing.T) {
	s := New()
	id := s.Generate()
	valid := s.Validate(id)
	assert.True(t, valid)
}

func TestULIDValidateInvalidID(t *testing.T) {
	s := New()
	invalidID := "invalidID"
	valid := s.Validate(invalidID)
	assert.False(t, valid)
}

func TestULIDSize(t *testing.T) {
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

	// Check that the default identifier is of type ULID
	if defaultIdentifier.Name() != "ulid" {
		t.Errorf("Expected default identifier to be 'ulid', but got '%s'", defaultIdentifier.Name())
	}
}

func TestGenerateID(t *testing.T) {
	// Generateerate an ID
	gen := identifier.DefaultString()
	generatedID := gen.GenerateString()

	// Check that the generated ID is valid
	if !gen.ValidateString(generatedID) {
		t.Errorf("Generateerated ID is not valid")
	}
}

func TestGenerateSize(t *testing.T) {
	identifier.SetDefaultString(New())
	gen := identifier.DefaultString()
	// Check that the size of the generated ID is correct
	if gen.Size() != bitSize {
		t.Errorf("Expected size of generated ID to be %d, but it was %d", bitSize, gen.Size())
	}
}

func TestValidate(t *testing.T) {
	// Create a new identifier
	generator := New()

	// Generateerate an ID
	generatedID := generator.GenerateString()

	// Check that the generated ID is valid
	if !generator.ValidateString(generatedID) {
		t.Errorf("Generateerated ID is not valid")
	}

	// Check that an invalid ID is not valid
	if generator.ValidateString("invalid") {
		t.Errorf("Invalid ID is valid")
	}
}
