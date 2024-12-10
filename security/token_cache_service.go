/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security is a toolkit for security check and authorization
package security

import (
	"context"
	"errors"
	"time"

	"github.com/goexts/generic/settings"

	"github.com/origadmin/toolkits/storage/cache"
)

const (
	TokenCacheNS = "security:token"
)

// TokenCacheService is the interface that TokenCacheService the token.
type TokenCacheService interface {
	// Store stores the token with a specific expiration time to TokenCacheService
	Store(context.Context, string, time.Duration) error
	// Validate checks if the token exists in the TokenCacheService
	Validate(context.Context, string) (bool, error)
	// Remove deletes the token from the TokenCacheService
	Remove(context.Context, string) error
	// Close closes the TokenCacheService
	Close(context.Context) error
}

type StorageSetting = func(*tokenCacheService)

func WithCache(c cache.Cache) StorageSetting {
	return func(o *tokenCacheService) {
		o.c = c
	}
}

func WithNamespace(ns string) StorageSetting {
	return func(o *tokenCacheService) {
		o.ns = ns
	}
}

// tokenCacheService is the implementation of tokenCacheService
type tokenCacheService struct {
	c  cache.Cache
	ns string
}

func (obj *tokenCacheService) Store(ctx context.Context, tokenStr string, duration time.Duration) error {
	return obj.c.Set(ctx, obj.tokenKey(tokenStr), tokenStr, duration)
}

func (obj *tokenCacheService) Validate(ctx context.Context, tokenStr string) (bool, error) {
	err := obj.c.Exists(ctx, obj.tokenKey(tokenStr))
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, cache.ErrNotFound):
		return false, nil
	default:
		return false, err
	}
	return false, nil
}

func (obj *tokenCacheService) Remove(ctx context.Context, tokenStr string) error {
	return obj.c.Delete(ctx, obj.tokenKey(tokenStr))
}

func (obj *tokenCacheService) Close(ctx context.Context) error {
	return nil
}

func (obj *tokenCacheService) tokenKey(key string) string {
	return obj.ns + ":" + key
}

// DefaultTokenCacheService creates a new tokenCacheService with a c and optional StoreOptions
func DefaultTokenCacheService(ss ...StorageSetting) TokenCacheService {
	service := settings.Apply(&tokenCacheService{
		ns: TokenCacheNS,
	}, ss)

	if service.c == nil {
		service.c = NewMapCache()
	}
	return service
}
