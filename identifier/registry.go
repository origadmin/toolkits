/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package identifier

import (
	"sync"
)

// registry manages the registration and retrieval of identifier generator providers.
// It is kept internal to the package.
type registry struct {
	sync.RWMutex
	providers             map[string]GeneratorProvider
	defaultStringProvider GeneratorProvider
	defaultNumberProvider GeneratorProvider
}

// globalRegistry is the singleton instance of the registry.
var globalRegistry = &registry{
	providers: make(map[string]GeneratorProvider),
}

// Register registers a generator provider, making it available via New().
// This function should be called from the init() function of each algorithm's package.
func Register(p GeneratorProvider) {
	globalRegistry.Lock()
	defer globalRegistry.Unlock()
	if p == nil {
		panic("identifier: cannot register a nil provider")
	}
	name := p.Name()
	if _, dup := globalRegistry.providers[name]; dup {
		panic("identifier: Register called twice for provider " + name)
	}
	globalRegistry.providers[name] = p
}

// New retrieves a typed generator by name in a single step.
// This is the primary, recommended entry point for getting a generator.
// It returns a ready-to-use, typed generator, or nil if the named
// provider doesn't exist or doesn't support the requested type (string or int64).
func New[T ~string | ~int64](name string) TypedGenerator[T] {
	globalRegistry.RLock()
	provider := globalRegistry.providers[name]
	globalRegistry.RUnlock()

	if provider == nil {
		return nil
	}

	// Use 'any' to check the desired type T and call the correct provider method.
	var t T
	switch any(t).(type) {
	case string:
		// The result of AsString() is TypedGenerator[string].
		// We cast it to 'any' and then to the generic return type TypedGenerator[T],
		// which is valid because in this case, T is string.
		if gen := provider.AsString(); gen != nil {
			var asAny any = gen
			return asAny.(TypedGenerator[T])
		}
	case int64:
		// Same logic for the number type.
		if gen := provider.AsNumber(); gen != nil {
			var asAny any = gen
			return asAny.(TypedGenerator[T])
		}
	}
	return nil // Return nil if the type is not supported by the provider
}

// SetDefaultString sets the default provider for string identifiers.
// Panics if the provider does not support string generation.
func SetDefaultString(p GeneratorProvider) {
	if p == nil || p.AsString() == nil {
		panic("identifier: provider cannot be nil or does not support string generation")
	}
	globalRegistry.Lock()
	defer globalRegistry.Unlock()
	globalRegistry.defaultStringProvider = p
}

// SetDefaultNumber sets the default provider for number identifiers.
// Panics if the provider does not support number generation.
func SetDefaultNumber(p GeneratorProvider) {
	if p == nil || p.AsNumber() == nil {
		panic("identifier: provider cannot be nil or does not support number generation")
	}
	globalRegistry.Lock()
	defer globalRegistry.Unlock()
	globalRegistry.defaultNumberProvider = p
}

// DefaultString returns the default string identifier generator.
func DefaultString() TypedGenerator[string] {
	globalRegistry.RLock()
	defer globalRegistry.RUnlock()
	if globalRegistry.defaultStringProvider != nil {
		return globalRegistry.defaultStringProvider.AsString()
	}
	return nil
}

// DefaultNumber returns the default number identifier generator.
func DefaultNumber() TypedGenerator[int64] {
	globalRegistry.RLock()
	defer globalRegistry.RUnlock()
	if globalRegistry.defaultNumberProvider != nil {
		return globalRegistry.defaultNumberProvider.AsNumber()
	}
	return nil
}
