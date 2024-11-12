/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package sloge implements the functions, types, and interfaces for the module.
package sloge

//go:generate stringer -type=Format -trimprefix=Format
type Format int

const (
	// FormatJSON json format
	FormatJSON Format = iota
	// FormatText text format
	FormatText
	// FormatTint tint format
	FormatTint
	// FormatDev dev format
	FormatDev
)
