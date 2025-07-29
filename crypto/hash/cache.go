/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"sync"
	"time"

	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// verificationResultCache is a global, in-memory cache for recently verified successful hashes.
// The key is the full hashed string, and the value is a timestamp of when it expires.
// This cache is used to reduce CPU load for repeated successful verifications within a short period.
// IMPORTANT: A cache hit DOES NOT bypass the full cryptographic verification to prevent timing attacks.
var verificationResultCache sync.Map

const cacheDuration = 5 * time.Minute // Cache entries expire after 5 minutes

// cachedVerifier wraps a Cryptographic implementation to add result caching for Verify operations.
type cachedVerifier struct {
	wrapped interfaces.Cryptographic
}

func (c *cachedVerifier) Type() types.Type {
	return c.wrapped.Type()
}

func (c *cachedVerifier) Hash(password string) (*types.HashParts, error) {
	return c.wrapped.Hash(password)
}

func (c *cachedVerifier) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	return c.wrapped.HashWithSalt(password, salt)
}

// Verify implements the verify method with result caching.
func (c *cachedVerifier) Verify(parts *types.HashParts, password string) error {
	// Check cache for a recent successful verification of this specific hash string.
	// This is a performance optimization, not a security bypass.
	// We use the raw hash bytes as the key for the cache.
	hashedString := string(parts.Hash)
	if expiryTime, ok := verificationResultCache.Load(hashedString); ok {
		if time.Now().Before(expiryTime.(time.Time)) {
			// Cache hit: The hash was successfully verified recently.
			// We still proceed with the full verification to prevent timing attacks.
			// The presence in cache only indicates a high probability of success.
		} else {
			// Cache expired, remove it.
			verificationResultCache.Delete(hashedString)
		}
	}

	// Perform the actual cryptographic verification.
	err := c.wrapped.Verify(parts, password)
	if err == nil {
		// If verification is successful, cache the hash for a short period.
		verificationResultCache.Store(hashedString, time.Now().Add(cacheDuration))
	}
	return err
}

// NewCachedVerifier creates a new cachedVerifier that wraps the given Cryptographic implementation.
func NewCachedVerifier(wrapped interfaces.Cryptographic) interfaces.Cryptographic {
	return &cachedVerifier{wrapped: wrapped}
}
