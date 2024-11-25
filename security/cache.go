// Package security implements the functions, types, and interfaces for the module.
package security

import (
	"sync"
	"time"

	"github.com/origadmin/toolkits/errors"

	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/storage/cache"
)

type element struct {
	value    string
	expireAt time.Time
}

type securityCache struct {
	maps sync.Map
}

func (s *securityCache) Get(ctx context.Context, key string) (string, error) {
	value, ok := s.maps.Load(key)
	if !ok {
		return "", cache.ErrNotFound
	}
	ele, ok := value.(*element)
	if !ok {
		return "", errors.New("invalid cache value")
	}
	if ele.expireAt.Before(time.Now()) {
		_ = s.Delete(ctx, key)
		return "", cache.ErrNotFound
	}
	return ele.value, nil
}

func (s *securityCache) GetAndDelete(ctx context.Context, key string) (string, error) {
	value, ok := s.maps.LoadAndDelete(key)
	if !ok {
		return "", cache.ErrNotFound
	}
	ele, ok := value.(*element)
	if !ok {
		return "", errors.New("invalid cache value")
	}
	return ele.value, nil
}

func (s *securityCache) Exists(ctx context.Context, key string) error {
	_, ok := s.maps.Load(key)
	if !ok {
		return cache.ErrNotFound
	}
	return nil
}

func (s *securityCache) Set(ctx context.Context, key string, value string, expiration ...time.Duration) error {
	var expireAt time.Time
	if len(expiration) > 0 {
		expireAt = time.Now().Add(expiration[0])
	} else {
		expireAt = time.Now().Add(time.Hour)
	}
	ele := &element{value: value, expireAt: expireAt}
	s.maps.Store(key, ele)
	return nil
}

func (s *securityCache) Delete(ctx context.Context, key string) error {
	s.maps.Delete(key)
	return nil
}

func NewSecurityCache() cache.Cache {
	return &securityCache{}
}

var _ cache.Cache = (*securityCache)(nil)
