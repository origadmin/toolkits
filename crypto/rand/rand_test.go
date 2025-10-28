/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package rand

import (
	"bytes"
	"fmt"
	"math"
	"sync"
	"testing"
)

func TestNewGenerator(t *testing.T) {
	testCases := []struct {
		name    string
		kind    Kind
		charset string
	}{
		{"KindDigit", KindDigit, Digits},
		{"KindLowerCase", KindLowerCase, Lowercase},
		{"KindUpperCase", KindUpperCase, Uppercase},
		{"KindSymbol", KindSymbol, Symbols},
		{"KindAlphanumeric", KindDigit | KindLowerCase | KindUpperCase, Digits + Lowercase + Uppercase},
		{"KindAllWithSymbols", KindDigit | KindLowerCase | KindUpperCase | KindSymbol, Digits + Lowercase + Uppercase + Symbols},
		{"Digit and Symbol", KindDigit | KindSymbol, Digits + Symbols}, // Test a non-pre-calculated combination
		{"UnknownKind", Kind(999), Digits + Lowercase + Uppercase}, // Fallback to KindAlphanumeric
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gen := NewGenerator(tc.kind).(*randGenerator) // Type assertion for internal access
			if gen.charset != tc.charset {
				t.Errorf("NewGenerator(%v) charset mismatch:\n got %q,\nwant %q", tc.kind, gen.charset, tc.charset)
			}
			expectedMaxByte := byte(256 - (256 % len(tc.charset)))
			if len(tc.charset) == 0 {
				expectedMaxByte = 0
			}
			if gen.maxByte != expectedMaxByte {
				t.Errorf("NewGenerator(%v) maxByte mismatch: got %d, want %d", tc.kind, gen.maxByte, expectedMaxByte)
			}
		})
	}
}

func TestNewGeneratorWithCharset(t *testing.T) {
	customCharset := "abcDE123!@#"
	gen := NewGeneratorWithCharset(customCharset).(*randGenerator)
	if gen.charset != customCharset {
		t.Errorf("NewGeneratorWithCharset charset mismatch: got %q, want %q", gen.charset, customCharset)
	}
	expectedMaxByte := byte(256 - (256 % len(customCharset)))
	if gen.maxByte != expectedMaxByte {
			t.Errorf("NewGeneratorWithCharset maxByte mismatch: got %d, want %d", gen.maxByte, expectedMaxByte)
	}

	emptyCharset := ""
	gen = NewGeneratorWithCharset(emptyCharset).(*randGenerator)
	if gen.charset != emptyCharset {
		t.Errorf("NewGeneratorWithCharset empty charset mismatch: got %q, want %q", gen.charset, emptyCharset)
	}
	if gen.maxByte != 0 {
		t.Errorf("NewGeneratorWithCharset empty maxByte mismatch: got %d, want %d", gen.maxByte, 0)
	}
}

func TestRandBytes(t *testing.T) {
	testCases := []struct {
		name    string
		kind    Kind
		size    int
		charset string
	}{
		{"DigitBytes", KindDigit, 10, Digits},
		{"LowerCaseBytes", KindLowerCase, 15, Lowercase},
		{"UpperCaseBytes", KindUpperCase, 20, Uppercase},
		{"SymbolBytes", KindSymbol, 5, Symbols},
		{"AlphanumericBytes", KindAlphanumeric, 25, Digits + Lowercase + Uppercase},
		{"AllWithSymbolsBytes", KindAllWithSymbols, 30, Digits + Lowercase + Uppercase + Symbols},
		{"ZeroSizeBytes", KindAlphanumeric, 0, Digits + Lowercase + Uppercase},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gen := NewGenerator(tc.kind)
			b, err := gen.RandBytes(tc.size)
			if err != nil {
				t.Fatalf("RandBytes returned an error: %v", err)
			}
			if len(b) != tc.size {
				t.Errorf("RandBytes length mismatch: got %d, want %d", len(b), tc.size)
			}
			for _, char := range b {
				if tc.size > 0 && !bytes.ContainsRune([]byte(tc.charset), rune(char)) {
					t.Errorf("RandBytes generated character not in charset: %q (charset: %q)", char, tc.charset)
				}
			}
		})
	}

	t.Run("CustomCharsetBytes", func(t *testing.T) {
		customCharset := "XYZ789"
		gen := NewGeneratorWithCharset(customCharset)
		b, err := gen.RandBytes(10)
		if err != nil {
			t.Fatalf("CustomRandBytes returned an error: %v", err)
		}
		if len(b) != 10 {
			t.Errorf("CustomRandBytes length mismatch: got %d, want %d", len(b), 10)
		}
		for _, char := range b {
			if !bytes.ContainsRune([]byte(customCharset), rune(char)) {
				t.Errorf("CustomRandBytes generated character not in charset: %q (charset: %q)", char, customCharset)
			}
		}
	})

	t.Run("EmptyCharsetBytes", func(t *testing.T) {
		gen := NewGeneratorWithCharset("")
		b, err := gen.RandBytes(10)
		if err != nil {
			t.Fatalf("EmptyCharsetBytes returned an error: %v", err)
		}
		if len(b) != 0 { // Should return an empty slice
			t.Errorf("EmptyCharsetBytes length mismatch: got %d, want %d", len(b), 0)
		}
	})
}

