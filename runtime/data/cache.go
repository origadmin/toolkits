/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package data implements the functions, types, and interfaces for the module.
package data

import (
	"github.com/origadmin/toolkits/errors"
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/storage/cache"
	"github.com/origadmin/toolkits/storage/cache/memory"
)

const (
	ErrCacheConfigNil = errors.String("cache: config is nil")
)

type (
	Cache = cache.Cache
)

func NewCache(cfg *configv1.Data_Cache) (Cache, error) {
	if cfg == nil {
		return nil, ErrCacheConfigNil
	}

	if c := cfg.GetMemory(); c != nil {
		memcache := memory.NewCache()
		//if c.Capacity > 0 {
		//	cache.Capacity = c.Capacity
		//}
		if exp := c.Expiration; exp != nil && exp.AsDuration() > 0 {
			memcache.DefaultExpiration = exp.AsDuration()
		}
		if interval := c.CleanupInterval; interval != nil && interval.AsDuration() > 0 {
			memcache.CleanupInterval = interval.AsDuration()
		}
		return memcache, nil

	}
	return nil, errors.String("cache: unknown cache type")
}
