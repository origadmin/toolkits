/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package interfaces

import "github.com/origadmin/toolkits/crypto/hash/types"

// HashEncoder defines the interface for hash encoding operations
type HashEncoder interface {
	// Encode encodes salt and hash into a string
	Encode(salt []byte, hash []byte, params ...string) string
}

// HashDecoder defines the interface for hash decoding operations
type HashDecoder interface {
	// Decode decodes a string into hash parts
	Decode(encoded string) (*types.HashParts, error)
}

// HashCodec defines the interface for hash encoding and decoding operations
type HashCodec interface {
	HashEncoder
	HashDecoder
}
