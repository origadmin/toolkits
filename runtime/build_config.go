/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package runtime implements the functions, types, and interfaces for the module.
package runtime

import (
	"github.com/origadmin/toolkits/runtime/config"
)

type configBuildRegistry interface {
	RegisterConfig(name string, configBuilder ConfigBuilder)
}

type ConfigBuilder interface {
	NewConfig(cfg *config.SourceConfig, opts ...config.Option) (config.Config, error)
}

// ConfigBuildFunc is a function type that takes a SourceConfig and a list of Options and returns a Config and an error.
type ConfigBuildFunc func(*config.SourceConfig, ...config.Option) (config.Config, error)

func (fn ConfigBuildFunc) NewConfig(cfg *config.SourceConfig, opts ...config.Option) (config.Config, error) {
	return fn(cfg, opts...)
}
