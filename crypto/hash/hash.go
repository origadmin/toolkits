/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides the hash functions
package hash

import (
	"fmt"
	"sync"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// --- Global Default Instances ---

// defaultFactory is the global, default instance of the factory.
var defaultFactory = NewFactory()

var (
	// globalCrypto is the global, default instance of Crypto.
	globalCrypto Crypto
	// globalCryptoMutex protects globalCrypto from concurrent access
	globalCryptoMutex sync.RWMutex
)

// uninitializedCrypto is a no-op Crypto implementation used when the module fails to initialize.
type uninitializedCrypto struct{}

func (u *uninitializedCrypto) Spec() types.Spec {
	return types.Spec{Name: types.UNKNOWN}
}

func (u *uninitializedCrypto) Hash(password string) (string, error) {
	return "", errors.ErrHashModuleNotInitialized
}

func (u *uninitializedCrypto) HashWithSalt(password string, salt []byte) (string, error) {
	return "", errors.ErrHashModuleNotInitialized
}

func (u *uninitializedCrypto) Verify(hashed, password string) error {
	return errors.ErrHashModuleNotInitialized
}

// --- Package-Level Convenience Functions ---

// Register is a convenience function that registers a factory to the default global factory.
func Register(factory scheme.Factory, canonicalSpec types.Spec, aliases ...string) {
	defaultFactory.Register(factory, canonicalSpec, aliases...)
}

// NewCrypto is the primary convenience function for creating a Crypto instance.
// It uses the default global factory.
func NewCrypto(defaultAlgName string, opts ...Option) (Crypto, error) {
	// It calls the internal core constructor, injecting the defaultFactory.
	return newCrypto(defaultFactory, defaultAlgName, opts)
}

// NewCryptoWithFactory creates a Crypto instance using a specific, provided factory.
// This is useful for testing or creating isolated instances with different configurations.
func NewCryptoWithFactory(factory *Factory, defaultAlgName string, opts ...Option) (Crypto, error) {
	// It calls the internal core constructor with the provided factory.
	return newCrypto(factory, defaultAlgName, opts)
}

// UseCrypto updates the active cryptographic instance
func UseCrypto(algName string, opts ...Option) error {
	if algName == "" {
		return errors.ErrInvalidAlgorithm
	}
	algSpec, exists := defaultFactory.GetSpec(algName)
	if !exists {
		return fmt.Errorf("hash: spec for algorithm '%s' not found", algName)
	}

	globalCryptoMutex.RLock()
	currentCrypto := globalCrypto
	globalCryptoMutex.RUnlock()

	if currentCrypto != nil && currentCrypto.Spec().Is(algSpec) {
		return nil
	}

	newCrypto, err := NewCrypto(algName, opts...)
	if err != nil {
		return err
	}
	globalCryptoMutex.Lock()
	globalCrypto = newCrypto
	globalCryptoMutex.Unlock()
	return nil
}

// Verify is a convenience function that uses the active global crypto instance.
func Verify(hashed, password string) error {
	globalCryptoMutex.RLock()
	defer globalCryptoMutex.RUnlock()
	return globalCrypto.Verify(hashed, password)
}

// Generate is a convenience function that uses the active global crypto instance.
func Generate(password string) (string, error) {
	globalCryptoMutex.RLock()
	defer globalCryptoMutex.RUnlock()
	return globalCrypto.Hash(password)
}

// GenerateWithSalt is a convenience function that uses the active global crypto instance.
func GenerateWithSalt(password string, salt []byte) (string, error) {
	globalCryptoMutex.RLock()
	defer globalCryptoMutex.RUnlock()
	return globalCrypto.HashWithSalt(password, salt)
}

// AvailableAlgorithms returns a list of all registered hash algorithm aliases from the default factory.
func AvailableAlgorithms() []string {
	return defaultFactory.AvailableAlgorithms()
}
