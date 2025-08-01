// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package config implements the functions, types, and interfaces for the module.
package config

// Mod set the predefined modules to be used
type Mod struct {
	Name   string
	Repo   string
	Tag    string
	Branch string
	Commit string
}
