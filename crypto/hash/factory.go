package hash

import (
	"fmt"
	"sync"

	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// --- Global Default Factory ---

// defaultFactory is the global, default instance of the factory.
var defaultFactory = NewFactory()

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

// Register registers a scheme factory to the default global factory.
// This is intended to be called from init() functions of algorithm packages.
func Register(name string, factory scheme.Factory) {
	defaultFactory.Register(name, factory)
}
