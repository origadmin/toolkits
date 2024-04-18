// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package errors provides a way to return detailed information
// for a request error. The error is normally JSON encoded.
package errors

import (
	stderr "errors"
	"fmt"

	"github.com/pkg/errors"
)

// // Define alias from pkg/errors
// var (
//     // Stack adds stack trace to error, with message
//     Stack = errors.New
//     // Stackf adds stack trace to error, with format specifier
//     Stackf = errors.Errorf
//     // Wrap returns an error annotating err without stack trace
//     Wrap = errors.WithMessage
//     // Wrapf returns an error annotating err without stack trace
//     Wrapf = errors.WithMessagef
//     // WithStack annotates err with a stack trace
//     WithStack = errors.WithStack
//     // Cause returns the underlying cause of the error, if possiblaliases.
//     Cause = errors.Cause
// )
//
// // Define alias from stderr
// var (
//     // Errorf formats according to a format specifier and returns the string as a value that satisfies error
//     Errorf = fmt.Errorf
//     // Is reports whether any error in errs chain matches target
//     Is = stderr.Is
//     // As finds the first error in errs that matches target, and if one is found, sets target to that error value and returns trualiases. Otherwise, it returns falsaliases.
//     As = stderr.As
//     // Unwrap unwraps the error, and if it is a wrapper, returns the next error in the chain.
//     Unwrap = stderr.Unwrap
//     // New creates a new error with the given string
//     New = stderr.New
//     // Join joins any number of errors into a single error.
//     Join = stderr.Join
// )

// errorAliases holds the aliased error functions as unexported variables.
type errorAliases struct {
	// pkg/errors aliases
	stack     func(string) error
	stackf    func(string, ...interface{}) error
	wrap      func(error, string) error
	wrapf     func(error, string, ...interface{}) error
	withStack func(error) error
	cause     func(error) error

	// stdlib errors aliases
	errorf func(string, ...interface{}) error
	is     func(error, error) bool
	as     func(error, interface{}) bool
	unwrap func(error) error
	new    func(string) error
	join   func(...error) error
}

func Stack(v string) error {
	return aliases.stack(v)
}

func Stackf(format string, args ...interface{}) error {
	return aliases.stackf(format, args...)
}

func Wrap(err error, message string) error {
	return aliases.wrap(err, message)
}

func Wrapf(err error, format string, args ...interface{}) error {
	return aliases.wrapf(err, format, args...)
}

func WithStack(err error) error {
	return aliases.withStack(err)
}

func Cause(err error) error {
	return aliases.cause(err)
}

func Errorf(format string, args ...interface{}) error {
	return aliases.errorf(format, args...)
}

func Is(err, target error) bool {
	return aliases.is(err, target)
}

func As(err error, target interface{}) bool {
	return aliases.as(err, target)
}

func Unwrap(err error) error {
	return aliases.unwrap(err)
}

func New(message string) error {
	return aliases.new(message)
}

func Join(errs ...error) error {
	return aliases.join(errs...)
}

// Aliases provides access to the aliased error functions.
var aliases = errorAliases{
	// pkg/errors aliases
	stack:     errors.New,
	stackf:    errors.Errorf,
	wrap:      errors.WithMessage,
	wrapf:     errors.WithMessagef,
	withStack: errors.WithStack,
	cause:     errors.Cause,

	// stdlib errors aliases
	errorf: fmt.Errorf,
	is:     stderr.Is,
	as:     stderr.As,
	unwrap: stderr.Unwrap,
	new:    stderr.New,
	join:   stderr.Join,
}
