/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package runtime implements the functions, types, and interfaces for the module.
package runtime

import (
	"github.com/origadmin/toolkits/runtime/config"
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
)

type (
	// configBuildRegistry is an interface that defines a method for registering a config builder.
	configBuildRegistry interface {
		// RegisterConfigBuilder registers a config builder with the given name.
		RegisterConfigBuilder(name string, configBuilder ConfigBuilder)
	}
	// ConfigBuilder is an interface that defines a method for creating a new config.
	ConfigBuilder interface {
		// NewConfig creates a new config using the given SourceConfig and a list of Options.
		NewConfig(cfg *configv1.SourceConfig, ss ...config.SettingFunc) (config.Config, error)
	}
)

// ConfigBuildFunc is a function type that takes a SourceConfig and a list of Options and returns a Config and an error.
type ConfigBuildFunc func(*configv1.SourceConfig, ...config.SettingFunc) (config.Config, error)

// NewConfig is a method that implements the ConfigBuilder interface for ConfigBuildFunc.
func (fn ConfigBuildFunc) NewConfig(cfg *configv1.SourceConfig, ss ...config.SettingFunc) (config.Config, error) {
	// Call the function with the given SourceConfig and a list of Options.
	return fn(cfg, ss...)
}

// NewConfig creates a new Config object based on the given SourceConfig and options.
func (b *builder) NewConfig(cfg *configv1.SourceConfig, ss ...config.SettingFunc) (config.Config, error) {
	b.configMux.RLock()
	defer b.configMux.RUnlock()
	configBuilder, ok := b.configs[cfg.Type]
	if !ok {
		return nil, ErrNotFound
	}

	return configBuilder.NewConfig(cfg, ss...)
}

// RegisterConfigBuilder registers a new ConfigBuilder with the given name.
func (b *builder) RegisterConfigBuilder(name string, configBuilder ConfigBuilder) {
	b.configMux.Lock()
	defer b.configMux.Unlock()
	b.configs[name] = configBuilder
}

// RegisterConfigFunc registers a new ConfigBuilder with the given name and function.
func (b *builder) RegisterConfigFunc(name string, configBuilder ConfigBuildFunc) {
	b.RegisterConfigBuilder(name, configBuilder)
}
