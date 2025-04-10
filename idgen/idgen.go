/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package idgen provides the helpers functions.
package idgen

import (
	"sync"
)

// NumberIdentifier is the interface of randNumber.
type NumberIdentifier interface {
	Name() string
	Number() int64
	ValidateNumber(id int64) bool
	Size() int
}

// StringIdentifier is the interface of randString.
type StringIdentifier interface {
	Name() string
	String() string
	ValidateString(id string) bool
	Size() int
}

var (
	numberGenerators = make(map[string]NumberIdentifier)
	numberOnce       sync.Once
	stringGenerators = make(map[string]StringIdentifier)
	stringsOnce      sync.Once
)

// RegisterStringIdentifier sets the defaultIdentifier randNumber.
func RegisterStringIdentifier(identifier StringIdentifier) {
	stringsOnce.Do(func() {
		stringGenerators[identifier.Name()] = identifier
	})
}

func RegisterNumberIdentifier(identifier NumberIdentifier) {
	numberOnce.Do(func() {
		numberGenerators[identifier.Name()] = identifier
	})
}

func GetNumberIdentifier(name string) NumberIdentifier {
	return numberGenerators[name]
}

func GetStringIdentifier(name string) StringIdentifier {
	return stringGenerators[name]
}

// GenStringID The function "GenID" generates a new unique identifier and returns it as a string.
func GenStringID(name string) string {
	if gen, ok := stringGenerators[name]; ok {
		return gen.String()
	}
	return ""
}

func GenNumberID(name string) int64 {
	if gen, ok := numberGenerators[name]; ok {
		return gen.Number()
	}
	return 0
}

// Size The function "Size" returns the size of the generated identifier
func Size(name string) int {
	if gen, ok := stringGenerators[name]; ok {
		return gen.Size()
	}
	if gen, ok := numberGenerators[name]; ok {
		return gen.Size()
	}
	return 0
}

// Validate The function "Validate" checks whether the given identifier is valid or not.
func Validate[T ~string | ~int64](name string, id T) bool {
	switch v := any(id).(type) {
	case string:
		if gen, ok := stringGenerators[name]; ok {
			return gen.ValidateString(v)
		}
	case int64:
		if gen, ok := numberGenerators[name]; ok {
			return gen.ValidateNumber(v)
		}
	}
	return false
}

var (
	_ = GenStringID
	_ = GenNumberID
	_ = Size
	_ = Validate
)
