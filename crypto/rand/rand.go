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
	KindDigit     Kind = 1 << iota // Digits (0-9)
	KindLowerCase Kind = 1 << iota // Lowercase letters (a-z)
	KindUpperCase Kind = 1 << iota // Uppercase letters (A-Z)
	KindSymbol    Kind = 1 << iota // Common symbols

	// Combinations
	KindAlphanumeric   = KindDigit | KindLowerCase | KindUpperCase
	KindAllWithSymbols = KindAlphanumeric | KindSymbol
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

// Rand is a generator for random data based on a specified character set.
type Rand struct {
	charset string
	length  int
}

// NewRand creates a new random data generator for the given kind of character set.
func NewRand(kind Kind) *Rand {
	charset, ok := charsets[kind]
	if !ok {
		// Fallback to a safe default if the kind is not pre-built.
		charset = charsets[KindAlphanumeric]
	}
	return &Rand{
		charset: charset,
		length:  len(charset),
	}
}

// CustomRand creates a new random data generator with a custom character set.
func CustomRand(charset string) *Rand {
	return &Rand{
		charset: charset,
		length:  len(charset),
	}
}

// RandBytes generates a random byte slice of a given size using the generator's character set.
// It uses a cryptographically secure random number generator.
// Note: This method has a slight modulo bias, which is acceptable for most non-uniformity-critical applications (like salt generation).
func (r *Rand) RandBytes(size int) ([]byte, error) {
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
func (r *Rand) RandString(size int) (string, error) {
	b, err := r.RandBytes(size)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// RandomBytes generates n cryptographically secure random bytes using the default alphanumeric character set.
func RandomBytes(n int) ([]byte, error) {
	return NewRand(KindAlphanumeric).RandBytes(n)
}

// RandomString generates a cryptographically secure random string of length n using the default alphanumeric character set.
func RandomString(n int) (string, error) {
	return NewRand(KindAlphanumeric).RandString(n)
}