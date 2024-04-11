package errors

import (
	"errors"
)

type StdError string

// HTTP returns the JSON representation of the error
func (obj StdError) Error() string {
	return obj.String()
}

func (obj StdError) String() string {
	return string(obj)
}

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
