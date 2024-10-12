// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package codec is the codec package for encoding and decoding
package codec

import (
	"github.com/BurntSushi/toml"
	ggb "github.com/goexts/ggb"
)

type tomlDecoder struct {
	dec *toml.Decoder
}

func (t tomlDecoder) Decode(obj any) error {
	return ggb.OrNil(t.dec.Decode(obj))
}
