// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/v1/metrics.proto

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

// Validate checks the field values on Metric with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Metric) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Metric with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in MetricMultiError, or nil if none found.
func (m *Metric) ValidateAll() error {
	return m.validate(true)
}

func (m *Metric) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	// no validation rules for Address

	if len(errors) > 0 {
		return MetricMultiError(errors)
	}

	return nil
}

// MetricMultiError is an error wrapping multiple validation errors returned by
// Metric.ValidateAll() if the designated constraints aren't met.
type MetricMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetricMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetricMultiError) AllErrors() []error { return m }

// MetricValidationError is the validation error returned by Metric.Validate if
// the designated constraints aren't met.
type MetricValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetricValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetricValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetricValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetricValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetricValidationError) ErrorName() string { return "MetricValidationError" }

// Error satisfies the builtin error interface
func (e MetricValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetric.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetricValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetricValidationError{}
