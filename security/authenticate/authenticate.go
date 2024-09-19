package authenticate

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// Authenticate represents an authorization.
type Authenticate struct {
	// Type is the type of authorization.
	Scheme AuthScheme
	// Credentials are the credentials associated with the authorization.
	Credentials string
}

// Encode encodes the Authenticate struct into a string.
//
// Returns the encoded string and an error if any.
func (obj Authenticate) Encode() string {
	return obj.Scheme.String() + " " + obj.Credentials
}

// ParseRequest parses the Authorization header from the provided HTTP request.
// If the header is empty, it returns nil and false.
func ParseRequest(req *http.Request) (*Authenticate, bool) {
	return NewRequestParser(&AuthorizationDecoder{
		Key: "Authorization",
	}).Parse(req)
}

// ParseWWWRequest parses the WWW-Authenticate header from the provided HTTP request.
// If the header is empty, it returns nil and false.
func ParseWWWRequest(req *http.Request) (*Authenticate, bool) {
	return NewRequestParser(&AuthorizationDecoder{
		Key: "WWW-Authenticate",
	}).Parse(req)
}

// ParseAuth parses the authorization and returns an Authenticate struct with the type and credentials extracted.
// If the header is empty, it returns TypeUnknown with the provided auth string and false.
// If the header contains only one token, it returns TypeAnonymous with the token as credentials and true.
// For headers with multiple tokens, it checks the type (basic, bearer, digest) and returns the corresponding to Authenticate struct with the credentials and true.
// If the type is not recognized, it returns TypeUnknown with the original auth string and false.
func ParseAuth(auth string) (*Authenticate, bool) {
	if len(auth) == 0 {
		return &Authenticate{Scheme: AuthSchemeUnknown, Credentials: auth}, false
	}

	tokens := strings.Split(auth, " ")
	if len(tokens) == 1 {
		return &Authenticate{Scheme: AuthSchemeAnonymous, Credentials: tokens[0]}, true
	}

	for scheme := AuthScheme(0); scheme < AuthSchemeUnknown; scheme++ {
		if strings.EqualFold(tokens[0], scheme.String()) {
			return &Authenticate{Scheme: scheme, Credentials: tokens[1]}, true
		}
	}
	return &Authenticate{Scheme: AuthSchemeUnknown, Credentials: auth}, false
}

// BasicAuthUserPass base64 encodes userid and password and returns the encoded string.:
// https://datatracker.ietf.org/doc/html/rfc2617#section-2
// user-pass   = userid ":" password
// userid      = *<TEXT excluding ":">
// password    = *TEXT
func BasicAuthUserPass(userid, password string) (string, error) {
	if strings.Contains(userid, ":") {
		return "", fmt.Errorf("RFC7617 user-id cannot include a colon (':') [%v]", userid)
	}
	userpass := userid + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(userpass)), nil
}

// BasicAuthHeader creates a basic auth header.
func BasicAuthHeader(userid, password string) (string, error) {
	userpass, err := BasicAuthUserPass(userid, password)
	if err != nil {
		return "", err
	}
	return AuthSchemeBasic.String() + " " + userpass, nil
}

// ParseBasicAuth parses the basic authentication and returns the username and password.
// It takes the authentication header string as input and returns the username, password, and a boolean indicating success.
// If the header is not in the correct format or decoding fails, it returns empty strings and false.
func ParseBasicAuth(header string) (string, string, bool) {
	prefix := AuthSchemeBasic.Lower() + " "
	if len(header) < len(prefix) || strings.EqualFold(header[:len(prefix)], prefix) {
		return "", "", false
	}
	c, err := base64.StdEncoding.DecodeString(header[len(prefix):])
	if err != nil {
		return "", "", false
	}
	cs := string(c)
	username, password, ok := strings.Cut(cs, ":")
	if !ok {
		return "", "", false
	}
	return username, password, true
}
