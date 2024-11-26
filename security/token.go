/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package security is a toolkit for security check and authorization
package security

import (
	"encoding/json"
	"time"
)

// Claims defines basic Token claims
type Claims struct {
	ID        string                 `json:"jti,omitempty"`
	Issuer    string                 `json:"iss,omitempty"`
	Subject   string                 `json:"sub,omitempty"`
	Audience  []string               `json:"aud,omitempty"`
	ExpiresAt int64                  `json:"exp,omitempty"`
	NotBefore int64                  `json:"nbf,omitempty"`
	IssuedAt  int64                  `json:"iat,omitempty"`
	Custom    map[string]interface{} `json:"custom,omitempty"`
}

// Token is an interface for getting token information
type Token struct {
	AccessToken  string    `json:"access_token"`
	TokenType    string    `json:"token_type,omitempty"`
	RefreshToken string    `json:"refresh_token,omitempty"`
	Expiry       time.Time `json:"expiry,omitempty"`
	ExpiresIn    int64     `json:"expires_in,omitempty"`
	Claims       *Claims   `json:"-"`
}

func (t Token) Encode(args ...any) (string, error) {
	v, _ := json.Marshal(t)
	return string(v), nil
}

var _ Encoder = (*Token)(nil)
