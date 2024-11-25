/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package runtime provides functions for loading configurations and registering services.
package runtime

import (
	"sync"

	"github.com/origadmin/toolkits/errors"
	"github.com/origadmin/toolkits/runtime/config"
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/runtime/middleware"
	"github.com/origadmin/toolkits/runtime/registry"
)

type Builder interface {
	ConfigBuilder
	RegistryBuilder
	ServiceBuilder
	MiddlewareBuilders

	configBuildRegistry
	registryBuildRegistry
	serviceBuildRegistry
	middlewareBuildRegistry
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

// Global returns the global instance of the builder.
func Global() Builder {
	return build
}

// NewConfig creates a new Config using the registered ConfigBuilder.
func NewConfig(cfg *configv1.SourceConfig, ss ...config.SettingFunc) (config.Config, error) {
	return build.NewConfig(cfg, ss...)
}

// RegisterConfig registers a ConfigBuilder with the builder.
func RegisterConfig(name string, configBuilder ConfigBuilder) {
	build.RegisterConfigBuilder(name, configBuilder)
}

// RegisterConfigFunc registers a ConfigBuilder with the builder.
func RegisterConfigFunc(name string, buildFunc ConfigBuildFunc) {
	build.RegisterConfigBuilder(name, buildFunc)
}

// NewDiscovery creates a new Discovery using the registered RegistryBuilder.
func NewDiscovery(cfg *configv1.Registry) (registry.Discovery, error) {
	return build.NewDiscovery(cfg)
}

// NewRegistrar creates a new Registrar using the registered RegistryBuilder.
func NewRegistrar(cfg *configv1.Registry) (registry.Registrar, error) {
	return build.NewRegistrar(cfg)
}

// RegisterRegistry registers a RegistryBuilder with the builder.
func RegisterRegistry(name string, registryBuilder RegistryBuilder) {
	build.RegisterRegistryBuilder(name, registryBuilder)
}

// NewMiddlewareClient creates a new Middleware with the builder.
func NewMiddlewareClient(name string, cm *configv1.Customize_Config) (middleware.Middleware, error) {
	return build.NewMiddlewareClient(name, cm)
}

// NewMiddlewareServer creates a new Middleware with the builder.
func NewMiddlewareServer(name string, cm *configv1.Customize_Config) (middleware.Middleware, error) {
	return build.NewMiddlewareServer(name, cm)
}

// NewMiddlewaresClient creates a new Middleware with the builder.
func NewMiddlewaresClient(cc *configv1.Customize) []middleware.Middleware {
	return build.NewMiddlewaresClient(nil, cc)
}

// NewMiddlewaresServer creates a new Middleware with the builder.
func NewMiddlewaresServer(cc *configv1.Customize) []middleware.Middleware {
	return build.NewMiddlewaresServer(nil, cc)
}

// RegisterMiddleware registers a MiddlewareBuilder with the builder.
func RegisterMiddleware(name string, middlewareBuilder MiddlewareBuilder) {
	build.RegisterMiddlewareBuilder(name, middlewareBuilder)
}

// New creates a new Builder.
func New() Builder {
	return newBuilder()
}
