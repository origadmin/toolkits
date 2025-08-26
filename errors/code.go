/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package errors

import (
	"fmt"
	"sync"
)

const (
	ErrorCodeSuccess = ErrorCode(0)
	ErrorCodeError   = ErrorCode(1)
)

type ErrorCode int

var (
	errCodes = map[ErrorCode]string{
		0: "success",
		1: "error",
	}
	mutCodes = sync.RWMutex{}
)

func RegisterCode(code ErrorCode, val string) {
	mutCodes.Lock()
	errCodes[code] = val
	mutCodes.Unlock()
}

func CodeString(code ErrorCode) string {
	mutCodes.Lock()
	v, ok := errCodes[code]
	mutCodes.Unlock()
	if ok {
		return v
	}
	return fmt.Sprintf("unknown code: %d", code)
}

func (obj ErrorCode) Error() string {
	return obj.String()
}

// Is checks whether the error is equal to the
func (obj ErrorCode) Is(err error) bool {
	if err == nil {
		return false
	}

	var e ErrorCode
	if As(err, &e) {
		return e == obj
	}

	return false
}

func (obj ErrorCode) Code() int {
	return int(obj)
}

func (obj ErrorCode) String() string {
	return CodeString(obj)
}

// ErrInteger creates a new error from a string
func ErrInteger(err int) ErrorCode {
	return ErrorCode(err)
}
