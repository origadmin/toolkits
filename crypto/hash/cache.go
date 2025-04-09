/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"sync"
	"time"

	"github.com/origadmin/toolkits/crypto/hash/interfaces"
)

// Memory cache optimization
type cachedCrypto struct {
	crypto interfaces.Cryptographic
	cache  sync.Map
}

func (c *cachedCrypto) Type() string {
	return c.crypto.Type()
}

type cacheItem struct {
	hash      string
	expiresAt time.Time
}

// Hash implements Cryptographic.
func (c *cachedCrypto) Hash(password string) (string, error) {
	return c.crypto.Hash(password)
}

// HashWithSalt implements Cryptographic.
func (c *cachedCrypto) HashWithSalt(password, salt string) (string, error) {
	return c.crypto.HashWithSalt(password, salt)
}

func (c *cachedCrypto) Verify(hashed, password string) error {
	// Retrieve from cache
	if item, ok := c.cache.Load(password); ok {
		cached := item.(cacheItem)
		if time.Now().Before(cached.expiresAt) {
			if cached.hash == hashed {
				return nil
			}
			return ErrPasswordNotMatch
		}
	}

	// Verify password
	err := c.crypto.Verify(hashed, password)
	if err != nil {
		return err
	}

	// Cache the result
	c.cache.Store(password, cacheItem{
		hash:      hashed,
		expiresAt: time.Now().Add(5 * time.Minute),
	})

	return nil
}

func NewCachedCrypto(crypto interfaces.Cryptographic) interfaces.Cryptographic {
	return &cachedCrypto{
		crypto: crypto,
	}
}
