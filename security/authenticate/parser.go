package authenticate

import (
	"context"
	"net/http"
	"strings"

	"github.com/origadmin/toolkits/storage/cache"
)

type AuthDecoder interface {
	Decode(req *http.Request) (string, bool)
}

type RequestParser interface {
	Parse(req *http.Request) (Authenticator, bool)
}

type ResponseWriter interface {
	Write(http.ResponseWriter, Authenticator) error
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
func (rp *requestParser) Parse(req *http.Request) (Authenticator, bool) {
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

// SessionDecoder decodes the session information.
type SessionDecoder struct {
	SessionIDKey   string
	SessionStorage cache.Cache
}

// Decode implements the AuthDecoder interface.
func (sd *SessionDecoder) Decode(req *http.Request) (string, bool) {
	// Gets the session ID from the request
	cookie, err := req.Cookie(sd.SessionIDKey)
	if err != nil || cookie == nil {
		return "", false
	}

	return sd.getSession(cookie.Value)
}

func (sd *SessionDecoder) getSession(id string) (string, bool) {
	if sd.SessionStorage == nil {
		return "", false
	}
	// Obtain authentication information based on the session ID
	session, err := sd.SessionStorage.Get(context.Background(), id)
	if err != nil {
		return "", false
	}
	return session, true
}

const defaultMaxMemory = int64(32 << 20)

type FormDecoder struct {
	Key       string
	MaxMemory int64
}

// Decode implements the AuthDecoder interface.
func (fd *FormDecoder) Decode(req *http.Request) (string, bool) {
	if req.PostForm != nil {
		maxMemory := defaultMaxMemory
		if fd.MaxMemory > 0 {
			maxMemory = fd.MaxMemory
		}
		err := req.ParseMultipartForm(maxMemory)
		if err != nil {
			return "", false
		}
	}
	value := req.PostForm.Get(fd.Key)
	if value == "" {
		return "", false
	}

	return value, true
}

type ParamDecoder struct {
	Key string
}

// Decode implements the AuthDecoder interface.
func (pd *ParamDecoder) Decode(req *http.Request) (string, bool) {
	value := req.PathValue(pd.Key)
	if value == "" {
		return "", false
	}

	return value, true
}

// ParseRequest parses the Authorization header from the provided HTTP request.
// If the header is empty, it returns nil and false.
func ParseRequest(req *http.Request) (Authenticator, bool) {
	return NewRequestParser(&AuthorizationDecoder{
		Key: "Authorization",
	}).Parse(req)
}

// WriteResponse writes the WWW-Authenticate header to the provided HTTP response.
func WriteResponse(resp http.ResponseWriter, auth Authenticator) {
	// TODO: implement WWW-Authenticate header
	resp.Header().Set("WWW-Authenticate", "")
}

// ParseAuth parses the authorization and returns an authenticate struct with the type and credentials extracted.
// If the header is empty, it returns TypeUnknown with the provided auth string and false.
// If the header contains only one token, it returns TypeAnonymous with the token as credentials and true.
// For headers with multiple tokens, it checks the type (basic, bearer, digest) and returns the corresponding to authenticate struct with the credentials and true.
// If the type is not recognized, it returns TypeUnknown with the original auth string and false.
func ParseAuth(auth string) (Authenticator, bool) {
	if len(auth) == 0 {
		return &httpAuthenticate{scheme: AuthSchemeUnknown, credentials: auth}, false
	}

	tokens := strings.Split(auth, " ")
	switch len(tokens) {
	case 0:
		return &httpAuthenticate{scheme: AuthSchemeUnknown, credentials: auth}, false
	case 1:
		return &httpAuthenticate{scheme: AuthSchemeAnonymous, credentials: tokens[0]}, true
	default:
	}
	for scheme := AuthSchemeBasic; scheme < authSchemeMax; scheme++ {
		if scheme.Equal(tokens[0]) {
			return &httpAuthenticate{scheme: scheme, credentials: tokens[1], extra: tokens}, true
		}
	}
	return &httpAuthenticate{scheme: AuthSchemeUnknown, credentials: tokens[1], extra: tokens}, true
}
