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

func TestShortidentifier(t *testing.T) {
	s := New()
	id := s.Generate()
	assert.NotEmpty(t, id)
}

func TestShortIDValidateValidID(t *testing.T) {
	s := New()
	id := s.Generate()
	valid := s.Validate(id)
	assert.True(t, valid)
}

func TestShortIDValidateInvalidID(t *testing.T) {
	s := New()
	invalidID := ""
	valid := s.Validate(invalidID)
	assert.False(t, valid)
}

func TestShortIDSize(t *testing.T) {
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

	// Check that the default identifier is of type ShortID
	if defaultIdentifier.Name() != "shortid" {
		t.Errorf("Expected default identifier to be 'shortid', but got '%s'", defaultIdentifier.Name())
	}
}

func TestGenerateString(t *testing.T) {
	// Generate an ID
	generatedID := identifier.GenerateString("shortid")
	t.Logf("Generated ID: %s", generatedID)
	// Check that the generated ID is valid
	if !identifier.Validate("shortid", generatedID) {
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
	generatedID := generator.Generate()
	fmt.Println(generatedID, len(generatedID), generator.Size())
	// Check that the generated ID is valid
	if !identifier.Validate("shortid", generatedID) {
		t.Errorf("Generated ID is not valid")
	}

	// Check that an invalid ID is not valid
	if !identifier.Validate("shortid", "valid") {
		t.Errorf("Invalid ID is invalid")
	}
}
