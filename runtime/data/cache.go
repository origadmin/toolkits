// Copyright (c) 2024 OrigAdmin. All rights reserved.

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

func NewCache(cfg *configv1.Data_Cache) (cache.Cache, error) {
	if cfg == nil {
		return nil, ErrCacheConfigNil
	}
	return memory.NewCache(cfg.Memory), nil
}
