/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package idgen

import (
	"math/rand/v2"
	"sync/atomic"
)

const (
	bitSize = 8 // bitSize is used to store the length of generated ID.
)

func init() {
	RegisterNumberIdentifier(NewNumber(StartNumber))
}

type NumberGenerator struct {
	start int64
}

// StartNumber is the start number of the randNumber.
var StartNumber = int64(rand.Int32()) << 32

// Name returns the name of the object identifier.
func (obj *NumberGenerator) Name() string {
	return "number"
}

// Number method generates an identifier.
func (obj *NumberGenerator) Number() int64 {
	num := atomic.AddInt64(&obj.start, 1)
	return num
}

// ValidateNumber method checks if the given identifier is valid.
func (obj *NumberGenerator) ValidateNumber(id int64) bool {
	return id > 0
}

// Size method returns the size of the identifier.
func (obj *NumberGenerator) Size() int {
	return bitSize
}

// NewNumber method creates a new instance of the randNumber type.
func NewNumber(sta int64) NumberIdentifier {
	return &NumberGenerator{
		start: sta,
	}
}
