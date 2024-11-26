/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package filter

import (
	"testing"
)

// Creates a linked list of Link objects from a valid path string
func TestCreateLinkedListFromValidPath(t *testing.T) {
	// Arrr, let's sail the seas of valid paths!
	path := "home/user/docs"
	link := NewLink(path)

	if link == nil {
		t.Fatal("Expected a Link object, got nil")
	}

	expectedPaths := []string{"home", "user", "docs"}
	current := link
	for _, expectedPath := range expectedPaths {
		if current.Path != expectedPath {
			t.Fatalf("Expected path %s, got %s", expectedPath, current.Path)
		}
		if len(current.Subs) > 0 {
			current = current.Subs[0]
		}
	}
}

// Returns a Link object with correct hierarchy when path has multiple segments
func TestLinkHierarchyWithMultipleSegments(t *testing.T) {
	// Ahoy! Let's check the hierarchy of the links!
	path := "root/branch/leaf"
	link := NewLink(path)

	if link == nil {
		t.Fatal("Expected a Link object, got nil")
	}

	if !link.Has("branch") || !link.Subs[0].Has("leaf") {
		t.Fatal("Hierarchy is not maintained correctly")
	}
}

// Adds a Terminator Link when the last path segment is a Terminator
func TestAddTerminatorLink(t *testing.T) {
	// Shiver me timbers! Let's see if the Terminator is added!
	path := "end/"
	link := NewLink(path)

	if link == nil || len(link.Subs) == 0 || !link.Subs[0].IsEnd() {
		t.Fatal("Expected a Terminator Link, but it was not added", link == nil)
	}
}

// Returns nil when given an empty path string
func TestReturnNilForEmptyPath(t *testing.T) {
	// Avast! An empty path be like a ship with no sails!
	path := ""
	link := NewLink(path)

	if link != nil {
		t.Fatal("Expected nil for an empty path, got a Link object")
	}
}

// Handles paths with only a Terminator correctly
func TestHandleOnlyTerminatorPath(t *testing.T) {
	// Arrr, only the Terminator be on this voyage!
	path := "/"
	link := NewLink(path)

	if link == nil || !link.IsEnd() {
		t.Fatal("Expected a Terminator Link for a path with only a Terminator")
	}
}

// Manages paths with consecutive slashes without creating empty Link objects
func TestManageConsecutiveSlashes(t *testing.T) {
	// Yo ho ho! Let's see how it handles the stormy slashes!
	path := "sea//ocean//wave"
	link := NewLink(path)

	if link != nil {
		t.Fatal("Expected a nil Link object for a path with consecutive slashes")
	}
}

// Correctly processes paths with special characters or spaces
func TestProcessSpecialCharactersAndSpaces(t *testing.T) {
	// Aye aye! Special characters and spaces be no match for us!
	path := "/folder/with special@chars"
	link := NewLink(path)

	if link == nil {
		t.Fatal("Expected for a Link object, but got nil")
	}

	expectedPaths := []string{"folder", "with special@chars"}
	current := link
	for _, expectedPath := range expectedPaths {
		if current.Path != expectedPath {
			t.Fatalf("Expected path %s, got %s", expectedPath, current.Path)
		}
		if len(current.Subs) > 0 {
			current = current.Subs[0]
		}
	}
}
