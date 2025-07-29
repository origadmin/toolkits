/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package blake2 implements the functions, types, and interfaces for the module.
package blake2

import (
	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func ResolveType(p types.Type) (types.Type, error) {
	p.Name = p.String()
	p.Underlying = ""
	switch p.Name {
	case constants.BLAKE2b:
		p.Name = constants.BLAKE2b_512
	case constants.BLAKE2s:
		p.Name = constants.BLAKE2s_256
	}
	return p, nil
}
