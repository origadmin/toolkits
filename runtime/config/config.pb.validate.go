// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: runtime/config/config.proto

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

// Validate checks the field values on SourceConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SourceConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SourceConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SourceConfigMultiError, or
// nil if none found.
func (m *SourceConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *SourceConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if _, ok := _SourceConfig_Type_InLookup[m.GetType()]; !ok {
		err := SourceConfigValidationError{
			field:  "Type",
			reason: "value must be in list [none consul etcd nacos apollo kubernetes polaris]",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for Name

	// no validation rules for EnvArgs

	if m.File != nil {

		if all {
			switch v := interface{}(m.GetFile()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SourceConfigValidationError{
						field:  "File",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SourceConfigValidationError{
						field:  "File",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetFile()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SourceConfigValidationError{
					field:  "File",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Consul != nil {

		if all {
			switch v := interface{}(m.GetConsul()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SourceConfigValidationError{
						field:  "Consul",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SourceConfigValidationError{
						field:  "Consul",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetConsul()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SourceConfigValidationError{
					field:  "Consul",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Etcd != nil {

		if all {
			switch v := interface{}(m.GetEtcd()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, SourceConfigValidationError{
						field:  "Etcd",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, SourceConfigValidationError{
						field:  "Etcd",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEtcd()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return SourceConfigValidationError{
					field:  "Etcd",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return SourceConfigMultiError(errors)
	}

	return nil
}

// SourceConfigMultiError is an error wrapping multiple validation errors
// returned by SourceConfig.ValidateAll() if the designated constraints aren't met.
type SourceConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SourceConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SourceConfigMultiError) AllErrors() []error { return m }

// SourceConfigValidationError is the validation error returned by
// SourceConfig.Validate if the designated constraints aren't met.
type SourceConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SourceConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SourceConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SourceConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SourceConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SourceConfigValidationError) ErrorName() string { return "SourceConfigValidationError" }

// Error satisfies the builtin error interface
func (e SourceConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSourceConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SourceConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SourceConfigValidationError{}

var _SourceConfig_Type_InLookup = map[string]struct{}{
	"none":       {},
	"consul":     {},
	"etcd":       {},
	"nacos":      {},
	"apollo":     {},
	"kubernetes": {},
	"polaris":    {},
}

// Validate checks the field values on SourceConfig_File with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SourceConfig_File) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SourceConfig_File with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SourceConfig_FileMultiError, or nil if none found.
func (m *SourceConfig_File) ValidateAll() error {
	return m.validate(true)
}

func (m *SourceConfig_File) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Path

	if len(errors) > 0 {
		return SourceConfig_FileMultiError(errors)
	}

	return nil
}

// SourceConfig_FileMultiError is an error wrapping multiple validation errors
// returned by SourceConfig_File.ValidateAll() if the designated constraints
// aren't met.
type SourceConfig_FileMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SourceConfig_FileMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SourceConfig_FileMultiError) AllErrors() []error { return m }

// SourceConfig_FileValidationError is the validation error returned by
// SourceConfig_File.Validate if the designated constraints aren't met.
type SourceConfig_FileValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SourceConfig_FileValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SourceConfig_FileValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SourceConfig_FileValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SourceConfig_FileValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SourceConfig_FileValidationError) ErrorName() string {
	return "SourceConfig_FileValidationError"
}

// Error satisfies the builtin error interface
func (e SourceConfig_FileValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSourceConfig_File.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SourceConfig_FileValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SourceConfig_FileValidationError{}

// Validate checks the field values on SourceConfig_Consul with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *SourceConfig_Consul) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SourceConfig_Consul with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SourceConfig_ConsulMultiError, or nil if none found.
func (m *SourceConfig_Consul) ValidateAll() error {
	return m.validate(true)
}

func (m *SourceConfig_Consul) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for Scheme

	// no validation rules for Token

	// no validation rules for Path

	if len(errors) > 0 {
		return SourceConfig_ConsulMultiError(errors)
	}

	return nil
}

// SourceConfig_ConsulMultiError is an error wrapping multiple validation
// errors returned by SourceConfig_Consul.ValidateAll() if the designated
// constraints aren't met.
type SourceConfig_ConsulMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SourceConfig_ConsulMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SourceConfig_ConsulMultiError) AllErrors() []error { return m }

// SourceConfig_ConsulValidationError is the validation error returned by
// SourceConfig_Consul.Validate if the designated constraints aren't met.
type SourceConfig_ConsulValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SourceConfig_ConsulValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SourceConfig_ConsulValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SourceConfig_ConsulValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SourceConfig_ConsulValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SourceConfig_ConsulValidationError) ErrorName() string {
	return "SourceConfig_ConsulValidationError"
}

// Error satisfies the builtin error interface
func (e SourceConfig_ConsulValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSourceConfig_Consul.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SourceConfig_ConsulValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SourceConfig_ConsulValidationError{}

// Validate checks the field values on SourceConfig_ETCD with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SourceConfig_ETCD) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SourceConfig_ETCD with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SourceConfig_ETCDMultiError, or nil if none found.
func (m *SourceConfig_ETCD) ValidateAll() error {
	return m.validate(true)
}

func (m *SourceConfig_ETCD) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SourceConfig_ETCDMultiError(errors)
	}

	return nil
}

// SourceConfig_ETCDMultiError is an error wrapping multiple validation errors
// returned by SourceConfig_ETCD.ValidateAll() if the designated constraints
// aren't met.
type SourceConfig_ETCDMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SourceConfig_ETCDMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SourceConfig_ETCDMultiError) AllErrors() []error { return m }

// SourceConfig_ETCDValidationError is the validation error returned by
// SourceConfig_ETCD.Validate if the designated constraints aren't met.
type SourceConfig_ETCDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SourceConfig_ETCDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SourceConfig_ETCDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SourceConfig_ETCDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SourceConfig_ETCDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SourceConfig_ETCDValidationError) ErrorName() string {
	return "SourceConfig_ETCDValidationError"
}

// Error satisfies the builtin error interface
func (e SourceConfig_ETCDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSourceConfig_ETCD.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SourceConfig_ETCDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SourceConfig_ETCDValidationError{}
