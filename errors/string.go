/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package errors

const (
	UnsupportedError = String("unsupported error")
)

type String string

// Error returns the JSON representation of the error
func (s String) Error() string {
	return string(s)
}

// String returns the JSON representation of the error
func (s String) String() string {
	return string(s)
}

// Is checks if the target error is a String error and has the same value.
// This allows `errors.Is(wrappedErr, someStringError)` to work correctly.
func (s String) Is(err error) bool {
	if err == nil {
		return false
	}

	var e String
	if As(err, &e) {
		return s == e
	}

	return false
}

// NewString creates a new error from a string
func NewString(message string) String {
	return String(message)
}
