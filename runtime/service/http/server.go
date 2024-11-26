/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package grpc implements the functions, types, and interfaces for the module.
package grpc

import (
	"net/url"

	transhttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/goexts/generic/settings"

	"github.com/origadmin/toolkits/helpers"
	"github.com/origadmin/toolkits/runtime/config"
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
	"github.com/origadmin/toolkits/runtime/middleware"
)

// NewServer Create an HTTP server instance.
func NewServer(cfg *configv1.Service, opts ...config.ServiceSetting) *transhttp.Server {
	var options []transhttp.ServerOption

	option := settings.Apply(&config.ServiceConfig{}, opts)
	var ms []middleware.Middleware
	ms = middleware.NewServer(cfg.GetMiddleware())
	if option.Middlewares != nil {
		ms = append(ms, option.Middlewares...)
	}
	options = append(options, transhttp.Middleware(ms...))

	if serviceHttp := cfg.GetHttp(); serviceHttp != nil {
		if serviceHttp.Network != "" {
			options = append(options, transhttp.Network(serviceHttp.Network))
		}
		if serviceHttp.Addr != "" {
			options = append(options, transhttp.Address(serviceHttp.Addr))
		}
		if serviceHttp.Timeout != nil {
			options = append(options, transhttp.Timeout(serviceHttp.Timeout.AsDuration()))
		}
		if cfg.Endpoint {
			var endpoint *url.URL
			var err error

			// Obtain an endpoint using the custom EndpointURL function or the default service discovery method
			if option.EndpointURL != nil {
				endpoint, err = option.EndpointURL(serviceHttp.Endpoint, "http", cfg.Host, serviceHttp.Addr)
			} else {
				endpointStr := helpers.ServiceDiscoveryEndpoint(serviceHttp.Endpoint, "http", cfg.Host, serviceHttp.Addr)
				endpoint, err = url.Parse(endpointStr)
			}

			// If there are no errors, add an endpoint to options
			if err == nil {
				options = append(options, transhttp.Endpoint(endpoint))
			} else {
				// Record errors for easy debugging
				// log.Printf("Failed to get or parse endpoint: %v", err)
			}
		}
	}

	srv := transhttp.NewServer(options...)
	return srv
}
