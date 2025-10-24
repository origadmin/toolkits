/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package blake2 implements the functions, types, and interfaces for the module.
package blake2

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func ResolveSpec(p types.Spec) (types.Spec, error) {
	p.Name = p.String()
	p.Underlying = ""
	switch p.Name {
	case types.BLAKE2b:
		p.Name = types.BLAKE2b_512
	case types.BLAKE2s:
		p.Name = types.BLAKE2s_256
	}
	return p, nil
}
