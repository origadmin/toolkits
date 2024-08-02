// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package context provides the context functions
package context

import (
	"context"
	"log/slog"
	"reflect"
)

// WithContext returns a new context with the provided context.Context value.
func WithContext(ctx context.Context) Context {
	if ctx == nil {
		return Background()
	}
	return ctx
}

type traceIDCtx struct{}

// NewTraceID returns a new context with the provided traceID value.
//
// It takes a context and a traceID string as parameters and returns a context.
func NewTraceID(ctx Context, traceID string) Context {
	return WithMapValue(ctx, traceIDCtx{}, traceID)
}

// FromTraceID returns the trace ID from the context.
//
// It takes a Context as a parameter and returns a string.
func FromTraceID(ctx Context) string {
	if v := Value(ctx, traceIDCtx{}); v != nil {
		return v.(string)
	}
	return ""
}

type transCtx struct{}

// NewTrans creates a new context with the provided dbx client value.
func NewTrans(ctx Context, db any) Context {
	return WithMapValue(ctx, transCtx{}, db)
}

// FromTrans retrieves a dbx client from the context.
func FromTrans(ctx Context) (any, bool) {
	if v := ctx.Value(transCtx{}); v != nil {
		return v, true
	}
	return nil, false
}

type rowLockCtx struct{}

// NewRowLock creates a new context with a row lock value.
func NewRowLock(ctx Context) Context {
	return WithValue(ctx, rowLockCtx{}, true)
}

// FromRowLock checks if the row is locked in the given context.
//
// It takes a Context as a parameter and returns a boolean.
func FromRowLock(ctx Context) bool {
	v := ctx.Value(rowLockCtx{})
	return v != nil && v.(bool)
}

type userIDCtx struct{}

// NewUserID returns a new context with the provided userID value.
//
// It takes a context and a userID string as parameters and returns a context.
func NewUserID(ctx Context, userID string) Context {
	return WithValue(ctx, userIDCtx{}, userID)
}

// FromUserID returns the user ID from the context.
//
// It takes a Context as a parameter and returns a string.
func FromUserID(ctx Context) string {
	v := ctx.Value(userIDCtx{})
	if v != nil {
		return v.(string)
	}
	return ""
}

type userTokenCtx struct{}

// NewUserToken returns a new context with the provided userToken value.
//
// It takes a context and a userToken string as parameters and returns a context.
func NewUserToken(ctx Context, userToken string) Context {
	return WithValue(ctx, userTokenCtx{}, userToken)
}

// FromUserToken returns the user token from the context.
//
// It takes a Context as a parameter and returns a string.
func FromUserToken(ctx Context) string {
	v := ctx.Value(userTokenCtx{})
	if v != nil {
		return v.(string)
	}
	return ""
}

type isRootUserCtx struct{}

// NewIsRootUser returns a new context with the provided isRootUser value.
//
// It takes a Context as a parameter and returns a context.
func NewIsRootUser(ctx Context) Context {
	return WithValue(ctx, isRootUserCtx{}, true)
}

// FromIsRootUser returns the isRootUser from the context.
//
// It takes a Context as a parameter and returns a boolean.
func FromIsRootUser(ctx Context) bool {
	v := ctx.Value(isRootUserCtx{})
	return v != nil && v.(bool)
}

type userCacheCtx struct{}

// NewUserCache returns a new context with the provided userCache value.
//
// It takes a Context and a userCache value as parameters and returns a context.
func NewUserCache(ctx Context, userCache any) Context {
	return WithValue(ctx, userCacheCtx{}, userCache)
}

// FromUserCache returns the userCache from the context.
//
// It takes a Context as a parameter and returns a userCache value.
func FromUserCache(ctx Context) (any, bool) {
	v := ctx.Value(userCacheCtx{})
	if v != nil {
		return v, true
	}
	return nil, false
}

type createdByCtx struct{}

// NewCreatedBy creates a new context with the provided 'by' value
//
// It takes a Context and a 'by' string as parameters and returns a context.
func NewCreatedBy(ctx Context, by string) Context {
	return WithValue(ctx, createdByCtx{}, by)
}

