/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package interfaces

import "github.com/origadmin/toolkits/crypto/hash/types"

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

// Codec defines the interface for hash encoding and decoding operations
type Codec interface {
	Encoder
	Decoder
	Version
}
