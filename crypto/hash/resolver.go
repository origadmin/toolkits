/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"sync"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

var (
	// algorithmResolvers stores registered resolvers for different algorithm names.
	algorithmResolvers = make(map[string]interfaces.AlgorithmResolver)
	resolversMu        sync.RWMutex
)

// RegisterAlgorithmResolver registers a resolver for a specific algorithm name.
func RegisterAlgorithmResolver(name string, resolver interfaces.AlgorithmResolver) {
	resolversMu.Lock()
	defer resolversMu.Unlock()
	algorithmResolvers[name] = resolver
}

func ResolveType(t types.Type) (types.Type, error) {
	resolversMu.RLock()
	defer resolversMu.RUnlock()
	resolver, ok := algorithmResolvers[t.Name]
	if !ok {
		return t, errors.ErrResolverNotRegistered
	}
	return resolver.ResolveType(t)
}