// FromCreatedBy retrieves the creator information from the given context.
//
// It takes a Context as a parameter and returns a string.
func FromCreatedBy(ctx Context) string {
	// Attempt to retrieve the creator information from the context.
	v := ctx.Value(createdByCtx{})
	if v != nil {
		// If found, return the creator information.
		return v.(string)
	}
	// If not found, return an empty string.
	return ""
}

type dbCtx struct{}

// NewDB creates a new context with the provided db client value.
//
// It takes a context and a db client as parameters and returns a context.
func NewDB(ctx Context, db any) Context {
	return WithValue(ctx, dbCtx{}, db)
}

// FromDB retrieves a db client from the context.
//
// It takes a Context as a parameter and returns a db client.
func FromDB(ctx Context) (any, bool) {
	v := ctx.Value(dbCtx{})
	if v != nil {
		return v, true
	}
	return nil, false
}

type loggerCtx struct{}

// NewLogger creates a new context with the provided logger value.
//
// It takes a Context and a logger as parameters and returns a context.
func NewLogger(ctx Context, logger *slog.Logger) Context {
	return WithValue(ctx, loggerCtx{}, logger)
}

// FromLogger retrieves a logger from the context.
//
// It takes a Context as a parameter and returns a logger.
func FromLogger(ctx Context) (*slog.Logger, bool) {
	v := ctx.Value(loggerCtx{})
	if v != nil {
		return v.(*slog.Logger), true
	}
	return nil, false
}

type tagCtx struct{}

// NewTag creates a new context with the provided tag value.
//
// It takes a Context and a tag string as parameters and returns a context.
func NewTag(ctx Context, tag string) Context {
	return WithValue(ctx, tagCtx{}, tag)
}

// FromTag retrieves the tag from the context.
//
// It takes a Context as a parameter and returns a string.
func FromTag(ctx Context) string {
	v := ctx.Value(tagCtx{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

type stackCtx struct{}

// NewStack creates a new context with the provided stack value.
//
// It takes a Context and a stack string as parameters and returns a context.
func NewStack(ctx Context, stack string) Context {
	return WithValue(ctx, stackCtx{}, stack)
}

// FromStack retrieves the stack from the context.
//
// It takes a Context as a parameter and returns a string.
func FromStack(ctx Context) string {
	v := ctx.Value(stackCtx{})
	if v != nil {
		if s, ok := v.(string); ok {
			return s
		}
	}
	return ""
}

// A mapValueCtx carries a key-value pair. It implements Value for that key and
// delegates all other calls to the embedded Context.
type mapValueCtx struct {
	Context
	keyValues map[any]any
}

// Value returns the value for the given key or nil if no value is present.
func (ctx *mapValueCtx) Value(key any) any {
	if val, ok := ctx.keyValues[key]; ok {
		return val
	}
	return Value(ctx.Context, key)
}

// WithMapValue creates a new context with the provided key-value pair.
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
	ctx := parent
	for ctx != nil {
		if mctx, ok := ctx.(*mapValueCtx); ok {
			mctx.keyValues[key] = val
			return parent
		}
		ctx = Parent(ctx)
	}

	return &mapValueCtx{parent, map[any]any{key: val}}
}

// FromMapContext retrieves all values from the context.
func FromMapContext(parent Context) map[any]any {
	ctx := parent
	for ctx != nil {
		if mctx, ok := ctx.(*mapValueCtx); ok {
			return mctx.keyValues
		}
		ctx = Parent(ctx)
	}
	return map[any]any{}
}

// Value retrieves the value for the given key or nil if no value is present.
func Value(ctx Context, key any) any {
	return ctx.Value(key)
}

// Parent retrieves the parent context.
func Parent(ptr Context) Context {
	if ptr == nil {
		return nil
	}
	kind := reflect.TypeOf(ptr).Kind()
	v := reflect.ValueOf(ptr)
	if kind == reflect.Ptr {
		v = v.Elem()
	}
	var field reflect.Value
	for i := 0; i < v.NumField(); i++ {
		field = v.Field(i)
		if field.Type().Name() == "Context" {
			return field.Interface().(Context)
		}
	}
	return nil
}
