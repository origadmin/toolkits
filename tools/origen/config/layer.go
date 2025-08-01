// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package config implements the functions, types, and interfaces for the module.
package config

type Layer struct {
	Type     Type
	Path     string
	Template Template
	Mods     []Mod
	Children []Layer
}
