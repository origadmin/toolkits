/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package runtime implements the functions, types, and interfaces for the module.
package runtime

import (
	"github.com/origadmin/toolkits/runtime/config"
	"github.com/origadmin/toolkits/runtime/registry"
)

type registryBuildRegistry interface {
	RegisterRegistry(name string, registryBuilder RegistryBuilder)
}

// RegistryBuilder is an interface that defines methods for creating a Discovery and a Registrar.
type RegistryBuilder interface {
	NewRegistrar(cfg *config.RegistryConfig) (registry.Registrar, error)
	NewDiscovery(cfg *config.RegistryConfig) (registry.Discovery, error)
}

type RegistrarBuildFunc func(cfg *config.RegistryConfig) (registry.Registrar, error)

func (fn RegistrarBuildFunc) NewRegistrar(cfg *config.RegistryConfig) (registry.Registrar, error) {
	return fn(cfg)
}

type DiscoveryBuildFunc func(cfg *config.RegistryConfig) (registry.Discovery, error)

func (fn DiscoveryBuildFunc) NewDiscovery(cfg *config.RegistryConfig) (registry.Discovery, error) {
	return fn(cfg)
}

type registryWrap struct {
	RegistrarBuildFunc
	DiscoveryBuildFunc
}

var _ RegistryBuilder = &registryWrap{}
