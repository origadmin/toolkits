/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package toml provides the toml functions
package toml

import (
	"bytes"
	"io"

	"github.com/BurntSushi/toml"
)

var (
	Marshal        = marshal
	Unmarshal      = toml.Unmarshal
	newDecoderTOML = toml.NewDecoder
	NewEncoder     = toml.NewEncoder
	DecodeFile     = toml.DecodeFile
	DecodeTOML     = toml.Decode
)

type (
	Value       = toml.Primitive
	DecoderTOML = toml.Decoder
)

type Decoder struct {
	dec *toml.Decoder
}

func (t Decoder) Decode(obj any) error {
	_, err := t.dec.Decode(obj)
	return err
}

func NewDecoder(ior io.Reader) *Decoder {
	return &Decoder{dec: toml.NewDecoder(ior)}
}

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

//func init() {
//	encoding.RegisterCodec(Codec)
//}
