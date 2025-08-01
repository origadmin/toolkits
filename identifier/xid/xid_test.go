/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package xid

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/identifier"
)

func TestXidentifier(t *testing.T) {
	s := New()
	id := s.Generate()
	assert.NotEmpty(t, id)
}

func TestXIDValidateValidID(t *testing.T) {
	s := New()
	id := s.Generate()
	valid := s.Validate(id)
	assert.True(t, valid)
}

func TestXIDValidateInvalidID(t *testing.T) {
	s := New()
	invalidID := "invalidID"
	valid := s.Validate(invalidID)
	assert.False(t, valid)
}

func TestXIDSize(t *testing.T) {
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

	// Check that the default identifier is of type XID
	if defaultIdentifier.Name() != "xid" {
		t.Errorf("Expected default identifier to be 'xid', but got '%s'", defaultIdentifier.Name())
	}
}

func TestGenerateID(t *testing.T) {
	identifier.SetDefaultString(New())
	// Generateerate an ID
	generatedID := identifier.DefaultString().GenerateString()

	// Check that the generated ID is valid
	if !identifier.DefaultString().ValidateString(generatedID) {
		t.Errorf("Generateerated ID is not valid")
	}
}

func TestGenerateSize(t *testing.T) {
	identifier.SetDefaultString(New())
	// Check that the size of the generated ID is correct
	if identifier.DefaultString().Size() != bitSize {
		t.Errorf("Expected size of generated ID to be %d, but it was %d", bitSize, identifier.DefaultString().Size())
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
	if generator.Validate("invalid") {
		t.Errorf("Invalid ID is valid")
	}
}
