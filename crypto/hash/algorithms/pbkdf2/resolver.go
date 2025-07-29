/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package pbkdf2 implements the functions, types, and interfaces for the module.
package pbkdf2

import (
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// ResolveType resolves the Type for PBKDF2, providing a default underlying hash if not specified.
func ResolveType(algType types.Type) (types.Type, error) {
	if algType.Underlying == "" {
		algType.Underlying = constants.SHA256 // Default to SHA256 for PBKDF2
	}
	// Validate the underlying hash algorithm
	_, err := stdhash.ParseHash(algType.Underlying)
	if err != nil {
		return types.Type{}, fmt.Errorf("unsupported underlying hash for PBKDF2: %s", algType.Underlying)
	}
	return algType, nil
}
