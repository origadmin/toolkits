/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package validator implements the functions, types, and interfaces for the module.
package validator

import (
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type Validator[T interfaces.Params] struct {
	params T
}

func (v Validator[T]) Validate(config *types.Config) error {
	if config.ParamConfig == "" {
		return v.params.Validate(config)
	}
	return v.params.Validate(config)
}

func (v Validator[T]) Params() T {
	return v.params
}

func WithParams[T interfaces.Params](params T) *Validator[T] {
	return &Validator[T]{
		params: params,
	}
}
