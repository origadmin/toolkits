/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"fmt"
	"sync"

	"github.com/goexts/generic/configure"

	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type internalFactory interface {
	// create now accepts types.Type directly
	create(algType types.Type) (interfaces.Cryptographic, error)
}

type algorithmFactory struct {
	cryptos map[string]interfaces.Cryptographic
	mux     sync.RWMutex
}

var (
	defaultFactory internalFactory
	once           sync.Once
)

func getFactory() internalFactory {
	once.Do(func() {
		defaultFactory = &algorithmFactory{
			cryptos: make(map[string]interfaces.Cryptographic),
		}
	})
	return defaultFactory
}

func (f *algorithmFactory) create(algType types.Type) (interfaces.Cryptographic, error) {
	// First, find the algorithm entry based on the initial algType.Name
	// This is needed to get the specific resolver for this algorithm.
	algEntry, exists := algorithmMap[algType.Name]
	if !exists {
		return nil, fmt.Errorf("unsupported algorithm: %s", algType.String())
	}

	// 1. Resolve the algorithm Type to its canonical form using the algorithm's specific resolver
	resolvedAlgType, err := algEntry.resolver.ResolveType(algType) // Use algEntry.resolver
	if err != nil {
		return nil, fmt.Errorf("failed to resolve algorithm type %s: %w", algType.String(), err)
	}

	// Use the resolved algorithm's string representation for caching
	algNameKey := resolvedAlgType.String()

	f.mux.RLock()
	if cachedAlg, exists := f.cryptos[algNameKey]; exists {
		f.mux.RUnlock()
		return cachedAlg, nil
	}
	f.mux.RUnlock()

	f.mux.Lock()
	defer f.mux.Unlock()

	// Double-check after acquiring lock
	if cachedAlg, exists := f.cryptos[algNameKey]; exists {
		return cachedAlg, nil
	}

	// Re-check algEntry after resolution, in case the resolved type's Name is different
	// and points to a different algorithm entry. This is important if resolvers can change the main Name.
	// For now, we assume algEntry is based on the initial algType.Name
	// and the resolver just canonicalizes the type itself.
	// If resolvedAlgType.Name could be different, we'd need to re-lookup here.
	// For simplicity and current design, we'll stick with the initial algEntry.

	// Always create with default config for verification
	defaultConfig := algEntry.defaultConfig()
	newAlg, err := algEntry.creator(resolvedAlgType, defaultConfig) // Pass resolvedAlgType
	if err != nil {
		return nil, fmt.Errorf("failed to create algorithm %s: %w", resolvedAlgType.String(), err)
	}

	f.cryptos[algNameKey] = newAlg
	return newAlg, nil
}

func (f *algorithmFactory) createConfig(algEntry algorithm, opts ...types.Option) *types.Config {
	cfg := &types.Config{}
	if algEntry.defaultConfig != nil {
		cfg = algEntry.defaultConfig()
	}
	return configure.Apply(cfg, opts)
}
