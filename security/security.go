/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.
package security

// TokenType represents the type of token.
type TokenType int

// ContextType constants represent the different types of context.
const (
	// ContextTypeContext represents the context type for the context.
	ContextTypeContext TokenType = iota
	// ContextTypeHeader represents the context type for the header.
	ContextTypeHeader
	// ContentTypeMetadata represents the context type for the metadata.
	ContentTypeMetadata
	// ContextTypeQuery represents the context type for the query.
	ContextTypeQuery
	// ContextTypeCookie represents the context type for the cookie.
	ContextTypeCookie
	// ContextTypeParam represents the context type for the parameter.
	ContextTypeParam
	// ContextTypeForm represents the context type for the form.
	ContextTypeForm
	// ContextTypeBody represents the context type for the body.
	ContextTypeBody
	// ContextTypeSession represents the context type for the session.
	ContextTypeSession
	// ContextTypeUnknown represents an unknown context type.
	ContextTypeUnknown
)

// HeaderAuthorize is the name of the authorization header.
const (
	// HeaderAuthorize is the name of the authorization header.
	HeaderAuthorize = "Authorization"
)
