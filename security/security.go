package security

import (
	"context"
	"fmt"
	"time"

	"github.com/goexts/generic/settings"

	"github.com/origadmin/toolkits/errors"
)

const (
	ErrInvalidToken = errors.String("invalid token")
	ErrNotFound     = errors.String("not found")
)

type Security interface {
	// GenerateToken Generate a JWT (JSON Web HeaderToken) with the provided subject.
	GenerateToken(ctx context.Context, subject string) (Token, error)
	// ValidateAccess Validate if a token is valid.
	ValidateAccess(ctx context.Context, tokenStr string) error
	// DestroyToken Invalidate a token by removing it from the token store.
	DestroyToken(ctx context.Context, tokenStr string) error
	// ParseSubject Parse the subject (or user identifier) from a given access token.
	ParseSubject(ctx context.Context, tokenStr string) (string, error)
	// Release any resources held by the authorize instance.
	Release(ctx context.Context) error
}

type security struct {
	serialize TokenSerializer
	storage   TokenStorage
}

func (s security) ValidateAccess(ctx context.Context, accessToken string) error {
	fmt.Println("token: ", accessToken)
	return s.storage.Validate(ctx, accessToken)
}

func (s security) GenerateToken(ctx context.Context, subject string) (Token, error) {
	token := s.serialize.Generate(subject)
	if token.AccessToken == "" {
		return Token{}, errors.New("generate token failed")
	}
	fmt.Println("token: ", token.AccessToken, "duration: ", token.ExpiresIn)
	if err := s.storage.Set(ctx, token.AccessToken, time.Duration(token.ExpiresIn)); err != nil {
		return Token{}, err
	}
	return token, nil
}

func (s security) DestroyToken(ctx context.Context, tokenStr string) error {
	return s.storage.Delete(ctx, tokenStr)
}

func (s security) ParseSubject(ctx context.Context, tokenStr string) (string, error) {
	if tokenStr == "" {
		return "", ErrInvalidToken
	}

	token, err := s.serialize.Parse(tokenStr)
	if err != nil {
		return "", err
	}

	err = s.storage.Validate(ctx, tokenStr)
	fmt.Println("token: ", tokenStr, "error: ", err)
	switch {
	case err == nil:
		if token.Claims == nil {
			return "", nil
		}
		return token.Claims.Subject, nil
	case !errors.Is(err, ErrNotFound):
		return "", err
	default:
		return "", ErrInvalidToken
	}
}

func (s security) Release(ctx context.Context) error {
	return s.storage.Close(ctx)
}

func (s security) ParseToken(ctx context.Context, tokenStr string) (Token, error) {
	return s.serialize.Parse(tokenStr)
}

type Setting = func(s *security)

func WithStorage(store TokenStorage) Setting {
	return func(s *security) {
		s.storage = store
	}
}

func NewSecurity(serializer TokenSerializer, ss ...Setting) Security {
	s := settings.Apply(&security{
		serialize: serializer,
	}, ss)

	if s.storage == nil {
		s.storage = NewTokenStorage()
	}

	return s
}
