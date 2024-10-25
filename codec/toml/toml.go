// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package toml provides the toml functions
package toml

import (
	"bytes"

	"github.com/BurntSushi/toml"
	"github.com/go-kratos/kratos/v2/encoding"
)

var (
	Marshal    = marshal
	Unmarshal  = toml.Unmarshal
	NewDecoder = toml.NewDecoder
	NewEncoder = toml.NewEncoder
	DecodeFile = toml.DecodeFile
	DecodeTOML = toml.Decode
	Codec      = codec{}
	Name       = "toml"
)

type (
	Value   = toml.Primitive
	Decoder = toml.Decoder
)

func marshal(v any) ([]byte, error) {
	buf := new(bytes.Buffer)
	err := toml.NewEncoder(buf).Encode(v)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// MarshalToString returns json string, and ignores error
func MarshalToString(v any) string {
	b, err := marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

// MustToString returns json string, or panic
func MustToString(v any) string {
	data, err := marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}

// Decode decodes the given JSON string into v
func Decode(data string, v any) error {
	_, err := DecodeTOML(data, v)
	if err != nil {
		return err
	}
	return nil
}

type codec struct{}

func (c codec) Marshal(v interface{}) ([]byte, error) {
	return toml.Marshal(v)
}

func (c codec) Unmarshal(data []byte, v interface{}) error {
	return toml.Unmarshal(data, v)
}

func (c codec) Name() string {
	return Name
}

func init() {
	encoding.RegisterCodec(Codec)
}
