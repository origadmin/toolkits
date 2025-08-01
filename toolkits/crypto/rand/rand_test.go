/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package rand

import (
	"bytes"
	"fmt" // Added for fmt.Errorf
	"sync"
	"testing"
)

func TestNewRand(t *testing.T) {
	testCases := []struct {
		name    string
		kind    Kind
		charset string
	}{
		{"KindDigit", KindDigit, Digits},
		{"KindLowerCase", KindLowerCase, Lowercase},
		{"KindUpperCase", KindUpperCase, Uppercase},
		{"KindSymbol", KindSymbol, Symbols},
		{"KindAlphanumeric", KindAlphanumeric, Digits + Lowercase + Uppercase},
		{"KindAllWithSymbols", KindAllWithSymbols, Digits + Lowercase + Uppercase + Symbols},
		{"UnknownKind", Kind(999), Digits + Lowercase + Uppercase}, // Fallback to KindAlphanumeric
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gen := NewRand(tc.kind).(*randGenerator) // Type assertion for internal access
			if gen.charset != tc.charset {
				t.Errorf("NewRand(%v) charset mismatch: got %q, want %q", tc.kind, gen.charset, tc.charset)
			}
			if gen.length != len(tc.charset) {
				t.Errorf("NewRand(%v) length mismatch: got %d, want %d", tc.kind, gen.length, len(tc.charset))
			}
		})
	}
}

func TestCustomRand(t *testing.T) {
	customCharset := "abcDE123!@#"
	gen := CustomRand(customCharset).(*randGenerator)
	if gen.charset != customCharset {
		t.Errorf("CustomRand charset mismatch: got %q, want %q", gen.charset, customCharset)
	}
	if gen.length != len(customCharset) {
		t.Errorf("CustomRand length mismatch: got %d, want %d", gen.length, len(customCharset))
	}

	emptyCharset := ""
	gen = CustomRand(emptyCharset).(*randGenerator)
	if gen.charset != emptyCharset {
		t.Errorf("CustomRand empty charset mismatch: got %q, want %q", gen.charset, emptyCharset)
	}
	if gen.length != 0 {
		t.Errorf("CustomRand empty length mismatch: got %d, want %d", gen.length, 0)
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
			gen := NewRand(tc.kind)
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
		gen := CustomRand(customCharset)
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
		gen := CustomRand("")
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
			gen := NewRand(tc.kind)
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
		gen := CustomRand(customCharset)
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
		gen := CustomRand("")
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
			gen := NewRand(tc.kind)
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
		gen := CustomRand(customCharset)
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
		gen := CustomRand("")
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

func TestConcurrentGeneration(t *testing.T) {
	gen := NewRand(KindAlphanumeric)
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
	if charsets[KindAlphanumeric] != Digits+Lowercase+Uppercase {
		t.Errorf("KindAlphanumeric charset not initialized correctly: got %q, want %q", charsets[KindAlphanumeric], Digits+Lowercase+Uppercase)
	}
	if charsets[KindAllWithSymbols] != Digits+Lowercase+Uppercase+Symbols {
		t.Errorf("KindAllWithSymbols charset not initialized correctly: got %q, want %q", charsets[KindAllWithSymbols], Digits+Lowercase+Uppercase+Symbols)
	}
}
