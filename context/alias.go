package context

import (
	"context"
	"time"
)

// All context types are defined in the context package.
type (
	Context         = context.Context
	CancelFunc      = context.CancelFunc
	CancelCauseFunc = context.CancelCauseFunc
)

type Interface interface {
	Background() Context
	TODO() Context
	Canceled() error
	DeadlineExceeded() error
	WithCancel(ctx Context) (Context, CancelFunc)
	WithCancelCause(parent Context) (ctx Context, cancel CancelCauseFunc)
	Cause(ctx Context) error
	AfterFunc(ctx Context, f func()) (stop func() bool)
	WithoutCancel(ctx Context) Context
	WithDeadline(ctx Context, t time.Time) (Context, CancelFunc)
	WithDeadlineCause(ctx Context, t time.Time, e error) (Context, CancelFunc)
	WithTimeout(ctx Context, t time.Duration) (Context, CancelFunc)
	WithTimeoutCause(ctx Context, t time.Duration, e error) (Context, CancelFunc)
	WithValue(ctx Context, key any, val any) Context
}

// contextAlias holds the aliased context functions and constants as unexported variables.
type contextAlias struct {
	background        func() Context
	todo              func() Context
	canceled          error
	deadlineExceeded  error
	withCancel        func(Context) (Context, CancelFunc)
	withCancelCause   func(Context) (Context, CancelCauseFunc)
	cause             func(Context) error
	afterFunc         func(Context, func()) func() bool
	withoutCancel     func(Context) Context
	withDeadline      func(Context, time.Time) (Context, CancelFunc)
	withDeadlineCause func(Context, time.Time, error) (Context, CancelFunc)
	withTimeout       func(Context, time.Duration) (Context, CancelFunc)
	withTimeoutCause  func(Context, time.Duration, error) (Context, CancelFunc)
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

func (c contextAlias) WithCancel(ctx Context) (Context, CancelFunc) {
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

func (c contextAlias) WithDeadline(ctx Context, t time.Time) (Context, CancelFunc) {
	return c.withDeadline(ctx, t)
}

func (c contextAlias) WithDeadlineCause(ctx Context, t time.Time, e error) (Context, CancelFunc) {
	return c.withDeadlineCause(ctx, t, e)
}

func (c contextAlias) WithTimeout(ctx Context, t time.Duration) (Context, CancelFunc) {
	return c.withTimeout(ctx, t)
}

func (c contextAlias) WithTimeoutCause(ctx Context, t time.Duration, e error) (Context, CancelFunc) {
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
func WithCancel(parent Context) (Context, CancelFunc) {
	return alias.withCancel(parent)
}

// WithCancelCause returns a copy of the parent Context with a new Done channel and an associated
// cause. The returned context's Done channel is closed when the returned cancel function is called,
// when the parent context's Done channel is closed, or when the provided cause is returned by the
// parent context's Err method, whichever happens first.
func WithCancelCause(parent Context) (Context, CancelCauseFunc) {
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
func WithDeadline(parent Context, deadline time.Time) (Context, CancelFunc) {
	return alias.withDeadline(parent, deadline)
}

// WithDeadlineCause returns a copy of the parent Context with a deadline adjusted to be no later than d
// and an associated cause. If the parent's deadline is already earlier than d, WithDeadlineCause(parent, d)
// is semantically equivalent to parent. The returned context's Done channel is closed when the deadline
// expires, when the returned cancel function is called, or when the provided cause is returned by the
// parent context's Err method, whichever happens first.
func WithDeadlineCause(parent Context, deadline time.Time, cause error) (Context, CancelFunc) {
	return alias.withDeadlineCause(parent, deadline, cause)
}

// WithTimeout returns a copy of the parent Context with a deadline adjusted to be no later than when
// the duration d elapses. It is equivalent to WithDeadline(parent, time.Now().Add(d)).
// The returned context's Done channel is closed when the timeout elapses, when the returned cancel
// function is called, or when the parent context's Done channel is closed, whichever happens first.
func WithTimeout(parent Context, timeout time.Duration) (Context, CancelFunc) {
	return alias.withTimeout(parent, timeout)
}

// WithTimeoutCause returns a copy of the parent Context with a deadline adjusted to be no later than
// when the duration d elapses and an associated cause. It is equivalent to WithDeadline(parent, time.Now().Add(d)).
// The returned context's Done channel is closed when the timeout elapses, when the returned cancel
// function is called, or when the provided cause is returned by the parent context's Err method,
// whichever happens first.
func WithTimeoutCause(parent Context, timeout time.Duration, cause error) (
	Context, CancelFunc,
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

// Alias returns an Interface implementation that provides access to the context methods.
func Alias() Interface {
	return alias
}
