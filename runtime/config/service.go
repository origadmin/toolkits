// Package config implements the functions, types, and interfaces for the module.
package config

import (
	"net/url"

	"github.com/origadmin/toolkits/runtime/middleware"
	"github.com/origadmin/toolkits/runtime/registry"
)

type (
	EndpointURLFunc = func(endpoint string, scheme string, host string, addr string) (*url.URL, error)
)

type ServiceOption = func(*ServiceConfig)

type ServiceConfig struct {
	Discovery   registry.Discovery
	Middlewares []middleware.Middleware
	EndpointURL func(endpoint string, scheme string, host string, addr string) (*url.URL, error)
}

func WithDiscovery(discovery registry.Discovery) func(*ServiceConfig) {
	return func(config *ServiceConfig) {
		config.Discovery = discovery
	}
}

func WithMiddlewares(middlewares ...middleware.Middleware) func(*ServiceConfig) {
	return func(config *ServiceConfig) {
		config.Middlewares = middlewares
	}
}

func WithEndpointURL(endpoint EndpointURLFunc) func(*ServiceConfig) {
	return func(config *ServiceConfig) {
		config.EndpointURL = endpoint
	}
}
