/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package authenticate

// authenticate represents an authorization.
type httpAuthenticate struct {
	// Type is the type of authorization.
	scheme AuthScheme
	// credentials are the credentials associated with the authorization.
	credentials string
	// extra contains other information about the authorization.
	extra []string
}

// Scheme returns the type of authorization.
func (obj httpAuthenticate) Scheme() AuthScheme {
	return obj.scheme
}

// Credentials returns the credentials associated with the authorization.
func (obj httpAuthenticate) Credentials() string {
	return obj.credentials
}

// Extra returns the extra information about the authorization.
func (obj httpAuthenticate) Extra() []string {
	return obj.extra
}

// Encode encodes the authenticate struct into a string.
//
// Returns the encoded string and an error if any.
func (obj httpAuthenticate) Encode(args ...any) (string, error) {
	return obj.scheme.String() + " " + obj.credentials, nil
}
