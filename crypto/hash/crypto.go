/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"fmt"
	"sync"

	"github.com/patrickmn/go-cache"

	"github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Crypto defines the interface for a cryptographic instance.
// It can be used to hash new passwords with its configured scheme and verify any hash.
type Crypto interface {
	Spec() types.Spec
	Hash(password string) (string, error)
	HashWithSalt(password string, salt []byte) (string, error)
	Verify(hashed, password string) error
}

// crypto is the internal implementation of the Crypto interface.
// It holds the default scheme for hashing and caches for verification.
type crypto struct {
	factory           *Factory
	defaultAlg        scheme.Scheme
	schemeCache       map[string]scheme.Scheme
	verificationCache *cache.Cache
	mu                sync.RWMutex
}

var globalCodec = codec.NewCodec()

func (c *crypto) Spec() types.Spec {
	return c.defaultAlg.Spec()
}

func (c *crypto) Hash(password string) (string, error) {
	hashParts, err := c.defaultAlg.Hash(password)
	if err != nil {
		return "", err
	}
	return globalCodec.Encode(hashParts)
}

func (c *crypto) HashWithSalt(password string, salt []byte) (string, error) {
	hashParts, err := c.defaultAlg.HashWithSalt(password, salt)
	if err != nil {
		return "", err
	}
	return globalCodec.Encode(hashParts)
}

// Verify checks if the given password matches the hashed value, with caching.
func (c *crypto) Verify(hashed, password string) error {
	if hashed == "" {
		return errors.ErrInvalidHash
	}

	// 1. Check verification cache first
	cacheKey := hashed + password
	if err, found := c.verificationCache.Get(cacheKey); found {
		if err == nil {
			return nil
		}
		return err.(error)
	}

	// 2. Decode the hash
	parts, err := globalCodec.Decode(hashed)
	if err != nil {
		return err
	}
	if parts == nil || parts.Hash == nil || parts.Salt == nil {
		return errors.ErrInvalidHashParts
	}

	// 3. Get the scheme (from cache or create new)
	schemeInstance, err := c.getScheme(parts)
	if err != nil {
		return err
	}

	// 4. Perform the actual verification
	verificationErr := schemeInstance.Verify(parts, password)

	// 5. Cache the result
	c.verificationCache.Set(cacheKey, verificationErr, cache.DefaultExpiration)

	return verificationErr
}

// getScheme retrieves a scheme from the cache or creates it if not present.
// It returns the scheme and the resolved spec.
func (c *crypto) getScheme(parts *types.HashParts) (scheme.Scheme, error) {
	// First, resolve the spec to get the correct algorithm name
	schemeFactory, exists := c.factory.GetFactory(parts.Spec.Name)
	if !exists {
		return nil, fmt.Errorf("hash: factory for algorithm '%s' not found", parts.Spec.Name)
	}

	// Resolve the spec to handle any aliases or default values
	resolvedSpec, err := schemeFactory.ResolveSpec(parts.Spec)
	if err != nil {
		return nil, fmt.Errorf("hash: failed to resolve spec for algorithm '%s': %w", parts.Spec.Name, err)
	}

	// Use the resolved spec for cache key
	specString := resolvedSpec.String()

	// Check cache with resolved spec
	c.mu.RLock()
	cachedScheme, exists := c.schemeCache[specString]
	c.mu.RUnlock()

	if exists {
		// Update the original parts with the resolved spec
		parts.Spec = resolvedSpec
		return cachedScheme, nil
	}

	// Not in cache, create new scheme
	cfg := ConfigFromHashParts(parts)

	// Create the scheme with the resolved spec
	newScheme, err := schemeFactory.Create(resolvedSpec, cfg)
	if err != nil {
		return nil, fmt.Errorf("hash: failed to create verification scheme for %s: %w",
			resolvedSpec.String(), err)
	}

	// Cache the newly created scheme
	c.mu.Lock()
	c.schemeCache[specString] = newScheme
	c.mu.Unlock()

	// Update the original parts with the resolved spec
	parts.Spec = resolvedSpec

	return newScheme, nil
}
