/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package shortid

import (
	"cmp"

	"github.com/sony/sonyflake"

	"github.com/origadmin/toolkits/idgen"
)

const (
	bitSize = 8 // bitSize is used to store the length of generated ID.
)

// init registers the Snowflake generator with the ident package and initializes bitSize.
func init() {
	s := New()
	idgen.RegisterNumberIdentifier(s)
}

type Sonyflake struct {
	generator *sonyflake.Sonyflake
}

func (s Sonyflake) Number() int64 {
	id, err := s.generator.NextID()
	if err != nil {
		return 0
	}
	return int64(id)
}

func (s Sonyflake) ValidateNumber(id int64) bool {
	return id != 0
}

// Name returns the name of the generator.
func (s Sonyflake) Name() string {
	return "sonyflake"
}

// Size returns the bit size of the generated Sonyflake ID.
func (s Sonyflake) Size() int {
	return bitSize
}

type Setting = sonyflake.Settings

// New creates a new Sonyflake generator with a unique node.
func New(ss ...*Setting) *Sonyflake {
	ss = append(ss, &Setting{})
	o := cmp.Or(ss...)
	generator, err := sonyflake.New(*o)
	if err != nil {
		panic(err)
	}
	return &Sonyflake{
		generator: generator,
	}
}
