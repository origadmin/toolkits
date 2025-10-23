/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

import (
	"fmt" // Needed for fmt.Stringer
)

// Config represents the configuration for hash algorithms
type Config struct {
	SaltLength  int    `env:"HASH_SALTLENGTH"`
	ParamConfig string `env:"HASH_PARAM_CONFIG"`
}

// Option is a function that modifies a Config
type Option func(*Config)

// ParamEncoderFunc defines a function type for encoding parameters (map[string]string) into a string.
// This is used for dependency injection to avoid circular dependencies.
type ParamEncoderFunc func(params map[string]string) string

// WithSaltLength sets the salt length
func WithSaltLength(length int) Option {
	return func(cfg *Config) {
		cfg.SaltLength = DefaultSaltLength // Default salt length
	}
}

// WithParams sets the parameters for the hash algorithm using a type that implements fmt.Stringer.
func WithParams(stringer fmt.Stringer) Option {
	return func(cfg *Config) {
		if stringer == nil {
			return
		}
		cfg.ParamConfig = stringer.String()
	}
}

// WithEncodedParams sets the parameters for the hash algorithm by encoding a map[string]string
// using the provided ParamEncoderFunc. This is used when the parameters are in map format
// and need to be converted to a string by an external encoder (e.g., from the codec package).
func WithEncodedParams(params map[string]string, encoder ParamEncoderFunc) Option {
	return func(cfg *Config) {
		if params == nil || encoder == nil {
			return
		}
		cfg.ParamConfig = encoder(params)
	}
}

// WithParamConfig allows direct manipulation of the ParamConfig string.
func WithParamConfig(fn func(string) string) Option {
	return func(cfg *Config) {
		cfg.ParamConfig = fn(cfg.ParamConfig)
	}
}

// DefaultConfig return to the default configuration
func DefaultConfig() *Config {
	return &Config{
		SaltLength: DefaultSaltLength, // Default salt length
	}
}
