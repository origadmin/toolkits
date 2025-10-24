/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package codec

import (
	"encoding/hex"
	"fmt"
	"sort"
	"strings"

	"github.com/goexts/generic/configure"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Encoder defines the interface for hash encoding operations
type Encoder interface {
	// Encode encodes salt and hash into a string
	Encode(parts *types.HashParts) (string, error)
}

// Decoder defines the interface for hash decoding operations
type Decoder interface {
	// Decode decodes a string into hash parts
	Decode(encoded string) (*types.HashParts, error)
}

// Version defines the interface for version operations
type Version interface {
	// Version returns the version of the codec
	Version() string
}

// codec defines the interface for hash encoding and decoding operations
type Codec interface {
	Encoder
	Decoder
	Version
}

// codec implements a generic hash codec
type codec struct {
	version string
}

func (c *codec) Version() string {
	return c.version
}

// Encode implements the core encoding method
func (c *codec) Encode(parts *types.HashParts) (string, error) {
	if parts.Version == "" {
		parts.Version = c.version
	}
	return fmt.Sprintf(
		"$%s$%s$%s$%s$%s",
		parts.Algorithm,
		parts.Version,
		EncodeParams(parts.Params),
		hex.EncodeToString(parts.Hash),
		hex.EncodeToString(parts.Salt),
	), nil
}

// Decode implements the core decoding method
func (c *codec) Decode(encoded string) (*types.HashParts, error) {
	parts := strings.Split(encoded, types.CodecSeparator)
	if len(parts) != 6 {
		return nil, errors.ErrInvalidHashFormat
	}

	algorithm, err := types.Parse(parts[1])
	if err != nil {
		return nil, err
	}
	version := parts[2]

	// Add version checks
	if version != c.version {
		return nil, fmt.Errorf("unsupported codec version: %s, expected %s", version, c.version)
	}
	paramsStr := parts[3]
	decodedParams, err := DecodeParams(paramsStr)
	if err != nil {
		return nil, fmt.Errorf("invalid params format: %v", err)
	}
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
		Version:   version,
		Params:    decodedParams,
		Hash:      hash,
		Salt:      salt,
	}, nil
}

// NewCodec creates a new codec
func NewCodec(opts ...Option) Codec {
	return configure.Apply(
		&codec{
			version: types.DefaultVersion,
		}, opts)
}

// Option defines configuration options for the codec
type Option func(*codec)

// WithVersion sets the version number
func WithVersion(version string) Option {
	return func(c *codec) {
		c.version = version
	}
}

func DecodeParams(params string) (map[string]string, error) {
	kv := make(map[string]string)
	if params == "" {
		return kv, nil
	}
	for _, param := range strings.Split(params, types.ParamSeparator) {
		parts := strings.Split(param, types.ParamValueSeparator)
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid param format: %s", param)
		}
		kv[parts[0]] = parts[1]
	}
	return kv, nil
}

// EncodeParams encodes a map of parameters into a string.
// It sorts the keys to ensure a consistent output string.
func EncodeParams(params map[string]string) string {
	if len(params) == 0 {
		return ""
	}

	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var parts []string
	for _, k := range keys {
		parts = append(parts, fmt.Sprintf("%s%s%s", k, types.ParamValueSeparator, params[k]))
	}
	return strings.Join(parts, types.ParamSeparator)
}
