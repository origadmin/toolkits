/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security is a toolkit for security check and authorization
package security

import (
	"context"
	"time"

	"github.com/goexts/generic/settings"

	"github.com/origadmin/toolkits/storage/cache"
)

const (
	TokenCacheNS = "security:token"
)

// TokenService is the interface that TokenService the token.
type TokenService interface {
	// Store stores the token with a specific expiration time to TokenService
	Store(context.Context, string, time.Duration) error
	// Validate checks if the token exists in the TokenService
	Validate(context.Context, string) (bool, error)
	// Remove deletes the token from the TokenService
	Remove(context.Context, string) error
	// Close closes the TokenService
	Close(context.Context) error
}

type StorageSetting = func(*tokenService)

func WithCache(c cache.Cache) StorageSetting {
	return func(o *tokenService) {
		o.c = c
	}
}

func WithNamespace(ns string) StorageSetting {
	return func(o *tokenService) {
		o.ns = ns
	}
}

// tokenService is the implementation of tokenService
type tokenService struct {
	c  cache.Cache
	ns string
}

func (obj *tokenService) Store(ctx context.Context, tokenStr string, duration time.Duration) error {
	return obj.c.Set(ctx, obj.tokenKey(tokenStr), tokenStr, duration)
}

func (obj *tokenService) Validate(ctx context.Context, tokenStr string) (bool, error) {
	ok, err := obj.c.Exists(ctx, obj.tokenKey(tokenStr))
	switch {
	case ok:
		return true, nil
	default:
		return false, err
	}
}

func (obj *tokenService) Remove(ctx context.Context, tokenStr string) error {
	return obj.c.Delete(ctx, obj.tokenKey(tokenStr))
}

func (obj *tokenService) Close(ctx context.Context) error {
	return obj.c.Close(ctx)
}

func (obj *tokenService) tokenKey(key string) string {
	return obj.ns + ":" + key
}

// DefaultTokenService creates a new tokenService with a c and optional StoreOptions
func DefaultTokenService(ss ...StorageSetting) TokenService {
	service := settings.Apply(&tokenService{
		ns: TokenCacheNS,
	}, ss)

	if service.c == nil {
		service.c = cache.NewMemoryCache()
	}
	return service
}
