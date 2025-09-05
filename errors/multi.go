/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package errors

import (
	"encoding/json"
	"errors"
	"sync"
)

// ThreadSafeMultiError represents a thread-safe collection of errors.
// It is a wrapper around hashicorp/go-multierror.
type ThreadSafeMultiError struct {
	ErrorFormat MultiErrorFormatFunc
	lock        sync.Mutex
	multiErr    MultiError
}

// Append adds an error to the MultiError collection
func (m *ThreadSafeMultiError) Append(err error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.multiErr.Errors = append(m.multiErr.Errors, err)
}

// HasErrors checks if the MultiError collection has any errors.
func (m *ThreadSafeMultiError) HasErrors() bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	return len(m.multiErr.Errors) > 0
}

// Contains checks if any error in the collection matches the target error
// using errors.Is. It returns true if a match is found.
func (m *ThreadSafeMultiError) Contains(target error) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	for _, err := range m.multiErr.Errors {
		if errors.Is(err, target) {
			return true
		}
	}
	return false
}

// Snapshot returns a new, non-thread-safe copy of the MultiError collection.
func (m *ThreadSafeMultiError) Snapshot() *MultiError {
	m.lock.Lock()
	defer m.lock.Unlock()
	me := new(MultiError)
	me.ErrorFormat = m.ErrorFormat
	me.Errors = append(me.Errors, m.multiErr.Errors...)
	return me
}

// Error returns the JSON representation of the MultiError collection
func (m *ThreadSafeMultiError) Error() string {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.multiErr.ErrorFormat = m.ErrorFormat
	return m.multiErr.Error()
}

// Errors returns a copy of the errors collection.
func (m *ThreadSafeMultiError) Errors() []error {
	m.lock.Lock()
	defer m.lock.Unlock()
	return append([]error{}, m.multiErr.Errors...)
}

// ThreadSafe creates a new ThreadSafeMultiError collection
func ThreadSafe(err error, fns ...MultiErrorFormatFunc) *ThreadSafeMultiError {
	if len(fns) == 0 {
		fns = append(fns, ErrorFormatJSON)
	}

	if err == nil {
		return &ThreadSafeMultiError{
			ErrorFormat: fns[0],
			multiErr: MultiError{
				ErrorFormat: fns[0],
			},
		}
	}

	var multiError *MultiError
	if As(err, &multiError) {
		return &ThreadSafeMultiError{
			ErrorFormat: fns[0],
			multiErr:    *multiError,
		}
	}
	return &ThreadSafeMultiError{
		ErrorFormat: fns[0],
		multiErr: MultiError{
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
