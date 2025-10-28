package hash

import (
	"fmt"
	"sync"

	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Factory holds the registration of all scheme factories and spec mappings.
type Factory struct {
	mu        sync.RWMutex
	factories map[string]scheme.Factory // name -> factory
	specs     map[string]types.Spec     // canonical_string -> spec_object
	aliases   map[string]string         // alias -> canonical_string
}

// NewFactory creates a new, empty factory.
func NewFactory() *Factory {
	return &Factory{
		factories: make(map[string]scheme.Factory),
		specs:     make(map[string]types.Spec),
		aliases:   make(map[string]string),
	}
}

// Register registers a single algorithm, its factory, its canonical spec, and any string aliases.
func (f *Factory) Register(factory scheme.Factory, canonicalSpec types.Spec, aliases ...string) {
	f.mu.Lock()
	defer f.mu.Unlock()

	if factory == nil {
		panic("hash: Register factory is nil")
	}

	// The algorithm's base name (e.g., "sha", "blake2") is the key for the factory.
	baseName := canonicalSpec.Name
	f.factories[baseName] = factory // Always register/overwrite the factory for the base name.

	// Register the canonical spec
	canonicalString := canonicalSpec.String()
	f.specs[canonicalString] = canonicalSpec
	f.aliases[canonicalString] = canonicalString // The canonical name is an alias for itself
	// Register all other string aliases
	for _, alias := range aliases {
		f.aliases[alias] = canonicalString
	}
}

func (f *Factory) GetFactory(name string) (scheme.Factory, bool) {
	f.mu.RLock()
	defer f.mu.RUnlock()

	factory, exists := f.factories[name]
	return factory, exists
}

// GetSpec resolves a spec string to a types.Spec object using the internal alias map.
func (f *Factory) GetSpec(specStr string) (types.Spec, bool) {
	f.mu.RLock()
	defer f.mu.RUnlock()
	canonicalName, exists := f.aliases[specStr]
	if !exists {
		canonicalName = specStr
	}
	parsed, err := types.Parse(canonicalName)
	if err != nil {
		return types.Spec{}, false
	}
	return parsed, true

	//spec, exists := f.specs[canonicalName]
	//return spec, exists
}

// GetConfig returns the default configuration for a given algorithm module.
func (f *Factory) GetConfig(name string) *types.Config {
	f.mu.RLock()
	schemeFactory, exists := f.factories[name]
	f.mu.RUnlock()

	if !exists {
		return types.DefaultConfig()
	}
	return schemeFactory.Config()
}

// Create uses the registered providers to create a new Scheme instance.
func (f *Factory) Create(spec types.Spec, cfg *types.Config) (scheme.Scheme, error) {
	f.mu.RLock()
	schemeFactory, exists := f.factories[spec.Name]
	f.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("hash: unsupported algorithm module: %s", spec.Name)
	}

	return schemeFactory.Create(spec, cfg)
}

// AvailableAlgorithms returns a list of all registered hash algorithm aliases.
func (f *Factory) AvailableAlgorithms() []string {
	f.mu.RLock()
	defer f.mu.RUnlock()
	algorithms := make([]string, 0, len(f.aliases))
	for name := range f.aliases {
		algorithms = append(algorithms, name)
	}
	return algorithms
}
