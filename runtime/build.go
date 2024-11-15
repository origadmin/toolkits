/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package runtime implements the functions, types, and interfaces for the module.
package runtime

import (
	"sync"

	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc"

	"github.com/origadmin/toolkits/runtime/bootstrap"
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/runtime/registry"
)

// builder is a struct that holds a map of ConfigBuilders and a map of RegistryBuilders.
type builder struct {
	configMux   sync.RWMutex
	configs     map[string]ConfigBuilder
	registryMux sync.RWMutex
	registries  map[string]RegistryBuilder
	serviceMux  sync.RWMutex
	services    map[string]ServiceBuilder
}

// NewConfig creates a new Config object based on the given SourceConfig and options.
func (b *builder) NewConfig(cfg *configv1.SourceConfig, opts ...bootstrap.Option) (bootstrap.Config, error) {
	b.configMux.RLock()
	defer b.configMux.RUnlock()
	configBuilder, ok := build.configs[cfg.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return configBuilder.NewConfig(cfg, opts...)
}

// RegisterConfigBuilder registers a new ConfigBuilder with the given name.
func (b *builder) RegisterConfigBuilder(name string, configBuilder ConfigBuilder) {
	b.configMux.Lock()
	defer b.configMux.Unlock()
	build.configs[name] = configBuilder
}

// RegisterConfigFunc registers a new ConfigBuilder with the given name and function.
func (b *builder) RegisterConfigFunc(name string, configBuilder ConfigBuildFunc) {
	b.RegisterConfigBuilder(name, configBuilder)
}

// NewRegistrar creates a new Registrar object based on the given RegistryConfig.
func (b *builder) NewRegistrar(cfg *configv1.Registry) (registry.Registrar, error) {
	b.registryMux.RLock()
	defer b.registryMux.RUnlock()
	registryBuilder, ok := build.registries[cfg.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return registryBuilder.NewRegistrar(cfg)
}

// NewDiscovery creates a new Discovery object based on the given RegistryConfig.
func (b *builder) NewDiscovery(cfg *configv1.Registry) (registry.Discovery, error) {
	b.registryMux.RLock()
	defer b.registryMux.RUnlock()
	registryBuilder, ok := build.registries[cfg.Type]
	if !ok {
		return nil, ErrNotFound
	}
	return registryBuilder.NewDiscovery(cfg)
}

// RegisterRegistryBuilder registers a new RegistryBuilder with the given name.
func (b *builder) RegisterRegistryBuilder(name string, registryBuilder RegistryBuilder) {
	b.registryMux.Lock()
	defer b.registryMux.Unlock()
	build.registries[name] = registryBuilder
}

// RegisterRegistryFunc registers a new RegistryBuilder with the given name and functions.
func (b *builder) RegisterRegistryFunc(name string, registryBuilder RegistrarBuildFunc, discoveryBuilder DiscoveryBuildFunc) {
	b.RegisterRegistryBuilder(name, &registryWrap{
		RegistrarBuildFunc: registryBuilder,
		DiscoveryBuildFunc: discoveryBuilder,
	})
}

// NewGRPCServer creates a new gRPC server based on the given ServiceConfig.
func (b *builder) NewGRPCServer(cfg *configv1.Service) (*transgrpc.Server, error) {
	b.serviceMux.RLock()
	defer b.serviceMux.RUnlock()
	serviceBuilder, ok := build.services[cfg.Name]
	if !ok {
		return nil, ErrNotFound
	}
	return serviceBuilder.NewGRPCServer(cfg)
}

// NewHTTPServer creates a new HTTP server based on the given ServiceConfig.
func (b *builder) NewHTTPServer(cfg *configv1.Service) (*transhttp.Server, error) {
	b.serviceMux.RLock()
	defer b.serviceMux.RUnlock()
	serviceBuilder, ok := build.services[cfg.Name]
	if !ok {
		return nil, ErrNotFound
	}
	return serviceBuilder.NewHTTPServer(cfg)
}

// NewGRPCClient creates a new gRPC client based on the given ServiceConfig.
func (b *builder) NewGRPCClient(cfg *configv1.Service) (*grpc.ClientConn, error) {
	b.serviceMux.RLock()
	defer b.serviceMux.RUnlock()
	serviceBuilder, ok := build.services[cfg.Name]
	if !ok {
		return nil, ErrNotFound
	}
	return serviceBuilder.NewGRPCClient(cfg)
}

// NewHTTPClient creates a new HTTP client based on the given ServiceConfig.
func (b *builder) NewHTTPClient(cfg *configv1.Service) (*transhttp.Client, error) {
	b.serviceMux.RLock()
	defer b.serviceMux.RUnlock()
	serviceBuilder, ok := build.services[cfg.Name]
	if !ok {
		return nil, ErrNotFound
	}
	return serviceBuilder.NewHTTPClient(cfg)
}

// RegisterServiceBuilder registers a new ServiceBuilder with the given service name.
func (b *builder) RegisterServiceBuilder(name string, builder ServiceBuilder) {
	b.serviceMux.Lock()
	defer b.serviceMux.Unlock()
	build.services[name] = builder
}
