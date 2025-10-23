/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package io

import (
	"bytes"
	"context"
	"io"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// setupTestDir creates a temporary directory for testing and returns its path.
// It also returns a cleanup function that should be deferred.
func setupTestDir(t *testing.T) (string, func()) {
	tempDir, err := os.MkdirTemp("", "io_test_*")
	assert.NoError(t, err)
	return tempDir, func() {
		_ = os.RemoveAll(tempDir)
	}
}

func TestExists(t *testing.T) {
	tempDir, cleanup := setupTestDir(t)
	defer cleanup()

	// Test with a non-existent file
	exists, err := Exists(filepath.Join(tempDir, "nonexistent.txt"))
	assert.NoError(t, err)
	assert.False(t, exists)

	// Test with an existing file
	testFile := filepath.Join(tempDir, "exists.txt")
	f, err := os.Create(testFile)
	assert.NoError(t, err)
	assert.NoError(t, f.Close()) // Close the file handle immediately

	exists, err = Exists(testFile)
	assert.NoError(t, err)
	assert.True(t, exists)

	// Test with an existing directory
	exists, err = Exists(tempDir)
	assert.NoError(t, err)
	assert.True(t, exists)
}

func TestDelete(t *testing.T) {
	tempDir, cleanup := setupTestDir(t)
	defer cleanup()

	// Test deleting an existing file
	testFile := filepath.Join(tempDir, "delete_me.txt")
	f, err := os.Create(testFile)
	assert.NoError(t, err)
	// This is the crucial fix: close the file handle to release the lock.
	assert.NoError(t, f.Close())

	err = Delete(testFile)
	assert.NoError(t, err)

	exists, err := Exists(testFile)
	assert.NoError(t, err)
	assert.False(t, exists)

	// Test deleting a non-existent file
	err = Delete(filepath.Join(tempDir, "nonexistent.txt"))
	assert.NoError(t, err)

	// Test deleting with an empty path
	err = Delete("")
	assert.ErrorIs(t, err, ErrEmptyPath)
}

func TestSaveAndStream(t *testing.T) {
	tempDir, cleanup := setupTestDir(t)
	defer cleanup()

	content := "hello world, this is a test for save and stream."
	srcReader := strings.NewReader(content)
	dstPath := filepath.Join(tempDir, "sub", "test.txt")

	// Test Save
	n, err := Save(context.Background(), dstPath, srcReader)
	assert.NoError(t, err)
	assert.Equal(t, int64(len(content)), n)

	// Verify file content
	fileBytes, err := os.ReadFile(dstPath)
	assert.NoError(t, err)
	assert.Equal(t, content, string(fileBytes))

	// Test Stream
	var dstBuffer bytes.Buffer
	n, err = Stream(context.Background(), dstPath, &dstBuffer)
	assert.NoError(t, err)
	assert.Equal(t, int64(len(content)), n)
	assert.Equal(t, content, dstBuffer.String())
}

func TestReadToBuffer(t *testing.T) {
	tempDir, cleanup := setupTestDir(t)
	defer cleanup()

	content := "reading to buffer"
	testFile := filepath.Join(tempDir, "read_buffer.txt")
	err := os.WriteFile(testFile, []byte(content), 0644)
	assert.NoError(t, err)

	var buf bytes.Buffer
	n, err := ReadToBuffer(testFile, &buf)
	assert.NoError(t, err)
	assert.Equal(t, int64(len(content)), n)
	assert.Equal(t, content, buf.String())
}

// slowReader is a reader that pauses between reads, useful for testing context cancellation.
type slowReader struct {
	reader io.Reader
	delay  time.Duration
}

func (r *slowReader) Read(p []byte) (n int, err error) {
	time.Sleep(r.delay)
	return r.reader.Read(p)
}

func TestCopy_ContextCancellation(t *testing.T) {
	content := "this is a long string that will be copied slowly"
	src := &slowReader{
		reader: strings.NewReader(content),
		delay:  100 * time.Millisecond,
	}

	var dst bytes.Buffer

	ctx, cancel := context.WithCancel(context.Background())

	// Cancel the context after a short time
	go func() {
		time.Sleep(50 * time.Millisecond)
		cancel()
	}()

	_, err := Copy(ctx, &dst, src)

	// We expect a context cancellation error
	assert.ErrorIs(t, err, context.Canceled)
}
