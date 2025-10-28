/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

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

// Crypto defines the primary interface for cryptographic hashing operations.
// An instance of Crypto is configured with a default algorithm for creating new hashes,
// but it can verify hashes from any algorithm registered in its factory.
type Crypto interface {
	// Spec returns the configured default algorithm specification for this crypto instance.
	Spec() types.Spec

	// Hash creates a new hash for the given password using the default algorithm.
	// The returned string is a fully encoded hash that includes all necessary metadata
	// for verification, such as the algorithm, parameters, and salt.
	Hash(password string) (string, error)

	// HashWithSalt creates a new hash with a user-provided salt.
	// This is useful for testing or specific use cases where salt generation is handled externally.
	HashWithSalt(password string, salt []byte) (string, error)

	// Verify checks if a password matches an encoded hash string.
	// It automatically detects the algorithm from the hash string, creates the appropriate
	// verification scheme, and performs the comparison. Results are cached for performance.
	Verify(hashed, password string) error
}

// crypto is the internal, concrete implementation of the Crypto interface.
// It manages a factory for creating algorithm schemes, a default scheme for new hashes,
// and caches for both scheme instances and verification results to optimize performance.
type crypto struct {
	factory           *Factory                 // The factory used to create new algorithm schemes.
	defaultAlg        scheme.Scheme            // The default scheme for hashing new passwords.
	schemeCache       map[string]scheme.Scheme // Caches scheme instances to avoid repeated creation.
	verificationCache *cache.Cache             // Caches verification results to speed up repeated checks.
	mu                sync.RWMutex             // Protects the schemeCache.
}

// globalCodec is the single, shared codec instance for encoding and decoding hash strings.
var globalCodec = codec.NewCodec()

// Spec returns the configured default algorithm specification for this crypto instance.
func (c *crypto) Spec() types.Spec {
	return c.defaultAlg.Spec()
}

// Hash creates a new hash for the given password using the default algorithm.
func (c *crypto) Hash(password string) (string, error) {
	hashParts, err := c.defaultAlg.Hash(password)
	if err != nil {
		return "", err
	}
	return globalCodec.Encode(hashParts)
}

// HashWithSalt creates a new hash with a user-provided salt.
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

	// Step 1: Check the verification cache first for a quick result.
	cacheKey := hashed + password
	if err, found := c.verificationCache.Get(cacheKey); found {
		if err == nil {
			return nil
		}
		return err.(error)
	}

	// Step 2: If not in cache, decode the full hash string into its constituent parts.
	parts, err := globalCodec.Decode(hashed)
	if err != nil {
		return err
	}
	if parts == nil || parts.Hash == nil || parts.Salt == nil {
		return errors.ErrInvalidHashParts
	}

	// Step 3: Get the appropriate algorithm scheme, either from cache or by creating a new one.
	schemeInstance, err := c.getScheme(parts)
	if err != nil {
		return err
	}

	// Step 4: Perform the actual, potentially expensive, verification.
	verificationErr := schemeInstance.Verify(parts, password)

	// Step 5: Cache the final result to speed up subsequent identical requests.
	c.verificationCache.Set(cacheKey, verificationErr, cache.DefaultExpiration)

	return verificationErr
}

// getScheme retrieves an algorithm scheme instance based on the provided HashParts.
// It uses a cache to avoid recreating scheme instances for the same algorithm and parameters.
func (c *crypto) getScheme(parts *types.HashParts) (scheme.Scheme, error) {
	// Resolve the spec from the parts to get the canonical algorithm name and handle aliases.
	schemeFactory, exists := c.factory.GetFactory(parts.Spec.Name)
	if !exists {
		return nil, fmt.Errorf("hash: factory for algorithm '%s' not found", parts.Spec.Name)
	}

	resolvedSpec, err := schemeFactory.ResolveSpec(parts.Spec)
	if err != nil {
		return nil, fmt.Errorf("hash: failed to resolve spec for algorithm '%s': %w", parts.Spec.Name, err)
	}

	// Use the resolved, canonical spec string as the cache key.
	specString := resolvedSpec.String()

	// Check the scheme cache first.
	c.mu.RLock()
	cachedScheme, exists := c.schemeCache[specString]
	c.mu.RUnlock()
	if exists {
		parts.Spec = resolvedSpec // Ensure the parts have the resolved spec for consistency.
		return cachedScheme, nil
	}

	// If not in cache, create a new scheme instance.
	cfg := ConfigFromHashParts(parts)
	newScheme, err := schemeFactory.Create(resolvedSpec, cfg)
	if err != nil {
		return nil, fmt.Errorf("hash: failed to create verification scheme for %s: %w", resolvedSpec.String(), err)
	}

	// Add the new scheme to the cache.
	c.mu.Lock()
	c.schemeCache[specString] = newScheme
	c.mu.Unlock()

	parts.Spec = resolvedSpec // Ensure the parts have the resolved spec.

	return newScheme, nil
}
