/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package interfaces implements the functions, types, and interfaces for the module.
package scheme

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type SpecResolver interface {
	ResolveSpec(algSpec types.Spec) (types.Spec, error)
}
type AlgorithmResolver func(algSpec types.Spec) (types.Spec, error)

func (r AlgorithmResolver) ResolveSpec(algSpec types.Spec) (types.Spec, error) {
	return r(algSpec)
}
