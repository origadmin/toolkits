// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package rand provides the random string
package rand

import (
	"bytes"
	"math/rand/v2"
	"strings"
	"sync"
)

const (
	randCharset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!@#$%^&*-_=+"
)

type Kind int

// Define a series of character type constants representing different categories of characters.
// KindDigit     - Represents digit characters.
// KindLowerCase - Represents lowercase letter characters.
// KindUpperCase - Represents uppercase letter characters.
// KindSymbol    - Represents symbol characters.
// KindAll       - Represents the collective set of all character types, including digits, lowercase, uppercase, and symbols.
const (
	KindDigit     Kind = 1 << iota
	KindLowerCase Kind = 1 << iota
	KindUpperCase Kind = 1 << iota
	KindSymbol    Kind = 1 << iota
	KindAll       Kind = KindDigit | KindLowerCase | KindUpperCase | KindSymbol
	KindCustom    Kind = 0xFF
)

const (
	randomDigit = iota
	randomLowerCase
	randomUpperCase
	randomSymbol
	randomAll
	randomMax
)

type Rand struct {
	kind    Kind
	length  int
	charset string
}

// predefined charsets for different random types
// Digit             - Represents digit characters.
// LowerCase         - Represents lowercase letter characters.
// UpperCase         - Represents uppercase letter characters.
// Symbol            - Represents symbol characters.
// LowerAndUpperCase - Represents lowercase and uppercase letter characters.
// DigitAndLowerCase - Represents digit and lowercase letter characters.
// DigitAndUpperCase - Represents digit and uppercase letter characters.

var (
	Digit             = NewRand(KindDigit)
	LowerCase         = NewRand(KindLowerCase)
	UpperCase         = NewRand(KindUpperCase)
	Symbol            = NewRand(KindSymbol)
	LowerAndUpperCase = NewRand(KindLowerCase | KindUpperCase)
	DigitAndLowerCase = NewRand(KindDigit | KindLowerCase)
	DigitAndUpperCase = NewRand(KindDigit | KindUpperCase)
	All               = NewRand(KindAll)
)

var randPool = sync.Pool{
	New: func() interface{} {
		return &Rand{}
	},
}

var stringIndex = [randomMax][2]int{
	randomDigit:     {0, 10},
	randomLowerCase: {10, 36},
	randomUpperCase: {36, 62},
	randomSymbol:    {62, 74},
	randomAll:       {0, 74},
}

func (k Kind) String() string {
	var kinds []string
	if k > KindAll {
		return "Custom"
	}
	if k&KindAll == KindAll {
		return "All"
	}
	if k&KindDigit == KindDigit {
		kinds = append(kinds, "Digit")
	}
	if k&KindLowerCase == KindLowerCase {
		kinds = append(kinds, "LowerCase")
	}
	if k&KindUpperCase == KindUpperCase {
		kinds = append(kinds, "UpperCase")
	}
	if k&KindSymbol == KindSymbol {
		kinds = append(kinds, "Symbol")
	}

	if len(kinds) == 0 {
		return "<nil>"
	}
	return strings.Join(kinds, "|")
}

// RandBytes generates a random byte slice of given size using the given charset.
//
// Parameters:
// - size: the length of the byte slice to be generated.
//
// Return:
// - []byte: the generated random byte slice.
func (r *Rand) RandBytes(size int) []byte {
	if r.length == 0 {
		return nil
	}
	ret := make([]byte, size)
	for ; size > 0; size-- {
		ret[size-1] = r.charset[rand.IntN(r.length)]
	}
	return ret
}

// RandString generates a random string of given size using the given charset.
//
// Parameters:
// - size: the length of the string to be generated.
//
// Return:
// - string: the generated random string.
func (r *Rand) RandString(size int) string {
	return string(r.RandBytes(size))
}

// The Read method populates the given byte slice by randomly selecting characters.
//
// Parameters:
// p []byte: The byte slice to be populated.
//
// Return values:
// n int: The number of bytes actually populated.
// err error: An error encountered during population, always nil in this implementation.
func (r *Rand) Read(p []byte) (n int, err error) {
	if r.length == 0 {
		return 0, nil
	}
	n = len(p) // Set n to the length of p, indicating the number of bytes to populate.
	for i := 0; i < n; i++ {
		// Randomly select a character for each position in p.
		p[i] = r.charset[rand.IntN(r.length)]
	}
	return n, nil // Return the populated byte count and a nil error.
}

// Close releases the resources associated with the Rand object.
func (r *Rand) Close() {
	r.Reset()
	randPool.Put(r)
}

// Reset resets the Rand object to its initial state.
func (r *Rand) Reset() {
	r.kind = 0
	r.length = 0
	r.charset = ""
}

func loadCharset(rand Kind) string {
	if rand >= KindAll {
		return randCharset
	}
	var sta, end int
	var buf bytes.Buffer
	for i := 0; i < randomMax-1; i++ {
		if rand&(1<<uint(i)) != 0 {
			sta, end = getStringIndex(i)
			buf.WriteString(randCharset[sta:end])
		}
	}
	return buf.String()
}

func getStringIndex(idx int) (int, int) {
	return stringIndex[idx][0], stringIndex[idx][1]
}

// NewRand generates a new Rand object based on the given random type.
//
// Parameters:
// - rndType: an integer representing the random type.
//
// Return:
// - a pointer to a Rand object.
func NewRand(kind Kind) *Rand {
	return newRand(kind)
}

// CustomRand creates a new Rand object with a custom charset.
//
// Parameters:
// - charset: a string representing the custom charset.
//
// Return:
// - a pointer to a Rand object.
func CustomRand(charset string) *Rand {
	r := randPool.Get().(*Rand)
	r.kind = KindCustom
	r.charset = charset
	r.length = len(charset)
	return r
}

func newRand(kind Kind) *Rand {
	r := randPool.Get().(*Rand)
	r.kind = kind
	r.charset = loadCharset(kind)
	r.length = len(r.charset)
	return r
}
