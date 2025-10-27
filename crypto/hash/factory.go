package hash

import (
	"fmt"
	"sync"

	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Factory holds the registration of all scheme factories. It is a stateless creator.
type Factory struct {
	mu        sync.RWMutex
	factories map[string]scheme.Factory
}

// NewFactory creates a new, empty factory.
func NewFactory() *Factory {
	return &Factory{
		factories: make(map[string]scheme.Factory),
	}
}

// Register registers a new scheme factory. If a factory with the same name already
// exists, it will be overwritten.
func (f *Factory) Register(name string, factory scheme.Factory) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if factory == nil {
		panic("hash: Register factory is nil")
	}
	f.factories[name] = factory
}

// GetConfig returns the default configuration for a given algorithm.
// If the specific algorithm's factory is not found, it gracefully falls back
// to the global default configuration from the types package.
func (f *Factory) GetConfig(name string) *types.Config {
	f.mu.RLock()
	schemeFactory, exists := f.factories[name]
	f.mu.RUnlock()

	if !exists {
		// Fallback to global default if specific factory is not registered.
		return types.DefaultConfig()
	}
	return schemeFactory.Config()
}

// Create uses the registered factories to create a new Scheme instance.
// This is the primary creation method of the factory.
func (f *Factory) Create(spec types.Spec, cfg *types.Config) (scheme.Scheme, error) {
	f.mu.RLock()
	schemeFactory, exists := f.factories[spec.Name]
	f.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("hash: unsupported algorithm: %s", spec.Name)
	}
	return schemeFactory.Create(spec, cfg)
}

// AvailableAlgorithms returns a list of all registered hash algorithms.
func (f *Factory) AvailableAlgorithms() []string {
	f.mu.RLock()
	defer f.mu.RUnlock()
	algorithms := make([]string, 0, len(f.factories))
	for name := range f.factories {
		algorithms = append(algorithms, name)
	}
	return algorithms
}
