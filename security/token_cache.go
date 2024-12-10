/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

import (
	"context"
	"sync"
	"time"

	"github.com/origadmin/toolkits/errors"

	"github.com/origadmin/toolkits/storage/cache"
)

type element struct {
	value    string
	expireAt time.Time
}

type mapCache struct {
	elems sync.Pool
	maps  sync.Map
}

func (s *mapCache) Get(ctx context.Context, key string) (string, error) {
	value, ok := s.maps.Load(key)
	if !ok {
		return "", cache.ErrNotFound
	}
	elem, ok := value.(*element)
	if !ok {
		return "", errors.New("invalid cache value")
	}
	if elem.expireAt.Before(time.Now()) {
		s.maps.Delete(key)
		s.elems.Put(elem)
		return "", cache.ErrNotFound
	}
	return elem.value, nil
}

func (s *mapCache) GetAndDelete(ctx context.Context, key string) (string, error) {
	value, ok := s.maps.LoadAndDelete(key)
	if !ok {
		return "", cache.ErrNotFound
	}
	elem, ok := value.(*element)
	if !ok {
		return "", errors.New("invalid cache value")
	}
	v := elem.value
	s.elems.Put(elem)
	return v, nil
}

func (s *mapCache) Exists(ctx context.Context, key string) error {
	if _, ok := s.maps.Load(key); ok {
		return nil
	}
	return cache.ErrNotFound
}

func (s *mapCache) Set(ctx context.Context, key string, value string, expiration ...time.Duration) error {
	var expireAt time.Time
	if len(expiration) > 0 {
		expireAt = time.Now().Add(expiration[0])
	} else {
		expireAt = time.Now().Add(time.Hour)
	}
	elem := s.getElement()
	elem.value = value
	elem.expireAt = expireAt
	s.maps.Store(key, elem)
	return nil
}

func (s *mapCache) Delete(ctx context.Context, key string) error {
	if v, ok := s.maps.LoadAndDelete(key); ok {
		s.elems.Put(v)
		return nil
	}
	return cache.ErrNotFound
}

func (s *mapCache) getElement() *element {
	return s.elems.Get().(*element)
}

func (s *mapCache) putElement(elem *element) {
	s.elems.Put(elem)
}

func NewMapCache() cache.Cache {
	return &mapCache{}
}

var _ cache.Cache = (*mapCache)(nil)
