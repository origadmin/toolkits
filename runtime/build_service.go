// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package runtime implements the functions, types, and interfaces for the module.
package runtime

import (
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/runtime/service"
)

type (
	// ServiceBuildRegistry is an interface that defines a method for registering a service builder.
	serviceBuildRegistry interface {
		RegisterServiceBuilder(name string, builder ServiceBuilder)
	}
	// ServiceBuilder is an interface that defines a method for creating a new service.
	ServiceBuilder interface {
		NewGRPCServer(cfg *configv1.Service) (*service.GRPCServer, error)
		NewHTTPServer(cfg *configv1.Service) (*service.HTTPServer, error)
		NewGRPCClient(cfg *configv1.Service) (*service.GRPCClient, error)
		NewHTTPClient(cfg *configv1.Service) (*service.HTTPClient, error)
	}
)
