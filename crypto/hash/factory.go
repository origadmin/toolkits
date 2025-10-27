package hash

import (
	"fmt"
	"sync"

	"github.com/goexts/generic/configure"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// --- Factory Implementation ---

// Factory holds the registration of all scheme factories.
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

// NewCrypto creates a new cryptographic instance configured with a specific algorithm.
func (f *Factory) NewCrypto(algName string, opts ...Option) (Crypto, error) {
	spec, err := types.Parse(algName)
	if err != nil {
		return nil, fmt.Errorf("hash: failed to parse algorithm name: %w", err)
	}

	f.mu.RLock()
	schemeFactory, exists := f.factories[spec.Name]
	f.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("hash: unsupported algorithm: %s", spec.Name)
	}

	// **CORRECTED**: Get algorithm-specific default config from the factory first.
	defaultCfg := schemeFactory.Config()
	// Apply user-provided options on top of the default config.
	cfg := configure.Apply(defaultCfg, opts)

	// Create the scheme instance using the factory.
	algImpl, err := schemeFactory.Create(spec, cfg)
	if err != nil {
		return nil, fmt.Errorf("hash: failed to create scheme for %s: %w", spec.Name, err)
	}

	return &crypto{
		algImpl: algImpl,
	}, nil
}

// Verify checks if the given password matches the hashed value.
// It automatically detects the algorithm and parameters from the hashed string.
func (f *Factory) Verify(hashed, password string) error {
	if hashed == "" {
		return errors.ErrInvalidHash
	}

	parts, err := globalCodec.Decode(hashed)
	if err != nil {
		return err
	}

	if parts == nil || parts.Hash == nil || parts.Salt == nil {
		return errors.ErrInvalidHashParts
	}

	f.mu.RLock()
	schemeFactory, exists := f.factories[parts.Spec.Name]
	f.mu.RUnlock()

	if !exists {
		return fmt.Errorf("hash: unsupported algorithm: %s", parts.Spec.Name)
	}

	// **CORRECTED**: Create a config from the factory's default.
	cfg := schemeFactory.Config()
	// Populate it directly from the hash parts. This overrides any defaults.
	cfg.Params = parts.Params

	// Create a temporary scheme instance for verification.
	tempScheme, err := schemeFactory.Create(parts.Spec, cfg)
	if err != nil {
		return fmt.Errorf("hash: failed to create verification scheme for %s: %w", parts.Spec.Name, err)
	}

	// Delegate the final validation to the scheme's Verify method.
	return tempScheme.Verify(parts, password)
}

// --- Global Default Factory ---

// defaultFactory is the global, default instance of the factory.
var defaultFactory = NewFactory()

// Register registers a scheme factory to the default global factory.
// This is intended to be called from init() functions of algorithm packages.
func Register(name string, factory scheme.Factory) {
	defaultFactory.Register(name, factory)
}

// NewCrypto creates a new crypto instance using the default global factory.
func NewCrypto(algName string, opts ...Option) (Crypto, error) {
	return defaultFactory.NewCrypto(algName, opts...)
}
