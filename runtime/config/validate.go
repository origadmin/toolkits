/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package config implements the functions, types, and interfaces for the module.
package config

import (
	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/proto"

	"github.com/origadmin/toolkits/errors"
)

// Validator is an interface for validating protobuf messages.
type Validator interface {
	Validate(message proto.Message) error
}

// validate is a struct that implements the Validator interface.
type validate struct {
	v *protovalidate.Validator
}

// Validate validates a protobuf message.
func (v validate) Validate(message proto.Message) error {
	return v.Validate(message)
}

// NewValidate creates a new Validator.
func NewValidate(opts ...ProtoValidatorOption) (Validator, error) {
	v, err := NewProtoValidate(opts...)
	if err != nil {
		return nil, errors.Wrap(err, "failed to initialize validator")
	}

	return &validate{
		v: v,
	}, nil
}
