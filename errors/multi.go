package errors

import (
	"encoding/json"
	"errors"
	"sync"

	merr "github.com/hashicorp/go-multierror"
)

type (
	// MultiError is an alias of `github.com/hashicorp/go-multierror`.Error
	MultiError = merr.Error
	// ErrorFormatFunc is a helper function that will format the merr
	ErrorFormatFunc = merr.ErrorFormatFunc
)

var (
	// Append is a helper function that will append more merr
	Append = merr.Append
	// ListFormatFunc is a helper function that will format the merr
	ListFormatFunc = merr.ListFormatFunc
)

// ThreadSafeMultiError  represents a collection of merr
type ThreadSafeMultiError struct {
	ErrorFormat ErrorFormatFunc
	lock        sync.Mutex
	merr        MultiError
}

// Append adds an error to the MultiError collection
func (e *ThreadSafeMultiError) Append(err error) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.merr.Errors = append(e.merr.Errors, err)
}

// HasErrors checks if the MultiError collection has any merr
func (e *ThreadSafeMultiError) HasErrors() bool {
	e.lock.Lock()
	defer e.lock.Unlock()
	return len(e.merr.Errors) > 0
}

// Has checks if the MultiError collection has the given merr or not
func (e *ThreadSafeMultiError) Has(err any) error {
	var idx = -1
	e.lock.Lock()
	defer e.lock.Unlock()
	for idx = range e.merr.Errors {
		if errors.As(e.merr.Errors[idx], &err) {
			return e.merr.Errors[idx]
		}
	}
	return nil
}

// Unsafe returns the MultiError collection
func (e *ThreadSafeMultiError) Unsafe() *MultiError {
	e.lock.Lock()
	defer e.lock.Unlock()
	me := new(MultiError)
	me.ErrorFormat = e.ErrorFormat
	me.Errors = append(me.Errors, e.merr.Errors...)
	return me
}

// Error returns the JSON representation of the MultiError collection
func (e *ThreadSafeMultiError) Error() string {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.merr.ErrorFormat = e.ErrorFormat
	return e.merr.Error()
}

// Errors returns the merr collection
func (e *ThreadSafeMultiError) Errors() []error {
	e.lock.Lock()
	defer e.lock.Unlock()
	return append([]error{}, e.merr.Errors...)
}

// ThreadSafe creates a new ThreadSafeMultiError collection
func ThreadSafe(err error, fns ...ErrorFormatFunc) *ThreadSafeMultiError {
	if len(fns) == 0 {
		fns = append(fns, ErrorFormatJSON)
	}

	if err == nil {
		return &ThreadSafeMultiError{
			ErrorFormat: fns[0],
			merr: MultiError{
				ErrorFormat: fns[0],
			},
		}
	}

	var multiError *MultiError
	if errors.As(err, &multiError) {
		return &ThreadSafeMultiError{
			ErrorFormat: fns[0],
			merr:        *multiError,
		}
	}
	return &ThreadSafeMultiError{
		ErrorFormat: fns[0],
		merr: MultiError{
			ErrorFormat: fns[0],
			Errors:      []error{err},
		},
	}
}

func ErrorFormatJSON(i []error) string {
	var errStrs []string
	for _, e := range i {
		errStrs = append(errStrs, e.Error())
	}
	bytes, _ := json.Marshal(errStrs)
	return string(bytes)
}
