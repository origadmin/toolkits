package security

type TokenSerializer interface {
	Generate(subject string) Token
	Parse(token string) (Token, error)
}
