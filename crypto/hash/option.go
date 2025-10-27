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
		cfg.SaltLength = types.DefaultSaltLength // Default salt length
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

// WithEncodedParams sets the parameters for the hash algorithm by encoding a map[string]string
// using the provided ParamEncoderFunc. This is used when the parameters are in map format
// and need to be converted to a string by an external encoder (e.g., from the codec package).
func WithEncodedParams(params map[string]string) Option {
	return func(cfg *types.Config) {
		if len(params) == 0 {
			return
		}
		cfg.Params = params
	}
}

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
