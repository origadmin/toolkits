/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash provides the hash functions
package hash

import (
	"os"

	"github.com/goexts/generic/settings"
	"github.com/origadmin/toolkits/errors"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
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
	defaultCrypto interfaces.Cryptographic
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
func UseCrypto(t types.Type, opts ...types.ConfigOption) error {
	if alg, ok := algorithms[t]; ok {
		cfg := &types.Config{}
		if alg.defaultConfig != nil {
			cfg = alg.defaultConfig()
		}
		cfg = settings.Apply(cfg, opts)
		crypto, err := alg.creator(cfg)
		if err != nil {
			return err
		}
		defaultCrypto = crypto
	}
	return errors.Errorf("unsupported hash type: %s", t)
}
