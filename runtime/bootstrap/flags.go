/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package bootstrap

import (
	"os"
	"time"
)

// Flags is a struct that holds the flags for the service
type Flags struct {
	ID          string
	Version     string
	ServiceName string
	StartTime   time.Time
	Metadata    map[string]string
}

// ServiceID returns the ID of the service
func (f Flags) ServiceID() string {
	return f.ID + "." + f.ServiceName
}

// DefaultFlags returns the default flags for the service
func DefaultFlags() Flags {
	id, _ := os.Hostname()
	return Flags{
		ID:        id,
		StartTime: time.Now(),
	}
}

// NewFlags returns a new set of flags for the service
func NewFlags(name string, version string) Flags {
	f := DefaultFlags()
	f.Version = version
	f.ServiceName = name
	return f
}
