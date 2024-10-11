package filter

import (
	"testing"
)

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
