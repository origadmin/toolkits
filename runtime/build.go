/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package runtime implements the functions, types, and interfaces for the module.
package runtime

import (
	"sync"

	"github.com/origadmin/toolkits/runtime/config"
	"github.com/origadmin/toolkits/runtime/registry"
)

// builder is a struct that holds a map of ConfigBuilders and a map of RegistryBuilders.
type builder struct {
	configs     map[string]ConfigBuilder
	configMux   sync.RWMutex
	registries  map[string]RegistryBuilder
	registryMux sync.RWMutex
}

// NewConfig creates a new Config object based on the given SourceConfig and options.
func (b *builder) NewConfig(cfg *config.SourceConfig, opts ...config.Option) (config.Config, error) {
	b.configMux.RLock()
	defer b.configMux.RUnlock()
	configBuilder, ok := build.configs[cfg.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return configBuilder.NewConfig(cfg, opts...)
}

// NewRegistrar creates a new Registrar object based on the given RegistryConfig.
func (b *builder) NewRegistrar(cfg *config.Registry) (registry.Registrar, error) {
	b.registryMux.RLock()
	defer b.registryMux.RUnlock()
	registryBuilder, ok := build.registries[cfg.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return registryBuilder.NewRegistrar(cfg)
}

// NewDiscovery creates a new Discovery object based on the given RegistryConfig.
func (b *builder) NewDiscovery(cfg *config.Registry) (registry.Discovery, error) {
	b.registryMux.RLock()
	defer b.registryMux.RUnlock()
	registryBuilder, ok := build.registries[cfg.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return registryBuilder.NewDiscovery(cfg)
}

// RegisterConfig registers a new ConfigBuilder with the given name.
func (b *builder) RegisterConfig(name string, configBuilder ConfigBuilder) {
	b.configMux.Lock()
	defer b.configMux.Unlock()
	build.configs[name] = configBuilder
}

// RegisterConfigFunc registers a new ConfigBuilder with the given name and function.
func (b *builder) RegisterConfigFunc(name string, configBuilder ConfigBuildFunc) {
	b.RegisterConfig(name, configBuilder)
}

// RegisterRegistry registers a new RegistryBuilder with the given name.
func (b *builder) RegisterRegistry(name string, registryBuilder RegistryBuilder) {
	b.registryMux.Lock()
	defer b.registryMux.Unlock()
	build.registries[name] = registryBuilder
}

// RegisterRegistryFunc registers a new RegistryBuilder with the given name and functions.
func (b *builder) RegisterRegistryFunc(name string, registryBuilder RegistrarBuildFunc, discoveryBuilder DiscoveryBuildFunc) {
	b.RegisterRegistry(name, &registryWrap{
		RegistrarBuildFunc: registryBuilder,
		DiscoveryBuildFunc: discoveryBuilder,
	})
}
