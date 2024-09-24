package authenticate

type Authenticator interface {
	Scheme() AuthScheme
	Credentials() string
	Extra() []string
	Encode(args ...any) (string, error)
}
