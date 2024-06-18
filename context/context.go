// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package context provides the context functions
package context

import (
	"context"
	"log/slog"
	"time"
)

type Interface interface {
	Background() Context
	TODO() Context
	Canceled() error
	DeadlineExceeded() error
	WithCancel(ctx Context) (Context, context.CancelFunc)
	WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc)
	Cause(ctx Context) error
	AfterFunc(ctx Context, f func()) (stop func() bool)
	WithoutCancel(ctx Context) Context
	WithDeadline(ctx Context, t time.Time) (Context, context.CancelFunc)
	WithDeadlineCause(ctx Context, t time.Time, e error) (Context, context.CancelFunc)
	WithTimeout(ctx Context, t time.Duration) (Context, context.CancelFunc)
	WithTimeoutCause(ctx Context, t time.Duration, e error) (Context, context.CancelFunc)
	WithValue(ctx Context, key any, val any) Context
}

// All context types are defined in the context package.
type (
	Context         = context.Context
	CancelFunc      = context.CancelFunc
	CancelCauseFunc = context.CancelCauseFunc
)

// contextAlias holds the aliased context functions and constants as unexported variables.
type contextAlias struct {
	background        func() Context
	todo              func() Context
	canceled          error
	deadlineExceeded  error
	withCancel        func(Context) (Context, context.CancelFunc)
	withCancelCause   func(Context) (Context, CancelCauseFunc)
	cause             func(Context) error
	afterFunc         func(Context, func()) func() bool
	withoutCancel     func(Context) Context
	withDeadline      func(Context, time.Time) (Context, context.CancelFunc)
	withDeadlineCause func(Context, time.Time, error) (Context, context.CancelFunc)
	withTimeout       func(Context, time.Duration) (Context, context.CancelFunc)
	withTimeoutCause  func(Context, time.Duration, error) (Context, context.CancelFunc)
	withValue         func(Context, any, any) Context
}

func (c contextAlias) Background() Context {
	return c.background()
}

func (c contextAlias) TODO() Context {
	return c.todo()
}

func (c contextAlias) Canceled() error {
	return c.canceled
}

func (c contextAlias) DeadlineExceeded() error {
	return c.deadlineExceeded
}

func (c contextAlias) WithCancel(ctx Context) (Context, context.CancelFunc) {
	return c.withCancel(ctx)
}

func (c contextAlias) WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc) {
	return c.withCancelCause(parent)
}

func (c contextAlias) Cause(ctx Context) error {
	return c.cause(ctx)
}

func (c contextAlias) AfterFunc(ctx Context, f func()) (stop func() bool) {
	return c.afterFunc(ctx, f)
}

func (c contextAlias) WithoutCancel(ctx Context) Context {
	return c.withoutCancel(ctx)
}

func (c contextAlias) WithDeadline(ctx Context, t time.Time) (Context, context.CancelFunc) {
	return c.withDeadline(ctx, t)
}

func (c contextAlias) WithDeadlineCause(ctx Context, t time.Time, e error) (Context, context.CancelFunc) {
	return c.withDeadlineCause(ctx, t, e)
}

func (c contextAlias) WithTimeout(ctx Context, t time.Duration) (Context, context.CancelFunc) {
	return c.withTimeout(ctx, t)
}

func (c contextAlias) WithTimeoutCause(ctx Context, t time.Duration, e error) (Context, context.CancelFunc) {
	return c.withTimeoutCause(ctx, t, e)
}

func (c contextAlias) WithValue(ctx Context, key any, val any) Context {
	return c.withValue(ctx, key, val)
}

// alias holds the exported alias for context functions and constants.
var alias = contextAlias{
	background:        context.Background,
	todo:              context.TODO,
	canceled:          context.Canceled,
	deadlineExceeded:  context.DeadlineExceeded,
	withCancel:        context.WithCancel,
	withCancelCause:   context.WithCancelCause,
	cause:             context.Cause,
	afterFunc:         context.AfterFunc,
	withoutCancel:     context.WithoutCancel,
	withDeadline:      context.WithDeadline,
	withDeadlineCause: context.WithDeadlineCause,
	withTimeout:       context.WithTimeout,
	withTimeoutCause:  context.WithTimeoutCause,
	withValue:         context.WithValue,
}

// TODO returns an empty Context. It is never canceled, has no deadline, and has no values.
func TODO() Context {
	return alias.todo()
}

// Background returns a non-nil, empty Context. It is never canceled, has no deadline, and has no values.
// Background is typically used by the main function, initialization, and tests, and as the top-level
// Context for incoming requests.
func Background() Context {
	return alias.background()
}

// WithCancel returns a copy of the parent Context with a new Done channel. The returned
// context's Done channel is closed when the returned cancel function is called or when the parent
// context's Done channel is closed, whichever happens first.
func WithCancel(parent Context) (Context, context.CancelFunc) {
	return alias.withCancel(parent)
}

// WithCancelCause returns a copy of the parent Context with a new Done channel and an associated
// cause. The returned context's Done channel is closed when the returned cancel function is called,
// when the parent context's Done channel is closed, or when the provided cause is returned by the
// parent context's Err method, whichever happens first.
func WithCancelCause(parent Context) (Context, context.CancelCauseFunc) {
	return alias.withCancelCause(parent)
}

