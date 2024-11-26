/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package codec is the codec package for encoding and decoding
package codec

type Codec interface {
	Marshal(v interface{}) ([]byte, error)
	Unmarshal(data []byte, v interface{}) error
	Name() string
}
