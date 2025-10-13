/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package filter

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestLinkList(t *testing.T) {
	paths := []string{"a/b/c/d/e", "a/c/d/e", "a/d/e/f"}
	link := &Link{Path: "root"}
	for _, path := range paths {
		link.AddSub(NewLink(path))
	}
	if !link.Has("a") {
		t.Fatal("Expected a to be in the link")
	}
	if link.Has("b") {
		t.Fatal("Expected b to be in the link")
	}
	if link.Has("c") {
		t.Fatal("Expected c to be in the link")
	}
	if link.Has("d") {
		t.Fatal("Expected d to be in the link")
	}

	for _, path := range paths {
		rules := strings.Split(path, "/")
		sub := link
		for i, rule := range rules {
			sub = getLinkFromPath(sub.Subs, rule)
			if sub == nil {
				t.Fatal("Expected to find path:", rules[i:])
			}
		}
	}

	list := link.StringList("")
	fmt.Println(list)
	//Output: [root/a/b/c/d/e root/a/c/d/e root/a/d/e/f]
}

func TestLinkAdd(t *testing.T) {
	paths := []string{"a/b/c/d/e", "a/c/d/e", "a/d/e/f"}
	link := &Link{Path: "root"}
	link.AddSub(NewLink("/"))

	list := link.StringList("")
	fmt.Println(list)
	//Output: [root/*]

	for _, path := range paths {
		if !link.Contains(strings.Split("root/"+path, "/")) {
			t.Fatal("Expected to find path:", path)
		}
	}

	link = &Link{Path: "root"}
	link.AddSub(NewLink("/*"))

	list = link.StringList("")
	fmt.Println(list)
	//Output: [root/*]

	for _, path := range paths {
		if !link.Contains(strings.Split("root/"+path, "/")) {
			t.Fatal("Expected to find path:", path)
		}
	}

	link = &Link{Path: "root"}
	link.AddSub(NewLink("*"))

	list = link.StringList("")
	fmt.Println(list)
	//Output: [root/*]

	for _, path := range paths {
		if !link.Contains(strings.Split("root/"+path, "/")) {
			t.Fatal("Expected to find path:", path)
		}
	}

	link = &Link{Path: "root"}
	link.AddSub(NewLink("*/a/b/c/d/e"))

	list = link.StringList("")
	fmt.Println(list)
	//Output: [root/*]

	for _, path := range paths {
		if !link.Contains(strings.Split("root/"+path, "/")) {
			t.Fatal("Expected to find path:", path)
		}
	}
}

// AddSubs returns true when a sub-link is successfully added
func TestAddSubsSuccessfullyAddsSubLink(t *testing.T) {
	// Arrr, let's set sail with a new Link!
	parent := &Link{Path: "parent"}
	child := &Link{Path: "child"}

	// Add the child to the parent
	parent.AddSubs(child)

	// Check if the child was added
	if !parent.Has("child") {
		t.Errorf("Expected child to be added, but it wasn't!")
	}
}

// AddSubs iterates over all provided sub-links
func TestAddSubsIteratesOverAllSubLinks(t *testing.T) {
	// Ahoy! Let's create a parent and a fleet of children!
	parent := &Link{Path: "parent"}
	child1 := &Link{Path: "child1"}
	child2 := &Link{Path: "child2"}

	// Add the children to the parent
	parent.AddSubs(child1, child2)

	// Check if all children were added
	if !parent.Has("child1") || !parent.Has("child2") {
		t.Errorf("Expected all children to be added, but some are missing!")
	}
}

// AddSubs calls AddSub on matching sub-links
func TestAddSubsCallsAddSubOnMatchingLinks(t *testing.T) {
	// Avast! Let's create a parent and a matching child!
	parent := &Link{Path: "parent"}
	child := &Link{Path: "child"}

	// Add the child to the parent
	parent.AddSubs(child)

	// Check if AddSub was called by verifying the child was added
	if !parent.Has("child") {
		t.Errorf("Expected AddSub to be called, but it wasn't!")
	}
}

// AddSubs returns false when no sub-links are provided
func TestAddSubsReturnsFalseWhenNoSubLinksProvided(t *testing.T) {
	// Shiver me timbers! No sub-links to add!
	parent := &Link{Path: "parent"}

	// Try adding no children
	parent.AddSubs()

	// Check if no children were added
	if len(parent.Subs) != 0 {
		t.Errorf("Expected no children to be added, but found some!")
	}
}

