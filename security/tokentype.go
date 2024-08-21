// Copyright (c) 2024 GodCong. All rights reserved.

//go:generate stringer -type=TokenType

// Package security is a toolkit for authorization.
package security

// TokenType is the type of token.
type TokenType int

const (
	AnonymousToken TokenType = iota
	BasicToken
	BearerToken
	DigestToken
	DPoPToken
	GNAPToken
	HOBAToken
	MutualToken
	NegotiateToken
	NTLMToken
	OAuthToken
	PrivateTokenToken
	AWS4HMACSHA256Token
	SCRAMToken
	SCRAMSHA1Token
	SCRAMSHA256Token
	VAPIDToken
	InvalidToken
	maxTokenType
)
