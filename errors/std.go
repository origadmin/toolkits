package errors

import (
	"errors"
)

type StdError string

// Error returns the JSON representation of the error
func (obj StdError) Error() string {
	return obj.String()
}

// String returns the JSON representation of the error
func (obj StdError) String() string {
	return string(obj)
}

// Is checks whether the error is equal to the
func (obj StdError) Is(err error) bool {
	if err == nil {
		return false
	}

	var e StdError
	if errors.As(err, &e) {
		return obj.String() == e.String()
	}

	return false
}
