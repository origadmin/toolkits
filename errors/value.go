package errors

// Valuer defines an interface for retrieving encapsulated values and error information.
// It provides a unified way to access different types of values and error information.
type Valuer[T any] interface {
	Value() T          // Value returns the encapsulated value.
	Error() string     // Error returns the string representation of the error, satisfying the error interface.
	Unwrap() error     // Unwrap returns the original error object.
	Is(err error) bool // Is determines if the current error was created by the given error.
}

// wrapped is a generic struct used to encapsulate any type of value and error.
type wrapped[T any] struct {
	v T     // v is the encapsulated value.
	e error // e is the encapsulated original error.
}

// Valued is a generic function that creates a Valuer interface object encapsulating a value and an error.
// Parameters:
//
//	v - The value to be encapsulated.
//	err - The original error to be encapsulated.
//
// Returns:
//
//	An object implementing the Valuer interface, encapsulating the given value and error.
func Valued[T any](v T, err error) Valuer[T] {
	return &wrapped[T]{v, err}
}

// Value implements the Value method of the Valuer interface, returning the encapsulated value.
func (obj wrapped[T]) Value() T {
	return obj.v
}

// Unwrap implements the Unwrap method of the Valuer interface, returning the encapsulated original error.
func (obj wrapped[T]) Unwrap() error {
	return obj.e
}

// Error implements the Error method of the Valuer interface, returning the string representation of the error.
func (obj wrapped[T]) Error() string {
	if obj.e == nil {
		return ""
	}
	return obj.e.Error()
}

// Is implements the Is method of the Valuer interface, determining if the current error was created by the given error.
func (obj wrapped[T]) Is(err error) bool {
	if obj.e == nil {
		return err == nil
	}

	return Is(obj.e, err)
}
