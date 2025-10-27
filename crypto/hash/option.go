// Package hash implements the functions, types, and interfaces for the module.
package hash

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Option is a function that modifies a Config
type Option func(*types.Config)

// ConfigFromHashParts creates a Config from a HashParts object.
func ConfigFromHashParts(parts *types.HashParts) *types.Config {
	return &types.Config{
		SaltLength: len(parts.Salt),
		Params:     parts.Params,
	}
}
