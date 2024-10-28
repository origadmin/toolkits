package ini

import (
	"testing"
)

// Successfully marshals a valid struct into INI format bytes
func TestMarshalValidStruct(t *testing.T) {
	// Arrr, let's see if we can turn this shipshape struct into bytes!
	type Config struct {
		Name string
		Age  int
	}
	config := &Config{Name: "Jack Sparrow", Age: 35}
	data, err := Marshal(config)
	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}
	if len(data) == 0 {
		t.Fatal("Expected non-empty byte slice")
	}
	const result = `Name = Jack Sparrow
Age  = 35`
	if string(data) != result {
		t.Errorf("Expected %s, but got %s", result, string(data))
	}
}
