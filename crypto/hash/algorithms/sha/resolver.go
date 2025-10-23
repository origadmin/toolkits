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
	
	// Map generic SHA3 name to default version SHA3-256
	switch algType.Name {
	case types.SHA3:
		algType.Name = types.SHA3_256
	}
	
	return algType, nil
}