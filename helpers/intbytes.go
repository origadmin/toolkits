// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package helpers for toolkits
package helpers

import (
	"encoding/binary"
)

// Uint64ToBytes converts a slice of uint64 values to a byte slice.
// Each uint64 value is encoded as 8 bytes in big-endian order.
// The resulting byte slice can be used to store the uint64 values in a binary format.
//
// Parameters:
// - ints: a slice of uint64 values to be converted to bytes.
//
// Returns:
// - a byte slice containing the encoded uint64 values.
func Uint64ToBytes(ints ...uint64) []byte {
	// Create a byte slice with a length equal to the number of uint64 values multiplied by 8.
	var buf = make([]byte, 8*len(ints))

	// Iterate over each uint64 value and encode it as 8 bytes in big-endian order.
	for i := range ints {
		// Use the binary.BigEndian.PutUint64 function to encode the uint64 value into the byte slice.
		binary.BigEndian.PutUint64(buf[i*8:i*8+8], ints[i])
	}

	// Return the byte slice containing the encoded uint64 values.
	return buf
}

// BytesToUint64 converts a byte slice to a slice of uint64 values.
// Each 8 bytes in the byte slice are decoded as a single uint64 value in big-endian order.
// The resulting slice of uint64 values can be used to retrieve the original uint64 values from the byte slice.
//
// Parameters:
// - buf: a byte slice containing the encoded uint64 values.
//
// Returns:
// - a slice of uint64 values decoded from the byte slice.
func BytesToUint64(buf []byte) []uint64 {
	// Get the length of the byte slice.
	size := len(buf)

	// If the byte slice is empty, return nil.
	if size == 0 {
		return nil
	}

	// Create a slice of uint64 values with a length equal to the byte slice length divided by 8.
	ints := make([]uint64, size/8)

	// Iterate over each 8 bytes in the byte slice and decode them as a single uint64 value in big-endian order.
	for i := 0; i < size/8; i++ {
		// Use the binary.BigEndian.Uint64 function to decode the 8 bytes as a single uint64 value.
		ints[i] = binary.BigEndian.Uint64(buf[i*8 : i*8+8])
	}

	// Return the slice of uint64 values decoded from the byte slice.
	return ints
}
