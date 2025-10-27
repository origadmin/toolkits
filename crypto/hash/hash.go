/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides the hash functions
package hash

import (
	"log/slog"
	"os"
	"sync"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

var (
	// activeCrypto is the currently active cryptographic instance
	activeCrypto Crypto
	// activeCryptoMu protects activeCrypto from concurrent access
	activeCryptoMu sync.RWMutex
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

func init() {
	algStr := os.Getenv(types.ENV)
	if algStr == "" {
		algStr = types.DefaultSpec
	}

	// Try to Create an encryption instance with the defined algorithm type
	crypto, err := NewCrypto(algStr)
	activeCryptoMu.Lock()
	defer activeCryptoMu.Unlock()
	if err != nil {
		slog.Error("hash: failed to initialize active crypto", "type", algStr, "error", err)
		// If the hash module fails to initialize, use a no-op implementation
		activeCrypto = &uninitializedCrypto{}
	} else {
		activeCrypto = crypto
	}
}

// UseCrypto updates the active cryptographic instance
func UseCrypto(algName string, opts ...Option) error {
	if algName == "" {
		return errors.ErrInvalidAlgorithm
	}
	algSpec, err := types.Parse(algName)
	if err != nil {
		return err
	}

	activeCryptoMu.RLock()
	currentCrypto := activeCrypto
	activeCryptoMu.RUnlock()

	if currentCrypto != nil && currentCrypto.Spec().Is(algSpec) {
		return nil
	}

	newCrypto, err := NewCrypto(algName, opts...)
	if err != nil {
		return err
	}
	activeCryptoMu.Lock()
	activeCrypto = newCrypto
	activeCryptoMu.Unlock()
	return nil
}

// Verify verifies a password using the active cryptographic instance.
func Verify(hashed, password string) error {
	activeCryptoMu.RLock()
	defer activeCryptoMu.RUnlock()
	return activeCrypto.Verify(hashed, password)
}

// Generate generates a hash for the given password using the active cryptographic instance.
func Generate(password string) (string, error) {
	activeCryptoMu.RLock()
	defer activeCryptoMu.RUnlock()
	return activeCrypto.Hash(password)
}

// GenerateWithSalt generates a hash for the given password with the specified salt using the active cryptographic instance.
func GenerateWithSalt(password string, salt []byte) (string, error) {
	activeCryptoMu.RLock()
	defer activeCryptoMu.RUnlock()
	return activeCrypto.HashWithSalt(password, salt)
}

// AvailableAlgorithms returns a list of all registered hash algorithms.
func AvailableAlgorithms() []types.Spec {
	var algorithms []types.Spec
	for _, algEntry := range algorithmMap {
		algorithms = append(algorithms, algEntry.algSpec)
	}
	return algorithms
}
