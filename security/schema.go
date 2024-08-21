// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package security is a toolkit for security check and authorization
package security

import (
	"encoding/base64"
	"fmt"
	"strings"
)

type authTypeErr struct {
	Type TokenType
}

func (a *authTypeErr) Error() string {
	return fmt.Sprintf("invalid s type: %d", a.Type)
}

func ErrAuthType(t TokenType) error {
	return &authTypeErr{t}
}

type Scheme struct {
	Type        TokenType
	Credentials string
}

func (ts *Scheme) String() string {
	return ts.Type.String() + " " + ts.Credentials
}

// ParseScheme parses a token string into a Scheme.
func ParseScheme(token string) Scheme {
	if len(token) == 0 {
		return Scheme{Type: InvalidToken, Credentials: token}
	}

	switch {
	case strings.HasPrefix(token, BearerToken.String()):
		tokens := strings.Split(token, " ")
		if len(tokens) == 2 {
			return Scheme{Type: BearerToken, Credentials: tokens[1]}
		}
	case strings.HasPrefix(token, BasicToken.String()):
		tokens := strings.Split(token, " ")
		if len(tokens) == 2 {
			return Scheme{Type: BasicToken, Credentials: tokens[1]}
		}
	}
	return Scheme{Type: AnonymousToken, Credentials: token}
}

// RFC7617UserPass base64 encodes a authUser-id and password per:
// https://tools.ietf.org/html/rfc7617#section-2
func RFC7617UserPass(userid, password string) (string, error) {
	if strings.Contains(userid, ":") && strings.Contains(password, ":") {
		return "", fmt.Errorf("RFC7617 authUser-id cannot include a colon (':') [%v]", userid)
	}

	return base64.StdEncoding.EncodeToString(
		[]byte(userid + ":" + password),
	), nil
}

// BasicAuth creates a basic s string.
func BasicAuth(userid, password string) (string, error) {
	apiKey, err := RFC7617UserPass(userid, password)
	if err != nil {
		return "", err
	}
	return BasicToken.String() + " " + apiKey, nil
}
