/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package interfaces

import "github.com/origadmin/toolkits/crypto/hash/types"

// Encoder defines the interface for hash encoding operations
type Encoder interface {
	// Encode encodes salt and hash into a string
	Encode(salt []byte, hash []byte, params ...string) string
}

// Decoder defines the interface for hash decoding operations
type Decoder interface {
	// Decode decodes a string into hash parts
	Decode(encoded string) (*types.HashParts, error)
}

// Codec defines the interface for hash encoding and decoding operations
type Codec interface {
	Encoder
	Decoder
}
