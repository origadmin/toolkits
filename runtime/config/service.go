// Package config implements the functions, types, and interfaces for the module.
package config

import (
	"net/url"

	"github.com/origadmin/toolkits/runtime/middleware"
)

type (
	EndpointURLFunc = func(endpoint string, scheme string, host string, addr string) (*url.URL, error)
)

type ServiceOption = func(*ServiceConfig)

type ServiceConfig struct {
	Middlewares []middleware.Middleware
	EndpointURL func(endpoint string, scheme string, host string, addr string) (*url.URL, error)
}

func WithEndpointURL(endpoint EndpointURLFunc) func(*ServiceConfig) {
	return func(config *ServiceConfig) {
		config.EndpointURL = endpoint
	}
}

func WithMiddlewares(middlewares ...middleware.Middleware) func(*ServiceConfig) {
	return func(config *ServiceConfig) {
		config.Middlewares = middlewares
	}
}
