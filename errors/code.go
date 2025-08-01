/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package errors

import (
	"fmt"
	"sync"
)

const (
	Success = Code(0)
	Error   = Code(1)
)

type Code int

var (
	errCodes = map[Code]string{
		0: "success",
		1: "error",
	}
	mutCodes = sync.RWMutex{}
)

func RegisterCode(code Code, val string) {
	mutCodes.Lock()
	errCodes[code] = val
	mutCodes.Unlock()
}

func CodeString(code Code) string {
	mutCodes.Lock()
	v, ok := errCodes[code]
	mutCodes.Unlock()
	if ok {
		return v
	}
	return fmt.Sprintf("unknown code: %d", code)
}

func (obj Code) Error() string {
	return obj.String()
}

// Is checks whether the error is equal to the
func (obj Code) Is(err error) bool {
	if err == nil {
		return false
	}

	var e Code
	if As(err, &e) {
		return e == obj
	}

	return false
}

func (obj Code) Code() int {
	return int(obj)
}

func (obj Code) String() string {
	return CodeString(obj)
}

// ErrInteger creates a new error from a string
func ErrInteger(err int) Code {
	return Code(err)
}
