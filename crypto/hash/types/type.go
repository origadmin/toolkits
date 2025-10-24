/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

import (
	"fmt"
	"strings"

	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
)

// ----------------------------------------------------------------------------
// Specs
// ----------------------------------------------------------------------------

// Spec represents a structured hash algorithm specification.
// It separates the main algorithm from its underlying hash function,
// allowing for clear and extensible handling of composite algorithms
// like HMAC and PBKDF2.
type Spec struct {
	// Name is the main algorithm's name, e.g., "hmac", "pbkdf2", "sha256".
	// This is the key field for logical dispatch.
	Name string

	// Underlying is the full string representation of the underlying hash algorithm.
	// For simple hashes, this field is empty.
	// For composite hashes, it specifies the hash function to be used,
	// e.g., "sha256" for "hmac-sha256", or "sha3-512" for "pbkdf2-sha3-512".
	Underlying string
}

// String returns the string representation of the spec.
func (s Spec) String() string {
	if s.Underlying != "" {
		return s.Name + "-" + s.Underlying
	}
	return s.Name
}

// Is compares two Spec instances for equality.
func (s Spec) Is(s2 Spec) bool {
	return s.Name == s2.Name && s.Underlying == s2.Underlying
}

// Parse parses an algorithm string into its structured Spec.
// It handles common aliases and composite algorithm formats.
func Parse(algorithm string) (Spec, error) {
	algorithm = strings.ToLower(algorithm)

	parts := strings.SplitN(algorithm, "-", 2)
	var s Spec
	if len(parts) == 2 {
		// This is a composite algorithm like "hmac-sha256" or "pbkdf2-sha512"
		s = Spec{Name: parts[0], Underlying: parts[1]}
	} else {
		// This is a simple algorithm like "sha256"
		s = Spec{Name: algorithm}
	}

	// If no specific resolver is registered, return the parsed spec as is.
	return s, nil
}

// New creates a new Spec instance with the specified name and underlying hash.
func New(name string, underlying ...string) Spec {
	s := Spec{Name: name}
	if len(underlying) > 0 {
		s.Underlying = underlying[0]
	}
	return s
}

// Hash is a helper function that might need to be refactored
// depending on how stdhash.ParseHash is updated to handle the new Spec struct.
// For now, it assumes subAlg is a simple string.
func Hash(alg string) (stdhash.Hash, error) {
	h, err := stdhash.ParseHash(alg)
	if err != nil {
		return 0, fmt.Errorf("unsupported hash type: %s", alg)
	}
	return h, nil
}
