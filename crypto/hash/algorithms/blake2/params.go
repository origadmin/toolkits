package blake2

import (
	"encoding/base64"
	"fmt"

	hashcodec "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/validator"
)

const (
	MinKeyLength = 16
	MaxKeyLength = 64
)

type Params struct {
	Key []byte
}

// String returns the string representation of parameters
func (p *Params) String() string {
	return hashcodec.EncodeParams(p.ToMap())
}

func (p *Params) IsNil() bool {
	return p == nil
}

// Validate checks the parameters for correctness. It now only validates the key length.
func (p *Params) Validate(config *types.Config) error {
	if len(p.Key) > 0 && (len(p.Key) < MinKeyLength || len(p.Key) > MaxKeyLength) {
		return fmt.Errorf("invalid key length: %d, must be between %d and %d", len(p.Key), MinKeyLength, MaxKeyLength)
	}
	return nil
}

func (p *Params) FromMap(params map[string]string) error {
	if v, ok := params["k"]; ok {
		key, err := base64.RawURLEncoding.DecodeString(v)
		if err != nil {
			return fmt.Errorf("invalid base64 key: %w", err)
		}
		p.Key = key
	}
	return nil
}

// ToMap converts Params to a map[string]string
func (p *Params) ToMap() map[string]string {
	m := make(map[string]string)
	if len(p.Key) > 0 {
		m["k"] = base64.RawURLEncoding.EncodeToString(p.Key)
	}
	return m
}

// WithKey is an option to set the key for blake2 hashing.
// It operates on the new Config model by directly modifying the Params map.
func WithKey(key []byte) func(config *types.Config) {
	return func(config *types.Config) {
		if config.Params == nil {
			config.Params = make(map[string]string)
		}
		config.Params["k"] = base64.RawURLEncoding.EncodeToString(key)
	}
}

func DefaultParams() *Params {
	return &Params{}
}

// Ensure Params implements the validator.Parameters interface for now.
var _ validator.Parameters = (*Params)(nil)
