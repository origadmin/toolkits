/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package ripemd160 implements the functions, types, and interfaces for the module.
package ripemd160

import (
	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// ResolveType resolves the Type for PBKDF2, providing a default underlying hash if not specified.
func ResolveType(algType types.Type) (types.Type, error) {
	algType.Name = constants.RIPEMD160
	algType.Underlying = ""
	return algType, nil
}
