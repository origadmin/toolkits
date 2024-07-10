package ident_test

import (
	"testing"

	"github.com/origadmin/toolkits/ident"
)

type mockIdentifier struct{}

func (m *mockIdentifier) Name() string {
	return "MockIdentifier"
}

func (m *mockIdentifier) Gen() string {
	return "mock-generated-id"
}

func (m *mockIdentifier) Validate(id string) bool {
	return id == "mock-generated-id"
}

func (m *mockIdentifier) Size() int {
	return 12 // Replace with the actual size of the generated identifier
}

func TestGenID(t *testing.T) {
	ident.Use(&mockIdentifier{})
	generatedID := ident.GenID()
	expectedID := "mock-generated-id"

	if generatedID != expectedID {
		t.Errorf("GenID() = %v, want %v", generatedID, expectedID)
	}
}

func TestGenSize(t *testing.T) {
	ident.Use(&mockIdentifier{})
	size := ident.GenSize()
	expectedSize := 12 // Replace with the actual size of the generated identifier

	if size != expectedSize {
		t.Errorf("GenSize() = %v, want %v", size, expectedSize)
	}
}

func TestValidate(t *testing.T) {
	ident.Use(&mockIdentifier{})
	validID := "mock-generated-id"
	invalidID := "invalid-id"

	if !ident.Validate(validID) {
		t.Errorf("Validate(%v) = %v, want %v", validID, false, true)
	}

	if ident.Validate(invalidID) {
		t.Errorf("Validate(%v) = %v, want %v", invalidID, true, false)
	}
}
