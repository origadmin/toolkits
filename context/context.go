/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package context provides the context functions
package context

import (
	"reflect"
)

// WithContext returns a new context with the provided context.Context value.
// Deprecated: Use runtime.context instead.
func WithContext(ctx Context) Context {
	if ctx == nil {
		return Background()
	}
	return ctx
}

type traceCtx struct{}

// NewTrace returns a new context with the provided trace value.
//
// It takes a context and a trace string as parameters and returns a context.
// Deprecated: Use runtime.context instead.
func NewTrace(ctx Context, trace string) Context {
	return WithValue(ctx, traceCtx{}, trace)
}

// FromTrace returns the trace id from the context.
//
// It takes a Context as a parameter and returns a string.
// Deprecated: Use runtime.context instead.
func FromTrace(ctx Context) string {
	if v, ok := Value(ctx, traceCtx{}).(string); ok {
		return v
	}
	return ""
}

type spanCtx struct{}

// NewSpan creates a new context with the provided span value.
//
// It takes a context and a span string as parameters and returns a context.
// Deprecated: Use runtime.context instead.
func NewSpan(ctx Context, span string) Context {
	return WithValue(ctx, spanCtx{}, span)
}

// FromSpan returns the span id from the context.
//
// It takes a Context as a parameter and returns a string.
// Deprecated: Use runtime.context instead.
func FromSpan(ctx Context) string {
	if v, ok := ctx.Value(spanCtx{}).(string); ok {
		return v
	}
	return ""
}

type dbCtx struct{}

// NewDB creates a new context with the provided db client value.
//
// It takes a context and a db client as parameters and returns a context.
// Deprecated: Use runtime.context instead.
func NewDB(ctx Context, db any) Context {
	return WithValue(ctx, dbCtx{}, db)
}

// FromDB retrieves a db client from the context.
//
// It takes a Context as a parameter and returns a db client.
// Deprecated: Use runtime.context instead.
func FromDB(ctx Context) (any, bool) {
	if v := ctx.Value(dbCtx{}); v != nil {
		return v, true
	}
	return nil, false
}

type transCtx struct{}

// NewTrans creates a new context with the provided tx client value.
// Deprecated: Use runtime.context instead.
func NewTrans(ctx Context, tx any) Context {
	return WithValue(ctx, transCtx{}, tx)
}

// FromTrans retrieves a tx client from the context.
// Deprecated: Use runtime.context instead.
func FromTrans(ctx Context) (any, bool) {
	if v := ctx.Value(transCtx{}); v != nil {
		return v, true
	}
	return nil, false
}

type rowLockCtx struct{}

// NewRowLock creates a new context with a row lock value.
// Deprecated: Use runtime.context instead.
func NewRowLock(ctx Context) Context {
	return WithValue(ctx, rowLockCtx{}, true)
}

// FromRowLock checks if the row is locked in the given context.
//
// It takes a Context as a parameter and returns a boolean.
// Deprecated: Use runtime.context instead.
func FromRowLock(ctx Context) bool {
	v, ok := ctx.Value(rowLockCtx{}).(bool)
	return ok && v
}

type idCtx struct{}

// NewID returns a new context with the provided userID value.
//
// It takes a context and a userID string as parameters and returns a context.
// Deprecated: Use runtime.context instead.
func NewID(ctx Context, id string) Context {
	return WithValue(ctx, idCtx{}, id)
}

// FromID returns the user ID from the context.
//
// It takes a Context as a parameter and returns a string.
// Deprecated: Use runtime.context instead.
func FromID(ctx Context) string {
	if v, ok := ctx.Value(idCtx{}).(string); ok {
		return v
	}
	return ""
}

type tokenCtx struct{}

// NewToken returns a new context with the provided userToken value.
//
// It takes a context and a userToken string as parameters and returns a context.
// Deprecated: Use runtime.context instead.
func NewToken(ctx Context, token string) Context {
	return WithValue(ctx, tokenCtx{}, token)
}

// FromToken returns the user token from the context.
//
// It takes a Context as a parameter and returns a string.
// Deprecated: Use runtime.context instead.
func FromToken(ctx Context) string {
	if v, ok := ctx.Value(tokenCtx{}).(string); ok {
		return v
	}
	return ""
}

type userCacheCtx struct{}

