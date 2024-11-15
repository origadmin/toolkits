// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package grpc implements the functions, types, and interfaces for the module.
package grpc

import (
	"net/url"

	transgrpc "github.com/go-kratos/kratos/v2/transport/grpc"

	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/runtime/middleware"
	"github.com/origadmin/toolkits/utils"
)

// NewServer Create a GRPC server
func NewServer(cfg *configv1.Service, m ...middleware.Middleware) *transgrpc.Server {
	var options []transgrpc.ServerOption

	var ms []middleware.Middleware

	ms = middleware.NewServer(cfg.GetMiddleware())
	ms = append(ms, m...)
	options = append(options, transgrpc.Middleware(ms...))

	serviceGrpc := cfg.GetGrpc()
	if serviceGrpc != nil {
		if serviceGrpc.Network != "" {
			options = append(options, transgrpc.Network(serviceGrpc.Network))
		}
		if serviceGrpc.Addr != "" {
			options = append(options, transgrpc.Address(serviceGrpc.Addr))
		}
		if serviceGrpc.Timeout != nil {
			options = append(options, transgrpc.Timeout(serviceGrpc.Timeout.AsDuration()))
		}
		if cfg.AutoEndpoint {
			endpoint := utils.DiscoveryEndpoint(serviceGrpc.Endpoint, "grpc", cfg.Host, serviceGrpc.Addr)
			if e, err := url.Parse(endpoint); err == nil {
				options = append(options, transgrpc.Endpoint(e))
			}
		}
	}

	srv := transgrpc.NewServer(options...)
	return srv
}
