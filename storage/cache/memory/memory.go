package memory

import (
	"context"
	"time"

	"github.com/coocood/freecache"

	"github.com/origadmin/toolkits/errors"
)

const (
	ErrNotFound = errors.String("not found")
)

type Cache struct {
	Delimiter         string
	Cache             *freecache.Cache
	DefaultExpiration time.Duration
	CleanupInterval   time.Duration
}

func (obj *Cache) Set(ctx context.Context, key, value string, expiration ...time.Duration) error {
	var exp time.Duration
	if len(expiration) > 0 {
		exp = expiration[0]
	}
	return obj.Cache.Set([]byte(key), []byte(value), int(exp))
}

func (obj *Cache) Get(ctx context.Context, key string) (string, error) {
	val, err := obj.Cache.Get([]byte(key))
	if err != nil {
		return "", ErrNotFound
	}
	return string(val), nil
}

func (obj *Cache) Exists(ctx context.Context, key string) error {
	_, err := obj.Cache.Get([]byte(key))
	if err != nil {
		return ErrNotFound
	}
	return nil
}

func (obj *Cache) Delete(ctx context.Context, key string) error {
	affected := obj.Cache.Del([]byte(key))
	if !affected {
		return ErrNotFound
	}
	return nil
}

func (obj *Cache) GetAndDelete(ctx context.Context, key string) (string, error) {
	value, err := obj.Get(ctx, key)
	if err != nil {
		return "", err
	}
	err = obj.Delete(ctx, key)
	if err != nil {
		return "", err
	}
	return value, nil
}

func (obj *Cache) Iterator(ctx context.Context, fn func(ctx context.Context, key, value string) bool) error {
	iter := obj.Cache.NewIterator()
	for entry := iter.Next(); entry != nil; entry = iter.Next() {
		if !fn(ctx, string(entry.Key), string(entry.Value)) {
			break
		}
	}
	return nil
}

func (obj *Cache) Close(_ context.Context) error {
	obj.Cache.Clear()
	return nil
}

func NewCache() *Cache {
	return &Cache{
		Delimiter:         ":",
		DefaultExpiration: time.Second * 60,
		CleanupInterval:   time.Second * 60,
		Cache:             newFreeCache(100 * 1024 * 1024),
	}
}

func newFreeCache(size int) *freecache.Cache {
	return freecache.NewCache(size)
}
