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

// Validate checks the field values on RegistryConfig with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegistryConfig) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegistryConfig with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegistryConfigMultiError,
// or nil if none found.
func (m *RegistryConfig) ValidateAll() error {
	return m.validate(true)
}

func (m *RegistryConfig) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Name

	if m.Consul != nil {

		if all {
			switch v := interface{}(m.GetConsul()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, RegistryConfigValidationError{
						field:  "Consul",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegistryConfigValidationError{
						field:  "Consul",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetConsul()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegistryConfigValidationError{
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
					errors = append(errors, RegistryConfigValidationError{
						field:  "Etcd",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, RegistryConfigValidationError{
						field:  "Etcd",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetEtcd()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return RegistryConfigValidationError{
					field:  "Etcd",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return RegistryConfigMultiError(errors)
	}

	return nil
}

// RegistryConfigMultiError is an error wrapping multiple validation errors
// returned by RegistryConfig.ValidateAll() if the designated constraints
// aren't met.
type RegistryConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegistryConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegistryConfigMultiError) AllErrors() []error { return m }

// RegistryConfigValidationError is the validation error returned by
// RegistryConfig.Validate if the designated constraints aren't met.
type RegistryConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegistryConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegistryConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegistryConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegistryConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegistryConfigValidationError) ErrorName() string { return "RegistryConfigValidationError" }

// Error satisfies the builtin error interface
func (e RegistryConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistryConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegistryConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegistryConfigValidationError{}

// Validate checks the field values on RegistryConfig_Consul with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegistryConfig_Consul) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegistryConfig_Consul with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegistryConfig_ConsulMultiError, or nil if none found.
func (m *RegistryConfig_Consul) ValidateAll() error {
	return m.validate(true)
}

func (m *RegistryConfig_Consul) validate(all bool) error {
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
				errors = append(errors, RegistryConfig_ConsulValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, RegistryConfig_ConsulValidationError{
					field:  "Timeout",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimeout()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return RegistryConfig_ConsulValidationError{
				field:  "Timeout",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for DeregisterCriticalServiceAfter

	if len(errors) > 0 {
		return RegistryConfig_ConsulMultiError(errors)
	}

	return nil
}

// RegistryConfig_ConsulMultiError is an error wrapping multiple validation
// errors returned by RegistryConfig_Consul.ValidateAll() if the designated
// constraints aren't met.
type RegistryConfig_ConsulMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegistryConfig_ConsulMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegistryConfig_ConsulMultiError) AllErrors() []error { return m }

// RegistryConfig_ConsulValidationError is the validation error returned by
// RegistryConfig_Consul.Validate if the designated constraints aren't met.
type RegistryConfig_ConsulValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegistryConfig_ConsulValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegistryConfig_ConsulValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegistryConfig_ConsulValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegistryConfig_ConsulValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegistryConfig_ConsulValidationError) ErrorName() string {
	return "RegistryConfig_ConsulValidationError"
}

// Error satisfies the builtin error interface
func (e RegistryConfig_ConsulValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistryConfig_Consul.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegistryConfig_ConsulValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegistryConfig_ConsulValidationError{}

// Validate checks the field values on RegistryConfig_ETCD with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RegistryConfig_ETCD) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegistryConfig_ETCD with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegistryConfig_ETCDMultiError, or nil if none found.
func (m *RegistryConfig_ETCD) ValidateAll() error {
	return m.validate(true)
}

func (m *RegistryConfig_ETCD) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return RegistryConfig_ETCDMultiError(errors)
	}

	return nil
}

// RegistryConfig_ETCDMultiError is an error wrapping multiple validation
// errors returned by RegistryConfig_ETCD.ValidateAll() if the designated
// constraints aren't met.
type RegistryConfig_ETCDMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegistryConfig_ETCDMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegistryConfig_ETCDMultiError) AllErrors() []error { return m }

// RegistryConfig_ETCDValidationError is the validation error returned by
// RegistryConfig_ETCD.Validate if the designated constraints aren't met.
type RegistryConfig_ETCDValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegistryConfig_ETCDValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegistryConfig_ETCDValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegistryConfig_ETCDValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegistryConfig_ETCDValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegistryConfig_ETCDValidationError) ErrorName() string {
	return "RegistryConfig_ETCDValidationError"
}

// Error satisfies the builtin error interface
func (e RegistryConfig_ETCDValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegistryConfig_ETCD.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegistryConfig_ETCDValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegistryConfig_ETCDValidationError{}