func TestRandString(t *testing.T) {
	testCases := []struct {
		name    string
		kind    Kind
		size    int
		charset string
	}{
		{"DigitString", KindDigit, 10, Digits},
		{"LowerCaseString", KindLowerCase, 15, Lowercase},
		{"UpperCaseString", KindUpperCase, 20, Uppercase},
		{"SymbolString", KindSymbol, 5, Symbols},
		{"AlphanumericString", KindAlphanumeric, 25, Digits + Lowercase + Uppercase},
		{"AllWithSymbolsString", KindAllWithSymbols, 30, Digits + Lowercase + Uppercase + Symbols},
		{"ZeroSizeString", KindAlphanumeric, 0, Digits + Lowercase + Uppercase},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gen := NewGenerator(tc.kind)
			s, err := gen.RandString(tc.size)
			if err != nil {
				t.Fatalf("RandString returned an error: %v", err)
			}
			if len(s) != tc.size {
				t.Errorf("RandString length mismatch: got %d, want %d", len(s), tc.size)
			}
			for _, char := range s {
				if tc.size > 0 && !bytes.ContainsRune([]byte(tc.charset), char) {
					t.Errorf("RandString generated character not in charset: %q (charset: %q)", char, tc.charset)
				}
			}
		})
	}

	t.Run("CustomCharsetString", func(t *testing.T) {
		customCharset := "XYZ789"
		gen := NewGeneratorWithCharset(customCharset)
		s, err := gen.RandString(10)
		if err != nil {
			t.Fatalf("CustomRandString returned an error: %v", err)
		}
		if len(s) != 10 {
			t.Errorf("CustomRandString length mismatch: got %d, want %d", len(s), 10)
		}
		for _, char := range s {
			if !bytes.ContainsRune([]byte(customCharset), char) {
				t.Errorf("CustomRandString generated character not in charset: %q (charset: %q)", char, customCharset)
			}
		}
	})

	t.Run("EmptyCharsetString", func(t *testing.T) {
		gen := NewGeneratorWithCharset("")
		s, err := gen.RandString(10)
		if err != nil {
			t.Fatalf("EmptyCharsetString returned an error: %v", err)
		}
		if len(s) != 0 {
			t.Errorf("EmptyCharsetString length mismatch: got %d, want %d", len(s), 0)
		}
	})
}

func TestRead(t *testing.T) {
	testCases := []struct {
		name    string
		kind    Kind
		size    int
		charset string
	}{
		{"DigitRead", KindDigit, 10, Digits},
		{"LowerCaseRead", KindLowerCase, 15, Lowercase},
		{"UpperCaseRead", KindUpperCase, 20, Uppercase},
		{"SymbolRead", KindSymbol, 5, Symbols},
		{"AlphanumericRead", KindAlphanumeric, 25, Digits + Lowercase + Uppercase},
		{"AllWithSymbolsRead", KindAllWithSymbols, 30, Digits + Lowercase + Uppercase + Symbols},
		{"ZeroSizeRead", KindAlphanumeric, 0, Digits + Lowercase + Uppercase},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gen := NewGenerator(tc.kind)
			p := make([]byte, tc.size)
			n, err := gen.Read(p)
			if err != nil {
				t.Fatalf("Read returned an error: %v", err)
			}
			if n != tc.size {
				t.Errorf("Read bytes count mismatch: got %d, want %d", n, tc.size)
			}
			for _, char := range p {
				if tc.size > 0 && !bytes.ContainsRune([]byte(tc.charset), rune(char)) {
					t.Errorf("Read generated character not in charset: %q (charset: %q)", char, tc.charset)
				}
			}
		})
	}

	t.Run("CustomCharsetRead", func(t *testing.T) {
		customCharset := "XYZ789"
		gen := NewGeneratorWithCharset(customCharset)
		p := make([]byte, 10)
		n, err := gen.Read(p)
		if err != nil {
			t.Fatalf("CustomCharsetRead returned an error: %v", err)
		}
		if n != 10 {
			t.Errorf("CustomCharsetRead bytes count mismatch: got %d, want %d", n, 10)
		}
		for _, char := range p {
			if !bytes.ContainsRune([]byte(customCharset), rune(char)) {
				t.Errorf("CustomCharsetRead generated character not in charset: %q (charset: %q)", char, customCharset)
			}
		}
	})

	t.Run("EmptyCharsetRead", func(t *testing.T) {
		gen := NewGeneratorWithCharset("")
		p := make([]byte, 10)
		n, err := gen.Read(p)
		if err != nil {
			t.Fatalf("EmptyCharsetRead returned an error: %v", err)
		}
		if n != 0 {
			t.Errorf("EmptyCharsetRead bytes count mismatch: got %d, want %d", n, 0)
		}
	})
}

