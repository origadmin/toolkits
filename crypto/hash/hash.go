/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides the hash functions
package hash

import (
	"fmt"
	"log"
	"os"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

const (
	// ENV environment variable name
	ENV = "ORIGADMIN_HASH_TYPE"
)

var (
	// defaultCrypto default cryptographic instance
	defaultCrypto Crypto
	// ErrHashModuleNotInitialized is returned when the hash module fails to initialize.
	ErrHashModuleNotInitialized = fmt.Errorf("hash module not initialized")
)

// uninitializedCrypto is a no-op Crypto implementation used when the module fails to initialize.
type uninitializedCrypto struct{}

func (u *uninitializedCrypto) Type() types.Type {
	return types.TypeUnknown
}

func (u *uninitializedCrypto) Hash(password string) (string, error) {
	return "", ErrHashModuleNotInitialized
}

func (u *uninitializedCrypto) HashWithSalt(password, salt string) (string, error) {
	return "", ErrHashModuleNotInitialized
}

func (u *uninitializedCrypto) Verify(hashed, password string) error {
	return ErrHashModuleNotInitialized
}

func init() {
	alg := os.Getenv(constants.ENV)
	if alg == "" {
		alg = constants.DefaultType
	}
	t := types.ParseType(alg)

	var initialized bool
	if t != types.TypeUnknown {
		crypto, err := NewCrypto(t)
		if err == nil {
			defaultCrypto = crypto
			initialized = true
		} else {
			log.Printf("hash: failed to initialize default crypto with type %s: %v", t, err)
		}
	}

	if !initialized {
		// Try to initialize with the hardcoded default type if the environment variable type failed or was unknown.
		cryptographic, err := NewCrypto(constants.DefaultType)
		if err != nil {
			log.Printf("hash: failed to initialize default crypto with hardcoded default type %s: %v", constants.DefaultType, err)
			// If even the hardcoded default fails, set to uninitializedCrypto to prevent panics.
			defaultCrypto = &uninitializedCrypto{}
		} else {
			defaultCrypto = cryptographic
		}
	}
}

// UseCrypto updates the default cryptographic instance
func UseCrypto(t types.Type, opts ...types.Option) error {
	if defaultCrypto != nil && defaultCrypto.Type() == t {
		return nil
	}
	newCrypto, err := NewCrypto(t, opts...)
	if err != nil {
		return err
	}
	defaultCrypto = newCrypto
	return nil
}

// Verify verifies a password using the default cryptographic instance.
func Verify(hashed, password string) error {
	return defaultCrypto.Verify(hashed, password)
}

// Generate generates a hash for the given password using the default cryptographic instance.
func Generate(password string) (string, error) {
	return defaultCrypto.Hash(password)
}

// GenerateWithSalt generates a hash for the given password with the specified salt using the default cryptographic instance.
func GenerateWithSalt(password, salt string) (string, error) {
	return defaultCrypto.HashWithSalt(password, salt)
}
