/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

import (
	"fmt"
	"strings"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
)

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

// ParseAlgorithm parses an algorithm string into its structured Type.
// It handles common aliases and composite algorithm formats.
func ParseAlgorithm(typedAlgorithm string) (Type, error) {
	typedAlgorithm = strings.ToLower(typedAlgorithm)

	parts := strings.SplitN(typedAlgorithm, "-", 2)
	if len(parts) == 2 {
		// This is a composite algorithm like "hmac-sha256" or "pbkdf2-sha512"
		return Type{Name: parts[0], Underlying: parts[1]}, nil
	}

	// This is a simple algorithm like "sha256"
	return Type{Name: typedAlgorithm}, nil
}

// AlgorithmTypeHash is a helper function that might need to be refactored
// depending on how stdhash.ParseHash is updated to handle the new Type struct.
// For now, it assumes subAlg is a simple string.
func AlgorithmTypeHash(subAlg string) (stdhash.Hash, error) {
	h, err := stdhash.ParseHash(subAlg)
	if err != nil {
		return 0, fmt.Errorf("unsupported hash type: %s", subAlg)
	}
	return h, nil
}

func Argon2() Type {
	return Type{Name: "argon2"}
}

func Bcrypt() Type {
	return Type{Name: "bcrypt"}
}

func Blake2b() Type {
	return Type{Name: "blake2b"}
}

func Blake2s() Type {
	return Type{Name: "blake2s"}
}

func Scrypt() Type {
	return Type{Name: constants.SCRYPT}
}

func NewHashParts(p Type, salt []byte, hash []byte) *HashParts {
	return &HashParts{
		Algorithm: p.String(),
		Salt:      []byte(salt),
		Hash:      hash,
		Params:    map[string]string{},
	}
}

func NewHashPartsWithParams(p Type, salt []byte, hash []byte, params map[string]string) *HashParts {
	return &HashParts{
		Algorithm: p.String(),
		Salt:      []byte(salt),
		Hash:      hash,
		Params:    params,
	}
}

func NewHashPartsFromUnderlying(name string, underlying string, salt []byte, hash []byte, ) *HashParts {
	t := Type{Name: name, Underlying: underlying}
	return &HashParts{
		Algorithm: t.String(),
		Salt:      salt,
		Hash:      hash,
		Params:    map[string]string{},
	}
}

func NewHashPartsWithParamsFromUnderlying(name string, underlying string, salt []byte, hash []byte, params map[string]string) *HashParts {
	t := Type{Name: name, Underlying: underlying}
	return &HashParts{
		Algorithm: t.String(),
		Salt:      salt,
		Hash:      hash,
		Params:    params,
	}
}