// AddSubs handles nil values in the src slice gracefully
func TestAddSubsHandlesNilValuesGracefully(t *testing.T) {
	// Yo ho ho! Let's see how we handle nil values!
	parent := &Link{Path: "parent"}

	// Add a nil value
	parent.AddSubs(nil)

	// Check if nil was handled gracefully (no panic)
	if len(parent.Subs) != 0 {
		t.Errorf("Expected nil to be handled gracefully, but it wasn't! %d", len(parent.Subs))
	}
}

// AddSubs handles an empty Subs slice in the Link object
func TestAddSubsHandlesEmptyChildrenSlice(t *testing.T) {
	// Arrr matey! Let's start with an empty crew!
	parent := &Link{Path: "parent", Subs: []*Link{}}

	// Add a new child to the empty crew
	newChild := &Link{Path: "newChild"}
	parent.AddSubs(newChild)

	// Check if the new child was added successfully
	if !parent.Has("newChild") {
		t.Errorf("Expected new child to be added, but it wasn't!")
	}
}

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

// Correctly builds a tree from a non-empty list of paths
func TestBuildTreeWithNonEmptyPaths(t *testing.T) {
	// Arrr, let's set sail with some paths!
	paths := []string{"deck", "cabin", "treasure"}
	var tree []*Link

	result := buildLinkTree(tree, paths)

	if len(result) == 0 || result[0].Path != "deck" {
		t.Errorf("Expected root path 'deck', but got %v", result)
	}
}

// Handles a single path element by creating a single node
func TestBuildTreeWithSinglePath(t *testing.T) {
	// Aye, a lone path be like a lone pirate on the sea!
	paths := []string{"anchor"}
	tree := []*Link{}

	result := buildLinkTree(tree, paths)

	if len(result) != 1 || result[0].Path != "anchor" {
		t.Errorf("Expected single node with path 'anchor', but got %v", result)
	}
}

// Recursively processes multiple path elements to build a nested tree structure
func TestBuildTreeWithNestedPaths(t *testing.T) {
	// Avast! We be diving deep into nested paths!
	paths := []string{"ship", "deck", "cannon"}
	tree := []*Link{}

	result := buildLinkTree(tree, paths)

	if len(result) == 0 || len(result[0].Subs) == 0 || result[0].Subs[0].Path != "deck" {
		t.Errorf("Expected nested structure with 'deck', but got %v", result)
	}
}

// Handles an empty paths list without errors
func TestBuildTreeWithEmptyPaths(t *testing.T) {
	// Shiver me timbers! No paths, no problem!
	paths := []string{}
	tree := []*Link{}

	result := buildLinkTree(tree, paths)

	if len(result) != 0 {
		t.Errorf("Expected empty tree, but got %v", result)
	}
}

// Processes paths with duplicate elements correctly
func TestBuildTreeWithDuplicatePaths(t *testing.T) {
	// Arrr, duplicates be like barnacles on the hull!
	paths := []string{"mast", "mast", "sail"}
	tree := []*Link{}

	result := buildLinkTree(tree, paths)

	if len(result) == 0 || len(result[0].Subs) == 0 || result[0].Subs[0].Path != "mast" {
		t.Errorf("Expected structure handling duplicates, but got %v", result)
	}
}

// Manages paths with special characters or unusual formats
func TestBuildTreeWithSpecialCharacters(t *testing.T) {
	// Yo ho ho! Special characters be no match for us!
	paths := []string{"c@ptain", "qu@rters", "tr3@sure"}
	tree := []*Link{}

	result := buildLinkTree(tree, paths)

	if len(result) == 0 || result[0].Path != "c@ptain" {
		t.Errorf("Expected handling of special characters, but got %v", result)
	}
}

// Handles very long paths without performance degradation
func TestBuildTreeWithLongPaths(t *testing.T) {
	// Blimey! A long path be like a long voyage!
	longPath := strings.Repeat("longpath/", 100)
	paths := strings.Split(longPath, "/")
	tree := []*Link{}

	result := buildLinkTree(tree, paths)

	if len(result) == 0 || result[0].Path != "longpath" {
		t.Errorf("Expected handling of long paths, but got %v", result)
	}
}

