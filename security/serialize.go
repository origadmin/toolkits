package security

import (
	"time"
)

type TokenSerializer interface {
	Generate(subject string, expires ...time.Duration) Token
	Parse(token string) (Token, error)
}
