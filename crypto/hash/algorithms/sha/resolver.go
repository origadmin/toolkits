/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package sha implements the functions, types, and interfaces for the module.
package sha

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// ResolveSpec resolves the Spec for PBKDF2, providing a default underlying hash if not specified.
func ResolveSpec(algSpec types.Spec) (types.Spec, error) {
	algSpec.Name = algSpec.String()
	algSpec.Underlying = ""

	// Map generic SHA3 name to default version SHA3-256
	switch algSpec.Name {
	case types.SHA3:
		algSpec.Name = types.SHA3_256
	}

	return algSpec, nil
}
