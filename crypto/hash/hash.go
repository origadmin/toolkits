/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides the hash functions
package hash

import (
	"os"

	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/errors"
)

const (
	// ENV environment variable name
	ENV = "ORIGADMIN_HASH_TYPE"
	// ErrPasswordNotMatch error when password not match
	ErrPasswordNotMatch = errors.String("password not match")
)

var (
	// defaultCrypto default cryptographic instance
	defaultCrypto interfaces.Cryptographic
)

func init() {
	// Use Argon2 as default algorithm
	crypto, err := NewCrypto(types.WithAlgorithm(types.TypeArgon2))
	if err != nil {
		panic(err)
	}
	defaultCrypto = crypto

	env := os.Getenv(ENV)
	if env == "" {
		return
	}
	t := types.ParseType(env)
	if creator, ok := algorithms[t]; ok {
		cfg := &types.Config{}
		crypto, err := creator(cfg)
		if err == nil {
			defaultCrypto = crypto
		}
	}
}

// UseCrypto updates the default cryptographic instance
func UseCrypto(t types.Type) {
	if creator, ok := algorithms[t]; ok {
		cfg := &types.Config{}
		crypto, err := creator(cfg)
		if err == nil {
			defaultCrypto = crypto
		}
	}
}

// Generate generates a password hash using the default cryptographic instance
func Generate(password string, salt string) (string, error) {
	return defaultCrypto.HashWithSalt(password, salt)
}

// Compare compares the given hashed password with the plaintext password using the default cryptographic instance
func Compare(hashpass, password, salt string) error {
	return defaultCrypto.Verify(hashpass, password)
}
