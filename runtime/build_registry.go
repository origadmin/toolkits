/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package runtime implements the functions, types, and interfaces for the module.
package runtime

import (
	"github.com/origadmin/toolkits/runtime/config"
	"github.com/origadmin/toolkits/runtime/registry"
)

// registryBuildRegistry is an interface that defines a method for registering a RegistryBuilder.
type registryBuildRegistry interface {
	RegisterRegistry(name string, registryBuilder RegistryBuilder)
}

// RegistryBuilder is an interface that defines methods for creating a Discovery and a Registrar.
type RegistryBuilder interface {
	NewRegistrar(cfg *config.Registry) (registry.Registrar, error)
	NewDiscovery(cfg *config.Registry) (registry.Discovery, error)
}

// RegistrarBuildFunc is a function type that takes a *config.RegistryConfig and returns a registry.Registrar and an error.
type RegistrarBuildFunc func(cfg *config.Registry) (registry.Registrar, error)

// NewRegistrar is a method that calls the RegistrarBuildFunc with the given config.
func (fn RegistrarBuildFunc) NewRegistrar(cfg *config.Registry) (registry.Registrar, error) {
	return fn(cfg)
}

// DiscoveryBuildFunc is a function type that takes a *config.RegistryConfig and returns a registry.Discovery and an error.
type DiscoveryBuildFunc func(cfg *config.Registry) (registry.Discovery, error)

// NewDiscovery is a method that calls the DiscoveryBuildFunc with the given config.
func (fn DiscoveryBuildFunc) NewDiscovery(cfg *config.Registry) (registry.Discovery, error) {
	return fn(cfg)
}

// registryWrap is a struct that embeds RegistrarBuildFunc and DiscoveryBuildFunc.
type registryWrap struct {
	RegistrarBuildFunc
	DiscoveryBuildFunc
}

// _ is a blank identifier that is used to satisfy the interface requirement for RegistryBuilder.
var _ RegistryBuilder = &registryWrap{}
