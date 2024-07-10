// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package errors provides a means to return detailed information
// for a request error, typically encoded in JSON format.
// This package wraps standard library error handling features and
// offers additional stack trace capabilities.
package errors

import (
	stderr "errors" // Import the standard library's error handling package
	"fmt"           // For formatted strings

	"github.com/pkg/errors" // Import third-party error handling package for richer error handling features
)

/**
 * errorAliases is a structure that holds aliased error functions as unexported variables.
 * It serves as a central point for accessing various error handling methods from both
 * the standard library and the 'pkg/errors' package.
 */
type errorAliases struct {
	// pkg/errors aliases
	stack     func(string) error                        // Creates an error with a stack trace
	stackf    func(string, ...interface{}) error        // Formats and creates an error with a stack trace
	wrap      func(error, string) error                 // Wraps an existing error with a new message
	wrapf     func(error, string, ...interface{}) error // Formats and wraps an existing error with a new message
	withStack func(error) error                         // Attaches a stack trace to an existing error
	cause     func(error) error                         // Returns the root cause of the error

	// stdlib errors aliases
	errorf func(string, ...interface{}) error // Creates a new error with a formatted message
	is     func(error, error) bool            // Checks if the error matches a target error
	as     func(error, interface{}) bool      // Determines if the error can be assigned to a specific type
	unwrap func(error) error                  // Unwraps the error to its underlying cause
	new    func(string) error                 // Creates a new error with the given message
	join   func(...error) error               // Joins multiple errors into one
}

// Stack creates an error with a stack trace from the provided string.
func Stack(v string) error {
	return aliases.stack(v)
}

// Stackf formats a string and creates an error with a stack trace.
func Stackf(format string, args ...interface{}) error {
	return aliases.stackf(format, args...)
}

// Wrap wraps an existing error with a new message.
func Wrap(err error, message string) error {
	return aliases.wrap(err, message)
}

// Wrapf formats a string and wraps an existing error with a new message.
func Wrapf(err error, format string, args ...interface{}) error {
	return aliases.wrapf(err, format, args...)
}

// WithStack attaches a stack trace to an existing error.
func WithStack(err error) error {
	return aliases.withStack(err)
}

// Cause returns the root cause of the error.
func Cause(err error) error {
	return aliases.cause(err)
}

// Errorf creates a new error with a formatted message.
func Errorf(format string, args ...interface{}) error {
	return aliases.errorf(format, args...)
}

// Is checks if the error matches a target error.
func Is(err, target error) bool {
	return aliases.is(err, target)
}

// As determines if the error can be assigned to a specific type.
func As(err error, target interface{}) bool {
	return aliases.as(err, target)
}

// Unwrap unwraps the error to its underlying cause.
func Unwrap(err error) error {
	return aliases.unwrap(err)
}

// New creates a new error with the given message.
func New(message string) error {
	return aliases.new(message)
}

// Join joins multiple errors into one.
func Join(errs ...error) error {
	return aliases.join(errs...)
}

// Aliases provides access to the aliased error functions from both the standard library and 'pkg/errors'.
var aliases = errorAliases{
	// pkg/errors aliases
	stack:     errors.New,       // Creates an error with a stack trace
	stackf:    errors.Errorf,    // Formats and creates an error with a stack trace
	wrap:      errors.Wrap,      // Wraps an existing error with a new message
	wrapf:     errors.Wrapf,     // Formats and wraps an existing error with a new message
	withStack: errors.WithStack, // Attaches a stack trace to an existing error
	cause:     errors.Cause,     // Returns the root cause of the error

	// stdlib errors aliases
	errorf: fmt.Errorf,    // Creates a new error with a formatted message
	is:     stderr.Is,     // Checks if the error matches a target error
	as:     stderr.As,     // Determines if the error can be assigned to a specific type
	unwrap: stderr.Unwrap, // Unwraps the error to its underlying cause
	new:    stderr.New,    // Creates a new error with the given message
	join:   stderr.Join,   // Joins multiple errors into one
}
