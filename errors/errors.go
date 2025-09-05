/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package errors provides enhanced error handling utilities for Go applications.
// It offers:
// - Thread-safe multi-error collection
// - Error chain traversal and inspection
// - Type-safe error assertions
// - Contextual error wrapping
// - Standard error interface compatibility
//
// The package is designed to work seamlessly with standard library errors
// while providing additional functionality for complex error handling scenarios.
package errors

import (
	"context"
	"errors"
	"fmt"
	"reflect"

	perr "github.com/pkg/errors"
)

//go:generate adptool .
//go:adapter:package github.com/pkg/errors perr
//go:adapter:package:func Is
//go:adapter:package:func:rename IsPkgError
//go:adapter:package:func Unwrap
//go:adapter:package:func:rename UnwrapPkgError
//go:adapter:package:func As
//go:adapter:package:func:rename AsPkgError
//go:adapter:package github.com/hashicorp/go-multierror merr
//go:adapter:package:type *
//go:adapter:package:type:prefix Multi
//go:adapter:package:func *
//go:adapter:package:func:prefix Multi
//go:adapter:package errors stderr
//go:adapter:package:func New
//go:adapter:package:func:rename NewStdError

// ErrorChain represents a chain of errors, where each error in the chain
// can be unwrapped to get the next error. This is compatible with Go 1.13+
// error wrapping.
type ErrorChain interface {
	// Unwrap returns the next error in the chain or nil if there are no more errors.
	Unwrap() error
	// Error returns the string representation of the error.
	Error() string
}

// ErrorWithCode is an error that includes an error code.
type ErrorWithCode interface {
	error
	// Code returns the error code.
	Code() int
}

// ErrorWithStack is an error that includes a stack trace.
type ErrorWithStack interface {
	error
	// StackTrace returns the stack trace associated with the error.
	StackTrace() perr.StackTrace
}

// contextualError is an error that carries a context.Context
type contextualError struct {
	err     error
	context context.Context
}

// Error implements the error interface
func (e *contextualError) Error() string { return e.err.Error() }

// Unwrap implements the error unwrapping interface
func (e *contextualError) Unwrap() error { return e.err }

// Context returns the context associated with the error
func (e *contextualError) Context() context.Context { return e.context }

// WalkFunc is the type of the function called for each error in the chain.
// If the function returns a non-nil error, Walk will stop and return that error.
type WalkFunc func(error) error

// Walk traverses the error chain and calls fn for each error in the chain.
// If fn returns an error, Walk stops and returns that error.
//
// Example:
//
//	err := fmt.Errorf("root error")
//	err = fmt.Errorf("wrapper: %w", err)
//
//	err = Walk(err, func(e error) error {
//	    fmt.Println(e)
//	    return nil // Continue walking
//	})
func Walk(err error, fn WalkFunc) error {
	for err != nil {
		if err = fn(err); err != nil {
			return err
		}

		// Check for standard unwrapping
		if unwrapper, ok := err.(interface{ Unwrap() error }); ok {
			err = unwrapper.Unwrap()
			continue
		}

		// Check for multi-error unwrapping
		if multi, ok := err.(interface{ Errors() []error }); ok {
			for _, e := range multi.Errors() {
				if walkErr := Walk(e, fn); walkErr != nil {
					return walkErr
				}
			}
		}

		// No more errors to unwrap
		break
	}

	return nil
}

// Find traverses the error chain and returns the first error for which the
// provided function returns true.
//
// Example:
//
//	// Find the first error of a specific type
//	var target *MyError
//	if found := Find(err, func(e error) bool {
//	    return As(e, &target)
//	}); found != nil {
//	    // Handle the found error
//	}
func Find(err error, fn func(error) bool) error {
	var result error

	Walk(err, func(e error) error {
		if fn(e) {
			result = e
			return fmt.Errorf("stop") // Signal to stop walking
		}
		return nil
	})

	return result
}

