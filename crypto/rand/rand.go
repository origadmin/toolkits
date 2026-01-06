/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package rand provides cryptographically secure, high-performance random string and byte generation.
package rand

import (
	"crypto/rand"
	"io"
	"sync"
)

// Kind defines the character set to be used for random string generation.
type Kind int

const (
	KindDigit     Kind = 1 << iota // KindDigit represents digit characters (0-9).
	KindLowerCase                  // KindLowerCase represents lowercase letters (a-z).
	KindUpperCase                  // KindUpperCase represents uppercase letters (A-Z).
	KindSymbol                     // KindSymbol represents common symbol characters.
)

const (
	// The maximum value for Kind, used to determine the size of the charset array.
	// This is calculated as (1 << iota) for the last Kind, which is KindSymbol (1 << 3 = 8).
	// The array size needs to be 2 * maxKind to hold all combinations.
	maxKind = KindSymbol * 2
)

const (
	Digits    = "0123456789"
	Lowercase = "abcdefghijklmnopqrstuvwxyz"
	Uppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Symbols   = "!@#$%^&*-_=+"
)

// KindAlphanumeric represents a combination of digits, lowercase, and uppercase letters.
const KindAlphanumeric = KindDigit | KindLowerCase | KindUpperCase

// KindAllWithSymbols represents a combination of digits, lowercase, uppercase letters, and symbols.
const KindAllWithSymbols = KindAlphanumeric | KindSymbol

// charsets is an array-based lookup table for performance. It's faster than a map.
var charsets [maxKind]string

func init() {
	// Initialize single kinds
	charsets[KindDigit] = Digits
	charsets[KindLowerCase] = Lowercase
	charsets[KindUpperCase] = Uppercase
	charsets[KindSymbol] = Symbols

	// Pre-calculate and initialize common combinations
	charsets[KindDigit|KindLowerCase] = Digits + Lowercase
	charsets[KindDigit|KindUpperCase] = Digits + Uppercase
	charsets[KindLowerCase|KindUpperCase] = Lowercase + Uppercase
	charsets[KindDigit|KindLowerCase|KindUpperCase] = Digits + Lowercase + Uppercase
	charsets[KindDigit|KindLowerCase|KindUpperCase|KindSymbol] = Digits + Lowercase + Uppercase + Symbols
}

// Rand defines the interface for generating cryptographically secure random data.
// Implementations of this interface are designed for high performance and statistical randomness.
type Rand interface {
	// Bytes generates a new random byte slice of the specified size.
	// This method allocates a new slice for each call. For maximum performance when
	// repeatedly generating random bytes into a pre-existing buffer, prefer the Read method.
	Bytes(size int) ([]byte, error)
	// String generates a new random string of the specified size.
	// This method involves internal byte slice allocation and conversion to string.
	// For scenarios requiring extreme performance and byte-level control, consider
	// using RandBytes or Read and managing string conversion manually if profiling
	// indicates this is a bottleneck.
	String(size int) (string, error)
	// Read populates the given byte slice with random bytes from the generator's character set.
	// It implements the io.Reader interface. This is the most performant method for
	// filling pre-allocated buffers, as it avoids internal allocations of the output slice.
	Read(p []byte) (n int, err error)
}

const (
	// Internal buffer size for reading from crypto/rand.Reader to optimize performance.
	randBufferSize = 512
)

// rand is the concrete implementation of Rand.
type randImpl struct {
	charset string
	maxByte byte // maxByte is the maximum byte value that can be used without introducing modulo bias.

	// Internal buffer for performance optimization
	buffer    [randBufferSize]byte
	bufferIdx int
	bufferN   int
	mu        sync.Mutex // Protects buffer access for concurrent Read/Bytes calls
}

func newRand(kind Kind) Rand {
	var charset string
	// Check bounds and if the specific combination is pre-calculated.
	if kind < maxKind && charsets[kind] != "" {
		charset = charsets[kind]
	} else {
		// Fallback for non-pre-calculated combinations or invalid kinds.
		if kind&KindDigit != 0 {
			charset += Digits
		}
		if kind&KindLowerCase != 0 {
			charset += Lowercase
		}
		if kind&KindUpperCase != 0 {
			charset += Uppercase
		}
		if kind&KindSymbol != 0 {
			charset += Symbols
		}
	}

	// If after all checks, charset is still empty (e.g., invalid kind), default to alphanumeric.
	if charset == "" {
		charset = charsets[KindDigit|KindLowerCase|KindUpperCase]
	}

	return NewWithCharset(charset)
}

