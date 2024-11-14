// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package runtime implements the functions, types, and interfaces for the module.
package runtime

import (
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc"

	"github.com/origadmin/toolkits/runtime/config"
)

type (
	// ServiceBuildRegistry is an interface that defines a method for registering a service builder.
	serviceBuildRegistry interface {
		RegisterServiceBuilder(name string, builder ServiceBuilder)
	}
	// ServiceBuilder is an interface that defines a method for creating a new service.
	ServiceBuilder interface {
		NewGRPCServer(cfg *config.Service) (*transgrpc.Server, error)
		NewHTTPServer(cfg *config.Service) (*transhttp.Server, error)
		NewGRPCClient(cfg *config.Service) (*grpc.ClientConn, error)
		NewHTTPClient(cfg *config.Service) (*transhttp.Client, error)
	}
)
