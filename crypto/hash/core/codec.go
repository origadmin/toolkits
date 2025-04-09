/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package core

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/goexts/generic/settings"

	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Codec implements a generic hash codec
type Codec struct {
	algorithm types.Type
	version   string
	params    string
}

func (c *Codec) Type() types.Type {
	return c.algorithm
}

func (c *Codec) Version() string {
	return c.version
}

// Encode implements the core encoding method
func (c *Codec) Encode(salt []byte, hash []byte, params ...string) string {
	var paramStr string
	if len(params) > 0 {
		paramStr = params[0]
	}
	return fmt.Sprintf(
		"$%s$%s$%s$%s$%s",
		c.algorithm.String(),
		c.version,
		paramStr,
		hex.EncodeToString(hash),
		hex.EncodeToString(salt),
	)
}

// Decode implements the core decoding method
func (c *Codec) Decode(encoded string) (*types.HashParts, error) {
	parts := strings.Split(encoded, "$")
	if len(parts) != 6 {
		return nil, ErrInvalidHashFormat
	}

	algorithm := types.Type(parts[1])
	varsion := parts[2]
	params := parts[3]

	hash, err := hex.DecodeString(parts[4])
	if err != nil {
		return nil, fmt.Errorf("invalid hash: %v", err)
	}
	salt, err := hex.DecodeString(parts[5])
	if err != nil {
		return nil, fmt.Errorf("invalid salt: %v", err)
	}
	return &types.HashParts{
		Algorithm: algorithm,
		Version:   varsion,
		Params:    params,
		Hash:      hash,
		Salt:      salt,
	}, nil
}

// NewCodec creates a new codec
func NewCodec(algorithm types.Type, opts ...CodecOption) interfaces.Codec {
	return settings.Apply(
		&Codec{
			algorithm: algorithm,
			version:   DefaultVersion,
		}, opts)
}

// CodecOption defines configuration options for the codec
type CodecOption func(*Codec)

// WithVersion sets the version number
func WithVersion(version string) CodecOption {
	return func(c *Codec) {
		c.version = version
	}
}

func ParseParams(params string) (map[string]string, error) {
	kv := make(map[string]string)
	if params == "" {
		return kv, nil
	}
	for _, param := range strings.Split(params, ParamSeparator) {
		parts := strings.Split(param, ParamValueSeparator)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid param format: %s", param)
		}
		kv[parts[0]] = parts[1]
	}
	return kv, nil
}