// New creates a new random data generator for the given kind of character set.
// It uses a high-performance array lookup instead of a map.
func New(kind Kind) Rand {
	return newRand(kind)
}

// NewGenerator creates a new random data generator for the given kind of character set.
// It uses a high-performance array lookup instead of a map.
//
// Decrypted: use New instead.
func NewGenerator(kind Kind) Rand {
	return newRand(kind)
}

// NewWithCharset creates a new random data generator with a custom character set.
func NewWithCharset(charset string) Rand {
	length := len(charset)
	if length == 0 {
		return &randImpl{charset: "", maxByte: 0}
	}
	// Calculate maxByte to eliminate modulo bias using rejection sampling.
	maxByte := byte(256 - (256 % length))
	return &randImpl{
		charset: charset,
		maxByte: maxByte,
	}
}

// refillBuffer reads a new chunk of random bytes from crypto/rand.Reader into the internal buffer.
func (r *randImpl) refillBuffer() error {
	n, err := io.ReadFull(rand.Reader, r.buffer[:])
	if err != nil {
		return err
	}
	r.bufferN = n
	r.bufferIdx = 0
	return nil
}

// Bytes generates a random byte slice of a given size using the generator's character set.
// It is highly optimized for performance and security (no modulo bias).
func (r *randImpl) Bytes(size int) ([]byte, error) {
	charsetLen := len(r.charset)
	if charsetLen == 0 {
		return nil, nil
	}

	ret := make([]byte, size)
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := 0; i < size; {
		if r.bufferIdx >= r.bufferN {
			if err := r.refillBuffer(); err != nil {
				return nil, err
			}
		}

		for r.bufferIdx < r.bufferN && i < size {
			randomByte := r.buffer[r.bufferIdx]
			r.bufferIdx++

			if randomByte < r.maxByte {
				ret[i] = r.charset[randomByte%byte(charsetLen)]
				i++
			}
		}
	}
	return ret, nil
}

// String generates a random string of a given size using the generator's character set.
func (r *randImpl) String(size int) (string, error) {
	b, err := r.Bytes(size)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Read populates the given byte slice with random bytes from the generator's character set.
// It implements the io.Reader interface and is highly optimized.
func (r *randImpl) Read(p []byte) (n int, err error) {
	charsetLen := len(r.charset)
	if charsetLen == 0 {
		return 0, nil
	}

	n = len(p)
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := 0; i < n; {
		if r.bufferIdx >= r.bufferN {
			if err := r.refillBuffer(); err != nil {
				if i > 0 {
					return i, err
				}
				return 0, err
			}
		}

		for r.bufferIdx < r.bufferN && i < n {
			randomByte := r.buffer[r.bufferIdx]
			r.bufferIdx++

			if randomByte < r.maxByte {
				p[i] = r.charset[randomByte%byte(charsetLen)]
				i++
			}
		}
	}
	return n, nil
}

// RandomBytes generates n cryptographically secure random bytes using the default alphanumeric character set.
// This function reuses a single, lazily initialized Rand instance for performance.
func RandomBytes(n int) ([]byte, error) {
	return NewWithCharset(charsets[KindAlphanumeric]).Bytes(n)
}

// RandomString generates a cryptographically secure random string of length n using the default alphanumeric character set.
// This function reuses a single, lazily initialized Rand instance for performance.
func RandomString(n int) (string, error) {
	return NewWithCharset(charsets[KindAlphanumeric]).String(n)
}

// RandomBytesWithSymbols generates n cryptographically secure random bytes using the alphanumeric and symbol character set.
// This function reuses a single, lazily initialized Rand instance for performance.
func RandomBytesWithSymbols(n int) ([]byte, error) {
	return NewWithCharset(charsets[KindAllWithSymbols]).Bytes(n)
}

// RandomStringWithSymbols generates a cryptographically secure random string of length n using the alphanumeric and symbol character set.
// This function reuses a single, lazily initialized Rand instance for performance.
func RandomStringWithSymbols(n int) (string, error) {
	return NewWithCharset(charsets[KindAllWithSymbols]).String(n)
}
