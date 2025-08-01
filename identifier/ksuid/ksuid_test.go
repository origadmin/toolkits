/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package shortid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
)

func init() {
	identifier.SetDefaultString(New())
}
func TestShortidentifier(t *testing.T) {
	s := New()
	id := s.Generate()
	assert.NotEmpty(t, id)
}

func TestShortIDValidateStringValidID(t *testing.T) {
	s := New()
	id := s.Generate()
	valid := s.ValidateString(id)
	assert.True(t, valid)
}

func TestShortIDValidateStringInvalidID(t *testing.T) {
	s := New()
	invalidID := "invalidID_shortid"
	valid := s.ValidateString(invalidID)
	assert.False(t, valid)
}

func TestShortIDSize(t *testing.T) {
	s := New()
	size := s.Size()
	assert.Equal(t, bitSize, size)
}

func TestRegister(t *testing.T) {
	if identifier.DefaultString() == nil {
		t.Errorf("Expected default identifier to be set, but it was not")
	}
	// Check that the default identifier has been updated
	if identifier.DefaultString().Name() != "ksuid" {
		t.Errorf("Expected default identifier to be updated, but it was %s", identifier.DefaultString().Name())
	}
}

func TestGenID(t *testing.T) {
	// Generate an ID
	gen := identifier.DefaultString()
	generatedID := gen.GenerateString()

	// Check that the generated ID is valid
	if !identifier.Validate(gen.Name(), generatedID) {
		t.Errorf("Generated ID is not valid")
	}
}

func TestGenSize(t *testing.T) {
	gen := identifier.DefaultString()
	// Check that the size of the generated ID is correct
	if gen.Size() != bitSize {
		t.Errorf("Expected size of generated ID to be %d, but it was %d", bitSize, gen.Size())
	}
}

func TestValidateString(t *testing.T) {
	// Create a new identifier
	generator := New()

	// Generate an ID
	generatedID := generator.GenerateString()
	fmt.Println(generatedID, len(generatedID), generator.Size())
	// Check that the generated ID is valid
	if !identifier.Validate(generator.Name(), generatedID) {
		t.Errorf("Generated ID is not valid")
	}

	// Check that an invalid ID is not valid
	if identifier.Validate("invalid", generator.Name()) {
		t.Errorf("Invalid ID is valid")
	}
}
