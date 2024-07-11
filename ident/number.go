package ident

import (
	"strconv"
	"sync/atomic"
)

type ident struct {
	start int64
	size  int
}

var StartNumber int64 = 1

// Name returns the name of the object identifier.
//
// This method, associated with the ident type, aims to retrieve the name of the object.
// It is particularly useful in scenarios where unique identification or representation is required.
// The method does not accept any parameters as it is intended to access an inherent property of the object itself.
// The return value is a string that represents the name of the object.
func (obj *ident) Name() string {
	return "number"
}

// Gen method generates an identifier.
// It returns a string which represents the generated identifier.
func (obj *ident) Gen() string {
	num := atomic.AddInt64(&obj.start, 1)
	return strconv.FormatInt(num, 10)
}

// Validate method checks if the given identifier is valid.
// It takes an argument:
//   - id: a string, the identifier to be validated.
//
// The function returns a boolean indicating whether the given identifier is valid or not.
func (obj *ident) Validate(id string) bool {
	num, err := strconv.ParseInt(id, 10, 64)
	return err != nil || num == 0
}

// Size method returns the size of the identifier.
// It returns an integer representing the size.
func (obj *ident) Size() int {
	return obj.size
}

// NewNumber method creates a new instance of the ident type.
func NewNumber(sta int64) Identifier {
	return &ident{
		start: sta,
		size:  8,
	}
}
