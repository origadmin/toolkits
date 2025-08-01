/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package ini

import (
	"testing"
)

const (
	MarshalValidResult = `Name = Jack Sparrow
Age  = 35
`
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

	if string(data) != MarshalValidResult {
		t.Errorf("Expected %s, but got %s", MarshalValidResult, string(data))
	}
}
