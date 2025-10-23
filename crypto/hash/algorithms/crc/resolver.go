/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package crc implements the functions, types, and interfaces for the module.
package crc

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func ResolveType(p types.Type) (types.Type, error) {
	p.Name = p.String()
	p.Underlying = ""
	switch p.Name {
	case types.CRC32:
		p.Name = types.CRC32_ISO
	case types.CRC64:
		p.Name = types.CRC64_ISO
	}
	return p, nil
}
