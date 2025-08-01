/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package types

import (
	"encoding/json"
)

// HashParts represents the parts of a hash
type HashParts struct {
	Algorithm string
	Version   string
	Params    map[string]string
	Hash      []byte
	Salt      []byte
}

func (h *HashParts) WithParams(params map[string]string) *HashParts {
	h.Params = params
	return h
}

func (h *HashParts) WithHash(hash []byte) *HashParts {
	h.Hash = hash
	return h
}

func (h *HashParts) WithVersion(version string) *HashParts {
	h.Version = version
	return h
}

func (h *HashParts) WithSalt(salt []byte) *HashParts {
	h.Salt = salt
	return h
}

func (h *HashParts) WithAlgorithm(algorithm string) *HashParts {
	h.Algorithm = algorithm
	return h
}

func (h *HashParts) WithHashSalt(hash []byte, salt []byte) *HashParts {
	h.Hash = hash
	h.Salt = salt
	return h
}

func (h *HashParts) String() string {
	b, err := json.Marshal(h)
	if err != nil {
		return ""
	}
	return string(b)
}

func NewHashParts(p Type) *HashParts {
	return &HashParts{
		Algorithm: p.String(),
		Params:    map[string]string{},
	}
}

func NewHashPartsWithHashSalt(p Type, hash []byte, salt []byte) *HashParts {
	return &HashParts{
		Algorithm: p.String(),
		Hash:      hash,
		Salt:      salt,
	}
}

func NewHashPartsFull(p Type, hash []byte, salt []byte, params map[string]string) *HashParts {
	return &HashParts{
		Algorithm: p.String(),
		Hash:      hash,
		Salt:      salt,
		Params:    params,
	}
}
