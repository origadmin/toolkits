/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package io provides I/O utility functions, extending the standard io package.
package io

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

//go:generate adptool .
//go:adapter:package io

// Errors returned by this package.
var (
	// ErrZeroSize indicates that a file's size is zero.
	ErrZeroSize = errors.New("file Size is zero")
	// ErrUnsupportedType indicates that a file's type is not supported.
	ErrUnsupportedType = errors.New("file type is not supported")
	// ErrTargetIsNotDir indicates that a target path for an operation is not a directory.
	ErrTargetIsNotDir = errors.New("target is not a directory")
	// ErrRead indicates a failure to read from a file.
	ErrRead = errors.New("file read error")
	// ErrWrite indicates a failure to write to a file.
	ErrWrite = errors.New("file write error")
	// ErrNotExist indicates that a file or directory does not exist.
	ErrNotExist = errors.New("file not exist")
	// ErrEmptyPath indicates that a file path is empty.
	ErrEmptyPath = errors.New("file path is empty")
	// ErrSizeNotMatch indicates that the number of bytes written does not match the expected size.
	ErrSizeNotMatch = errors.New("file size not match")
)

// Standard library fs types.
type (
	ReadDirFile = fs.ReadDirFile
	DirEntry    = fs.DirEntry
	FileInfo    = fs.FileInfo
)

// Copy copies data from src to dst until EOF, context cancellation, or an error occurs.
// It returns the number of bytes copied.
func Copy(ctx context.Context, dst io.Writer, src io.Reader) (int64, error) {
	var (
		err   error
		n     int
		total int64
	)

	buf := make([]byte, 1024)
	for {
		select {
		case <-ctx.Done():
			return 0, ctx.Err()
		default:
			n, err = src.Read(buf)
			if err != nil {
				if err == io.EOF {
					return total, nil
				}
				return 0, err
			}
			if n > 0 {
				if n, err = dst.Write(buf[:n]); err != nil {
					return 0, err
				}
				total += int64(n)
			}
		}
	}
}

// ReadToBuffer reads the entire file at the given path into the provided buffer.
func ReadToBuffer(path string, buf *bytes.Buffer) (int64, error) {
	reader, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = reader.Close()
	}()
	return io.Copy(buf, reader)
}

// Delete removes the file at the specified path.
// If the path is empty, it returns ErrEmptyPath.
// If the file does not exist, it returns no error.
func Delete(path string) error {
	path = filepath.FromSlash(path)
	if path == "" {
		return ErrEmptyPath
	}
	abspath, err := filepath.Abs(path)
	if err != nil {
		return err
	}
	if err := os.Remove(abspath); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

// Save reads all data from a reader and saves it to the specified destination path.
// It automatically creates any necessary parent directories for the destination path.
func Save(ctx context.Context, dstPath string, src io.Reader) (int64, error) {
	dir := filepath.Dir(dstPath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return 0, err
		}
	}

	f, err := os.Create(dstPath)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = f.Close()
	}()

	return Copy(ctx, f, src)
}

// Stream opens the file at the source path and streams its content to the destination writer.
// This function is memory-efficient and suitable for large files.
func Stream(ctx context.Context, srcPath string, dst io.Writer) (int64, error) {
	f, err := os.Open(srcPath)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = f.Close()
	}()

	return Copy(ctx, dst, f)
}

// Exists checks if a file or directory exists at the given path.
// It returns true if the path exists, and false if it does not.
// An error is returned only if the check itself fails due to permission issues or other system errors.
func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
