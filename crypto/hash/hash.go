/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides the hash functions
package hash

import (
	"os"

	"github.com/origadmin/toolkits/errors"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

const (
	// ENV environment variable name
	ENV = "ORIGADMIN_HASH_TYPE"
	// ErrPasswordNotMatch error when password not match
	ErrPasswordNotMatch = errors.String("password not match")
)

var (
	// defaultCrypto default cryptographic instance
	defaultCrypto Crypto
)

func init() {
	alg := os.Getenv(ENV)
	if alg == "" {
		alg = core.DefaultType
	}
	t := types.ParseType(alg)
	if t != types.TypeUnknown {
		crypto, err := NewCrypto(t)
		if err == nil {
			defaultCrypto = crypto
		}
	}
	if defaultCrypto == nil {
		cryptographic, err := NewCrypto(core.DefaultType)
		if err != nil {
			panic(err)
		}
		defaultCrypto = cryptographic
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
