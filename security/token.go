/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

import (
	"context"
)

type tokenCtx struct{}

func NewToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenCtx{}, token)
}

func FromToken(ctx context.Context) string {
	if token, ok := ctx.Value(tokenCtx{}).(string); ok {
		return token
	}
	return ""
}
