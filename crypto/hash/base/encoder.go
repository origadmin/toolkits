/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package base

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// BaseHashCodec implements the hash codec interface
type BaseHashCodec struct {
	Algorithm types.Type
	Version   string
}

// NewBaseHashCodec creates a new base hash codec
func NewBaseHashCodec(algorithm types.Type) interfaces.HashCodec {
	return &BaseHashCodec{
		Algorithm: algorithm,
		Version:   "1",
	}
}

// Encode implements the encode method
func (e *BaseHashCodec) Encode(salt []byte, hash []byte, params ...string) string {
	var paramStr string
	if len(params) > 0 {
		paramStr = params[0]
	}
	return fmt.Sprintf("$%s$%s$%s$%s$%s",
		e.Algorithm,
		e.Version,
		paramStr,
		hex.EncodeToString(salt),
		hex.EncodeToString(hash))
}

// Decode implements the decode method
func (e *BaseHashCodec) Decode(encoded string) (*types.HashParts, error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 {
		return nil, fmt.Errorf("invalid hash format")
	}

	if parts[1] != string(e.Algorithm) {
		return nil, fmt.Errorf("algorithm mismatch")
	}

	salt, err := hex.DecodeString(parts[4])
	if err != nil {
		return nil, fmt.Errorf("invalid salt format")
	}

	hash, err := hex.DecodeString(parts[5])
	if err != nil {
		return nil, fmt.Errorf("invalid hash format")
	}

	return &types.HashParts{
		Algorithm: types.Type(parts[1]),
		Version:   parts[2],
		Params:    parts[3],
		Salt:      salt,
		Hash:      hash,
	}, nil
}

// GetCodec returns a codec instance for the given algorithm
func GetCodec(algorithm types.Type) interfaces.HashCodec {
	return NewBaseHashCodec(algorithm)
}
