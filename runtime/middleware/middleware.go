/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package middlewares implements the functions, types, and interfaces for the module.
package middleware

import (
	"github.com/go-kratos/kratos/v2/middleware"

	"github.com/origadmin/toolkits/runtime/config"
)

type (
	Handler    = middleware.Handler
	Middleware = middleware.Middleware
)

// Chain returns a middleware that executes a chain of middleware.
func Chain(m ...Middleware) Middleware {
	return middleware.Chain(m...)
}

func NewClient(cfg *config.Middleware) []Middleware {
	var middlewares []Middleware

	if cfg == nil {
		return middlewares
	}
	middlewares = Recovery(middlewares, cfg.EnableRecovery)
	middlewares = Validate(middlewares, cfg.EnableValidate)
	middlewares = TracingClient(middlewares, cfg.EnableTracing)
	middlewares = MetadataClient(middlewares, cfg.EnableMetadata, cfg.Metadata)
	middlewares = CircuitBreakerClient(middlewares, cfg.EnableCircuitBreaker)
	return middlewares
}

func NewServer(cfg *config.Middleware) []Middleware {
	var middlewares []Middleware

	if cfg == nil {
		return middlewares
	}
	middlewares = Recovery(middlewares, cfg.EnableRecovery)
	middlewares = Validate(middlewares, cfg.EnableValidate)
	middlewares = TracingServer(middlewares, cfg.EnableTracing)
	middlewares = MetadataServer(middlewares, cfg.EnableMetadata, cfg.Metadata)
	middlewares = RateLimitServer(middlewares, cfg.RateLimiter)
	return middlewares
}
