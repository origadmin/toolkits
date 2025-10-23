/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package identifier

import (
	"crypto/rand"
	"fmt"
	"io"
)

// --- Built-in Default Provider for string (Fallback) ---

// defaultStringProvider is a minimal, dependency-free UUIDv4 generator.
// It is used as a fallback when no other provider for "uuid" is registered.
type defaultStringProvider struct{}

func (p *defaultStringProvider) Name() string { return "uuid" }
func (p *defaultStringProvider) Size() int    { return 128 }
func (p *defaultStringProvider) AsString() Generator[string] { return p }
func (p *defaultStringProvider) AsNumber() Generator[int64] { return nil }

func (p *defaultStringProvider) Generate() string {
	var b [16]byte
	if _, err := io.ReadFull(rand.Reader, b[:]); err != nil {
		// This is a critical failure of the OS's entropy source, panic is acceptable.
		panic(fmt.Sprintf("identifier: failed to read random data for default uuid: %v", err))
	}
	// Set version 4 and variant (RFC 4122)
	b[6] = (b[6]&0x0f) | 0x40
	b[8] = (b[8]&0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// Validate provides a basic validation for UUID format.
func (p *defaultStringProvider) Validate(id string) bool {
	// A simple format check. A full-featured validation is provided by the external uuid package.
	if len(id) != 36 {
		return false
	}
	if id[8] != '-' || id[13] != '-' || id[18] != '-' || id[23] != '-' {
		return false
	}
	return true
}

// builtinString is the singleton instance of our built-in fallback for string.
var builtinString Provider = &defaultStringProvider{}
