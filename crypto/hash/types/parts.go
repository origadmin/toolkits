/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

import (
	"encoding/json"
)

// ParamContainer defines the interface for any object that can be converted into a parameter map.
// This avoids a circular dependency on the validator package, as any type that implements
// these methods will implicitly satisfy the interface.
type ParamContainer interface {
	ToMap() map[string]string
	IsNil() bool
}

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
func (h *HashParts) WithParams(params ParamContainer) *HashParts {
	if params != nil && !params.IsNil() {
		h.Params = params.ToMap()
	}
	return h
}

func (h *HashParts) WithMapParams(params map[string]string) *HashParts {
	if params == nil {
		return h
	}
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

// AddParam adds a single parameter to the hash parts and returns the modified HashParts instance.
func (h *HashParts) AddParam(key, value string) *HashParts {
	if h.Params == nil {
		h.Params = make(map[string]string)
	}
	h.Params[key] = value
	return h
}

// DeleteParam removes a parameter from the hash parts and returns the modified HashParts instance.
func (h *HashParts) DeleteParam(key string) *HashParts {
	if h.Params != nil {
		delete(h.Params, key)
	}
	return h
}

func (h *HashParts) IsEmpty() bool {
	return h.Spec.Name == "" && len(h.Hash) == 0 && len(h.Salt) == 0
}

func (h *HashParts) Clone() *HashParts {
	clone := &HashParts{
		Spec:    New(h.Spec.Name, h.Spec.Underlying),
		Version: h.Version,
		Hash:    make([]byte, len(h.Hash)),
		Salt:    make([]byte, len(h.Salt)),
	}
	copy(clone.Hash, h.Hash)
	copy(clone.Salt, h.Salt)

	if h.Params != nil {
		clone.Params = make(map[string]string, len(h.Params))
		for k, v := range h.Params {
			clone.Params[k] = v
		}
	}

	return clone
}

// ToJSON returns the JSON string representation of HashParts.
// If marshaling fails, it returns an error message indicating the failure.
func (h *HashParts) ToJSON() (string, error) {
	data, err := json.MarshalIndent(h, "", "  ")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// FromJSON creates a HashParts instance from a JSON string.
func FromJSON(jsonStr string) (*HashParts, error) {
	var parts HashParts
	if err := json.Unmarshal([]byte(jsonStr), &parts); err != nil {
		return nil, err
	}
	return &parts, nil
}

// NewHashParts is the primary constructor for creating a fully-formed HashParts object.
// It accepts a ParamContainer interface and handles the conversion to a map internally.
func NewHashParts(spec Spec, hash, salt []byte, params ParamContainer) *HashParts {
	var paramsMap map[string]string
	if params != nil && !params.IsNil() {
		paramsMap = params.ToMap()
	} else {
		paramsMap = make(map[string]string)
	}

	return &HashParts{
		Spec:   spec,
		Hash:   hash,
		Salt:   salt,
		Params: paramsMap,
	}
}

func NewHashPartsWithSpec(spec Spec) *HashParts {
	return &HashParts{
		Spec:   spec,
		Params: make(map[string]string),
	}
}
