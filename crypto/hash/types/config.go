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

func (c *Config) String() string {
	return fmt.Sprintf("SaltLength: %d, ParamConfig: %s", c.SaltLength, c.ParamConfig)
}

// ParamEncoderFunc defines a function type for encoding parameters (map[string]string) into a string.
// This is used for dependency injection to avoid circular dependencies.
type ParamEncoderFunc func(params map[string]string) string

// DefaultConfig return to the default configuration
func DefaultConfig() *Config {
	return &Config{
		SaltLength: DefaultSaltLength, // Default salt length
	}
}