func TestRandomBytes(t *testing.T) {
	size := 10
	b, err := RandomBytes(size)
	if err != nil {
		t.Fatalf("RandomBytes returned an error: %v", err)
	}
	if len(b) != size {
		t.Errorf("RandomBytes length mismatch: got %d, want %d", len(b), size)
	}
	// Verify characters are from alphanumeric charset
	alphanumericCharset := Digits + Lowercase + Uppercase
	for _, char := range b {
		if !bytes.ContainsRune([]byte(alphanumericCharset), rune(char)) {
			t.Errorf("RandomBytes generated character not in alphanumeric charset: %q", char)
		}
	}

	// Test zero size
	b, err = RandomBytes(0)
	if err != nil {
		t.Fatalf("RandomBytes (size 0) returned an error: %v", err)
	}
	if len(b) != 0 {
		t.Errorf("RandomBytes (size 0) length mismatch: got %d, want %d", len(b), 0)
	}
}

func TestRandomString(t *testing.T) {
	size := 10
	s, err := RandomString(size)
	if err != nil {
		t.Fatalf("RandomString returned an error: %v", err)
	}
	if len(s) != size {
		t.Errorf("RandomString length mismatch: got %d, want %d", len(s), size)
	}
	// Verify characters are from alphanumeric charset
	alphanumericCharset := Digits + Lowercase + Uppercase
	for _, char := range s {
		if !bytes.ContainsRune([]byte(alphanumericCharset), char) {
			t.Errorf("RandomString generated character not in alphanumeric charset: %q", char)
		}
	}

	// Test zero size
	s, err = RandomString(0)
	if err != nil {
		t.Fatalf("RandomString (size 0) returned an error: %v", err)
	}
	if len(s) != 0 {
		t.Errorf("RandomString (size 0) length mismatch: got %d, want %d", len(s), 0)
	}
}

func TestRandomBytesWithSymbols(t *testing.T) {
	size := 10
	b, err := RandomBytesWithSymbols(size)
	if err != nil {
		t.Fatalf("RandomBytesWithSymbols returned an error: %v", err)
	}
	if len(b) != size {
		t.Errorf("RandomBytesWithSymbols length mismatch: got %d, want %d", len(b), size)
	}
	// Verify characters are from alphanumeric + symbols charset
	charset := Digits + Lowercase + Uppercase + Symbols
	for _, char := range b {
		if !bytes.ContainsRune([]byte(charset), rune(char)) {
			t.Errorf("RandomBytesWithSymbols generated character not in charset: %q", char)
		}
	}
}

func TestRandomStringWithSymbols(t *testing.T) {
	size := 10
	s, err := RandomStringWithSymbols(size)
	if err != nil {
		t.Fatalf("RandomStringWithSymbols returned an error: %v", err)
	}
	if len(s) != size {
		t.Errorf("RandomStringWithSymbols length mismatch: got %d, want %d", len(s), size)
	}
	// Verify characters are from alphanumeric + symbols charset
	charset := Digits + Lowercase + Uppercase + Symbols
	for _, char := range s {
		if !bytes.ContainsRune([]byte(charset), char) {
			t.Errorf("RandomStringWithSymbols generated character not in charset: %q", char)
		}
	}
}

