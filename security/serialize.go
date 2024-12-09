/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security provides interfaces and types for security-related operations
package security

// Serializer is an interface that defines the methods for a serializer
type Serializer interface {
	// Serialize serializes the given data into a byte slice
	Serialize(Claims) ([]byte, error)
	// Deserialize deserializes the given byte slice into the given data
	Deserialize([]byte) (Claims, error)
}
