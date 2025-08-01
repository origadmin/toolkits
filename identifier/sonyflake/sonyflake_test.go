/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package shortid

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
)

func TestSonyflakeGenerate(t *testing.T) {
	s := New()
	id := s.Generate()
	assert.NotEmpty(t, id)
}

func TestSonyflakeValidateValidID(t *testing.T) {
	s := New()
	id := s.Generate()
	valid := s.Validate(id)
	assert.Truef(t, valid, "ID %s is not valid", id)
}

func TestSonyflakeValidateInvalidID(t *testing.T) {
	s := New()
	invalidID := int64(0)
	valid := s.Validate(invalidID)
	assert.False(t, valid)
}

func TestSonyflakeSize(t *testing.T) {
	s := New()
	size := s.Size()
	assert.Equal(t, bitSize, size)
}

func TestRegister(t *testing.T) {
	identifier.SetDefaultNumber(New())
	// Check that the default identifier is set
	defaultIdentifier := identifier.DefaultNumber()
	if defaultIdentifier == nil {
		t.Fatal("Expected default identifier to be set, but it was not")
	}

	// Check that the default identifier is of type Sonyflake
	if defaultIdentifier.Name() != "sonyflake" {
		t.Errorf("Expected default identifier to be 'sonyflake', but got '%s'", defaultIdentifier.Name())
	}
}

func TestGenerateID(t *testing.T) {
	identifier.SetDefaultNumber(New())
	// Generateerate an ID
	generatedID := identifier.DefaultNumber().GenerateNumber()

	// Check that the generated ID is valid
	if !identifier.DefaultNumber().ValidateNumber(generatedID) {
		t.Errorf("Generateerated ID is not valid")
	}
}

func TestGenerateSize(t *testing.T) {
	identifier.SetDefaultNumber(New())
	// Check that the size of the generated ID is correct
	if identifier.DefaultNumber().Size() != bitSize {
		t.Errorf("Expected size of generated ID to be %d, but it was %d", bitSize, identifier.DefaultNumber().Size())
	}
}

func TestValidate(t *testing.T) {
	// Create a new identifier
	generator := New()

	// Generateerate an ID
	generatedID := generator.Generate()
	// Check that the generated ID is valid
	if !generator.Validate(generatedID) {
		t.Errorf("Generateerated ID is not valid")
	}

	// Check that an invalid ID is not valid
	if generator.Validate(0) {
		t.Errorf("Invalid ID is valid")
	}
}
