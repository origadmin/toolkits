package hash

import (
	"fmt"
	"time"

	"github.com/goexts/generic/configure"
	"github.com/patrickmn/go-cache"

	"github.com/origadmin/toolkits/crypto/hash/scheme"
)

// newCrypto is the internal core constructor for creating a Crypto instance.
// It requires a factory to create schemes.
func newCrypto(factory *Factory, defaultAlgName string, opts []Option) (Crypto, error) {
	spec, exists := factory.GetSpec(defaultAlgName)
	if !exists {
		return nil, fmt.Errorf("hash: spec for algorithm '%s' not found or not registered", defaultAlgName)
	}

	// Get the factory for the algorithm once
	schemeFactory, exists := factory.GetFactory(spec.Name)
	if !exists {
		return nil, fmt.Errorf("hash: factory for algorithm '%s' not found", spec.Name)
	}

	// Use the same factory instance to get config and create the scheme
	defaultCfg := schemeFactory.Config()
	cfg := configure.Apply(defaultCfg, opts)
	var err error
	nspec, err := schemeFactory.ResolveSpec(spec)
	if err != nil {
		return nil, fmt.Errorf("hash: failed to resolve spec for algorithm '%s': %w", spec.Name, err)
	}

	defaultAlg, err := schemeFactory.Create(nspec, cfg)
	if err != nil {
		return nil, fmt.Errorf("hash: failed to create default scheme: %w", err)
	}

	return &crypto{
		factory:           factory,
		defaultAlg:        defaultAlg,
		schemeCache:       make(map[string]scheme.Scheme),
		verificationCache: cache.New(5*time.Minute, 10*time.Minute),
	}, nil
}
