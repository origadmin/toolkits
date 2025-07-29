/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hmac implements the functions, types, and interfaces for the module.
package hmac

import (
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// ResolveType resolves the Type for HMAC, providing a default underlying hash if not specified.
func ResolveType(p types.Type) (types.Type, error) {
	if p.Underlying == "" {
		p.Underlying = constants.SHA256 // Default to SHA256 for HMAC
	}

	// Validate the underlying hash algorithm
	hashHash, err := stdhash.ParseHash(p.Underlying)
	if err != nil {
		return types.Type{}, fmt.Errorf("unsupported underlying hash for HMAC: %s", p.Underlying)
	}
	// Explicitly check for unsuitable hash types for HMAC
	// MAPHASH, ADLER32, CRC32, FNV are not cryptographically secure and should not be used with HMAC
	switch hashHash {
	case stdhash.MAPHASH, stdhash.ADLER32, stdhash.CRC32, stdhash.CRC32_ISO, stdhash.CRC32_CAST, stdhash.CRC32_KOOP,
		stdhash.CRC64_ISO, stdhash.CRC64_ECMA, stdhash.FNV32, stdhash.FNV32A, stdhash.FNV64, stdhash.FNV64A,
		stdhash.FNV128, stdhash.FNV128A:
		return types.Type{}, errors.ErrUnsupportedHashForHMAC
	default:
	}

	return p, nil
}
