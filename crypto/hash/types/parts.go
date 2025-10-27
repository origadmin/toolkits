/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

import (
	"encoding/json"
	"fmt"
)

// HashParts represents the parts of a hash, designed to be a portable data container
// that stores parsed algorithm information.
// It is suitable for serialization (e.g., to JSON) for debugging or transfer.
type HashParts struct {
	// Spec is the structured identifier for the hash algorithm (e.g., {Name: "pbkdf2", Underlying: "sha256"}).
	// It is of type Spec, representing the parsed algorithm definition.
	Spec Spec `json:"spec"`

	// Version indicates a specific version of the algorithm or its parameters, if applicable.
	Version string `json:"version,omitempty"`

	// Params holds algorithm-specific parameters, such as cost, rounds, or memory usage.
	Params map[string]string `json:"params,omitempty"`

	// Hash is the raw computed hash value.
	Hash []byte `json:"hash,omitempty"`

	// Salt is the salt used during the hashing process.
	Salt []byte `json:"salt,omitempty"`
}

// WithParams sets the parameters for the hash parts and returns the modified HashParts instance.
func (h *HashParts) WithParams(params map[string]string) *HashParts {
	h.Params = params
	return h
}

// WithVersion sets the version string for the hash parts and returns the modified HashParts instance.
func (h *HashParts) WithVersion(version string) *HashParts {
	h.Version = version
	return h
}

// WithHash sets the hash bytes for the hash parts and returns the modified HashParts instance.
func (h *HashParts) WithHash(hash []byte) *HashParts {
	h.Hash = hash
	return h
}

// WithSalt sets the salt bytes for the hash parts and returns the modified HashParts instance.
func (h *HashParts) WithSalt(salt []byte) *HashParts {
	h.Salt = salt
	return h
}

// WithSpec sets the algorithm Spec for the hash parts and returns the modified HashParts instance.
func (h *HashParts) WithSpec(spec Spec) *HashParts {
	h.Spec = spec
	return h
}

// WithHashSalt sets both the hash and salt bytes for the hash parts and returns the modified HashParts instance.
func (h *HashParts) WithHashSalt(hash []byte, salt []byte) *HashParts {
	h.Hash = hash
	h.Salt = salt
	return h
}

// String returns the JSON string representation of HashParts.
// If marshaling fails, it returns an error message indicating the failure.
func (h *HashParts) String() string {
	b, err := json.Marshal(h)
	if err != nil {
		return fmt.Sprintf("Error marshaling HashParts to JSON: %v", err)
	}
	return string(b)
}

// NewHashParts creates a new HashParts instance using the provided algorithm Spec.
// It initializes Params to an empty map to prevent nil map panics.
func NewHashParts(spec Spec) *HashParts {
	return &HashParts{
		Spec:   spec,
		Params: make(map[string]string), // Ensure Params is initialized
	}
}

// NewHashPartsWithHashSalt creates a new HashParts instance with the given algorithm Spec, hash, and salt.
// It initializes Params to an empty map to prevent nil map panics.
func NewHashPartsWithHashSalt(spec Spec, hash []byte, salt []byte) *HashParts {
	return &HashParts{
		Spec:   spec,
		Hash:   hash,
		Salt:   salt,
		Params: make(map[string]string), // Ensure Params is initialized
	}
}

// NewHashPartsFull creates a new HashParts instance with all fields provided.
// If the provided params map is nil, it initializes it to an empty map.
func NewHashPartsFull(spec Spec, hash []byte, salt []byte, params map[string]string) *HashParts {
	if params == nil {
		params = make(map[string]string) // Ensure Params is initialized if nil
	}
	return &HashParts{
		Spec:   spec,
		Hash:   hash,
		Salt:   salt,
		Params: params,
	}
}
