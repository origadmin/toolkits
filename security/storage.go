// Copyright (c) 2024 KasaAdmin. All rights reserved.

// Package security is a toolkit for security check and authorization
package security

import (
	"context"
	"time"

	"github.com/goexts/generic/settings"

	"github.com/origadmin/toolkits/storage/cache"
	"github.com/origadmin/toolkits/storage/cache/memory"
)

const (
	TokenNamespace = "security:jwt:token"
)

// TokenStorage is the interface that tokenStorage the token.
type TokenStorage interface {
	// Set stores the token with a specific expiration time
	Set(ctx context.Context, tokenStr string, expiration time.Duration) error
	// Delete deletes the token from the tokenStorage
	Delete(ctx context.Context, tokenStr string) error
	// Validate checks if the token exists in the tokenStorage
	Validate(ctx context.Context, tokenStr string) error
	// Close closes the tokenStorage
	Close(ctx context.Context) error
}

type StorageSetting = func(*StorageOption)

// StorageOption contains options for the JWT datacache
type StorageOption struct {
	Cache     cache.Cache
	Namespace string
}

func WithCache(c cache.Cache) StorageSetting {
	return func(o *StorageOption) {
		o.Cache = c
	}
}

func WithNamespace(ns string) StorageSetting {
	return func(o *StorageOption) {
		o.Namespace = ns
	}
}

// tokenStorage is the implementation of TokenStorage
type tokenStorage struct {
	*StorageOption
	Namespace string
}

// NewTokenStorage creates a new TokenStorage with a Cache and optional StoreOptions
func NewTokenStorage(ss ...StorageSetting) TokenStorage {
	opt := settings.Apply(&StorageOption{
		Namespace: TokenNamespace,
	}, ss)

	if opt.Cache == nil {
		opt.Cache = memory.NewCache(memory.Config{
			CleanupInterval:   5 * time.Minute,
			DefaultExpiration: 24 * time.Hour,
		})
	}

	s := &tokenStorage{
		StorageOption: opt,
	}

	return s
}

func (s *tokenStorage) tokenKey(key string) string {
	return s.Namespace + ":" + key
}

// Set stores the token with a specific expiration time
func (s *tokenStorage) Set(ctx context.Context, tokenStr string, expiration time.Duration) error {
	return s.Cache.Set(ctx, s.tokenKey(tokenStr), "", expiration)
}

// Delete deletes the token from the tokenStorage
func (s *tokenStorage) Delete(ctx context.Context, tokenStr string) error {
	return s.Cache.Delete(ctx, s.tokenKey(tokenStr))
}

// Validate checks if the token exists in the tokenStorage
func (s *tokenStorage) Validate(ctx context.Context, tokenStr string) error {
	return s.Cache.Exists(ctx, s.tokenKey(tokenStr))
}

// Close closes the tokenStorage
func (s *tokenStorage) Close(ctx context.Context) error {
	return nil
	//return s.Cache.Close(ctx)
}
