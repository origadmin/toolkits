/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

type TokenType int

const (
	ContextTypeContext TokenType = iota
	ContextTypeHeader
)

const (
	HeaderAuthorize = "Authorization"
)
