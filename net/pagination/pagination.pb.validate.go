// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: pagination.proto

package pagination

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

// Validate checks the field values on PageRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageRequestMultiError, or
// nil if none found.
func (m *PageRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *PageRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for PageToken

	// no validation rules for OnlyCount

	if all {
		switch v := interface{}(m.GetFieldMask()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PageRequestValidationError{
					field:  "FieldMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PageRequestValidationError{
					field:  "FieldMask",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetFieldMask()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PageRequestValidationError{
				field:  "FieldMask",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.Current != nil {
		// no validation rules for Current
	}

	if m.PageSize != nil {
		// no validation rules for PageSize
	}

	if m.NoPaging != nil {
		// no validation rules for NoPaging
	}

	if len(errors) > 0 {
		return PageRequestMultiError(errors)
	}

	return nil
}

// PageRequestMultiError is an error wrapping multiple validation errors
// returned by PageRequest.ValidateAll() if the designated constraints aren't met.
type PageRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageRequestMultiError) AllErrors() []error { return m }

// PageRequestValidationError is the validation error returned by
// PageRequest.Validate if the designated constraints aren't met.
type PageRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageRequestValidationError) ErrorName() string { return "PageRequestValidationError" }

// Error satisfies the builtin error interface
func (e PageRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageRequestValidationError{}

// Validate checks the field values on PageResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *PageResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on PageResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in PageResponseMultiError, or
// nil if none found.
func (m *PageResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *PageResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Success

	// no validation rules for Total

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PageResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PageResponseValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PageResponseValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for NextPageToken

	if all {
		switch v := interface{}(m.GetExtra()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PageResponseValidationError{
					field:  "Extra",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PageResponseValidationError{
					field:  "Extra",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetExtra()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PageResponseValidationError{
				field:  "Extra",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.Current != nil {
		// no validation rules for Current
	}

	if m.PageSize != nil {
		// no validation rules for PageSize
	}

	if len(errors) > 0 {
		return PageResponseMultiError(errors)
	}

	return nil
}

// PageResponseMultiError is an error wrapping multiple validation errors
// returned by PageResponse.ValidateAll() if the designated constraints aren't met.
type PageResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PageResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PageResponseMultiError) AllErrors() []error { return m }

// PageResponseValidationError is the validation error returned by
// PageResponse.Validate if the designated constraints aren't met.
type PageResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PageResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PageResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PageResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PageResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PageResponseValidationError) ErrorName() string { return "PageResponseValidationError" }

// Error satisfies the builtin error interface
func (e PageResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPageResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PageResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PageResponseValidationError{}
