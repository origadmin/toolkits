// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package cache provides a set of caching utilities for Go applications.
package cache

import (
	"context"
	"time"

	"github.com/goexts/ggb/settings"
)

// DefaultDelimiter is the default delimiter used in cache keys.
const DefaultDelimiter = ":"

// Cache is the interface that wraps the basic Get, Set, and Delete methods.
type Cache[T any] interface {
	// Get retrieves the value associated with the given key.
	Get(ctx context.Context, key string) (T, error)

	// GetAndDelete retrieves the value associated with the given key and deletes it.
	GetAndDelete(ctx context.Context, key string) (T, error)

	// Exists checks if a value exists for the given key.
	Exists(ctx context.Context, key string) error

	// Set sets the value for the given key.
	Set(ctx context.Context, key string, value T, expiration ...time.Duration) error

	// Delete deletes the value associated with the given key.
	Delete(ctx context.Context, key string) error

	// Iterator iterates over all key-value pairs in the cache.
	Iterator(ctx context.Context, fn func(ctx context.Context, key string, value T) bool) error

	// Close releases any resources associated with the cache.
	Close(ctx context.Context) error
}

// ObjectCache is the interface that wraps the basic Get, Set, and Delete methods for object values.
type ObjectCache interface {
	// Get retrieves the value associated with the given key.
	Get(ctx context.Context, key string) (any, error)

	// GetAndDelete retrieves the value associated with the given key and deletes it.
	GetAndDelete(ctx context.Context, key string) (any, error)

	// Exists checks if a value exists for the given key.
	Exists(ctx context.Context, key string) error

	// Set sets the value for the given key.
	Set(ctx context.Context, key string, value any, expiration ...time.Duration) error

	// Delete deletes the value associated with the given key.
	Delete(ctx context.Context, key string) error

	// Iterator iterates over all key-value pairs in the cache.
	Iterator(ctx context.Context, fn func(ctx context.Context, key string, value any) bool) error

	// Close releases any resources associated with the cache.
	Close(ctx context.Context) error
}

// NSCache is the interface that wraps the basic Get, Set, and Delete methods for namespaced values.
type NSCache[T any] interface {
	// Get retrieves the value associated with the given key in the specified namespace.
	Get(ctx context.Context, ns, key string) (T, error)

	// GetAndDelete retrieves the value associated with the given key in the specified namespace and deletes it.
	GetAndDelete(ctx context.Context, ns, key string) (T, error)

	// Exists checks if a value exists for the given key in the specified namespace.
	Exists(ctx context.Context, ns, key string) error

	// Set sets the value for the given key in the specified namespace.
	Set(ctx context.Context, ns, key string, value T, expiration ...time.Duration) error

	// Delete deletes the value associated with the given key in the specified namespace.
	Delete(ctx context.Context, ns, key string) error

	// Iterator iterates over all key-value pairs in the specified namespace.
	Iterator(ctx context.Context, ns string, fn func(ctx context.Context, key string, value T) bool) error

	// Close releases any resources associated with the cache.
	Close(ctx context.Context) error
}

// ObjectNSCache is the interface that wraps the basic Get, Set, and Delete methods for namespaced object values.
type ObjectNSCache interface {
	// Get retrieves the value associated with the given key in the specified namespace.
	Get(ctx context.Context, ns, key string) (any, error)

	// GetAndDelete retrieves the value associated with the given key in the specified namespace and deletes it.
	GetAndDelete(ctx context.Context, ns, key string) (any, error)

	// Exists checks if a value exists for the given key in the specified namespace.
	Exists(ctx context.Context, ns, key string) error

	// Set sets the value for the given key in the specified namespace.
	Set(ctx context.Context, ns, key string, value any, expiration ...time.Duration) error

	// Delete deletes the value associated with the given key in the specified namespace.
	Delete(ctx context.Context, ns, key string) error

	// Iterator iterates over all key-value pairs in the specified namespace.
	Iterator(ctx context.Context, ns string, fn func(ctx context.Context, key string, value any) bool) error

	// Close releases any resources associated with the cache.
	Close(ctx context.Context) error
}

// Option specifies configuration options for the cache.
type Option struct {
	// Joint is the separator used in cache key generation.
	Concat func(ns, key string) string
}

// WithConcat sets the delimiter option.
func WithConcat(concat func(ns, key string) string) settings.Setting[Option] {
	return func(o *Option) {
		o.Concat = concat
	}
}

// DefaultOption returns the default option.
func DefaultOption() Option {
	return Option{
		Concat: func(ns, key string) string {
			return ns + DefaultDelimiter + key
		},
	}
}
