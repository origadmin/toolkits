/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package runtime implements the functions, types, and interfaces for the module.
package runtime

import (
	"github.com/origadmin/toolkits/runtime/config"
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
		NewConfig(cfg *config.SourceConfig, opts ...config.Option) (config.Config, error)
	}
)

// ConfigBuildFunc is a function type that takes a SourceConfig and a list of Options and returns a Config and an error.
type ConfigBuildFunc func(*config.SourceConfig, ...config.Option) (config.Config, error)

// NewConfig is a method that implements the ConfigBuilder interface for ConfigBuildFunc.
func (fn ConfigBuildFunc) NewConfig(cfg *config.SourceConfig, opts ...config.Option) (config.Config, error) {
	// Call the function with the given SourceConfig and a list of Options.
	return fn(cfg, opts...)
}
