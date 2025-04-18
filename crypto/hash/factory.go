/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"fmt"

	"github.com/goexts/generic/settings"

	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type internalFactory interface {
	create(algType types.Type, opts ...types.Option) (interfaces.Cryptographic, error)
}

type algorithmFactory struct {
	cryptos map[types.Type]interfaces.Cryptographic
}

func (f *algorithmFactory) create(algType types.Type, opts ...types.Option) (interfaces.Cryptographic, error) {
	if alg, exists := f.cryptos[algType]; exists {
		return alg, nil
	}

	algorithm, exists := algorithms[algType]
	if !exists {
		return nil, fmt.Errorf("unsupported algorithm: %s", algType)
	}

	cfg := f.createConfig(algorithm, opts...)
	alg, err := algorithm.creator(cfg)
	if err != nil {
		return nil, err
	}

	f.cryptos[algType] = alg
	return alg, nil
}

// 统一配置创建逻辑
func (f *algorithmFactory) createConfig(algorithm algorithm, opts ...types.Option) *types.Config {
	cfg := &types.Config{}
	if algorithm.defaultConfig != nil {
		cfg = algorithm.defaultConfig()
	}
	return settings.Apply(cfg, opts)
}

func createFactory() internalFactory {
	return &algorithmFactory{
		cryptos: make(map[types.Type]interfaces.Cryptographic),
	}
}
