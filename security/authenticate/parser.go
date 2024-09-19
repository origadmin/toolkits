package authenticate

import (
	"net/http"
)

type AuthDecoder interface {
	Decode(req *http.Request) (string, bool)
}

type RequestParser interface {
	Parse(req *http.Request) (*Authenticate, bool)
}

// RequestParser is used to parse authentication information from different parts of an HTTP request.
type requestParser struct {
	// Fields specifies which fields to parse for authentication information.
	Decodes []AuthDecoder
}

// NewRequestParser creates a new RequestParser with specified fields.
func NewRequestParser(decodes ...AuthDecoder) RequestParser {
	return &requestParser{Decodes: decodes}
}

// Parse parses the authentication information from the provided HTTP request.
func (rp *requestParser) Parse(req *http.Request) (*Authenticate, bool) {
	for i := range rp.Decodes {
		if auth, ok := rp.Decodes[i].Decode(req); ok {
			return ParseAuth(auth)
		}
	}
	return nil, false
}

// AuthorizationDecoder decodes the "Authorization" header.
type AuthorizationDecoder struct {
	Key string
}

// Decode implements the AuthDecoder interface.
func (ad *AuthorizationDecoder) Decode(req *http.Request) (string, bool) {
	auth := req.Header.Get(ad.Key)
	if auth != "" {
		return auth, true
	}
	return "", false
}

// QueryDecoder decodes the query parameters.
type QueryDecoder struct {
	Key string
}

// Decode implements the AuthDecoder interface.
func (qd *QueryDecoder) Decode(req *http.Request) (string, bool) {
	query := req.URL.Query()
	auth := query.Get(qd.Key)
	if auth != "" {
		return auth, true
	}
	return "", false
}

// CookieDecoder decodes the cookies.
type CookieDecoder struct {
	Key string
}

// Decode implements the AuthDecoder interface.
func (cd *CookieDecoder) Decode(req *http.Request) (string, bool) {
	cookies, err := req.Cookie(cd.Key)
	if err == nil && cookies != nil {
		return cookies.Value, true
	}
	return "", false
}