// Has checks if any error in the chain matches the target error using errors.Is.
// It's similar to errors.Is but works with error chains.
//
// Example:
//
//	if Has(err, io.EOF) {
//	    // Handle EOF error
//	}
func Has(err, target error) bool {
	return Find(err, func(e error) bool {
		return errors.Is(e, target)
	}) != nil
}

// WithContext adds a context.Context to an error.
// If the error is nil or context is nil, the original error is returned unchanged.
//
// Example:
//
//	ctx := context.WithValue(context.Background(), "request_id", "123")
//	err := errors.New("operation failed")
//	err = WithContext(ctx, err)
func WithContext(ctx context.Context, err error) error {
	if err == nil || ctx == nil {
		return err
	}
	return &contextualError{
		err:     err,
		context: ctx,
	}
}

// ContextFrom retrieves the context.Context from an error if it exists.
// Returns the context and true if found, nil and false otherwise.
//
// Example:
//
//	if ctx, ok := ContextFrom(err); ok {
//	    // Use ctx
//	}
func ContextFrom(err error) (context.Context, bool) {
	if e, ok := err.(*contextualError); ok {
		return e.context, true
	}
	return nil, false
}

// Value retrieves a value from the error's context.
// It traverses the error chain to find a context that contains the key.
//
// Example:
//
//	if val := Value(err, "request_id"); val != nil {
//	    // Use val
//	}
func Value(err error, key interface{}) interface{} {
	var result interface{}
	Walk(err, func(e error) error {
		if e, ok := e.(*contextualError); ok {
			if val := e.context.Value(key); val != nil {
				result = val
				return fmt.Errorf("stop") // Stop walking after first match
			}
		}
		return nil
	})
	return result
}

// AssertType checks if the error is of the specified type and returns it.
// It works with both value and pointer types.
//
// Example:
//
//	// For non-pointer types
//	if e, ok := AssertType[MyError](err); ok {
//	    // Use e which is of type MyError
//	}
//
//	// For pointer types
//	if e, ok := AssertType[*MyError](err); ok {
//	    // Use e which is of type *MyError
//	}
func AssertType[T error](err error) (T, bool) {
	var zero T

	// Handle nil error
	if err == nil {
		return zero, false
	}

	// Check if the error is directly of type T
	if e, ok := err.(T); ok {
		return e, true
	}

	// Handle pointer vs value type comparison
	if reflect.TypeOf(zero) != nil && reflect.TypeOf(err).AssignableTo(reflect.TypeOf(zero)) {
		return reflect.ValueOf(err).Convert(reflect.TypeOf(zero)).Interface().(T), true
	}

	// Check if the error is in the chain
	var target T
	if errors.As(err, &target) {
		return target, true
	}

	return zero, false
}

// MustAssertType is like AssertType but panics if the error is not of the specified type.
// Use with caution, only when you're certain about the error type.
//
// Example:
//
//	// Will panic if err is not of type *MyError
//	e := MustAssertType[*MyError](err)
func MustAssertType[T error](err error) T {
	e, ok := AssertType[T](err)
	if !ok {
		panic(fmt.Sprintf("expected error of type %T, got %T", *new(T), err))
	}
	return e
}

// IsType checks if the error is of the specified type.
// It's a type-safe alternative to errors.As.
//
// Example:
//
//	if IsType[MyError](err) {
//	    // Handle MyError
//	}
func IsType[T error](err error) bool {
	_, ok := AssertType[T](err)
	return ok
}

// Helper function to create a new error with context
//
// Example:
//
//	err := NewWithContext(context.Background(), "operation failed")
//	// Add more context
//	err = fmt.Errorf("additional context: %w", err)
func NewWithContext(ctx context.Context, text string) error {
	if ctx == nil {
		return errors.New(text)
	}
	return &contextualError{
		err:     errors.New(text),
		context: ctx,
	}
}
