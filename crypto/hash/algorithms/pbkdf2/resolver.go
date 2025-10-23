/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package pbkdf2 implements the functions, types, and interfaces for the module.
package pbkdf2

import (
	"fmt"
	"strings"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// ResolveType resolves the Type for PBKDF2, providing a default underlying hash if not specified.
func ResolveType(algType types.Type) (types.Type, error) {
	// If the name is a composite HMAC type (e.g., "hmac-sha256"), parse it.
	// This handles cases where types.NewType might not fully parse the composite name into Name and Underlying.
	if strings.HasPrefix(algType.Name, types.PBKDF2_PREFIX) {
		algType.Underlying = strings.TrimPrefix(algType.Name, types.PBKDF2_PREFIX)
		algType.Name = types.PBKDF2
	}
	if algType.Underlying == "" {
		algType.Underlying = types.SHA256 // Default to SHA256 for PBKDF2
	}

	resolvedUnderlying := algType.Underlying
	if strings.HasPrefix(algType.Underlying, types.HMAC_PREFIX) {
		resolvedUnderlying = strings.TrimPrefix(algType.Underlying, types.HMAC_PREFIX)
	}

	// Validate the underlying hash algorithm
	hashHash, err := stdhash.ParseHash(resolvedUnderlying)
	if err != nil {
		return types.Type{}, fmt.Errorf("unsupported underlying hash for PBKDF2: %s", algType.Underlying)
	}

	// Explicitly check for unsuitable hash types for PBKDF2 (non-cryptographic or weak hashes)
	switch hashHash {
	case stdhash.MAPHASH, stdhash.ADLER32, stdhash.CRC32, stdhash.CRC32_ISO, stdhash.CRC32_CAST, stdhash.CRC32_KOOP,
		stdhash.CRC64_ISO, stdhash.CRC64_ECMA, stdhash.FNV32, stdhash.FNV32A, stdhash.FNV64, stdhash.FNV64A,
		stdhash.FNV128, stdhash.FNV128A:
		return types.Type{}, errors.ErrUnsupportedHashForPBKDF2 // Assuming this error exists or needs to be created
	default:
	}

	// Update algType.Underlying to the resolved (stripped) version if it was an HMAC composite
	// Comment out or remove this line if not needed
	//algType.Underlying = resolvedUnderlying

	return algType, nil
}
