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
	// AuthSchemeNegotiate represents a negotiate authorization.
	AuthSchemeNegotiate
	// AuthSchemeVapid represents a VAPID authorization.
	AuthSchemeVapid
	// AuthSchemeSCRAM represents a SCRAM authorization.
	AuthSchemeSCRAM
	// AuthSchemeAWS4HMAC256 represents an AWS4-HMAC-SHA256 authorization.
	AuthSchemeAWS4HMAC256
	// AuthSchemeDPoP represents a DPoP authorization.
	AuthSchemeDPoP
	// AuthSchemeGNAP represents a GNAP authorization.
	AuthSchemeGNAP
	// AuthSchemePrivate represents a private authorization.
	AuthSchemePrivate
	// AuthSchemeOAuth represents an OAuth authorization.
	AuthSchemeOAuth
	// AuthSchemeUnknown represents an unknown authorization.
	AuthSchemeUnknown
	authSchemeMax
)

const (
	// AuthSchemeNTLM represents an NTLM authorization.
	AuthSchemeNTLM = AuthSchemeNegotiate
)

// Lower returns the lowercase string representation of the Type.
func (t AuthScheme) Lower() string {
	return strings.ToLower(t.String())
}

func (t AuthScheme) Equal(other string) bool {
	schemeType := t.String() + " "
	if len(other) < len(schemeType) {
		return false
	}

	return strings.EqualFold(other[:len(schemeType)], schemeType)
}
