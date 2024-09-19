package authenticate

import (
	"strings"
)

// AuthScheme represents the type of authorization.
//
//go:generate go run -mod=mod golang.org/x/tools/cmd/stringer@latest -type=AuthScheme -trimprefix=AuthScheme
type AuthScheme int

const (
	// AuthSchemeAnonymous represents an anonymous authorization.
	AuthSchemeAnonymous AuthScheme = iota
	// AuthSchemeBasic represents a basic authorization.
	AuthSchemeBasic
	// AuthSchemeBearer represents a bearer authorization.
	AuthSchemeBearer
	// AuthSchemeDigest represents a digest authorization.
	AuthSchemeDigest
	// AuthSchemeHOBA represents a HTTP Origin-Bound Authentication (HOBA) authorization.
	AuthSchemeHOBA
	// AuthSchemeMutual represents a mutual authentication.
	AuthSchemeMutual
	// AuthSchemeSCRAMSHA1 represents a SCRAM-SHA-1 authorization.
	AuthSchemeSCRAMSHA1
	// AuthSchemeSCRAMSHA256 represents a SCRAM-SHA-256 authorization.
	AuthSchemeSCRAMSHA256
	// AuthSchemeAWS4HMAC256 represents an AWS4-HMAC-SHA256 authorization.
	AuthSchemeAWS4HMAC256
	// AuthSchemeNegotiate represents a negotiate authorization.
	AuthSchemeNegotiate
	// AuthSchemeNTLM represents an NTLM authorization.
	AuthSchemeNTLM
	// AuthSchemeOAuth represents an OAuth authorization.
	AuthSchemeOAuth
	// AuthSchemeVapid represents a VAPID authorization.
	AuthSchemeVapid
	// AuthSchemeUnknown represents an unknown authorization.
	AuthSchemeUnknown
)

// Lower returns the lowercase string representation of the Type.
func (t AuthScheme) Lower() string {
	return strings.ToLower(string(t))
}
