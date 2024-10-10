package filter

import (
	"reflect"
	"testing"
)

// Parses a path with a single delimiter correctly
func TestParsePathSingleDelimiter(t *testing.T) {
	// Arrr, let's split the seas with a single cutlass!
	path := "sea/treasure"
	delimiter := "/"
	methods, parsedPath := parsePath(path, delimiter)

	if len(methods) != 1 || methods[0] != "sea" {
		t.Errorf("Expected wildcard method, got %v", methods)
	}

	if parsedPath != "treasure" {
		t.Errorf("Expected path 'treasure', got %s", parsedPath)
	}

	// Arrr, let's split the seas with a single cutlass!
	path = "get,post:sea/treasure"
	delimiter = ":"
	methods, parsedPath = parsePath(path, delimiter)

	if len(methods) != 2 || methods[0] != "get" || methods[1] != "post" {
		t.Errorf("Expected wildcard method, got %v", methods)
	}

	if parsedPath != "sea/treasure" {
		t.Errorf("Expected path 'treasure', got %s", parsedPath)
	}

	// Arrr, let's split the seas with a single cutlass!
	path = "get:sea/treasure"
	delimiter = ":"
	methods, parsedPath = parsePath(path, delimiter)

	if len(methods) != 1 || methods[0] != "get" {
		t.Errorf("Expected wildcard method, got %v", methods)
	}

	if parsedPath != "sea/treasure" {
		t.Errorf("Expected path 'treasure', got %s", parsedPath)
	}

	// Arrr, let's split the seas with a single cutlass!
	path = ":sea/treasure"
	delimiter = ":"
	methods, parsedPath = parsePath(path, delimiter)

	if len(methods) != 1 || methods[0] != "*" {
		t.Errorf("Expected wildcard method, got %v", methods)
	}

	if parsedPath != "sea/treasure" {
		t.Errorf("Expected path 'treasure', got %s", parsedPath)
	}

	// Arrr, let's split the seas with a single cutlass!
	path = ":/sea/treasure"
	delimiter = ":"
	methods, parsedPath = parsePath(path, delimiter)

	if len(methods) != 1 || methods[0] != "*" {
		t.Errorf("Expected wildcard method, got %v", methods)
	}

	if parsedPath != "sea/treasure" {
		t.Errorf("Expected path 'treasure', got %s", parsedPath)
	}

	// Arrr, let's split the seas with a single cutlass!
	path = ":/sea/treasure"
	delimiter = ":"
	methods, parsedPath = parsePath(path, delimiter)

	if len(methods) != 1 || methods[0] != "*" {
		t.Errorf("Expected wildcard method, got %v", methods)
	}

	if parsedPath != "sea/treasure" {
		t.Errorf("Expected path 'treasure', got %s", parsedPath)
	}
}

// Splits a path with two segments and a single delimiter
func TestParsePathTwoSegments(t *testing.T) {
	// Ahoy! Two segments be better than one!
	path := "sail/ship"
	delimiter := "/"
	methods, parsedPath := parsePath(path, delimiter)

	if len(methods) != 1 || methods[0] != "sail" {
		t.Errorf("Expected method 'sail', got %v", methods)
	}

	if parsedPath != "ship" {
		t.Errorf("Expected path 'ship', got %s", parsedPath)
	}
}

// Returns a wildcard method when only one path segment is provided
func TestParsePathSingleSegment(t *testing.T) {
	// Yarrr! A lone island in the sea!
	path := "island"
	delimiter := "/"
	methods, parsedPath := parsePath(path, delimiter)

	if len(methods) != 1 || methods[0] != "*" {
		t.Errorf("Expected wildcard method, got %v", methods)
	}

	if parsedPath != "island" {
		t.Errorf("Expected path 'island', got %s", parsedPath)
	}
}

// Handles an empty path string gracefully
func TestParsePathEmptyString(t *testing.T) {
	// Avast! The sea be empty!
	path := ""
	delimiter := "/"
	methods, parsedPath := parsePath(path, delimiter)

	if methods != nil || parsedPath != "" {
		t.Errorf("Expected nil methods and empty path, got %v and %s", methods, parsedPath)
	}
}

// Manages a path with multiple delimiters by returning nil and an empty string
func TestParsePathMultipleDelimiters(t *testing.T) {
	// Shiver me timbers! Too many slashes!
	path := "sea//treasure"
	delimiter := "/"
	methods, parsedPath := parsePath(path, delimiter)

	if methods != nil || parsedPath != "" {
		t.Errorf("Expected nil methods and empty path, got %v and %s", methods, parsedPath)
	}
}

// Processes a path with a single segment containing a comma correctly
func TestParsePathSegmentWithComma(t *testing.T) {
	// Arrr! A list of treasures in one spot!
	path := "gold,silver/treasure"
	delimiter := "/"
	methods, parsedPath := parsePath(path, delimiter)

	expectedMethods := []string{"gold", "silver"}

	if !reflect.DeepEqual(methods, expectedMethods) {
		t.Errorf("Expected methods %v, got %v", expectedMethods, methods)
	}

	if parsedPath != "treasure" {
		t.Errorf("Expected path 'treasure', got %s", parsedPath)
	}
}

// Handles paths with leading or trailing delimiters
func TestParsePathLeadingTrailingDelimiters(t *testing.T) {
	// Yo ho ho! Mind the edges of the map!
	path := "/treasure/"
	delimiter := "/"
	methods, parsedPath := parsePath(path, delimiter)

	if methods != nil || parsedPath != "" {
		t.Errorf("Expected nil methods and empty path, got %v and %s", methods, parsedPath)
	}
}
