// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/v1/middleware.proto

package config

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Middleware with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Middleware) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Middleware with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MiddlewareMultiError, or
// nil if none found.
func (m *Middleware) ValidateAll() error {
	return m.validate(true)
}

func (m *Middleware) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, MiddlewareValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, MiddlewareValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return MiddlewareValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return MiddlewareMultiError(errors)
	}

	return nil
}

// MiddlewareMultiError is an error wrapping multiple validation errors
// returned by Middleware.ValidateAll() if the designated constraints aren't met.
type MiddlewareMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MiddlewareMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MiddlewareMultiError) AllErrors() []error { return m }

// MiddlewareValidationError is the validation error returned by
// Middleware.Validate if the designated constraints aren't met.
type MiddlewareValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MiddlewareValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MiddlewareValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MiddlewareValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MiddlewareValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MiddlewareValidationError) ErrorName() string { return "MiddlewareValidationError" }

// Error satisfies the builtin error interface
func (e MiddlewareValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMiddleware.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MiddlewareValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MiddlewareValidationError{}

// Validate checks the field values on Middleware_RateLimiter with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Middleware_RateLimiter) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Middleware_RateLimiter with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Middleware_RateLimiterMultiError, or nil if none found.
func (m *Middleware_RateLimiter) ValidateAll() error {
	return m.validate(true)
}

func (m *Middleware_RateLimiter) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Period

	// no validation rules for XRatelimitLimit

	// no validation rules for XRatelimitRemaining

	// no validation rules for XRatelimitReset

	// no validation rules for RetryAfter

	// no validation rules for StoreType

	if all {
		switch v := interface{}(m.GetMemory()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Middleware_RateLimiterValidationError{
					field:  "Memory",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Middleware_RateLimiterValidationError{
					field:  "Memory",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetMemory()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Middleware_RateLimiterValidationError{
				field:  "Memory",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRedis()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Middleware_RateLimiterValidationError{
					field:  "Redis",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Middleware_RateLimiterValidationError{
					field:  "Redis",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRedis()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Middleware_RateLimiterValidationError{
				field:  "Redis",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Middleware_RateLimiterMultiError(errors)
	}

	return nil
}

// Middleware_RateLimiterMultiError is an error wrapping multiple validation
// errors returned by Middleware_RateLimiter.ValidateAll() if the designated
// constraints aren't met.
type Middleware_RateLimiterMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Middleware_RateLimiterMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Middleware_RateLimiterMultiError) AllErrors() []error { return m }

// Middleware_RateLimiterValidationError is the validation error returned by
// Middleware_RateLimiter.Validate if the designated constraints aren't met.
type Middleware_RateLimiterValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Middleware_RateLimiterValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Middleware_RateLimiterValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Middleware_RateLimiterValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Middleware_RateLimiterValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Middleware_RateLimiterValidationError) ErrorName() string {
	return "Middleware_RateLimiterValidationError"
}

// Error satisfies the builtin error interface
func (e Middleware_RateLimiterValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMiddleware_RateLimiter.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Middleware_RateLimiterValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Middleware_RateLimiterValidationError{}

// Validate checks the field values on Middleware_RateLimiter_Redis with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Middleware_RateLimiter_Redis) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Middleware_RateLimiter_Redis with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Middleware_RateLimiter_RedisMultiError, or nil if none found.
func (m *Middleware_RateLimiter_Redis) ValidateAll() error {
	return m.validate(true)
}

func (m *Middleware_RateLimiter_Redis) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Addr

	// no validation rules for Username

	// no validation rules for Password

	// no validation rules for Db

	if len(errors) > 0 {
		return Middleware_RateLimiter_RedisMultiError(errors)
	}

	return nil
}

// Middleware_RateLimiter_RedisMultiError is an error wrapping multiple
// validation errors returned by Middleware_RateLimiter_Redis.ValidateAll() if
// the designated constraints aren't met.
type Middleware_RateLimiter_RedisMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Middleware_RateLimiter_RedisMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Middleware_RateLimiter_RedisMultiError) AllErrors() []error { return m }

// Middleware_RateLimiter_RedisValidationError is the validation error returned
// by Middleware_RateLimiter_Redis.Validate if the designated constraints
// aren't met.
type Middleware_RateLimiter_RedisValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Middleware_RateLimiter_RedisValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Middleware_RateLimiter_RedisValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Middleware_RateLimiter_RedisValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Middleware_RateLimiter_RedisValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Middleware_RateLimiter_RedisValidationError) ErrorName() string {
	return "Middleware_RateLimiter_RedisValidationError"
}

// Error satisfies the builtin error interface
func (e Middleware_RateLimiter_RedisValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMiddleware_RateLimiter_Redis.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Middleware_RateLimiter_RedisValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Middleware_RateLimiter_RedisValidationError{}

// Validate checks the field values on Middleware_RateLimiter_Memory with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *Middleware_RateLimiter_Memory) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Middleware_RateLimiter_Memory with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// Middleware_RateLimiter_MemoryMultiError, or nil if none found.
func (m *Middleware_RateLimiter_Memory) ValidateAll() error {
	return m.validate(true)
}

func (m *Middleware_RateLimiter_Memory) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetExpiration()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Middleware_RateLimiter_MemoryValidationError{
					field:  "Expiration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Middleware_RateLimiter_MemoryValidationError{
					field:  "Expiration",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExpiration()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Middleware_RateLimiter_MemoryValidationError{
				field:  "Expiration",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetCleanupInterval()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Middleware_RateLimiter_MemoryValidationError{
					field:  "CleanupInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Middleware_RateLimiter_MemoryValidationError{
					field:  "CleanupInterval",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCleanupInterval()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Middleware_RateLimiter_MemoryValidationError{
				field:  "CleanupInterval",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Middleware_RateLimiter_MemoryMultiError(errors)
	}

	return nil
}

// Middleware_RateLimiter_MemoryMultiError is an error wrapping multiple
// validation errors returned by Middleware_RateLimiter_Memory.ValidateAll()
// if the designated constraints aren't met.
type Middleware_RateLimiter_MemoryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Middleware_RateLimiter_MemoryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Middleware_RateLimiter_MemoryMultiError) AllErrors() []error { return m }

// Middleware_RateLimiter_MemoryValidationError is the validation error
// returned by Middleware_RateLimiter_Memory.Validate if the designated
// constraints aren't met.
type Middleware_RateLimiter_MemoryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Middleware_RateLimiter_MemoryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Middleware_RateLimiter_MemoryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Middleware_RateLimiter_MemoryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Middleware_RateLimiter_MemoryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Middleware_RateLimiter_MemoryValidationError) ErrorName() string {
	return "Middleware_RateLimiter_MemoryValidationError"
}

// Error satisfies the builtin error interface
func (e Middleware_RateLimiter_MemoryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMiddleware_RateLimiter_Memory.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Middleware_RateLimiter_MemoryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Middleware_RateLimiter_MemoryValidationError{}
