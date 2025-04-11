/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package snowflake provides a unique ID generation based on Twitter Snowflake algorithm.
package snowflake

import (
	"math/rand/v2"

	"github.com/bwmarrin/snowflake"
	"github.com/goexts/generic/settings"

	"github.com/origadmin/toolkits/identifier"
)

// nodeNumber range from 0 to 1023 is used for generating unique generator ID.
var (
	bitSize = 0 // bitSize is used to store the length of generated ID.
)

func init() {
	s := New()
	bitSize = len(s.GenerateString())
	identifier.RegisterNumberIdentifier(s)
	identifier.RegisterStringIdentifier(s)
}

// Snowflake represents a Snowflake generator with a unique generator.
type Snowflake struct {
	generator *snowflake.Node
}

// Name returns the name of the generator.
func (s Snowflake) Name() string {
	return "snowflake"
}

// GenerateString generates a new Snowflake ID as a string.
func (s Snowflake) GenerateString() string {
	return s.generator.Generate().String()
}

// GenerateNumber generates a new Snowflake ID as an int64.
func (s Snowflake) GenerateNumber() int64 {
	return s.generator.Generate().Int64()
}

// ValidateString checks if the provided ID is a valid Snowflake ID (string).
func (s Snowflake) ValidateString(id string) bool {
	if len(id) != bitSize {
		return false
	}
	_, err := snowflake.ParseString(id)
	return err == nil
}

// ValidateNumber checks if the provided ID is a valid Snowflake ID (int64).
func (s Snowflake) ValidateNumber(id int64) bool {
	return id > 0
}

// Size returns the bit size of the generated Snowflake ID.
func (s Snowflake) Size() int {
	return bitSize
}

type Options struct {
	Node int64
}

type Option = func(options *Options)

// New creates a new Snowflake generator with a unique generator.
func New(opts ...Option) *Snowflake {
	o := settings.Apply(&Options{}, opts)
	if o.Node < 0 || o.Node > 1023 {
		o.Node = rand.Int64N(1023)
	}
	generator, err := snowflake.NewNode(o.Node)
	if err != nil {
		panic(err)
	}
	return &Snowflake{
		generator: generator,
	}
}

type NumberSnowflake struct {
	*Snowflake
}

func (n NumberSnowflake) Generate() int64 {
	return n.GenerateNumber()
}

func (n NumberSnowflake) Validate(t int64) bool {
	return n.ValidateNumber(t)
}

func NewNumber(opts ...Option) *NumberSnowflake {
	return &NumberSnowflake{
		Snowflake: New(opts...),
	}
}

type StringSnowflake struct {
	*Snowflake
}

func (s StringSnowflake) Generate() string {
	return s.GenerateString()
}

func (s StringSnowflake) Validate(id string) bool {
	return s.ValidateString(id)
}

func NewString(opts ...Option) *StringSnowflake {
	return &StringSnowflake{
		Snowflake: New(opts...),
	}
}

var _ identifier.TypedIdentifier[int64] = &NumberSnowflake{}
var _ identifier.TypedIdentifier[string] = &StringSnowflake{}
