package authorize

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"
)

// Type represents the type of authorization.
type Type string

const (
	// TypeAnonymous represents an anonymous authorization.
	TypeAnonymous Type = "anonymous"
	// TypeBasic represents a basic authorization.
	TypeBasic Type = "basic"
	// TypeBearer represents a bearer authorization.
	TypeBearer Type = "bearer"
	// TypeDigest represents a digest authorization.
	TypeDigest Type = "digest"
	// TypeUnknown represents an unknown authorization.
	TypeUnknown Type = "unknown"
)

// String returns the string representation of the Type.
func (t Type) String() string {
	return string(t)
}

// Lower returns the lowercase string representation of the Type.
func (t Type) Lower() string {
	return strings.ToLower(string(t))
}

// Authorize represents an authorization.
type Authorize struct {
	// Type is the type of authorization.
	Type Type
	// Credentials are the credentials associated with the authorization.
	Credentials string
}

// Encode encodes the Authorize struct into a string.
//
// Returns the encoded string and an error if any.
func (a Authorize) Encode() (string, error) {
	// TODO: Implement encoding logic.
	return "", nil
}

// ParseHTTPRequest parses the Authorization header from the provided HTTP request.
// If the header is empty, it returns nil and false.
func ParseHTTPRequest(req *http.Request) (*Authorize, bool) {
	auth := req.Header.Get("Authorization")
	if auth == "" {
		return nil, false
	}
	return ParseHeader(auth)
}

// ParseWWWRequest parses the WWW-Authenticate header from the provided HTTP request.
// If the header is empty, it returns nil and false.
// TODO: Implement the parsing logic based on https://tools.ietf.org/html/rfc7235#section-4.1
func ParseWWWRequest(req *http.Request) (*Authorize, bool) {
	auth := req.Header.Get("WWW-Authenticate")
	if auth == "" {
		return nil, false
	}
	// TODO: parse WWW-Authenticate
	// https://tools.ietf.org/html/rfc7235#section-4.1
	return nil, false
}

// SetRequestHeader sets the Authorization header of the provided http.
// Request using the encoded credentials from the given Authorize struct.
// It returns an error if there is an issue encoding the credentials.
func SetRequestHeader(req *http.Request, auth *Authorize) error {
	encode, err := auth.Encode()
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", encode)
	return nil
}

// ParseHeader parses the authorization header and returns an Authorize struct with the type and credentials extracted.
// If the header is empty, it returns TypeUnknown with the provided auth string and false.
// If the header contains only one token, it returns TypeAnonymous with the token as credentials and true.
// For headers with multiple tokens, it checks the type (basic, bearer, digest) and returns the corresponding to Authorize struct with the credentials and true.
// If the type is not recognized, it returns TypeUnknown with the original auth string and false.
func ParseHeader(auth string) (*Authorize, bool) {
	if len(auth) == 0 {
		return &Authorize{Type: TypeUnknown, Credentials: auth}, false
	}
	tokens := strings.Split(auth, " ")
	if len(tokens) == 1 {
		return &Authorize{Type: TypeAnonymous, Credentials: tokens[0]}, true
	}
	switch strings.ToLower(tokens[0]) {
	case TypeBasic.Lower():
		return &Authorize{Type: TypeBasic, Credentials: tokens[1]}, true
	case TypeBearer.Lower():
		return &Authorize{Type: TypeBearer, Credentials: tokens[1]}, true
	case TypeDigest.Lower():
		return &Authorize{Type: TypeDigest, Credentials: tokens[1]}, true
	default:
		return &Authorize{Type: TypeUnknown, Credentials: auth}, false
	}
}

// BasicAuthUserPass base64 encodes userid and password and returns the encoded string.:
// https://datatracker.ietf.org/doc/html/rfc2617#section-2
// user-pass   = userid ":" password
// userid      = *<TEXT excluding ":">
// password    = *TEXT
func BasicAuthUserPass(userid, password string) string {
	if strings.Contains(userid, ":") {
		panic(fmt.Errorf("RFC7617 user-id cannot include a colon (':') [%v]", userid))
	}
	userpass := userid + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(userpass))
}

// BasicAuthHeader creates a basic auth header.
func BasicAuthHeader(userid, password string) string {
	userpass := BasicAuthUserPass(userid, password)
	return TypeBasic.String() + " " + userpass
}

// ParseBasicAuthHeader parses the basic authentication header and returns the username and password.
// It takes the authentication header string as input and returns the username, password, and a boolean indicating success.
// If the header is not in the correct format or decoding fails, it returns empty strings and false.
func ParseBasicAuthHeader(header string) (string, string, bool) {
	prefix := TypeBasic.Lower() + " "
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
