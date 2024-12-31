/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security implements the functions, types, and interfaces for the module.

// Package security is a package that provides security-related functions and types.
package security

import (
	"context"
)

// TokenType represents the type of token.
//
//go:generate go run -mod=mod golang.org/x/tools/cmd/stringer@latest -type=TokenType -trimprefix=ContextType -output=security_string.go
type TokenType int

// ContextType constants represent the different types of context.
const (
	// ContextTypeContext represents the context type for the context.
	ContextTypeContext TokenType = iota
	// ContextTypeClientHeader represents the context type for the header.
	ContextTypeClientHeader
	// ContextTypeServerHeader represents the context type for the header.
	ContextTypeServerHeader
	// ContextTypeMetadataClient represents the context type for the metadata.
	ContextTypeMetadataClient
	// ContextTypeMetadata represents the context type for the metadata.
	ContextTypeMetadata
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

const (
	// HeaderAuthorize is the name of the authorization header.
	HeaderAuthorize = "Authorization"
	// HeaderContentType is the name of the content type header.
	HeaderContentType = "Content-Type"
	// HeaderContentLength is the name of the content length header.
	HeaderContentLength = "Content-Length"
	// HeaderUserAgent is the name of the user agent header.
	HeaderUserAgent = "User-Agent"
	// HeaderReferer is the name of the referer header.
	HeaderReferer = "Referer"
	// HeaderOrigin is the name of the origin header.
	HeaderOrigin = "Origin"
)

// Security represents the security interface.
type Security interface {
	Authenticator
	Authorizer
}

type security struct {
	authenticator Authenticator
	authorizer    Authorizer
}

func (s security) Authenticate(ctx context.Context, s2 string) (Claims, error) {
	//TODO implement me
	panic("implement me")
}

func (s security) AuthenticateContext(ctx context.Context, tokenType TokenType) (Claims, error) {
	//TODO implement me
	panic("implement me")
}

func (s security) SetPolicies(ctx context.Context, policies map[string]any, roles map[string]any) error {
	//TODO implement me
	panic("implement me")
}

func (s security) Authorized(ctx context.Context, claims Claims, object string, action string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s security) AuthorizedWithDomain(ctx context.Context, claims Claims, domain string, object string, action string) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func (s security) xxxxAuthorized(ctx context.Context, policy Policy) (bool, error) {
	//TODO implement me
	panic("implement me")
}

func NewSecurity(authenticator Authenticator, authorizer Authorizer) Security {
	return &security{
		authenticator: authenticator,
		authorizer:    authorizer,
	}
}
