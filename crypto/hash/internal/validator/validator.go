/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package validator implements the functions, types, and interfaces for the module.
package validator

import (
	"github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type Validator[T interfaces.Params] struct {
	config *types.Config
	params T
}

func (v *Validator[T]) Validate(config *types.Config) error {
	// 1) Use internal config if provided config is nil
	if config == nil {
		if v.config == nil {
			v.config = types.DefaultConfig()
		}
	} else {
		v.config = config
	}

	// 2) Decode params if needed
	if v.config.ParamConfig != "" {
		params, err := codec.DecodeParams(v.config.ParamConfig)
		if err != nil {
			return err
		}
		if v.params == nil {
			v.params = *new(T)
		}
		if err := v.params.FromMap(params); err != nil {
			return err
		}
	}

	// 3) Skip validation if params is nil
	if v.params == nil {
		return nil
	}
	return v.params.Validate(v.config)
}

func (v *Validator[T]) Params() T {
	return v.params
}

func (v *Validator[T]) Config() *types.Config {
	return v.config
}

func (v *Validator[T]) WithConfig(cfg *types.Config) *Validator[T] {
	v.config = cfg
	return v
}

func (v *Validator[T]) Apply(opts ...func(*Validator[T])) *Validator[T] {
	for _, opt := range opts {
		opt(v)
	}
	return v
}

func WithParams[T interfaces.Params](params T) func(*Validator[T]) {
	return func(v *Validator[T]) {
		v.params = params
	}
}

func WithConfig[T interfaces.Params](cfg *types.Config) func(*Validator[T]) {
	return func(v *Validator[T]) {
		v.config = cfg
	}
}

func New[T interfaces.Params](params T, opts ...func(*Validator[T])) *Validator[T] {
	return (&Validator[T]{
		params: params,
	}).Apply(opts...)
}
