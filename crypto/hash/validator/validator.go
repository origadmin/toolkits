/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package validator provides a common interface for algorithm-specific parameters.
package validator

import (
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type ValidateFunc func(cfg *types.Config) error

func (v ValidateFunc) Validate(cfg *types.Config) error {
	return v(cfg)
}

type Validator interface {
	Validate(cfg *types.Config) error
}

// Parameters defines the interface for algorithm-specific parameter validation and handling.
// This interface needs to be implemented by parameter structs for different hashing algorithms.
type Parameters interface {
	// Validate checks if the parameters are valid within the context of the given configuration.
	Validate(cfg *types.Config) error
	// FromMap populates the parameter struct from a map of key-value strings.
	// IMPORTANT: This method must have a pointer receiver to ensure that the calling
	// struct is modified. This also ensures that only pointer types satisfy this interface.
	FromMap(params map[string]string) error
	// ToMap converts the parameter struct back into a map of key-value strings.
	ToMap() map[string]string
	// String returns a string representation of the parameters.
	String() string
	// IsNil checks if the underlying pointer is nil. This method must also have
	// a pointer receiver to be invoked on a nil receiver without panicking.
	IsNil() bool
}

// Validated holds the results of a successful validation operation on parameters.
type Validated[T Parameters] struct {
	// Params holds the validated and populated algorithm-specific parameters.
	// It will be the zero value of its type (e.g., nil) if the input was nil.
	Params T
	// Config holds the validated and potentially default-filled configuration.
	Config *types.Config
}

// validateConfig performs basic validation on the configuration.
func validateConfig(cfg *types.Config) error {
	if cfg.SaltLength < 8 {
		return fmt.Errorf("invalid salt length: %d, must be at least 8", cfg.SaltLength)
	}
	return nil
}

// ValidateConfig validates the configuration without any algorithm-specific parameters.
// It returns a validated and potentially default-filled configuration.
func ValidateConfig(cfg *types.Config) (*types.Config, error) {
	// Use default config if the provided one is nil.
	if cfg == nil {
		cfg = types.DefaultConfig()
	}

	if err := validateConfig(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

// ValidateParams decodes the parameter string from the configuration, populates the
// provided parameters object, and then runs the algorithm-specific validation.
// It returns a Validated struct containing the final parameters and configuration.
// The provided parameter 'p' must be a pointer to a struct that implements the
// Parameters interface. Use ValidateConfig for parameter-less validation.
func ValidateParams[T Parameters](cfg *types.Config, p T) (*Validated[T], error) {
	// Use default config if the provided one is nil.
	if cfg == nil {
		cfg = types.DefaultConfig()
	}

	// If p is a typed nil, perform basic config validation and return.
	// This relies on the IsNil() method being implemented with a pointer receiver.
	if p.IsNil() {
		if err := validateConfig(cfg); err != nil {
			return nil, err
		}
		var zero T // zero will be the typed nil
		return &Validated[T]{Params: zero, Config: cfg}, nil
	}

	// If a parameter string is present in the config, decode and apply it to p.
	if cfg.ParamConfig != "" {
		paramsMap, err := codec.DecodeParams(cfg.ParamConfig)
		if err != nil {
			return nil, err
		}

		if err := p.FromMap(paramsMap); err != nil {
			return nil, err
		}
	}

	// Run the final, algorithm-specific validation logic on p.
	if err := p.Validate(cfg); err != nil {
		return nil, err
	}

	return &Validated[T]{Params: p, Config: cfg}, nil
}

// WithValidator runs the provided Validator on the configuration.
func WithValidator(cfg *types.Config, v func(cfg *types.Config) error) (*types.Config, error) {
	// Use default config if the provided one is nil.
	if cfg == nil {
		cfg = types.DefaultConfig()
	}

	if err := v(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
