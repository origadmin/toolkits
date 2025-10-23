/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package argon2 implements the functions, types, and interfaces for the module.
package argon2

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func ResolveType(p types.Type) (types.Type, error) {
	if p.Name == types.ARGON2 {
		return types.Type{Name: types.ARGON2i, Underlying: ""}, nil
	}
	return p, nil
}
