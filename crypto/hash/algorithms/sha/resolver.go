/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package sha implements the functions, types, and interfaces for the module.
package sha

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// ResolveType resolves the Type for PBKDF2, providing a default underlying hash if not specified.
func ResolveType(algType types.Type) (types.Type, error) {
	algType.Name = algType.String()
	algType.Underlying = ""
	return algType, nil
}