// Cause returns the underlying cause of the context's cancellation or nil if it was not caused
// by another error.
func Cause(ctx Context) error {
	return alias.cause(ctx)
}

// AfterFunc waits for the duration to elapse and then calls f in its own goroutine. If the context
// is canceled before the duration elapses, f will not be called.
func AfterFunc(ctx Context, f func()) func() bool {
	return alias.afterFunc(ctx, f)
}

// WithoutCancel returns a non-cancellable Context derived from parent. It is never canceled, has no
// deadline, and has the same values as the parent. If the parent is already non-cancellable,
// WithoutCancel returns the parent.
func WithoutCancel(parent Context) Context {
	return alias.withoutCancel(parent)
}

// WithDeadline returns a copy of the parent Context with a deadline adjusted to be no later than d.
// If the parent's deadline is already earlier than d, WithDeadline(parent, d) is semantically
// equivalent to parent. The returned context's Done channel is closed when the deadline expires,
// when the returned cancel function is called, or when the parent context's Done channel is closed,
// whichever happens first.
func WithDeadline(parent Context, deadline time.Time) (Context, context.CancelFunc) {
	return alias.withDeadline(parent, deadline)
}

// WithDeadlineCause returns a copy of the parent Context with a deadline adjusted to be no later than d
// and an associated cause. If the parent's deadline is already earlier than d, WithDeadlineCause(parent, d)
// is semantically equivalent to parent. The returned context's Done channel is closed when the deadline
// expires, when the returned cancel function is called, or when the provided cause is returned by the
// parent context's Err method, whichever happens first.
func WithDeadlineCause(parent Context, deadline time.Time, cause error) (Context, context.CancelFunc) {
	return alias.withDeadlineCause(parent, deadline, cause)
}

// WithTimeout returns a copy of the parent Context with a deadline adjusted to be no later than when
// the duration d elapses. It is equivalent to WithDeadline(parent, time.Now().Add(d)).
// The returned context's Done channel is closed when the timeout elapses, when the returned cancel
// function is called, or when the parent context's Done channel is closed, whichever happens first.
func WithTimeout(parent Context, timeout time.Duration) (Context, context.CancelFunc) {
	return alias.withTimeout(parent, timeout)
}

// WithTimeoutCause returns a copy of the parent Context with a deadline adjusted to be no later than
// when the duration d elapses and an associated cause. It is equivalent to WithDeadline(parent, time.Now().Add(d)).
// The returned context's Done channel is closed when the timeout elapses, when the returned cancel
// function is called, or when the provided cause is returned by the parent context's Err method,
// whichever happens first.
func WithTimeoutCause(parent Context, timeout time.Duration, cause error) (
	Context, context.CancelFunc,
) {
	return alias.withTimeoutCause(parent, timeout, cause)
}

// WithValue returns a copy of the parent Context that carries the given key-value pair. The value's
// type must be comparable for equality; otherwise, WithValue will panic. Use context Values only
// for request-scoped data that transits processes and API boundaries, not for passing optional parameters
// to functions.
func WithValue(parent Context, key, val any) Context {
	return alias.withValue(parent, key, val)
}

type traceIDCtx struct{}

// NewTraceID returns a new context with the provided traceID value.
//
// It takes a context and a traceID string as parameters and returns a context.
func NewTraceID(ctx Context, traceID string) Context {
	return context.WithValue(ctx, traceIDCtx{}, traceID)
}

// FromTraceID returns the trace ID from the context.
//
// It takes a Context as a parameter and returns a string.
func FromTraceID(ctx Context) string {
	v := ctx.Value(traceIDCtx{})
	if v != nil {
		return v.(string)
	}
	return ""
}

type transCtx struct{}

// NewTrans creates a new context with the provided dbx client value.
func NewTrans(ctx Context, db any) Context {
	return context.WithValue(ctx, transCtx{}, db)
}

// FromTrans retrieves a dbx client from the context.
func FromTrans(ctx Context) (any, bool) {
	v := ctx.Value(transCtx{})
	if v != nil {
		return v, true
	}
	return nil, false
}

type rowLockCtx struct{}

// NewRowLock creates a new context with a row lock value.
func NewRowLock(ctx Context) Context {
	return context.WithValue(ctx, rowLockCtx{}, true)
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
	return context.WithValue(ctx, userIDCtx{}, userID)
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
	return context.WithValue(ctx, userTokenCtx{}, userToken)
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
	return context.WithValue(ctx, isRootUserCtx{}, true)
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
	return context.WithValue(ctx, userCacheCtx{}, userCache)
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
	return context.WithValue(ctx, createdByCtx{}, by)
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
	return context.WithValue(ctx, dbCtx{}, db)
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
	return context.WithValue(ctx, loggerCtx{}, logger)
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
	return context.WithValue(ctx, tagCtx{}, tag)
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
	return context.WithValue(ctx, stackCtx{}, stack)
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

// Alias returns an Interface implementation that provides access to the context methods.
func Alias() Interface {
	return alias
}
