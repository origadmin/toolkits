/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package runtime provides functions for loading configurations and registering services.
package runtime

import (
	"sync"

	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/runtime/config"
	"github.com/origadmin/toolkits/runtime/registry"
)

type Builder interface {
	ConfigBuilder
	RegistryBuilder

	configBuildRegistry
	registryBuildRegistry
}

// build is a global variable that holds an instance of the builder struct.
var (
	once  = &sync.Once{}
	build = &builder{}
)

// ErrNotFound is an error that is returned when a ConfigBuilder or RegistryBuilder is not found.
var ErrNotFound = errors.String("not found")

// init initializes the builder struct.
func init() {
	once.Do(func() {
		build.init()
	})
}

// init initializes the builder struct.
func (b *builder) init() {
	b.configs = make(map[string]ConfigBuilder)
	b.registries = make(map[string]RegistryBuilder)
}

// RegisterConfig registers a ConfigBuilder with the builder.
func RegisterConfig(name string, configBuilder ConfigBuilder) {
	build.RegisterConfig(name, configBuilder)
}

// NewConfig creates a new Config using the registered ConfigBuilder.
func NewConfig(cfg *config.SourceConfig, opts ...config.Option) (config.Config, error) {
	return build.NewConfig(cfg, opts...)
}

// RegisterRegistry registers a RegistryBuilder with the builder.
func RegisterRegistry(name string, registryBuilder RegistryBuilder) {
	build.RegisterRegistry(name, registryBuilder)
}

// NewDiscovery creates a new Discovery using the registered RegistryBuilder.
func NewDiscovery(cfg *config.RegistryConfig) (registry.Discovery, error) {
	return build.NewDiscovery(cfg)
}

// NewRegistrar creates a new Registrar using the registered RegistryBuilder.
func NewRegistrar(cfg *config.RegistryConfig) (registry.Registrar, error) {
	return build.NewRegistrar(cfg)
}

func New() Builder {
	b := &builder{}
	b.init()
	return b
}
