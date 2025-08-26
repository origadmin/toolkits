/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package identifier

import (
	"sync"
)

// registry manages the registration and retrieval of identifier providers.
// It is kept internal to the package.
type registry struct {
	sync.RWMutex
	providers map[string]Provider
}

// globalRegistry is the singleton instance of the registry.
var globalRegistry = &registry{
	providers: make(map[string]Provider),
}

// Register registers a provider, making it available via Get().
// This function should be called from the init() function of each algorithm's package.
func Register(p Provider) {
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

// Get retrieves a typed generator by name from the registry.
// This is the primary, recommended entry point for getting a shared generator instance.
// It returns a ready-to-use, typed generator, or nil if the named
// provider doesn't exist or doesn't support the requested type (string or int64).
func Get[T ~string | ~int64](name string) Generator[T] {
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
		// The result of AsString() is Generator[string].
		// We cast it to 'any' and then to the generic return type Generator[T],
		// which is valid because in this case, T is string.
		if gen := provider.AsString(); gen != nil {
			var asAny any = gen
			return asAny.(Generator[T])
		}
	case int64:
		// Same logic for the number type.
		if gen := provider.AsNumber(); gen != nil {
			var asAny any = gen
			return asAny.(Generator[T])
		}
	}
	return nil // Return nil if the type is not supported by the provider
}
