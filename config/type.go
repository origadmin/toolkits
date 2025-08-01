// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package config implements the functions, types, and interfaces for the module.
package config

//go:generate stringer -type=Type -trimprefix=Type
type Type int

const (
	TypeProject Type = iota
	TypeMod
	TypeService
	TypeBiz
	TypeDal
)
