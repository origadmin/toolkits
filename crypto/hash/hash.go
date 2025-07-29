/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides the hash functions
package hash

import (
	"log/slog"
	"os"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

const (
	// ENV environment variable name
	ENV = "ORIGADMIN_HASH_TYPE"
)

var (
	// activeCrypto is the currently active cryptographic instance
	activeCrypto Crypto
	// ErrHashModuleNotInitialized is returned when the hash module fails to initialize.
	ErrHashModuleNotInitialized = errors.ErrHashModuleNotInitialized
)

// uninitializedCrypto is a no-op Crypto implementation used when the module fails to initialize.
type uninitializedCrypto struct{}

func (u *uninitializedCrypto) Type() types.Type {
	return types.Type{Name: constants.UNKNOWN}
}

func (u *uninitializedCrypto) Hash(password string) (string, error) {
	return "", ErrHashModuleNotInitialized
}

func (u *uninitializedCrypto) HashWithSalt(password string, salt []byte) (string, error) {
	return "", ErrHashModuleNotInitialized
}

func (u *uninitializedCrypto) Verify(hashed, password string) error {
	return ErrHashModuleNotInitialized
}

func init() {
	algStr := os.Getenv(constants.ENV)
	if algStr == "" {
		algStr = constants.DefaultType
	}

	// Try to create an encryption instance with the defined algorithm type
	crypto, err := NewCrypto(algStr)
	if err != nil {
		slog.Error("hash: failed to initialize active crypto", "type", algStr, "error", err)
		// If the hash module fails to initialize, use a no-op implementation
		activeCrypto = &uninitializedCrypto{}
	} else {
		activeCrypto = crypto
	}
}

// UseCrypto updates the active cryptographic instance
func UseCrypto(algName string, opts ...types.Option) error {
	if algName == "" {
		return errors.ErrInvalidAlgorithm
	}
	algType, err := types.ParseType(algName)
	if err != nil {
		return err
	}
	if activeCrypto != nil && activeCrypto.Type() == algType {
		return nil
	}
	newCrypto, err := NewCrypto(algName, opts...)
	if err != nil {
		return err
	}
	activeCrypto = newCrypto
	return nil
}

// Verify verifies a password using the active cryptographic instance.
func Verify(hashed, password string) error {
	return activeCrypto.Verify(hashed, password)
}

// Generate generates a hash for the given password using the active cryptographic instance.
func Generate(password string) (string, error) {
	return activeCrypto.Hash(password)
}

// GenerateWithSalt generates a hash for the given password with the specified salt using the active cryptographic instance.
func GenerateWithSalt(password string, salt []byte) (string, error) {
	return activeCrypto.HashWithSalt(password, salt)
}

// AvailableAlgorithms returns a list of all registered hash algorithms.
func AvailableAlgorithms() []types.Type {
	var algorithms []types.Type
	for algName := range algorithmMap {
		algType, err := types.ParseType(algName)
		if err == nil {
			algorithms = append(algorithms, algType)
		}
	}
	return algorithms
}
