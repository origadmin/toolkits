// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/v1/customize.proto

package configv1

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

// Validate checks the field values on Customize with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Customize) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Customize with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CustomizeMultiError, or nil
// if none found.
func (m *Customize) ValidateAll() error {
	return m.validate(true)
}

func (m *Customize) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	{
		sorted_keys := make([]string, len(m.GetConfigs()))
		i := 0
		for key := range m.GetConfigs() {
			sorted_keys[i] = key
			i++
		}
		sort.Slice(sorted_keys, func(i, j int) bool { return sorted_keys[i] < sorted_keys[j] })
		for _, key := range sorted_keys {
			val := m.GetConfigs()[key]
			_ = val

			// no validation rules for Configs[key]

			if all {
				switch v := interface{}(val).(type) {
				case interface{ ValidateAll() error }:
					if err := v.ValidateAll(); err != nil {
						errors = append(errors, CustomizeValidationError{
							field:  fmt.Sprintf("Configs[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				case interface{ Validate() error }:
					if err := v.Validate(); err != nil {
						errors = append(errors, CustomizeValidationError{
							field:  fmt.Sprintf("Configs[%v]", key),
							reason: "embedded message failed validation",
							cause:  err,
						})
					}
				}
			} else if v, ok := interface{}(val).(interface{ Validate() error }); ok {
				if err := v.Validate(); err != nil {
					return CustomizeValidationError{
						field:  fmt.Sprintf("Configs[%v]", key),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}

		}
	}

	if len(errors) > 0 {
		return CustomizeMultiError(errors)
	}

	return nil
}

// CustomizeMultiError is an error wrapping multiple validation errors returned
// by Customize.ValidateAll() if the designated constraints aren't met.
type CustomizeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CustomizeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CustomizeMultiError) AllErrors() []error { return m }

// CustomizeValidationError is the validation error returned by
// Customize.Validate if the designated constraints aren't met.
type CustomizeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CustomizeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CustomizeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CustomizeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CustomizeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CustomizeValidationError) ErrorName() string { return "CustomizeValidationError" }

// Error satisfies the builtin error interface
func (e CustomizeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomize.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CustomizeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CustomizeValidationError{}

// Validate checks the field values on Customize_Config with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Customize_Config) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Customize_Config with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Customize_ConfigMultiError, or nil if none found.
func (m *Customize_Config) ValidateAll() error {
	return m.validate(true)
}

func (m *Customize_Config) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Enabled

	// no validation rules for Type

	if all {
		switch v := interface{}(m.GetValue()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Customize_ConfigValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Customize_ConfigValidationError{
					field:  "Value",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetValue()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Customize_ConfigValidationError{
				field:  "Value",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Customize_ConfigMultiError(errors)
	}

	return nil
}

// Customize_ConfigMultiError is an error wrapping multiple validation errors
// returned by Customize_Config.ValidateAll() if the designated constraints
// aren't met.
type Customize_ConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Customize_ConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Customize_ConfigMultiError) AllErrors() []error { return m }

// Customize_ConfigValidationError is the validation error returned by
// Customize_Config.Validate if the designated constraints aren't met.
type Customize_ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Customize_ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Customize_ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Customize_ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Customize_ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Customize_ConfigValidationError) ErrorName() string { return "Customize_ConfigValidationError" }

// Error satisfies the builtin error interface
func (e Customize_ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCustomize_Config.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Customize_ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Customize_ConfigValidationError{}