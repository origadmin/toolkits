/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package shortid

import (
	"cmp"
	"math/rand/v2"

	"github.com/teris-io/shortid"

	"github.com/origadmin/toolkits/identifier"
)

var (
	bitSize = 0 // bitSize is used to store the length of generated ID.
)

// init registers the Snowflake generator with the ident package and initializes bitSize.
func init() {
	s := New()
	bitSize = len(s.GenerateString())
	identifier.RegisterStringIdentifier(s)
}

type ShortID struct {
	generator *shortid.Shortid
}

func (s ShortID) Generate() string {
	return s.GenerateString()
}

func (s ShortID) Validate(id string) bool {
	return s.ValidateString(id)
}

// Name returns the name of the generator.
func (s ShortID) Name() string {
	return "shortid"
}

// GenerateString generates a new ShortID ID as a string.
func (s ShortID) GenerateString() string {
	ret, err := s.generator.Generate()
	if err != nil {
		return ""
	}
	return ret
}

// ValidateString checks if the provided ID is a valid ShortID ID.
func (s ShortID) ValidateString(id string) bool {
	return id != ""
}

// Size returns the bit size of the generated ShortID ID.
func (s ShortID) Size() int {
	return bitSize
}

type Options struct {
	Worker   uint8
	Alphabet string
	Seed     uint64
}

// New creates a new ShortID generator with a unique node.
func New(ss ...*Options) *ShortID {
	ss = append(ss, &Options{})
	o := cmp.Or(ss...)
	if o.Worker > 31 {
		o.Worker = uint8(rand.Uint32N(31))
	}
	if o.Seed == 0 {
		o.Seed = rand.Uint64()
	}
	if o.Alphabet == "" {
		o.Alphabet = shortid.DefaultABC
	}
	generator, err := shortid.New(o.Worker, o.Alphabet, o.Seed)
	if err != nil {
		panic(err)
	}
	return &ShortID{
		generator: generator,
	}
}

var _ identifier.TypedIdentifier[string] = &ShortID{}
