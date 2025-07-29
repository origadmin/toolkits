/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

import (
	"fmt"
	"strings"

	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
)

// AlgorithmResolver defines an interface for resolving and normalizing algorithm types.
type AlgorithmResolver interface {
	ResolveType(t Type) (Type, error)
}

// Type represents a structured hash algorithm definition.
// It separates the main algorithm from its underlying hash function,
// allowing for clear and extensible handling of composite algorithms
// like HMAC and PBKDF2.
type Type struct {
	// Name is the main algorithm's name, e.g., "hmac", "pbkdf2", "sha256".
	// This is the key field for logical dispatch.
	Name string

	// Underlying is the full string representation of the underlying hash algorithm.
	// For simple hashes, this field is empty.
	// For composite hashes, it specifies the hash function to be used,
	// e.g., "sha256" for "hmac-sha256", or "sha3-512" for "pbkdf2-sha3-512".
	Underlying string
}

// String returns the string representation of the type
func (t Type) String() string {
	if t.Underlying != "" {
		return t.Name + "-" + t.Underlying
	}
	return t.Name
}

// Is compares two Type instances for equality.
func (t Type) Is(t2 Type) bool {
	return t.Name == t2.Name && t.Underlying == t2.Underlying
}

// ParseType parses an algorithm string into its structured Type.
// It handles common aliases and composite algorithm formats.
func ParseType(algorithm string) (Type, error) {
	algorithm = strings.ToLower(algorithm)

	parts := strings.SplitN(algorithm, "-", 2)
	var t Type
	if len(parts) == 2 {
		// This is a composite algorithm like "hmac-sha256" or "pbkdf2-sha512"
		t = Type{Name: parts[0], Underlying: parts[1]}
	} else {
		// This is a simple algorithm like "sha256"
		t = Type{Name: algorithm}
	}

	// If no specific resolver is registered, return the parsed type as is.
	return t, nil
}

// NewType creates a new Type instance with the specified name and underlying hash.
func NewType(name string, underlying ...string) Type {
	t := Type{Name: name}
	if len(underlying) > 0 {
		t.Underlying = underlying[0]
	}
	return t
}

// TypeHash is a helper function that might need to be refactored
// depending on how stdhash.ParseHash is updated to handle the new Type struct.
// For now, it assumes subAlg is a simple string.
func TypeHash(subAlg string) (stdhash.Hash, error) {
	h, err := stdhash.ParseHash(subAlg)
	if err != nil {
		return 0, fmt.Errorf("unsupported hash type: %s", subAlg)
	}
	return h, nil
}
