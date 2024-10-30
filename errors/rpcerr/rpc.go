// Package rpcerr provides a way to return detailed information
// for an RPC request error. The error is normally JSON encoded.
package rpcerr

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
)

// Define  constant IDs for various Error status codes
//
//go:generate protoc -I../../third_party --go_out=paths=source_relative:../../ ../../third_party/errors/rpcerr/rpcerr.proto
const (
	UnknownID               = "http.response.status.unknown"                  // ID for Unknown status
	BadRequestID            = "http.response.status.bad_request"              // ID for Bad Request status
	UnauthorizedID          = "http.response.status.unauthorized"             // ID for Unauthorized status
	ForbiddenID             = "http.response.status.forbidden"                // ID for Forbidden status
	NotFoundID              = "http.response.status.not_found"                // ID for Not Found status
	MethodNotAllowedID      = "http.response.status.method_not_allowed"       // ID for Method Not Allowed status
	TooManyRequestsID       = "http.response.status.too_many_requests"        // ID for Too Many Requests status
	RequestEntityTooLargeID = "http.response.status.request_entity_too_large" // ID for Request Entity Too Large status
	InternalServerErrorID   = "http.response.status.internal_server_error"    // ID for Internal Server Error status
	ConflictID              = "http.response.status.conflict"                 // ID for Conflict status
	RequestTimeoutID        = "http.response.status.request_timeout"          // ID for Request Timeout status
)

var (
	ids = map[int32]string{
		http.StatusBadRequest:            BadRequestID,
		http.StatusUnauthorized:          UnauthorizedID,
		http.StatusForbidden:             ForbiddenID,
		http.StatusNotFound:              NotFoundID,
		http.StatusMethodNotAllowed:      MethodNotAllowedID,
		http.StatusTooManyRequests:       TooManyRequestsID,
		http.StatusRequestEntityTooLarge: RequestEntityTooLargeID,
		http.StatusInternalServerError:   InternalServerErrorID,
		http.StatusConflict:              ConflictID,
		http.StatusRequestTimeout:        RequestTimeoutID,
	}
	mutex sync.RWMutex
)

// RegisterCode register a new error code with the given ID,or overwriting any existing one
func RegisterCode(code int32, id string) {
	mutex.Lock()
	ids[code] = id
	mutex.Unlock()
}

// LookupID looks up the identifier corresponding to the given code.
// This function uses a read lock to ensure safe concurrent access to shared resources.
// Parameters: code - The integer value of the code to look up.
//
// Returns: The string identifier corresponding to the code.
func LookupID(code int32) (string, bool) {
	// Acquire the read lock to ensure data consistency during concurrent reads
	mutex.RLock()
	// Look up the identifier corresponding to the code
	id, ok := ids[code]
	// Release the read lock
	mutex.RUnlock()
	// Return the found identifier
	return id, ok
}

// Error returns the JSON representation of the error
func (obj *Error) Error() string {
	v, _ := json.Marshal(obj)
	return string(v)
}

func (obj *Error) Status() string {
	return http.StatusText(int(obj.Code))
}

func (obj *Error) Is(err error) bool {
	// If both obj and err are nil, return true.
	if obj == nil && err == nil {
		return true
	}

	if err == nil {
		return false
	}

	var target *Error
	if obj != nil && errors.As(err, &target) {
		return obj.Code == target.Code
	}
	return false
}

// Parse tries to parse a JSON string into an error. If that
// fails, it will set the given string as the error detail.
func Parse(err string) *Error {
	e := new(Error)
	errr := json.Unmarshal([]byte(err), e)
	if errr != nil {
		e.Detail = err
	}
	return e
}

// FromError tries to convert a Go error to *Error
func FromError(err error) *Error {
	if err == nil {
		return nil
	}
	var target *Error
	if errors.As(err, &target) && target != nil {
		return target
	}

	return Parse(err.Error())
}

