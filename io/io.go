/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package io is the input/output package
package io

import (
	"bytes"
	"context"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/origadmin/toolkits/errors"
)

const (
	SeekStart   = io.SeekStart
	SeekCurrent = io.SeekCurrent
	SeekEnd     = io.SeekEnd
)

const (
	ErrFileSize       = errors.String("file Size is zero")
	ErrFileType       = errors.String("file type is not supported")
	ErrTargetIsNotDir = errors.String("target is not a directory")
	ErrFileRead       = errors.String("file read error")
	ErrFileWrite      = errors.String("file write error")
	ErrFileNotExist   = errors.String("file not exist")
	ErrFileName       = errors.String("file path is empty")
	ErrSizeNotMatch   = errors.String("file size not match")
)

var (
	EOF              = io.EOF
	ErrNoProgress    = io.ErrNoProgress
	ErrUnexpectedEOF = io.ErrUnexpectedEOF
	ErrShortBuffer   = io.ErrShortBuffer
	ErrShortWrite    = io.ErrShortWrite
)

var (
	Discard        = io.Discard
	StdWriteString = io.WriteString
	StdCopy        = io.Copy
	StdCopyBuffer  = io.CopyBuffer
	StdCopyN       = io.CopyN
	StdLimitReader = io.LimitReader
	StdReadAll     = io.ReadAll
	StdReadFull    = io.ReadFull
	StdReadAtLeast = io.ReadAtLeast
	StdMultiReader = io.MultiReader
	StdMultiWriter = io.MultiWriter
	StdNopCloser   = io.NopCloser
	StdPipe        = io.Pipe
)

type (
	OffsetWriter  = io.OffsetWriter
	SectionReader = io.SectionReader
)

type (
	Reader          = io.Reader
	Writer          = io.Writer
	ReaderFrom      = io.ReaderFrom
	WriterTo        = io.WriterTo
	ReaderAt        = io.ReaderAt
	WriterAt        = io.WriterAt
	Closer          = io.Closer
	ByteReader      = io.ByteReader
	ByteScanner     = io.ByteScanner
	ByteWriter      = io.ByteWriter
	RuneReader      = io.RuneReader
	RuneScanner     = io.RuneScanner
	Seeker          = io.Seeker
	ReadCloser      = io.ReadCloser
	WriteCloser     = io.WriteCloser
	ReadWriteCloser = io.ReadWriteCloser
	ReadWriteSeeker = io.ReadWriteSeeker
	ReadSeeker      = io.ReadSeeker
	ReadSeekCloser  = io.ReadSeekCloser
	StringWriter    = io.StringWriter
	ReadDirFile     = fs.ReadDirFile
	DirEntry        = fs.DirEntry
	FileInfo        = fs.FileInfo
)

// CopyContext copy data from src to dst until EOF or error, returning the number of bytes copied.
func CopyContext(ctx context.Context, dst io.Writer, src io.Reader) (int64, error) {
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

// ReadBuffer read file to buffer
func ReadBuffer(path string, buf *bytes.Buffer) (int64, error) {
	reader, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = reader.Close()
	}()
	return io.Copy(buf, reader)
}

// DeleteFile deletes files related to the provided id and param.
func DeleteFile(path string) error {
	path = filepath.FromSlash(path)
	if path == "" {
		return ErrFileName
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
