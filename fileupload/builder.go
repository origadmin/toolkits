/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package fileupload implements the functions, types, and interfaces for the module.
package fileupload

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/goexts/generic/settings"
)

// uploadBuilder implements the Builder interface
type uploadBuilder struct {
	buildType string
	uri       string
	hash      func(string) string
	client    *http.Client
	bufPool   *sync.Pool
	bufSize   int
}

func (b uploadBuilder) Free(buf []byte) {
	buf = buf[:0]
	b.bufPool.Put(buf)
}

func (b uploadBuilder) NewBuffer() []byte {
	buf := b.bufPool.Get().([]byte)
	return buf
}

type BuildSetting = func(o *uploadBuilder)

func WithURI(uri string) BuildSetting {
	return func(o *uploadBuilder) {
		o.uri = uri
	}
}

func WithHash(hash func(name string) string) BuildSetting {
	return func(o *uploadBuilder) {
		o.hash = hash
	}
}

func GenerateFileNameHash(data string) string {
	hash := sha256.New()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

// GenerateContentHash generates a content-addressed hash for a file.
func GenerateContentHash(file io.Reader) (string, error) {
	hash := sha256.New()
	_, err := io.Copy(hash, file)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

// NewBuilder creates a new httpBuilder with the given options
func NewBuilder(ss ...BuildSetting) Builder {
	b := settings.Apply(&uploadBuilder{
		hash:    GenerateFileNameHash,
		bufSize: bufSize, // default 32kb
	}, ss)
	// initialize the buffer pool
	b.bufPool = &sync.Pool{
		New: func() interface{} {
			return make([]byte, b.bufSize)
		},
	}
	return b
}

func WithBufferSize(size int) BuildSetting {
	return func(o *uploadBuilder) {
		o.bufSize = size
	}
}

func NewUploader(ctx context.Context, ss ...BuildSetting) (Uploader, error) {
	b := settings.Apply(&uploadBuilder{}, ss)
	return &httpUploader{
		builder: b,
		uri:     b.uri,
	}, nil
}

func NewDownloader(req *http.Request, resp http.ResponseWriter, ss ...BuildSetting) (Downloader, error) {
	b := settings.Apply(&uploadBuilder{}, ss)
	// Read the file header from the request
	file, header, err := req.FormFile("file")
	if err != nil {
		return nil, err
	}

	fileheader := make(map[string]string, len(header.Header))
	for k, v := range header.Header {
		fileheader[k] = v[0]
	}
	modTime := uint32(time.Now().Unix())
	if mod, err := time.Parse(ModTimeFormat, header.Header.Get("Last-Modified")); err != nil {
		modTime = uint32(mod.Unix())
	}

	// Create a FileHeader struct and populate it with the file information
	fileHeader := &httpFileHeader{
		Filename:    header.Filename,
		Size:        uint32(header.Size),
		ContentType: header.Header.Get("Content-Type"),
		Header:      fileheader,
		ModTime:     modTime,
	}

	return &httpDownloader{
		builder:  b,
		file:     file,
		response: resp,
		header:   fileHeader,
	}, nil
}
