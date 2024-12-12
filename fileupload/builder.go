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

	"github.com/goexts/generic/settings"
	"github.com/google/uuid"
)

// Builder interface defines methods for creating uploaders and receivers
type Builder interface {
	// NewUploader creates a new uploader instance
	NewUploader(ctx context.Context) Uploader

	// NewReceiver creates a new receiver instance
	NewReceiver(ctx context.Context) Receiver

	// Free releases resources
	Free(buf []byte)

	// NewBuffer allocates a new buffer
	NewBuffer() []byte
}

// uploadBuilder implements the Builder interface
type uploadBuilder struct {
	services    map[ServiceType]Builder
	uri         string
	hash        func(string) string
	client      *http.Client
	bufPool     *sync.Pool
	bufSize     int
	serviceType ServiceType
}

func (b *uploadBuilder) NewUploader(ctx context.Context) Uploader {
	switch b.serviceType {
	case ServiceTypeGRPC:
		return NewGRPCUploader(ctx, b.uri)
	default:
		return NewHTTPUploader(ctx, b.uri)
	}
}

func (b *uploadBuilder) NewReceiver(ctx context.Context) Receiver {
	switch b.serviceType {
	case ServiceTypeGRPC:
		return NewGRPCReceiver(ctx)
	default:
		return NewHTTPReceiver(ctx)
	}
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

func GenerateRandomHash() string {
	id := uuid.Must(uuid.NewRandom())
	return hex.EncodeToString(id[:])
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

func WithServiceType(st ServiceType) BuildSetting {
	return func(o *uploadBuilder) {
		o.serviceType = st
	}
}