func TestConcurrentGeneration(t *testing.T) {
	gen := NewGenerator(KindAlphanumeric)
	numGoroutines := 100
	length := 100
	var wg sync.WaitGroup
	errors := make(chan error, numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			s, err := gen.RandString(length)
			if err != nil {
				errors <- err
				return
			}
			if len(s) != length {
				errors <- fmt.Errorf("concurrent RandString length mismatch: got %d, want %d", len(s), length)
				return
			}
			alphanumericCharset := Digits + Lowercase + Uppercase
			for _, char := range s {
				if !bytes.ContainsRune([]byte(alphanumericCharset), char) {
					errors <- fmt.Errorf("concurrent RandString generated character not in alphanumeric charset: %q", char)
					return
				}
			}
		}()
	}

	wg.Wait()
	close(errors)

	for err := range errors {
		if err != nil {
			t.Error(err)
		}
	}
}

func TestCharsetsInitialization(t *testing.T) {
	// This test implicitly checks the init() function
	if charsets[KindDigit] != Digits {
		t.Errorf("KindDigit charset not initialized correctly: got %q, want %q", charsets[KindDigit], Digits)
	}
	if charsets[KindLowerCase] != Lowercase {
		t.Errorf("KindLowerCase charset not initialized correctly: got %q, want %q", charsets[KindLowerCase], Lowercase)
	}
	if charsets[KindUpperCase] != Uppercase {
		t.Errorf("KindUpperCase charset not initialized correctly: got %q, want %q", charsets[KindUpperCase], Uppercase)
	}
	if charsets[KindSymbol] != Symbols {
		t.Errorf("KindSymbol charset not initialized correctly: got %q, want %q", charsets[KindSymbol], Symbols)
	}

	// Test pre-calculated combined charsets
	alphanumeric := Digits + Lowercase + Uppercase
	if charsets[KindDigit|KindLowerCase|KindUpperCase] != alphanumeric {
		t.Errorf("Alphanumeric charset not initialized correctly: got %q, want %q", charsets[KindDigit|KindLowerCase|KindUpperCase], alphanumeric)
	}

	allWithSymbols := alphanumeric + Symbols
	if charsets[KindDigit|KindLowerCase|KindUpperCase|KindSymbol] != allWithSymbols {
		t.Errorf("AllWithSymbols charset not initialized correctly: got %q, want %q", charsets[KindDigit|KindLowerCase|KindUpperCase|KindSymbol], allWithSymbols)
	}
}

// TestRandDistribution checks for a reasonably uniform distribution of generated characters.
// This is a statistical test and might occasionally fail due to pure chance, but should generally pass.
func TestRandDistribution(t *testing.T) {
	charset := Digits + Lowercase + Uppercase + Symbols
	gen := NewGeneratorWithCharset(charset)

	sampleSize := 100000 // Generate a large number of characters
	generated, err := gen.RandString(sampleSize)
	if err != nil {
		t.Fatalf("RandString failed: %v", err)
	}

	counts := make(map[rune]int)
	for _, r := range generated {
		counts[r]++
	}

	expectedAvg := float64(sampleSize) / float64(len(charset))
	// Allow for a deviation of +/- 20% from the expected average for this statistical test.
	// A more rigorous test, like a chi-squared test, would be better but is more complex.
	deviationTolerance := 0.20 * expectedAvg

	for _, r := range charset {
		count := counts[r]
		if math.Abs(float64(count)-expectedAvg) > deviationTolerance {
			t.Errorf("Character %q count %d deviates too much from expected average %.2f (tolerance %.2f)", r, count, expectedAvg, deviationTolerance)
		}
	}
}

// TestGeneratorProducesDifferentValues verifies that successive calls to RandString on the same Generator instance produce different outputs.
func TestGeneratorProducesDifferentValues(t *testing.T) {
	gen := NewGenerator(KindAlphanumeric)
	length := 32 // Sufficiently long to make collisions highly improbable

	val1, err := gen.RandString(length)
	if err != nil {
		t.Fatalf("First RandString call failed: %v", err)
	}

	val2, err := gen.RandString(length)
	if err != nil {
		t.Fatalf("Second RandString call failed: %v", err)
	}

	if val1 == val2 {
		t.Errorf("Generator produced identical values on successive calls: %q", val1)
	}

	// Also test with RandBytes
	bytes1, err := gen.RandBytes(length)
	if err != nil {
		t.Fatalf("First RandBytes call failed: %v", err)
	}
	bytes2, err := gen.RandBytes(length)
	if err != nil {
		t.Fatalf("Second RandBytes call failed: %v", err)
	}

	if bytes.Equal(bytes1, bytes2) {
		t.Errorf("Generator produced identical byte slices on successive calls: %v", bytes1)
	}
}
