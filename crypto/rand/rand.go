// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package rand provides the random string
package rand

import (
	"math/rand/v2"
	"strings"
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
	randType Kind
	length   int
	charset  string
}

// predefined charsets for different random types
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
func (r Rand) RandBytes(size int) []byte {
	ret := make([]byte, 0, size)
	for ; size > 0; size-- {
		ret[size-1] = r.charset[rand.IntN(r.length)]
	}
	return ret
}

func loadCharset(rand Kind) string {
	var charset string
	if rand >= KindAll {
		return randCharset
	}
	var sta, end int
	for i := 0; i < randomMax-1; i++ {
		if rand&(1<<uint(i)) != 0 {
			sta, end = getStringIndex(i)
			charset += randCharset[sta:end]
		}
	}
	return charset
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
func NewRand(rndType Kind) *Rand {
	charset := loadCharset(rndType)
	return &Rand{
		randType: rndType,
		length:   len(charset),
		charset:  charset,
	}
}

// CustomRand creates a new Rand object with a custom charset.
//
// Parameters:
// - charset: a string representing the custom charset.
//
// Return:
// - a pointer to a Rand object.
func CustomRand(charset string) *Rand {
	return &Rand{
		randType: KindCustom,
		length:   len(charset),
		charset:  charset,
	}
}
