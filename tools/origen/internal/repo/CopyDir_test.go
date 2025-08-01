package repo

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/origadmin/toolkits/context"
)

// Successfully copies all files and directories from source to destination
func TestCopyDirSuccess(t *testing.T) {
	// Arrr, set sail for the source and destination!
	src := "test_cache_dir/github.com/goexts/ggb"
	dst := "test_data_dir/github.com/goexts/ggb"

	// Hoist the CopyDir function!
	err := CopyDir(context.Background(), src, dst)
	if err != nil {
		t.Fatalf("Failed to copy directory: %v", err)
	}

	// Check if the file be copied to the destination
	_, err = os.Stat(filepath.Join(dst, "go.mod"))
	if os.IsNotExist(err) {
		t.Fatalf("File not copied to destination")
	}
}

// Skips copying the .git directory
func TestCopyDirSkipGit(t *testing.T) {
	// Avast! Prepare the source with a .git directory!
	src := "test_cache_dir/github.com/goexts/ggb"
	dst := "test_data_dir/github.com/goexts/ggb"

	// Hoist the CopyDir function!
	err := CopyDir(context.Background(), src, dst)
	if err != nil {
		t.Fatalf("Failed to copy directory: %v", err)
	}

	// Check if the .git directory be not copied
	_, err = os.Stat(filepath.Join(dst, ".git"))
	if !os.IsNotExist(err) {
		t.Fatalf(".git directory should not be copied")
	}
}

// Creates directories in the destination as needed
func TestCopyDirCreateDirectories(t *testing.T) {
	// Arrr, prepare the source with nested treasures!
	src := "test_cache_dir/github.com/goexts/ggb"
	dst := "test_data_dir/github.com/goexts/ggb"

	// Hoist the CopyDir function!
	err := CopyDir(context.Background(), src, dst)
	if err != nil {
		t.Fatalf("Failed to copy directory: %v", err)
	}

	// Check if the nested directories be created in the destination
	_, err = os.Stat(filepath.Join(dst, "settings"))
	if os.IsNotExist(err) {
		t.Fatalf("Nested directories not created in destination")
	}
}

// Source directory does not exist
func TestCopyDirSourceNotExist(t *testing.T) {
	// Shiver me timbers! The source be missing!
	src := "test_cache_dir/github.com/goexts/noggb"
	dst := "test_data_dir/github.com/goexts/ggb"

	// Hoist the CopyDir function!
	err := CopyDir(context.Background(), src, dst)

	// Expect an error when the source be missing
	if err == nil {
		t.Fatalf("Expected error when source does not exist")
	}
}

// Destination directory does not exist
func TestCopyDirDestinationNotExist(t *testing.T) {
	// Arrr, prepare the source with a file!
	src := "test_cache_dir/github.com/goexts/ggb"
	dst := "test_data_dir/github.com/goexts/noggb"

	// Hoist the CopyDir function!
	err := CopyDir(context.Background(), src, dst)

	// Check if it creates the destination and copies the file
	if err != nil {
		t.Fatalf("Failed to copy directory: %v", err)
	}

	_, err = os.Stat(filepath.Join(dst, "go.mod"))
	if os.IsNotExist(err) {
		t.Fatalf("File not copied to destination")
	}
}
