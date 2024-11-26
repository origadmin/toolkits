/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package config implements the functions, types, and interfaces for the module.
package config

import (
	"github.com/go-kratos/kratos/contrib/config/consul/v2"
	"github.com/hashicorp/consul/api"

	"github.com/origadmin/toolkits/context"
	"github.com/origadmin/toolkits/runtime/config"
)

const (
	DefaultPathName = "bootstrap.json"
)

type (
	Option = consul.Option
)

// New returns a new consul config source
func New(client *api.Client, opts ...Option) (config.Source, error) {
	return consul.New(client, opts...)
}

// WithContext with registry context
func WithContext(ctx context.Context) Option {
	return consul.WithContext(ctx)
}

// WithPath with registry path
func WithPath(p string) Option {
	return consul.WithPath(p)
}
