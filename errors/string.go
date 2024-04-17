package errors

import (
	"errors"
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
	if errors.As(err, &e) {
		return strings.Compare(obj.String(), e.String()) == 0
	}

	return false
}

// ErrString creates a new error from a string
func ErrString(err string) String {
	return String(err)
}
