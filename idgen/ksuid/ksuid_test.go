package shortid

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/origadmin/toolkits/idgen"
)

func TestShortIDGen(t *testing.T) {
	s := New()
	id := s.Gen()
	assert.NotEmpty(t, id)
}

func TestShortIDValidateValidID(t *testing.T) {
	s := New()
	id := s.Gen()
	valid := s.Validate(id)
	assert.True(t, valid)
}

func TestShortIDValidateInvalidID(t *testing.T) {
	s := New()
	invalidID := "invalidID_shortid"
	valid := s.Validate(invalidID)
	assert.False(t, valid)
}

func TestShortIDSize(t *testing.T) {
	s := New()
	size := s.Size()
	assert.Equal(t, bitSize, size)
}

func TestRegister(t *testing.T) {
	if idgen.Default() == nil {
		t.Errorf("Expected default identifier to be set, but it was not")
	}
	// Check that the default identifier has been updated
	if idgen.Default().Name() != "shortid" {
		t.Errorf("Expected default identifier to be updated, but it was not")
	}
}

func TestGenID(t *testing.T) {
	// Generate an ID
	generatedID := idgen.GenID()

	// Check that the generated ID is valid
	if !idgen.Validate(generatedID) {
		t.Errorf("Generated ID is not valid")
	}
}

func TestGenSize(t *testing.T) {
	// Check that the size of the generated ID is correct
	if idgen.Size() != bitSize {
		t.Errorf("Expected size of generated ID to be %d, but it was %d", bitSize, idgen.Size())
	}
}

func TestValidate(t *testing.T) {
	// Create a new identifier
	generator := New()

	// Generate an ID
	generatedID := generator.Gen()
	fmt.Println(generatedID, len(generatedID), generator.Size())
	// Check that the generated ID is valid
	if !idgen.Validate(generatedID) {
		t.Errorf("Generated ID is not valid")
	}

	// Check that an invalid ID is not valid
	if idgen.Validate("invalid") {
		t.Errorf("Invalid ID is valid")
	}
}
