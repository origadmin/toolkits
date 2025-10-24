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

func (p *Params) IsNil() bool {
	return p == nil
}

func (p *Params) Validate(config *types.Config) error {

	//if config.SaltLength < 8 {
	//	return fmt.Errorf("salt length must be at least 8 bytes")
	//}
	if len(p.Key) < MinKeyLength || len(p.Key) > MaxKeyLength {
		return fmt.Errorf("invalid key length: %d", len(p.Key))
	}
	return nil
}

func (p *Params) FromMap(params map[string]string) error {
	if v, ok := params["k"]; ok {
		key, err := base64.RawURLEncoding.DecodeString(v)
		if err != nil {
			return fmt.Errorf("invalid key: %w", err)
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

func (p *Params) String() string {
	return hashcodec.EncodeParams(p.ToMap())
}

func FromMap(m map[string]string) (params *Params, err error) {
	params = &Params{}
	if v, ok := m["k"]; ok {
		key, err := base64.RawURLEncoding.DecodeString(v)
		if err != nil {
			return nil, fmt.Errorf("invalid key: %w", err)
		}
		params.Key = key
	}
	return params, nil
}

func WithKey(key []byte) types.Option {
	p := &Params{
		Key: key,
	}
	return func(config *types.Config) {
		if config.ParamConfig == "" {
			config.ParamConfig = p.String()
		} else {
			config.ParamConfig = hashcodec.MergeParams(config.ParamConfig, p.ToMap())
		}
	}
}

func DefaultParams() *Params {
	return &Params{}
}

var _ validator.Parameters = (*Params)(nil)
