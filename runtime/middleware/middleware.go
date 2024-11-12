/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package middlewares implements the functions, types, and interfaces for the module.
package middleware

import (
	"github.com/go-kratos/kratos/v2/middleware"
)

type (
	Handler    = middleware.Handler
	Middleware = middleware.Middleware
)

// Chain returns a middleware that executes a chain of middleware.
func Chain(m ...middleware.Middleware) middleware.Middleware {
	return middleware.Chain(m...)
}
