/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package errors

import (
	"strings"
)

const (
	UnsupportedError = String("unsupported error")
)

type String string

// Error returns the JSON representation of the error
func (obj String) Error() string {
	return obj.String()
}

// String returns the JSON representation of the error
func (obj String) String() string {
	return string(obj)
}

// Is checks whether the error is equal to the
func (obj String) Is(err error) bool {
	if err == nil {
		return false
	}

	var e String
	if As(err, &e) {
		return strings.Compare(obj.String(), e.String()) == 0
	}

	return false
}

// ErrString creates a new error from a string
//
// Deprecated: use errors.NewString instead
func ErrString(err string) String {
	return String(err)
}

// NewString creates a new error from a string
func NewString(message string) String {
	return String(message)
}
