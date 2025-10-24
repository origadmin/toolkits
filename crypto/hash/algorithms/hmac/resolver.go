/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hmac implements the functions, types, and interfaces for the module.
package hmac

import (
	"fmt"
	"strings"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// ResolveSpec resolves the Spec for HMAC, providing a default underlying hash if not specified.
func ResolveSpec(algSpec types.Spec) (types.Spec, error) {
	// If the name is a composite HMAC type (e.g., "hmac-sha256"), parse it.
	// This handles cases where types.New might not fully parse the composite name into Name and Underlying.
	if strings.HasPrefix(algSpec.Name, types.HMAC_PREFIX) {
		algSpec.Underlying = strings.TrimPrefix(algSpec.Name, types.HMAC_PREFIX)
		algSpec.Name = types.HMAC
	}

	if algSpec.Name != types.HMAC {
		return types.Spec{}, fmt.Errorf("hmac: invalid algorithm name: %s", algSpec.Name)
	}

	if algSpec.Underlying == "" {
		algSpec.Underlying = types.SHA256 // Default to SHA256 for HMAC
	}

	// Validate the underlying hash algorithm
	hashHash, err := stdhash.ParseHash(algSpec.Underlying)
	if err != nil {
		return types.Spec{}, fmt.Errorf("unsupported underlying hash for HMAC: %s", algSpec.Underlying)
	}
	// Explicitly check for unsuitable hash types for HMAC
	// MAPHASH, ADLER32, CRC32, FNV are not cryptographically secure and should not be used with HMAC
	switch hashHash {
	case stdhash.MAPHASH, stdhash.ADLER32, stdhash.CRC32, stdhash.CRC32_ISO, stdhash.CRC32_CAST, stdhash.CRC32_KOOP,
		stdhash.CRC64_ISO, stdhash.CRC64_ECMA, stdhash.FNV32, stdhash.FNV32A, stdhash.FNV64, stdhash.FNV64A,
		stdhash.FNV128, stdhash.FNV128A:
		return types.Spec{}, errors.ErrUnsupportedHashForHMAC
	default:
	}

	return algSpec, nil
}
