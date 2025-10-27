/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"fmt"
	"sync"
	"time"

	"github.com/goexts/generic/configure"
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

// newCrypto is the internal core constructor for creating a Crypto instance.
// It requires a factory to create schemes.
func newCrypto(factory *Factory, defaultAlgName string, opts []Option) (Crypto, error) {
	spec, exists := factory.GetSpec(defaultAlgName)
	if !exists {
		return nil, fmt.Errorf("hash: spec for algorithm '%s' not found or not registered", defaultAlgName)
	}

	// Get the factory for the algorithm once
	schemeFactory, exists := factory.GetFactory(spec.Name)
	if !exists {
		return nil, fmt.Errorf("hash: factory for algorithm '%s' not found", spec.Name)
	}

	// Use the same factory instance to get config and create the scheme
	defaultCfg := schemeFactory.Config()
	cfg := configure.Apply(defaultCfg, opts)
	var err error
	nspec, err := schemeFactory.ResolveSpec(spec)
	if err != nil {
		return nil, fmt.Errorf("hash: failed to resolve spec for algorithm '%s': %w", spec.Name, err)
	}

	defaultAlg, err := schemeFactory.Create(nspec, cfg)
	if err != nil {
		return nil, fmt.Errorf("hash: failed to create default scheme: %w", err)
	}

	return &crypto{
		factory:           factory,
		defaultAlg:        defaultAlg,
		schemeCache:       make(map[string]scheme.Scheme),
		verificationCache: cache.New(5*time.Minute, 10*time.Minute),
	}, nil
}

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
func (c *crypto) getScheme(parts *types.HashParts) (scheme.Scheme, error) {
	specString := parts.Spec.String()

	c.mu.RLock()
	cachedScheme, exists := c.schemeCache[specString]
	c.mu.RUnlock()

	if exists {
		return cachedScheme, nil
	}

	// Not in cache, create it.
	// The Config for verification MUST be built from the parts.
	cfg := ConfigFromHashParts(parts)

	schemeFactory, exists := c.factory.GetFactory(parts.Spec.Name)
	if !exists {
		return nil, fmt.Errorf("hash: factory for algorithm '%s' not found", parts.Spec.Name)
	}
	nspec, err := schemeFactory.ResolveSpec(parts.Spec)
	if err != nil {
		return nil, fmt.Errorf("hash: failed to resolve spec for algorithm '%s': %w", parts.Spec.Name, err)
	}
	newScheme, err := schemeFactory.Create(nspec, cfg)
	if err != nil {
		return nil, fmt.Errorf("hash: failed to create verification scheme for %s: %w", parts.Spec.String(), err)
	}

	// Cache the newly created scheme.
	c.mu.Lock()
	c.schemeCache[specString] = newScheme
	c.mu.Unlock()

	return newScheme, nil
}
