/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

import (
	"context"
)

type tokenCtx struct{}

func NewTokenContext(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, tokenCtx{}, token)
}

func TokenFromContext(ctx context.Context) string {
	if token, ok := ctx.Value(tokenCtx{}).(string); ok {
		return token
	}
	return ""
}

type claimCtx struct{}

func NewClaimsContext(ctx context.Context, claims Claims) context.Context {
	return context.WithValue(ctx, claimCtx{}, claims)
}

func ClaimsFromContext(ctx context.Context) Claims {
	if claims, ok := ctx.Value(claimCtx{}).(Claims); ok {
		return claims
	}
	return nil
}

type userClaimsCtx struct{}

func NewUserClaimsContext(ctx context.Context, claims UserClaims) context.Context {
	return context.WithValue(ctx, userClaimsCtx{}, claims)
}

func UserClaimsFromContext(ctx context.Context) UserClaims {
	if claims, ok := ctx.Value(userClaimsCtx{}).(UserClaims); ok {
		return claims
	}
	return nil
}
