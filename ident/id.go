// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package ident provides the helpers functions.
package ident

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/rs/xid"
)

var generator Ident

// Ident is the interface of ident.
type Ident interface {
	Gen() string
	Validate(id string) bool
	Size() int
}

type ident struct {
	gen      func() string
	validate func(id string) bool
	size     int
}

// Gen method generates an identifier.
// It returns a string which represents the generated identifier.
func (obj ident) Gen() string {
	return obj.gen()
}

// Validate method checks if the given identifier is valid.
// It takes an argument:
//   - id: a string, the identifier to be validated.
//
// The function returns a boolean indicating whether the given identifier is valid or not.
func (obj ident) Validate(id string) bool {
	return obj.validate(id)
}

// Size method returns the size of the identifier.
// It returns an integer representing the size.
func (obj ident) Size() int {
	return obj.size
}

func init() {
	generator = ident{
		gen:      MustNewULID,
		validate: ValidateULID,
		size:     len(MustNewULID()),
	}
}

// Use sets the generator ident.
func Use(ident Ident) {
	generator = ident
}

// Default method returns the default generator ident.
func Default() Ident {
	return generator
}

// GenID The function "GenID" generates a new unique identifier (XID) and returns it as a string.
func GenID() string {
	return generator.Gen()
}

// GenSize The function "GenSize" returns the size of the generated identifier
func GenSize() int {
	return generator.Size()
}

// Validate The function "Validate" checks whether the given identifier is valid or not.
func Validate(id string) bool {
	return generator.Validate(id)
}

// NewXID The function "NewXID" generates a new unique identifier (XID) and returns it as a string.
func NewXID() string {
	return xid.New().String()
}

// ParseXID The function "ParseXID" parses a string into an XID and returns it.
func ParseXID(id string) (xid.ID, error) {
	return xid.FromString(id)
}

func ValidateXID(id string) bool {
	_, err := ParseXID(id)
	return err == nil
}

// MustNewUUID The function generates a new UUID and panics if there is an error.
func MustNewUUID() string {
	return uuid.Must(uuid.NewRandom()).String()
}

// ParseUUID The function "ParseUUID" parses a UUID string and returns the UUID value and an error.
func ParseUUID(id string) (uuid.UUID, error) {
	return uuid.Parse(id)
}

// ValidateUUID The function validates a UUID.
func ValidateUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

// MustNewULID The function generates a new ULID and returns it as a string.
func MustNewULID() string {
	return ulid.Make().String()
}

// ParseULID The function parses a ULID from a string.
func ParseULID(id string) (ulid.ULID, error) {
	return ulid.Parse(id)
}

// ValidateULID The function validates a ULID.
func ValidateULID(id string) bool {
	_, err := ulid.Parse(id)
	return err == nil
}

var (
	_ = Use
	_ = GenID
	_ = GenSize
	_ = Validate
	_ = ParseXID
	_ = ParseUUID
	_ = ParseULID
	_ = ValidateXID
	_ = ValidateUUID
	_ = ValidateULID
)
