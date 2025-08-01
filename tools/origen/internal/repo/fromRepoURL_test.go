package repo

import (
	"testing"
)

// Successfully parses a valid URL and returns the correct GOPATH format
func TestParseValidURL(t *testing.T) {
	// Arrr, let's see if this URL be parsed correctly!
	remoteURL := "https://github.com/user/repo"
	expected := "github.com/user/repo"

	result, err := fromRepoURL(remoteURL)

	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Successfully parses a valid URL and returns the correct GOPATH format
func TestParseValidURL2(t *testing.T) {
	// Arrr, let's see if this URL be parsed correctly!
	remoteURL := "https://github.com/user/repo.git"
	expected := "github.com/user/repo"

	result, err := fromRepoURL(remoteURL)

	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Successfully parses a valid URL and returns the correct GOPATH format
func TestParseValidURL3(t *testing.T) {
	// Arrr, let's see if this URL be parsed correctly!
	remoteURL := "https://github.com/User/Repo.git"
	expected := "github.com/!user/!repo"

	result, err := fromRepoURL(remoteURL)

	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Correctly handles URLs with standard path structures
func TestStandardPathStructure(t *testing.T) {
	// Aye, let's see if the standard path structure be handled!
	remoteURL := "https://bitbucket.org/user/repo"
	expected := "bitbucket.org/user/repo"

	result, err := fromRepoURL(remoteURL)

	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Returns the expected GOPATH format for typical repository URLs
func TestTypicalRepoURL(t *testing.T) {
	// Avast! Let's check if the typical URL returns the right GOPATH!
	remoteURL := "https://gitlab.com/user/repo"
	expected := "gitlab.com/user/repo"

	result, err := fromRepoURL(remoteURL)

	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Handles URLs with less than three path segments gracefully
func TestLessThanThreeSegments(t *testing.T) {
	// Shiver me timbers! What happens with less than three segments?
	remoteURL := "https://github.com/user"

	_, err := fromRepoURL(remoteURL)

	if err == nil {
		t.Fatal("Expected an error, but got none")
	}
}

// Returns an error for URLs with invalid formats
func TestInvalidURLFormat(t *testing.T) {
	// Arrr! This URL be invalid, matey!
	remoteURL := "htp://invalid-url"

	_, err := fromRepoURL(remoteURL)

	if err == nil {
		t.Fatal("Expected an error, but got none")
	}
}

// Manages URLs with unusual but valid characters in the path
func TestUnusualCharactersInPath(t *testing.T) {
	// Yo ho ho! Let's see how it handles unusual characters!
	remoteURL := "https://github.com/user/repo-with_special.chars"
	expected := "github.com/user/repo-with_special.chars"

	result, err := fromRepoURL(remoteURL)

	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

// Processes URLs with trailing slashes correctly
func TestTrailingSlashInURL(t *testing.T) {
	// Ahoy! Let's see if it handles trailing slashes like a true sailor!
	remoteURL := "https://github.com/user/repo/"
	expected := "github.com/user/repo"

	result, err := fromRepoURL(remoteURL)

	if err != nil {
		t.Fatalf("Expected no error, but got %v", err)
	}

	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
