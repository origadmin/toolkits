/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package errors

import (
	"fmt"
	"sync"
)

const (
	// Deprecated: ErrorCodeSuccess represents a success code. In Go, `nil` is the idiomatic way to indicate success.
	// Returning a non-nil error for a successful operation (even with code 0) is an anti-pattern and can lead to incorrect error handling
	// with standard `if err != nil` checks. This should be used with caution, primarily at application boundaries (e.g., API responses)
	// and not in general business logic. Prefer returning `nil` for success.
	ErrorCodeSuccess = ErrorCode(0)

	// ErrorCodeError represents a generic error code.
	ErrorCodeError = ErrorCode(1)
)

type ErrorCode int

var (
	errCodes = map[ErrorCode]string{
		ErrorCodeSuccess: "success",
		ErrorCodeError:   "error",
	}
	mutCodes = sync.RWMutex{}
)

func RegisterCode(code ErrorCode, val string) {
	mutCodes.Lock()
	errCodes[code] = val
	mutCodes.Unlock()
}

func CodeString(code ErrorCode) string {
	mutCodes.RLock()
	v, ok := errCodes[code]
	mutCodes.RUnlock()
	if ok {
		return v
	}
	return fmt.Sprintf("unknown code: %d", code)
}

func (ec ErrorCode) Error() string {
	return CodeString(ec)
}

// Is checks if the target error is an ErrorCode and has the same value.
// This allows `errors.Is(wrappedErr, someErrorCode)` to work correctly.
func (ec ErrorCode) Is(err error) bool {
	if err == nil {
		return false
	}

	var e ErrorCode
	if As(err, &e) {
		return e == ec
	}

	return false
}

func (ec ErrorCode) Code() int {
	return int(ec)
}

func (ec ErrorCode) String() string {
	return CodeString(ec)
}

// NewCode creates a new ErrorCode from an integer.
func NewCode(code int) ErrorCode {
	return ErrorCode(code)
}
