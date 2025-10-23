/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package identifier

import (
	"fmt"
	"sync"
)

// --- Registry Logic ---

const (
	// defaultString defines the name of the default string-based provider.
	defaultString = "uuid"
	// defaultNumber defines the name of the default number-based provider.
	defaultNumber = "snowflake"
)

type registry struct {
	sync.RWMutex
	providers         map[string]Provider
	defaultStringName string
	defaultNumberName string
}

// globalRegistry is pre-populated with built-in, dependency-free providers.
// External packages can override these defaults by re-registering the same name.
var globalRegistry = &registry{
	providers: map[string]Provider{
		defaultString: builtinString, // Pre-register the built-in string provider
		defaultNumber: builtinNumber, // Pre-register the built-in number provider
	},
	defaultStringName: defaultString, // Set the initial default
	defaultNumberName: defaultNumber, // Set the initial default
}

// Register registers a provider, overriding any existing provider with the same name.
func Register(p Provider) {
	globalRegistry.Lock()
	defer globalRegistry.Unlock()
	if p == nil {
		panic("identifier: cannot register a nil provider")
	}
	globalRegistry.providers[p.Name()] = p
}

// SetDefaultString sets the global default for string-based identifiers.
func SetDefaultString(name string) {
	globalRegistry.Lock()
	defer globalRegistry.Unlock()
	if _, ok := globalRegistry.providers[name]; !ok {
		panic(fmt.Sprintf("identifier: SetDefaultString called with unregistered provider \"%s\"", name))
	}
	globalRegistry.defaultStringName = name
}

// SetDefaultNumber sets the global default for number-based identifiers.
func SetDefaultNumber(name string) {
	globalRegistry.Lock()
	defer globalRegistry.Unlock()
	if _, ok := globalRegistry.providers[name]; !ok {
		panic(fmt.Sprintf("identifier: SetDefaultNumber called with unregistered provider \"%s\"", name))
	}
	globalRegistry.defaultNumberName = name
}

// Get retrieves a typed generator by name.
// It returns nil if no provider with the given name is registered.
func Get[T ~string | ~int64](name string) Generator[T] {
	globalRegistry.RLock()
	provider, ok := globalRegistry.providers[name]
	globalRegistry.RUnlock()

	if !ok {
		return nil
	}

	var t T
	switch any(t).(type) {
	case string:
		if gen := provider.AsString(); gen != nil {
			var asAny any = gen
			return asAny.(Generator[T])
		}
	case int64:
		if gen := provider.AsNumber(); gen != nil {
			var asAny any = gen
			return asAny.(Generator[T])
		}
	}
	return nil
}

// GenerateString generates a string ID using the configured default provider.
func GenerateString() string {
	globalRegistry.RLock()
	name := globalRegistry.defaultStringName
	globalRegistry.RUnlock()

	gen := Get[string](name)
	if gen == nil {
		// This should theoretically not happen if configured correctly,
		// as the default is pre-registered.
		panic("identifier: no default string provider available")
	}
	return gen.Generate()
}

// GenerateNumber generates a number ID using the configured default provider.
func GenerateNumber() int64 {
	globalRegistry.RLock()
	name := globalRegistry.defaultNumberName
	globalRegistry.RUnlock()

	gen := Get[int64](name)
	if gen == nil {
		// This should theoretically not happen if configured correctly,
		// as the default is pre-registered.
		panic("identifier: no default number provider available")
	}
	return gen.Generate()
}
