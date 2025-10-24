// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Option is a function that modifies a Config
type Option func(*types.Config)

// WithSaltLength sets the salt length
func WithSaltLength(length int) Option {
	return func(cfg *types.Config) {
		cfg.SaltLength = types.DefaultSaltLength // Default salt length
	}
}

// WithParams sets the parameters for the hash algorithm using a type that implements fmt.Stringer.
func WithParams(stringer fmt.Stringer) Option {
	return func(cfg *types.Config) {
		if stringer == nil {
			return
		}
		cfg.ParamConfig = stringer.String()
	}
}

// WithEncodedParams sets the parameters for the hash algorithm by encoding a map[string]string
// using the provided ParamEncoderFunc. This is used when the parameters are in map format
// and need to be converted to a string by an external encoder (e.g., from the codec package).
func WithEncodedParams(params map[string]string, encoder types.ParamEncoderFunc) Option {
	return func(cfg *types.Config) {
		if params == nil || encoder == nil {
			return
		}
		cfg.ParamConfig = encoder(params)
	}
}

// WithParamConfig allows direct manipulation of the ParamConfig string.
func WithParamConfig(fn func(string) string) Option {
	return func(cfg *types.Config) {
		cfg.ParamConfig = fn(cfg.ParamConfig)
	}
}