// BadRequest generates a 400 error.
func BadRequest(format string, obj ...any) error {
	return NewFormat(BadRequestID, http.StatusBadRequest, format, obj...)
}

// Unauthorized generates a 401 error.
func Unauthorized(format string, obj ...any) error {
	return NewFormat(UnauthorizedID, http.StatusUnauthorized, format, obj...)
}

// Forbidden generates a 403 error.
func Forbidden(format string, obj ...any) error {
	return NewFormat(ForbiddenID, http.StatusForbidden, format, obj...)
}

// NotFound generates a 404 error.
func NotFound(format string, obj ...any) error {
	return NewFormat(NotFoundID, http.StatusNotFound, format, obj...)
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(format string, obj ...any) error {
	return NewFormat(MethodNotAllowedID, http.StatusMethodNotAllowed, format, obj...)
}

// TooManyRequests generates a 429 error.
func TooManyRequests(format string, obj ...any) error {
	return NewFormat(TooManyRequestsID, http.StatusTooManyRequests, format, obj...)
}

// Timeout generates a 408 error.
func Timeout(format string, obj ...any) error {
	return NewFormat(RequestTimeoutID, http.StatusRequestTimeout, format, obj...)
}

// Conflict generates a 409 error.
func Conflict(format string, obj ...any) error {
	return NewFormat(ConflictID, http.StatusConflict, format, obj...)
}

// RequestEntityTooLarge generates a 413 error.
func RequestEntityTooLarge(format string, obj ...any) error {
	return NewFormat(RequestEntityTooLargeID, http.StatusRequestEntityTooLarge, format, obj...)
}

// InternalServerError generates a 500 error.
func InternalServerError(format string, obj ...any) error {
	return NewFormat(InternalServerErrorID, http.StatusInternalServerError, format, obj...)
}

// New generates a custom error.
func New(id string, code int32, detail string) error {
	return &Error{
		Id:     id,
		Code:   code,
		Detail: detail,
	}
}

// NewFormat generates an Error error
func NewFormat(id string, code int32, format string, args ...any) error {
	return &Error{
		Id:     id,
		Code:   code,
		Detail: fmt.Sprintf(format, args...),
	}
}

// NewCode generates an error for a given code
func NewCode(code int32, detail string) error {
	return &Error{
		Id:     lookupID(UnknownID, code),
		Code:   code,
		Detail: detail,
	}
}

// NewCodeFormat generates an error for a given code
func NewCodeFormat(code int32, format string, args ...any) error {
	return &Error{
		Id:     lookupID(UnknownID, code),
		Code:   code,
		Detail: fmt.Sprintf(format, args...),
	}
}

// Code returns the code of an error
func Code(err error) int {
	e := FromError(err)
	if e == nil {
		return http.StatusInternalServerError
	}
	return int(e.Code)
}

// Cause returns the underlying cause of an error
//func Cause(err error) error {
//	e := FromError(err)
//	if e == nil {
//		return nil
//	}
//	return e.Cause
//}

func NewMultiError() *MultiError {
	return &MultiError{
		Errors: make([]*Error, 0),
	}
}

func (e *MultiError) Append(err *Error) {
	e.Errors = append(e.Errors, err)
}

func (e *MultiError) HasErrors() bool {
	return len(e.Errors) > 0
}

func (e *MultiError) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

// Equal tries to compare errors, which are equal if they have the same NewCode
func Equal(err1 error, err2 error) bool {
	if err1 == nil {
		return err2 == nil
	}

	var verr1 *Error
	ok1 := errors.As(err1, &verr1)
	var verr2 *Error
	ok2 := errors.As(err2, &verr2)

	if ok1 && ok2 {
		return verr1.Code == verr2.Code
	}

	return errors.Is(err1, err2)
}

func lookupID(id string, code int32) string {
	if v, ok := LookupID(code); ok {
		return v
	}
	return id
}
