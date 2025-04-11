/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package shortid

import (
	"github.com/goexts/generic/settings"
	"github.com/sony/sonyflake"

	"github.com/origadmin/toolkits/identifier"
)

const (
	bitSize = 64 // bitSize is used to store the length of generated ID.
)

// init registers the Snowflake generator with the ident package and initializes bitSize.
func init() {
	s := New()
	identifier.RegisterNumberIdentifier(s)
}

type Sonyflake struct {
	generator *sonyflake.Sonyflake
}

// Name returns the name of the generator.
func (s Sonyflake) Name() string {
	return "sonyflake"
}

// GenerateNumber generates a new Sonyflake ID as an int64.
func (s Sonyflake) GenerateNumber() int64 {
	id, err := s.generator.NextID()
	if err != nil {
		return 0
	}
	return int64(id)
}

// ValidateNumber checks if the provided ID is a valid Sonyflake ID.
func (s Sonyflake) ValidateNumber(id int64) bool {
	return id != 0
}

// Size returns the bit size of the generated Sonyflake ID.
func (s Sonyflake) Size() int {
	return bitSize
}

type Options = sonyflake.Settings

type Option = func(*Options)

// New creates a new Sonyflake generator with a unique node.
func New(opts ...Option) *Sonyflake {
	o := settings.Apply(&Options{}, opts)
	generator, err := sonyflake.New(*o)
	if err != nil {
		panic(err)
	}
	return &Sonyflake{
		generator: generator,
	}
}
