// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package runtime implements the functions, types, and interfaces for the module.
package runtime

import (
	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"
	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"google.golang.org/grpc"

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

// NewGRPCServer creates a new gRPC server based on the given ServiceConfig.
func (b *builder) NewGRPCServer(cfg *configv1.Service) (*transgrpc.Server, error) {
	b.serviceMux.RLock()
	defer b.serviceMux.RUnlock()
	if serviceBuilder, ok := b.services[cfg.Name]; ok {
		return serviceBuilder.NewGRPCServer(cfg)
	}
	return nil, ErrNotFound
}

// NewHTTPServer creates a new HTTP server based on the given ServiceConfig.
func (b *builder) NewHTTPServer(cfg *configv1.Service) (*transhttp.Server, error) {
	b.serviceMux.RLock()
	defer b.serviceMux.RUnlock()
	if serviceBuilder, ok := b.services[cfg.Name]; ok {
		return serviceBuilder.NewHTTPServer(cfg)
	}
	return nil, ErrNotFound
}

// NewGRPCClient creates a new gRPC client based on the given ServiceConfig.
func (b *builder) NewGRPCClient(cfg *configv1.Service) (*grpc.ClientConn, error) {
	b.serviceMux.RLock()
	defer b.serviceMux.RUnlock()
	if serviceBuilder, ok := b.services[cfg.Name]; ok {
		return serviceBuilder.NewGRPCClient(cfg)
	}
	return nil, ErrNotFound
}

// NewHTTPClient creates a new HTTP client based on the given ServiceConfig.
func (b *builder) NewHTTPClient(cfg *configv1.Service) (*transhttp.Client, error) {
	b.serviceMux.RLock()
	defer b.serviceMux.RUnlock()
	if serviceBuilder, ok := b.services[cfg.Name]; ok {
		return serviceBuilder.NewHTTPClient(cfg)
	}
	return nil, ErrNotFound
}

// RegisterServiceBuilder registers a new ServiceBuilder with the given service name.
func (b *builder) RegisterServiceBuilder(name string, builder ServiceBuilder) {
	b.serviceMux.Lock()
	defer b.serviceMux.Unlock()
	b.services[name] = builder
}
