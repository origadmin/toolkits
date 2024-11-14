// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package middleware implements the functions, types, and interfaces for the module.
package middleware

import (
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"

	"github.com/origadmin/toolkits/runtime/config"
)

func RateLimitServer(ms []Middleware, cfg *config.Middleware_RateLimiter) []Middleware {
	if cfg == nil {
		return ms
	}
	var options []ratelimit.Option
	switch cfg.GetName() {
	case "redis":
		// TODO:
	case "memory":
		// TODO:
	//case "bbr":
	// default is bbr
	// options = append(options, middlewareRateLimit.WithLimiter(bbr.NewLimiter()))
	default:
		// do nothing
	}
	return append(ms, ratelimit.Server(options...))
}
