/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// TestThreadSafeMultiError tests the ThreadSafeMultiError type
package errors

import (
	"errors"
	"reflect"
	"testing"
)

func TestThreadSafeMultiError(t *testing.T) {
	// Create a new ThreadSafeMultiError instance
	tsme := ThreadSafe(errors.New("error 1"))

	// Append additional errors
	tsme.Append(errors.New("error 2"))
	tsme.Append(errors.New("error 3"))

	// Check if there are errors
	if !tsme.HasErrors() {
		t.Error("Expected errors, but got none")
	}

	// Check if a specific error exists
	if err := tsme.Has(errors.New("error 2")); err == nil {
		t.Error("Expected error 2, but got nil")
	}

	// Get the MultiError collection
	me := tsme.Unsafe()

	// Check if the MultiError collection has the expected number of errors
	if len(me.Errors) != 3 {
		t.Errorf("Expected 3 errors, but got %d", len(me.Errors))
	}

	// Check if the MultiError collection contains the expected errors
	expectedErrors := []string{"error 1", "error 2", "error 3"}
	for i, err := range me.Errors {
		if err.Error() != expectedErrors[i] {
			t.Errorf("Expected error %s, but got %s", expectedErrors[i], err.Error())
		}
	}

	// Check the JSON representation of the MultiError collection
	expectedJSON := `["error 1","error 2","error 3"]`
	if me.Error() != expectedJSON {
		t.Errorf("Expected JSON error representation %s, but got %s", expectedJSON, me.Error())
	}

	// Check the errors slice of the MultiError collection
	expectedErrorsSlice := []error{
		errors.New("error 1"),
		errors.New("error 2"),
		errors.New("error 3"),
	}
	if !reflect.DeepEqual(me.Errors, expectedErrorsSlice) {
		t.Errorf("Expected errors slice %v, but got %v", expectedErrorsSlice, me.Errors)
	}
}
