/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package crc implements the functions, types, and interfaces for the module.
package crc

import (
	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

func ResolveType(p types.Type) (types.Type, error) {
	p.Name = p.String()
	p.Underlying = ""
	switch p.Name {
	case constants.CRC32:
		p.Name = constants.CRC32_ISO
	case constants.CRC64:
		p.Name = constants.CRC64_ISO
	}
	return p, nil
}
