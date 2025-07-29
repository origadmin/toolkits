/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package interfaces implements the functions, types, and interfaces for the module.
package interfaces

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type TypeResolver interface {
	ResolveType(t types.Type) (types.Type, error)
}
type AlgorithmResolver func(t types.Type) (types.Type, error)

func (r AlgorithmResolver) ResolveType(t types.Type) (types.Type, error) {
	return r(t)
}
