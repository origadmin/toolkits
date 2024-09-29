// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package ginctx provides the context functions
package ginctx

import (
	"github.com/gin-gonic/gin"

	"github.com/origadmin/toolkits/context"
)

type ginContext struct{}

// WithContext returns a new context with the provided context.Context value.
func WithContext(ginCtx *gin.Context) context.Context {
	return context.WithValue(ginCtx.Request.Context(), ginContext{}, ginCtx)
}

// FromContext returns the gin.Context from the context.
//
// It takes a Context as a parameter and returns a gin.Context.
func FromContext(ctx context.Context) (*gin.Context, bool) {
	if v, ok := ctx.Value(ginContext{}).(*gin.Context); ok {
		return v, true
	}
	return nil, false
}
