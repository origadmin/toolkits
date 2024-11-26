/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package filter

import (
	"strings"
	"testing"
)

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
