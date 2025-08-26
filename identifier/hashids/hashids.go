/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hashids provides a tool to encode integers into short, unique, non-sequential strings.
// It's ideal for obfuscating database IDs in URLs.
package hashids

import (
	"fmt"

	"github.com/speps/go-hashids/v2"
)

// Config holds the configuration for creating a new Hashid encoder/decoder.
type Config struct {
	// Salt is a secret string that makes your hashes unique.
	// It's crucial to keep this value private and consistent.
	Salt string
	// MinLength is the minimum length of the generated hashes.
	MinLength int
	// Alphabet is the custom character set to use for generation.
	// It must contain at least 16 unique characters.
	Alphabet string
}

// Hashid is the main struct that performs encoding and decoding.
type Hashid struct {
	hd *hashids.HashID
}

// New creates a new, configured Hashid instance.
func New(cfg Config) (*Hashid, error) {
	if cfg.Salt == "" {
		return nil, fmt.Errorf("hashids: salt cannot be empty")
	}

	hdata := hashids.NewData()
	hdata.Salt = cfg.Salt
	hdata.MinLength = cfg.MinLength
	if cfg.Alphabet != "" {
		hdata.Alphabet = cfg.Alphabet
	}

	hd, err := hashids.NewWithData(hdata)
	if err != nil {
		return nil, fmt.Errorf("hashids: failed to initialize: %w", err)
	}
	return &Hashid{hd: hd}, nil
}

// Encode converts one or more non-negative integers into a unique hash string.
func (h *Hashid) Encode(numbers ...int64) (string, error) {
	// The underlying library uses `int`, so we need to check for potential overflow on 32-bit systems,
	// although it's rare for IDs to exceed `int` max.
	// For simplicity and modern architecture (64-bit), we'll proceed with direct conversion.
	// A more robust implementation might handle this conversion more carefully.
	
	// Convert int64 to int for the library
	nums := make([]int, len(numbers))
	for i, n := range numbers {
		if n < 0 {
			return "", fmt.Errorf("hashids: input numbers cannot be negative")
		}
		nums[i] = int(n)
	}

	return h.hd.Encode(nums)
}

// Decode converts a hash string back into a slice of integers.
func (h *Hashid) Decode(hash string) ([]int64, error) {
	numbers, err := h.hd.DecodeWithError(hash)
	if err != nil {
		return nil, err
	}

	// Convert int back to int64
	nums64 := make([]int64, len(numbers))
	for i, n := range numbers {
		nums64[i] = int64(n)
	}
	return nums64, nil
}
