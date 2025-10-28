// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Option is a function that modifies a Config
type Option func(*types.Config)

// WithSaltLength sets the salt length
func WithSaltLength(length int) Option {
	return func(cfg *types.Config) {
		cfg.SaltLength = length // Use the provided length, not DefaultSaltLength
	}
}

// WithParamString sets the parameters for the hash algorithm using a type that implements fmt.Stringer.
func WithParamString(params string) Option {
	return func(cfg *types.Config) {
		decoded, err := codec.DecodeParams(params)
		if err == nil {
			cfg.Params = decoded
		}
	}
}

// WithParams sets the parameters for the hash algorithm directly from a map[string]string.
// This is the preferred way to set algorithm-specific parameters when they are already in map format.
func WithParams(params map[string]string) Option {
	return func(cfg *types.Config) {
		if len(params) == 0 {
			return
		}
		// Directly assign the map, as Config.Params is already map[string]string
		cfg.Params = params
	}
}

// WithHashParts sets the salt length and parameters from a HashParts object.
func WithHashParts(parts *types.HashParts) Option {
	return func(cfg *types.Config) {
		if parts == nil {
			return
		}
		cfg.SaltLength = len(parts.Salt)
		cfg.Params = parts.Params
	}
}

// ConfigFromHashParts creates a Config from a HashParts object.
func ConfigFromHashParts(parts *types.HashParts) *types.Config {
	return &types.Config{
		SaltLength: len(parts.Salt),
		Params:     parts.Params,
	}
}

// DefaultConfig return to the default configuration
func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: types.DefaultSaltLength, // Default salt length
		Params:     make(map[string]string),
	}
}
