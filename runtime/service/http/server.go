// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package grpc implements the functions, types, and interfaces for the module.
package grpc

import (
	"net/url"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"

	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/runtime/middleware"
	"github.com/origadmin/toolkits/utils"
)

// NewServer Create a HTTP server
func NewServer(cfg *configv1.Service, m ...middleware.Middleware) *transhttp.Server {
	var options []transhttp.ServerOption

	var ms []middleware.Middleware

	ms = middleware.NewServer(cfg.GetMiddleware())
	ms = append(ms, m...)
	options = append(options, transhttp.Middleware(ms...))

	serviceHttp := cfg.GetHttp()
	if serviceHttp != nil {
		if serviceHttp.Network != "" {
			options = append(options, transhttp.Network(serviceHttp.Network))
		}
		if serviceHttp.Addr != "" {
			options = append(options, transhttp.Address(serviceHttp.Addr))
		}
		if serviceHttp.Timeout != nil {
			options = append(options, transhttp.Timeout(serviceHttp.Timeout.AsDuration()))
		}
		if cfg.AutoEndpoint {
			endpoint := utils.DiscoveryEndpoint(serviceHttp.Endpoint, "http", cfg.Host, serviceHttp.Addr)
			if e, err := url.Parse(endpoint); err == nil {
				options = append(options, transhttp.Endpoint(e))
			}
		}
	}

	srv := transhttp.NewServer(options...)
	return srv
}
