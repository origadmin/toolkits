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
	create(algName string) (interfaces.Cryptographic, error)
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

func (f *algorithmFactory) create(algName string) (interfaces.Cryptographic, error) {
	f.mux.RLock()
	if cachedAlg, exists := f.cryptos[algName]; exists {
		f.mux.RUnlock()
		return cachedAlg, nil
	}
	f.mux.RUnlock()

	f.mux.Lock()
	defer f.mux.Unlock()

	// Double-check after acquiring lock
	if cachedAlg, exists := f.cryptos[algName]; exists {
		return cachedAlg, nil
	}

	algType, err := types.ParseType(algName)
	if err != nil {
		return nil, err
	}
	algEntry, exists := algorithmMap[algType.Name]
	if !exists {
		return nil, fmt.Errorf("unsupported algorithm: %s", algType)
	}

	// Always create with default config for verification
	defaultConfig := algEntry.defaultConfig()
	newAlg, err := algEntry.creator(algType, defaultConfig)
	if err != nil {
		return nil, err
	}

	f.cryptos[algName] = newAlg
	return newAlg, nil
}

func (f *algorithmFactory) createConfig(algEntry algorithm, opts ...types.Option) *types.Config {
	cfg := &types.Config{}
	if algEntry.defaultConfig != nil {
		cfg = algEntry.defaultConfig()
	}
	return settings.Apply(cfg, opts)
}
