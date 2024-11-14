// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/v1/registry.proto

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

// Validate checks the field values on Registry with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Registry) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegistryMultiError, or nil
// if none found.
func (m *Registry) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for ServiceName

	// no validation rules for Debug

	if m.Consul != nil {

		if all {
			switch v := interface{}(m.GetConsul()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RegistryValidationError{
						field:  "Consul",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegistryValidationError{
						field:  "Consul",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetConsul()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegistryValidationError{
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
					errors = append(errors, RegistryValidationError{
						field:  "Etcd",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegistryValidationError{
						field:  "Etcd",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEtcd()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegistryValidationError{
					field:  "Etcd",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if m.Custom != nil {

		if all {
			switch v := interface{}(m.GetCustom()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RegistryValidationError{
						field:  "Custom",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegistryValidationError{
						field:  "Custom",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetCustom()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegistryValidationError{
					field:  "Custom",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RegistryMultiError(errors)
	}

	return nil
}

// RegistryMultiError is an error wrapping multiple validation errors returned
// by Registry.ValidateAll() if the designated constraints aren't met.
type RegistryMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegistryMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegistryMultiError) AllErrors() []error { return m }

// RegistryValidationError is the validation error returned by
// Registry.Validate if the designated constraints aren't met.
type RegistryValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegistryValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegistryValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegistryValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegistryValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegistryValidationError) ErrorName() string { return "RegistryValidationError" }

// Error satisfies the builtin error interface
func (e RegistryValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegistryValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegistryValidationError{}

// Validate checks the field values on Registry_Consul with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Registry_Consul) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry_Consul with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Registry_ConsulMultiError, or nil if none found.
func (m *Registry_Consul) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry_Consul) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Address

	// no validation rules for Scheme

	// no validation rules for Token

	// no validation rules for HeartBeat

	// no validation rules for HealthCheck

	// no validation rules for Datacenter

	// no validation rules for HealthCheckInterval

	if all {
		switch v := interface{}(m.GetTimeout()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Registry_ConsulValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Registry_ConsulValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Registry_ConsulValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DeregisterCriticalServiceAfter

	if len(errors) > 0 {
		return Registry_ConsulMultiError(errors)
	}

	return nil
}

// Registry_ConsulMultiError is an error wrapping multiple validation errors
// returned by Registry_Consul.ValidateAll() if the designated constraints
// aren't met.
type Registry_ConsulMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Registry_ConsulMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Registry_ConsulMultiError) AllErrors() []error { return m }

// Registry_ConsulValidationError is the validation error returned by
// Registry_Consul.Validate if the designated constraints aren't met.
type Registry_ConsulValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Registry_ConsulValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Registry_ConsulValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Registry_ConsulValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Registry_ConsulValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Registry_ConsulValidationError) ErrorName() string { return "Registry_ConsulValidationError" }

// Error satisfies the builtin error interface
func (e Registry_ConsulValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry_Consul.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Registry_ConsulValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Registry_ConsulValidationError{}

// Validate checks the field values on Registry_ETCD with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Registry_ETCD) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry_ETCD with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in Registry_ETCDMultiError, or
// nil if none found.
func (m *Registry_ETCD) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry_ETCD) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return Registry_ETCDMultiError(errors)
	}

	return nil
}

// Registry_ETCDMultiError is an error wrapping multiple validation errors
// returned by Registry_ETCD.ValidateAll() if the designated constraints
// aren't met.
type Registry_ETCDMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Registry_ETCDMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Registry_ETCDMultiError) AllErrors() []error { return m }

// Registry_ETCDValidationError is the validation error returned by
// Registry_ETCD.Validate if the designated constraints aren't met.
type Registry_ETCDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Registry_ETCDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Registry_ETCDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Registry_ETCDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Registry_ETCDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Registry_ETCDValidationError) ErrorName() string { return "Registry_ETCDValidationError" }

// Error satisfies the builtin error interface
func (e Registry_ETCDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry_ETCD.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Registry_ETCDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Registry_ETCDValidationError{}

// Validate checks the field values on Registry_Custom with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *Registry_Custom) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Registry_Custom with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// Registry_CustomMultiError, or nil if none found.
func (m *Registry_Custom) ValidateAll() error {
	return m.validate(true)
}

func (m *Registry_Custom) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetConfig()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, Registry_CustomValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, Registry_CustomValidationError{
					field:  "Config",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetConfig()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return Registry_CustomValidationError{
				field:  "Config",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return Registry_CustomMultiError(errors)
	}

	return nil
}

// Registry_CustomMultiError is an error wrapping multiple validation errors
// returned by Registry_Custom.ValidateAll() if the designated constraints
// aren't met.
type Registry_CustomMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m Registry_CustomMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m Registry_CustomMultiError) AllErrors() []error { return m }

// Registry_CustomValidationError is the validation error returned by
// Registry_Custom.Validate if the designated constraints aren't met.
type Registry_CustomValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Registry_CustomValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Registry_CustomValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Registry_CustomValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Registry_CustomValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Registry_CustomValidationError) ErrorName() string { return "Registry_CustomValidationError" }

// Error satisfies the builtin error interface
func (e Registry_CustomValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistry_Custom.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Registry_CustomValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Registry_CustomValidationError{}
