package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/pkg/errors"
)

// Define  constant IDs for various Error status codes
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

// Error customize the error structure for implementation errors.Error interface
type Error struct {
	ID     string `json:"id,omitempty"`
	Code   int32  `json:"code,omitempty"`
	Detail string `json:"detail,omitempty"`
}

// Error returns the JSON representation of the error
func (obj *Error) Error() string {
	return obj.String()
}

func (obj *Error) String() string {
	v, _ := json.Marshal(obj)
	return string(v)
}

func (obj *Error) Status() string {
	return http.StatusText(int(obj.Code))
}

func (obj *Error) Is(err error) bool {
	if err == nil {
		return false
	}
	var target *Error
	if errors.As(err, &target) {
		return obj.ID == target.ID && obj.Code == target.Code
	}
	return false
}

// FromStringToHTTP tries to parse a JSON string into an error. If that
// fails, it will set the given string as the error detail.
func FromStringToHTTP(err string) *Error {
	obj := new(Error)
	err2 := json.Unmarshal([]byte(err), obj)
	if err2 != nil {
		obj.Detail = err
	}
	return obj
}

// FromErrorToHTTP tries to convert a Go error to *Error
func FromErrorToHTTP(err error) *Error {
	if err == nil {
		return nil
	}
	var target *Error
	if errors.As(err, &target) && target != nil {
		return target
	}

	return FromStringToHTTP(err.Error())
}

// BadRequest generates a 400 error.
func BadRequest(id, format string, obj ...any) error {
	if id == "" {
		id = BadRequestID
	}
	return &Error{
		ID:     id,
		Code:   http.StatusBadRequest,
		Detail: fmt.Sprintf(format, obj...),
	}
}

// Unauthorized generates a 401 error.
func Unauthorized(id, format string, obj ...any) error {
	if id == "" {
		id = UnauthorizedID
	}
	return &Error{
		ID:     id,
		Code:   http.StatusUnauthorized,
		Detail: fmt.Sprintf(format, obj...),
	}
}

// Forbidden generates a 403 error.
func Forbidden(id, format string, obj ...any) error {
	if id == "" {
		id = ForbiddenID
	}
	return &Error{
		ID:     id,
		Code:   http.StatusForbidden,
		Detail: fmt.Sprintf(format, obj...),
	}
}

// NotFound generates a 404 error.
func NotFound(id, format string, obj ...any) error {
	if id == "" {
		id = NotFoundID
	}
	return &Error{
		ID:     id,
		Code:   http.StatusNotFound,
		Detail: fmt.Sprintf(format, obj...),
	}
}

// MethodNotAllowed generates a 405 error.
func MethodNotAllowed(id, format string, obj ...any) error {
	if id == "" {
		id = MethodNotAllowedID
	}
	return &Error{
		ID:     id,
		Code:   http.StatusMethodNotAllowed,
		Detail: fmt.Sprintf(format, obj...),
	}
}

// TooManyRequests generates a 429 error.
func TooManyRequests(id, format string, obj ...any) error {
	if id == "" {
		id = TooManyRequestsID
	}
	return &Error{
		ID:     id,
		Code:   http.StatusTooManyRequests,
		Detail: fmt.Sprintf(format, obj...),
	}
}

// Timeout generates a 408 error.
func Timeout(id, format string, obj ...any) error {
	if id == "" {
		id = RequestTimeoutID
	}
	return &Error{
		ID:     id,
		Code:   http.StatusRequestTimeout,
		Detail: fmt.Sprintf(format, obj...),
	}
}

// Conflict generates a 409 error.
func Conflict(id, format string, obj ...any) error {
	// Set  ID if not provided
	if id == "" {
		id = ConflictID
	}
	return &Error{
		ID:     id,
		Code:   http.StatusConflict,
		Detail: fmt.Sprintf(format, obj...),
	}
}

// RequestEntityTooLarge generates a 413 error.
func RequestEntityTooLarge(id, format string, obj ...any) error {
	// Set  ID if not provided
	if id == "" {
		id = RequestEntityTooLargeID
	}
	return &Error{
		ID:     id,
		Code:   http.StatusRequestEntityTooLarge,
		Detail: fmt.Sprintf(format, obj...),
	}
}

// InternalServerError generates a 500 error.
func InternalServerError(id, format string, obj ...any) error {
	// Set  ID if not provided
	if id == "" {
		id = InternalServerErrorID
	}
	return &Error{
		ID:     id,
		Code:   http.StatusInternalServerError,
		Detail: fmt.Sprintf(format, obj...),
	}
}

// New generates an Error error
func New(id string, code int32, detail string) error {
	return &Error{
		ID:     id,
		Code:   code,
		Detail: detail,
	}
}

// Newf generates an Error error
func Newf(id string, code int32, format string, args ...any) error {
	return &Error{
		ID:     id,
		Code:   code,
		Detail: fmt.Sprintf(format, args...),
	}
}

// Code generates an error for a given code
func Code(code int32, detail string) error {
	id, ok := ids[code]
	if !ok {
		id = UnknownID
	}
	return &Error{
		ID:     id,
		Code:   code,
		Detail: detail,
	}
}

// Codef generates an error for a given code
func Codef(code int32, format string, args ...any) error {
	id, ok := ids[code]
	if !ok {
		id = UnknownID
	}
	return &Error{
		ID:     id,
		Code:   code,
		Detail: fmt.Sprintf(format, args...),
	}
}
