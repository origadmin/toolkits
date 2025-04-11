/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package uuid provides a UUID (Universally Unique Identifier) implementation.
package uuid

import (
	"github.com/goexts/generic/settings"
	"github.com/google/uuid"

	"github.com/origadmin/toolkits/identifier"
)

// UUID represents a UUID generator
type UUID struct {
	generator func() (uuid.UUID, error)
}

var (
	bitSize = len(uuid.New().String())
)

// init registers UUID as an identifier
func init() {
	identifier.RegisterStringIdentifier(New())
}

// Name returns the name of the UUID generator
func (u UUID) Name() string {
	return "uuid"
}

// GenerateString generates a new UUID.
func (u UUID) GenerateString() string {
	return uuid.Must(u.generator()).String()
}

// ValidateString checks if the provided ID is a valid UUID.
func (u UUID) ValidateString(id string) bool {
	if len(id) != bitSize {
		return false
	}
	_, err := uuid.Parse(id)
	return err == nil
}

// Size returns the size of the UUID in bits
func (u UUID) Size() int {
	return bitSize
}

type Options struct {
	Version   int
	Namespace uuid.UUID
	Name      []byte
	Domain    interface{}
}

type Option = func(*Options)

// New creates a new UUID instance
func New(opts ...Option) *UUID {
	o := settings.Apply(&Options{}, opts)
	var generator func() (uuid.UUID, error)
	switch o.Version {
	case 1:
		generator = uuid.NewUUID
	case 2:
		switch o.Domain {
		case uuid.Person:
			generator = uuid.NewDCEPerson
		case uuid.Group:
			generator = uuid.NewDCEGroup
		default:
			panic("uuid: invalid domain")
		}
	case 3:
		generator = func() (uuid.UUID, error) {
			id := uuid.NewMD5(o.Namespace, o.Name)
			return id, nil
		}
	case 5:
		generator = func() (uuid.UUID, error) {
			return uuid.NewSHA1(o.Namespace, o.Name), nil
		}
	case 6:
		generator = uuid.NewV6
	case 7:
		generator = uuid.NewV7
	default:
		generator = uuid.NewRandom
	}

	return &UUID{
		generator: generator,
	}
}
