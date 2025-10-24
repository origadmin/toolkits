/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package argon2 implements the functions, types, and interfaces for the module.
package argon2

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func ResolveSpec(p types.Spec) (types.Spec, error) {
	if p.Name == types.ARGON2 {
		return types.Spec{Name: types.ARGON2i, Underlying: ""}, nil
	}
	return p, nil
}
