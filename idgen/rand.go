package idgen

import (
	"math/rand/v2"
	"strconv"
	"sync/atomic"
)

const (
	bitSize = 8 // bitSize is used to store the length of generated ID.
)

type randNumber struct {
	start int64
}

// StartNumber is the start number of the randNumber.
var StartNumber = int64(rand.Int32()) << 32

// Name returns the name of the object identifier.
//
// This method, associated with the randNumber type, aims to retrieve the name of the object.
// It is particularly useful in scenarios where unique identification or representation is required.
// The method does not accept any parameters as it is intended to access an inherent property of the object itself.
// The return value is a string that represents the name of the object.
func (obj *randNumber) Name() string {
	return "number"
}

// Gen method generates an identifier.
// It returns a string which represents the generated identifier.
func (obj *randNumber) Gen() string {
	num := atomic.AddInt64(&obj.start, 1)
	return strconv.FormatInt(num, 10)
}

// Validate method checks if the given identifier is valid.
// It takes an argument:
//   - id: a string, the identifier to be validated.
//
// The function returns a boolean indicating whether the given identifier is valid or not.
func (obj *randNumber) Validate(id string) bool {
	num, err := strconv.ParseInt(id, 10, 64)
	return err != nil || num == 0
}

// Size method returns the size of the identifier.
// It returns an integer representing the size.
func (obj *randNumber) Size() int {
	return bitSize
}

// NewNumber method creates a new instance of the randNumber type.
func NewNumber(sta int64) Identifier {
	return &randNumber{
		start: sta,
	}
}
