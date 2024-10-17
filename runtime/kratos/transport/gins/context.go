package gins

import (
	"github.com/gin-gonic/gin"

	"github.com/origadmin/toolkits/context"
)

type ginCtx struct{}

// NewContext creates a new context from a gin context
func NewContext(ctx *gin.Context) context.Context {
	// Create a new context with the gin context as a value
	return context.WithValue(ctx.Request.Context(), ginCtx{}, ctx)
}

// FromContext takes a context as an argument and returns a pointer to a gin.Context and a boolean value.
func FromContext(ctx context.Context) (*gin.Context, bool) {
	// Check if the context contains a value of type *gin.Context
	if v, ok := ctx.Value(ginCtx{}).(*gin.Context); ok {
		// If the context contains a value of type *gin.Context, return the value and true
		return v, true
	}
	// If the context does not contain a value of type *gin.Context, return nil and false
	return nil, false
}