// NewUserCache returns a new context with the provided userCache value.
//
// It takes a Context and a userCache value as parameters and returns a context.
// Deprecated: Use runtime.context instead.
func NewUserCache(ctx Context, userCache any) Context {
	return WithValue(ctx, userCacheCtx{}, userCache)
}

// FromUserCache returns the userCache from the context.
//
// It takes a Context as a parameter and returns a userCache value.
// Deprecated: Use runtime.context instead.
func FromUserCache(ctx Context) (any, bool) {
	if v := ctx.Value(userCacheCtx{}); v != nil {
		return v, true
	}
	return nil, false
}

type createdByCtx struct{}

// NewCreatedBy creates a new context with the provided 'by' value
//
// It takes a Context and a 'by' string as parameters and returns a context.
// Deprecated: Use runtime.context instead.
func NewCreatedBy(ctx Context, by string) Context {
	return WithValue(ctx, createdByCtx{}, by)
}

// FromCreatedBy retrieves the creator information from the given context.
//
// It takes a Context as a parameter and returns a string.
// Deprecated: Use runtime.context instead.
func FromCreatedBy(ctx Context) string {
	// Attempt to retrieve the creator information from the context.
	if v, ok := ctx.Value(createdByCtx{}).(string); ok {
		// If found, return the creator information.
		return v
	}
	// If not found, return an empty string.
	return ""
}

type tagCtx struct{}

// NewTag creates a new context with the provided tag value.
//
// It takes a Context and a tag string as parameters and returns a context.
// Deprecated: Use runtime.context instead.
func NewTag(ctx Context, tag string) Context {
	return WithValue(ctx, tagCtx{}, tag)
}

// FromTag retrieves the tag from the context.
//
// It takes a Context as a parameter and returns a string.
// Deprecated: Use runtime.context instead.
func FromTag(ctx Context) string {
	if v, ok := ctx.Value(tagCtx{}).(string); ok {
		return v
	}
	return ""
}

type stackCtx struct{}

// NewStack creates a new context with the provided stack value.
//
// It takes a Context and a stack string as parameters and returns a context.
// Deprecated: Use runtime.context instead.
func NewStack(ctx Context, stack string) Context {
	return WithValue(ctx, stackCtx{}, stack)
}

// FromStack retrieves the stack from the context.
//
// It takes a Context as a parameter and returns a string.
// Deprecated: Use runtime.context instead.
func FromStack(ctx Context) string {
	if v, ok := ctx.Value(stackCtx{}).(string); ok {
		return v
	}
	return ""
}

type mapCtx struct{}

// A mapValueCtx carries a key-value pair. It implements Value for that key and
// delegates all other calls to the embedded Context.
// Deprecated: Use runtime.context instead.
type mapValueCtx struct {
	Context
	keyValues map[any]any
}

// Value returns the value for the given key or nil if no value is present.
// Deprecated: Use runtime.context instead.
func (ctx *mapValueCtx) Value(key any) any {
	if any(mapCtx{}) == key {
		return ctx.keyValues
	}
	if val, ok := ctx.keyValues[key]; ok {
		return val
	}
	return Value(ctx.Context, key)
}

// WithMapValue creates a new context with the provided key-value pair.
// If the context saved over than 500 keys, use WithMapValue instead.
// Deprecated: Use runtime.context instead.
func WithMapValue(parent Context, key, val any) Context {
	if parent == nil {
		panic("cannot create context from nil parent")
	}
	if key == nil {
		panic("nil key")
	}
	if !reflect.TypeOf(key).Comparable() {
		panic("key is not comparable")
	}

	if v := parent.Value(mapCtx{}); v != nil {
		if kv, ok := v.(map[any]any); ok {
			kv[key] = val
			return parent
		}
	}
	return &mapValueCtx{parent, map[any]any{key: val}}
}

// FromMapContext retrieves all values from the context.
// Deprecated: Use runtime.context instead.
func FromMapContext(parent Context) map[any]any {
	if v := parent.Value(mapCtx{}); v != nil {
		if kv, ok := v.(map[any]any); ok {
			return kv
		}
	}
	return map[any]any{}
}

// Value retrieves the value for the given key or nil if no value is present.
// Deprecated: Use runtime.context instead.
func Value(ctx Context, key any) any {
	return ctx.Value(key)
}
