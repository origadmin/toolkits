/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package rand provides cryptographically secure, high-performance random string and byte generation.
package rand

import (
	"crypto/rand"
	"io"
)

// Kind defines the character set to be used for random string generation.
type Kind int

const (
	KindDigit     Kind = 1 << iota // KindDigit represents digit characters (0-9).
	KindLowerCase Kind = 1 << iota // KindLowerCase represents lowercase letters (a-z).
	KindUpperCase Kind = 1 << iota // KindUpperCase represents uppercase letters (A-Z).
	KindSymbol    Kind = 1 << iota // KindSymbol represents common symbol characters.

	KindAlphanumeric   = KindDigit | KindLowerCase | KindUpperCase // KindAlphanumeric represents a combination of digits, lowercase, and uppercase letters.
	KindAllWithSymbols = KindAlphanumeric | KindSymbol             // KindAllWithSymbols represents a combination of digits, lowercase, uppercase letters, and symbols.
)

const (
	Digits    = "0123456789"
	Lowercase = "abcdefghijklmnopqrstuvwxyz"
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Symbols   = "!@#$%^&*-_=+"
)

var charsets = make(map[Kind]string)

func init() {
	charsets[KindDigit] = Digits
	charsets[KindLowerCase] = Lowercase
	charsets[KindUpperCase] = Uppercase
	charsets[KindSymbol] = Symbols
	charsets[KindAlphanumeric] = Digits + Lowercase + Uppercase
	charsets[KindAllWithSymbols] = Digits + Lowercase + Uppercase + Symbols
}

// Generator defines the interface for generating cryptographically secure random data.
type Generator interface {
	// RandBytes generates a random byte slice of the specified size using the generator's character set.
	RandBytes(size int) ([]byte, error)
	// RandString generates a random string of the specified size using the generator's character set.
	RandString(size int) (string, error)
	// Read populates the given byte slice with random bytes from the generator's character set.
	// It implements the io.Reader interface.
	Read(p []byte) (n int, err error)
}

// randGenerator is the concrete implementation of RandGenerator.
// It is unexported to hide implementation details.
type randGenerator struct {
	charset string
	length  int
}

// NewRand creates a new random data generator for the given kind of character set.
// It returns the RandGenerator interface.
func NewRand(kind Kind) Generator {
	charset, ok := charsets[kind]
	if !ok {
		// Fallback to a safe default if the kind is not pre-built.
		charset = charsets[KindAlphanumeric]
	}
	return &randGenerator{
		charset: charset,
		length:  len(charset),
	}
}

// CustomRand creates a new random data generator with a custom character set.
// It returns the RandGenerator interface.
func CustomRand(charset string) Generator {
	return &randGenerator{
		charset: charset,
		length:  len(charset),
	}
}

// RandBytes generates a random byte slice of a given size using the generator's character set.
// It uses a cryptographically secure random number generator.
// Note: This method has a slight modulo bias, which is acceptable for most non-uniformity-critical applications (like salt generation).
func (r *randGenerator) RandBytes(size int) ([]byte, error) {
	if r.length == 0 {
		return nil, nil
	}
	ret := make([]byte, size)
	randomBytes := make([]byte, size)
	_, err := io.ReadFull(rand.Reader, randomBytes)
	if err != nil {
		return nil, err
	}

	for i := 0; i < size; i++ {
		ret[i] = r.charset[randomBytes[i]%byte(r.length)]
	}
	return ret, nil
}

// RandString generates a random string of a given size using the generator's character set.
func (r *randGenerator) RandString(size int) (string, error) {
	b, err := r.RandBytes(size)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Read populates the given byte slice with random bytes from the generator's character set.
// It implements the io.Reader interface.
func (r *randGenerator) Read(p []byte) (n int, err error) {
	if r.length == 0 {
		return 0, nil
	}
	n = len(p)
	randomBytes := make([]byte, n)
	_, err = io.ReadFull(rand.Reader, randomBytes)
	if err != nil {
		return 0, err
	}

	for i := 0; i < n; i++ {
		p[i] = r.charset[randomBytes[i]%byte(r.length)]
	}
	return n, nil
}

// RandomBytes generates n cryptographically secure random bytes using the default alphanumeric character set.
func RandomBytes(n int) ([]byte, error) {
	return NewRand(KindAlphanumeric).RandBytes(n)
}

// RandomString generates a cryptographically secure random string of length n using the default alphanumeric character set.
func RandomString(n int) (string, error) {
	return NewRand(KindAlphanumeric).RandString(n)
}
