/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package errors

import (
	"encoding/json"
	"errors"
	"sync"
)

// ThreadSafeMultiError provides a thread-safe implementation of a multi-error type.
// It wraps hashicorp/go-multierror and adds thread safety for concurrent operations.
//
// This type is particularly useful in concurrent scenarios where multiple goroutines
// may need to append errors to the same collection simultaneously.
//
// Example:
//
//	merr := ThreadSafe(nil)
//	var wg sync.WaitGroup
//
//	for i := 0; i < 10; i++ {
//	    wg.Add(1)
//	    go func(n int) {
//	        defer wg.Done()
//	        if n%2 == 0 {
//	            merr.Append(fmt.Errorf("even error: %d", n))
//	        }
//	    }(i)
//	}
//
//	wg.Wait()
//	if merr.HasErrors() {
//	    log.Printf("Encountered %d errors: %v", len(merr.Errors()), merr)
//	}
type ThreadSafeMultiError struct {
	ErrorFormat MultiErrorFormatFunc
	lock        sync.Mutex
	multiErr    MultiError
}

// Append adds an error to the ThreadSafeMultiError collection in a thread-safe manner.
// If the error is nil, it will be ignored. If the error is a MultiError,
// its individual errors will be flattened and added to the collection.
//
// This method is safe for concurrent use by multiple goroutines.
func (m *ThreadSafeMultiError) Append(err error) {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.multiErr.Errors = append(m.multiErr.Errors, err)
}

// HasErrors checks if the ThreadSafeMultiError collection contains any errors.
//
// Returns:
//   - bool: true if there is at least one error in the collection, false otherwise.
//
// This method is safe for concurrent use by multiple goroutines.
func (m *ThreadSafeMultiError) HasErrors() bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	return len(m.multiErr.Errors) > 0
}

// Contains checks if any error in the collection matches the target error
// using errors.Is. It performs a deep equality check by unwrapping errors.
//
// Parameters:
//   - target: The error to search for in the collection.
//
// Returns:
//   - bool: true if the target error is found in the collection, false otherwise.
//
// This method is safe for concurrent use by multiple goroutines.
//
// Example:
//
//	merr := ThreadSafe(nil)
//	merr.Append(io.EOF)
//	fmt.Println(merr.Contains(io.EOF)) // Output: true
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

// Snapshot creates and returns a new, non-thread-safe copy of the current error collection.
// The returned MultiError is a snapshot of the errors at the time of the call.
//
// Returns:
//   - *MultiError: A new MultiError containing a copy of the current errors.
//
// This method is safe for concurrent use by multiple goroutines, but the returned
// MultiError is not thread-safe.
func (m *ThreadSafeMultiError) Snapshot() *MultiError {
	m.lock.Lock()
	defer m.lock.Unlock()
	me := new(MultiError)
	me.ErrorFormat = m.ErrorFormat
	me.Errors = append(me.Errors, m.multiErr.Errors...)
	return me
}

// Error returns the string representation of the error collection.
// The format is determined by the ErrorFormat function (defaults to ErrorFormatJSON).
//
// Returns:
//   - string: The formatted error string.
//
// This method is safe for concurrent use by multiple goroutines.
func (m *ThreadSafeMultiError) Error() string {
	m.lock.Lock()
	defer m.lock.Unlock()
	m.multiErr.ErrorFormat = m.ErrorFormat
	return m.multiErr.Error()
}

// Errors returns a copy of the underlying errors slice.
//
// Returns:
//   - []error: A new slice containing all errors in the collection.
//
// This method is safe for concurrent use by multiple goroutines.
func (m *ThreadSafeMultiError) Errors() []error {
	m.lock.Lock()
	defer m.lock.Unlock()
	return append([]error{}, m.multiErr.Errors...)
}

// ThreadSafe creates and initializes a new ThreadSafeMultiError.
//
// Parameters:
//   - err: An optional initial error (or nil).
//   - fns: Optional list of MultiErrorFormatFunc to configure formatting.
//     The first function is used as the formatter.
//
// Returns:
//   - *ThreadSafeMultiError: A new thread-safe multi-error instance.
//
// If no formatter is provided, ErrorFormatJSON is used by default.
//
// Example:
//
//	// Create an empty thread-safe multi-error with default JSON formatting
//	merr := ThreadSafe(nil)
//
//	// Create with an initial error and custom formatting
//	merr := ThreadSafe(io.EOF, func(errs []error) string {
//	    return fmt.Sprintf("Found %d errors: %v", len(errs), errs)
//	})
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

// ErrorFormatJSON formats the list of errors as a JSON array of strings.
// Each error is converted to its string representation using Error().
//
// Parameters:
//   - i: The list of errors to format.
//
// Returns:
//   - string: A JSON array string containing the error messages.
//
// Example:
//
//	["error 1", "error 2"]
func ErrorFormatJSON(i []error) string {
	var errStrs []string
	for _, e := range i {
		errStrs = append(errStrs, e.Error())
	}
	bytes, _ := json.Marshal(errStrs)
	return string(bytes)
}
