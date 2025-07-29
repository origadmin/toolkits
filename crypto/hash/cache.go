/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"crypto/subtle"
	"fmt"
	"sync"
	"time"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Memory cache optimization
type cachedCrypto struct {
	crypto Crypto
	cache  sync.Map
}

func (c *cachedCrypto) Type() types.Type {
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
func (c *cachedCrypto) HashWithSalt(password string, salt []byte) (string, error) {
	return c.crypto.HashWithSalt(password, salt)
}

func (c *cachedCrypto) Verify(hashed string, password string) error {
	// Retrieve from cache
	if item, ok := c.cache.Load(hashed); ok {
		cached := item.(cacheItem)
		if time.Now().Before(cached.expiresAt) {
			if subtle.ConstantTimeCompare([]byte(cached.hash), []byte(hashed)) != 1 {
				fmt.Printf("compare: %s | %s\n", cached.hash, hashed)
				return errors.ErrPasswordNotMatch
			}
			return nil
		}
	}

	// Verify password
	err := c.crypto.Verify(hashed, password)
	if err != nil {
		return err
	}

	// Cache the result
	c.cache.Store(hashed, cacheItem{
		hash:      hashed,
		expiresAt: time.Now().Add(5 * time.Minute),
	})

	return nil
}

func CachedCrypto(crypto Crypto) Crypto {
	return &cachedCrypto{
		crypto: crypto,
	}
}
