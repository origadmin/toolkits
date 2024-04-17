// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package errors provides a way to return detailed information
// for a request error. The error is normally JSON encoded.
package errors

import (
	stderr "errors"
	"fmt"

	"github.com/pkg/errors"
)

// Define alias from pkg/errors
var (
	// Stack adds stack trace to error, with message
	Stack = errors.New
	// Stackf adds stack trace to error, with format specifier
	Stackf = errors.Errorf
	// Wrap returns an error annotating err with a stack trace
	Wrap = errors.WithMessage
	// Wrapf returns an error annotating err with a stack trace
	Wrapf = errors.WithMessagef
	// WithStack annotates err with a stack trace
	WithStack = errors.WithStack
)

// Define alias from stderr
var (
	// Errorf formats according to a format specifier and returns the string as a value that satisfies error
	Errorf = fmt.Errorf
	// Is reports whether any error in errs chain matches target
	Is = stderr.Is
	// As finds the first error in errs that matches target, and if one is found, sets target to that error value and returns true. Otherwise, it returns false.
	As = stderr.As
	// Unwrap unwraps the error, and if it is a wrapper, returns the next error in the chain.
	Unwrap = stderr.Unwrap
	// New creates a new error with the given string
	New = stderr.New
	// Join joins any number of errors into a single error.
	Join = stderr.Join
)
