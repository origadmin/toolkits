// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package cache provides a set of caching utilities for Go applications.
package cache

import (
	"context"
	"time"
)

const DefaultDelimiter = ":"

// Cache is the interface that wraps the basic Get, Set, and Delete methods.
type Cache[T any] interface {
	Get(ctx context.Context, key string) (T, error)
	GetAndDelete(ctx context.Context, key string) (T, error)
	Exists(ctx context.Context, key string) error
	Set(ctx context.Context, key string, value T, expiration ...time.Duration) error
	Delete(ctx context.Context, key string) error
	Iterator(ctx context.Context, fn func(ctx context.Context, key string, value T) bool) error
	Close(ctx context.Context) error
}

// ObjectCache is the interface that wraps the basic Get, Set, and Delete methods.
type ObjectCache interface {
	Get(ctx context.Context, key string) (any, error)
	GetAndDelete(ctx context.Context, key string) (any, error)
	Exists(ctx context.Context, key string) error
	Set(ctx context.Context, key string, value any, expiration ...time.Duration) error
	Delete(ctx context.Context, key string) error
	Iterator(ctx context.Context, fn func(ctx context.Context, key string, value any) bool) error
	Close(ctx context.Context) error
}

// NSCache is the interface that wraps the basic Get, Set, and Delete methods.
type NSCache[T any] interface {
	Get(ctx context.Context, ns, key string) (T, error)
	GetAndDelete(ctx context.Context, ns, key string) (T, error)
	Exists(ctx context.Context, ns, key string) error
	Set(ctx context.Context, ns, key string, value T, expiration ...time.Duration) error
	Delete(ctx context.Context, ns, key string) error
	Iterator(ctx context.Context, ns string, fn func(ctx context.Context, key string, value T) bool) error
	Close(ctx context.Context) error
}

// ObjectNSCache is the interface that wraps the basic Get, Set, and Delete methods.
type ObjectNSCache interface {
	Get(ctx context.Context, ns, key string) (any, error)
	GetAndDelete(ctx context.Context, ns, key string) (any, error)
	Exists(ctx context.Context, ns, key string) error
	Set(ctx context.Context, ns, key string, value any, expiration ...time.Duration) error
	Delete(ctx context.Context, ns, key string) error
	Iterator(ctx context.Context, ns string, fn func(ctx context.Context, key string, value any) bool) error
	Close(ctx context.Context) error
}

type Options struct {
	Delimiter string
}

type Option func(*Options)

func WithDelimiter(delimiter string) Option {
	return func(o *Options) {
		o.Delimiter = delimiter
	}
}
