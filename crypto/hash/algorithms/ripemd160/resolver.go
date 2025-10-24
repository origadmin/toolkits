/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package ripemd160 implements the functions, types, and interfaces for the module.
package ripemd160

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// ResolveSpec resolves the Spec for PBKDF2, providing a default underlying hash if not specified.
func ResolveSpec(algSpec types.Spec) (types.Spec, error) {
	algSpec.Name = types.RIPEMD160
	algSpec.Underlying = ""
	return algSpec, nil
}
