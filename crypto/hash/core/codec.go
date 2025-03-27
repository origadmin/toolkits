/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package core

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/origadmin/toolkits/crypto/hash/types"
)

// codec provides core implementation for hash codec
type codec struct {
	Algorithm types.Type
	Version   string
}

// Encode implements the core encoding method
func (e *codec) Encode(salt []byte, hash []byte, params ...string) string {
	var paramStr string
	if len(params) > 0 {
		paramStr = params[0]
	}
	return fmt.Sprintf("$%s$%s$%s$%s$%s",
		e.Algorithm.String(),
		e.Version,
		paramStr,
		hex.EncodeToString(salt),
		hex.EncodeToString(hash))
}

// Decode implements the core decoding method
func (e *codec) Decode(encoded string) (*types.HashParts, error) {
	parts := strings.Split(encoded, "$")
	if len(parts) < 5 {
		return nil, fmt.Errorf("invalid hash format")
	}

	algorithm := types.Type(parts[1])
	if algorithm != e.Algorithm {
		return nil, ErrAlgorithmMismatch
	}

	version := parts[2]
	if version != e.Version {
		return nil, fmt.Errorf("unsupported version: %s", version)
	}

	params := parts[3]
	salt, err := hex.DecodeString(parts[4])
	if err != nil {
		return nil, fmt.Errorf("invalid salt: %v", err)
	}

	hash, err := hex.DecodeString(parts[5])
	if err != nil {
		return nil, fmt.Errorf("invalid hash: %v", err)
	}

	return &types.HashParts{
		Algorithm: algorithm,
		Version:   version,
		Params:    params,
		Salt:      salt,
		Hash:      hash,
	}, nil
}

// NewCodec creates a new hash codec
func NewCodec(algorithm types.Type) *codec {
	return &codec{
		Algorithm: algorithm,
		Version:   DefaultVersion,
	}
}