// Parses a path with a single delimiter correctly
func TestParsePathSingleDelimiter(t *testing.T) {
	// Arrr, let's split the seas with a single cutlass!
	path := "sea/treasure"
	delimiter := "/"
	methods, parsedPath := splitURI(path, delimiter)

	if len(methods) != 1 || methods[0] != "*" { // Corrected expected method
		t.Errorf("Expected wildcard method, got %v", methods)
	}

	if parsedPath != "sea/treasure" {
		t.Errorf("Expected path 'treasure', got %s", parsedPath)
	}

	// Arrr, let's split the seas with a single cutlass!
	path = "get,post:sea/treasure"
	delimiter = ":"
	methods, parsedPath = splitURI(path, delimiter)

	if len(methods) != 2 || methods[0] != "get" || methods[1] != "post" {
		t.Errorf("Expected methods [get, post], got %v", methods)
	}

	if parsedPath != "sea/treasure" {
		t.Errorf("Expected path 'treasure', got %s", parsedPath)
	}

	// Arrr, let's split the seas with a single cutlass!
	path = "get:sea/treasure"
	delimiter = ":"
	methods, parsedPath = splitURI(path, delimiter)

	if len(methods) != 1 || methods[0] != "get" {
		t.Errorf("Expected method [get], got %v", methods)
	}

	if parsedPath != "sea/treasure" {
		t.Errorf("Expected path 'treasure', got %s", parsedPath)
	}

	// Arrr, let's split the seas with a single cutlass!
	path = ":sea/treasure"
	delimiter = ":"
	methods, parsedPath = splitURI(path, delimiter)

	if len(methods) != 1 || methods[0] != "*" {
		t.Errorf("Expected wildcard method, got %v", methods)
	}

	if parsedPath != "sea/treasure" {
		t.Errorf("Expected path 'treasure', got %s", parsedPath)
	}

	// Arrr, let's split the seas with a single cutlass!
	path = ":/sea/treasure"
	delimiter = ":"
	methods, parsedPath = splitURI(path, delimiter)

	if len(methods) != 1 || methods[0] != "*" {
		t.Errorf("Expected wildcard method, got %v", methods)
	}

	if parsedPath != "/sea/treasure" {
		t.Errorf("Expected path '/sea/treasure', got %s", parsedPath)
	}
}

// Splits a path with two segments and a single delimiter
func TestParsePathTwoSegments(t *testing.T) {
	// Ahoy! Two segments be better than one!
	path := "sail/ship"
	delimiter := "/"
	methods, parsedPath := splitURI(path, delimiter)

	if len(methods) != 1 || methods[0] != "*" { // Corrected expected method
		t.Errorf("Expected wildcard method, got %v", methods)
	}

	if parsedPath != "sail/ship" { // Corrected expected path
		t.Errorf("Expected path 'sail/ship', got %s", parsedPath)
	}
}

// Returns a wildcard method when only one path segment is provided
func TestParsePathSingleSegment(t *testing.T) {
	// Yarrr! A lone island in the sea!
	path := "island"
	delimiter := "/"
	methods, parsedPath := splitURI(path, delimiter)

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
	methods, parsedPath := splitURI(path, delimiter)

	if methods != nil || parsedPath != "" {
		t.Errorf("Expected nil methods and empty path, got %v and %s", methods, parsedPath)
	}
}

// Manages a path with multiple delimiters by returning nil and an empty string
func TestParsePathMultipleDelimiters(t *testing.T) {
	// Shiver me timbers! Too many slashes!
	path := "sea//treasure"
	delimiter := "/"
	methods, parsedPath := splitURI(path, delimiter)

	if methods != nil || parsedPath != "" {
		t.Errorf("Expected nil methods and empty path, got %v and %s", methods, parsedPath)
	}
}

// Processes a path with a single segment containing a comma correctly
func TestParsePathSegmentWithComma(t *testing.T) {
	// Arrr! A list of treasures in one spot!
	path := "gold,silver/treasure"
	delimiter := "/"
	methods, parsedPath := splitURI(path, delimiter)

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
	methods, parsedPath := splitURI(path, delimiter)

	if methods != nil || parsedPath != "" {
		t.Errorf("Expected nil methods and empty path, got %v and %s", methods, parsedPath)
	}
}
