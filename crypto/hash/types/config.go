/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

import (
	"encoding/json"
	"fmt"
)

// Config represents the configuration for hash algorithms
type Config struct {
	SaltLength int               `env:"HASH_SALTLENGTH"`
	Params     map[string]string `env:"HASH_PARAMS"`
}

func (c *Config) String() string {
	b, err := json.Marshal(c.Params)
	if err != nil {
		return fmt.Sprintf("SaltLength: %d, Params: (unmarshallable: %v)", c.SaltLength, err)
	}
	return fmt.Sprintf("SaltLength: %d, Params: %s", c.SaltLength, string(b))
}

// ConfigFromHashParts creates a Config from a HashParts object.
func ConfigFromHashParts(parts *HashParts) *Config {
	return &Config{
		SaltLength: len(parts.Salt),
		Params:     parts.Params,
	}
}

// DefaultConfig return to the default configuration
func DefaultConfig() *Config {
	return &Config{
		SaltLength: DefaultSaltLength, // Default salt length
		Params:     make(map[string]string),
	}
}
