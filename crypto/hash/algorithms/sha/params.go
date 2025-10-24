package sha

import (
	hashcodec "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Params represents parameters for SHA algorithms
type Params struct {
}

func (p *Params) IsNil() bool {
	return p == nil
}

func (p *Params) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return errors.ErrSaltLengthTooShort
	}

	return nil
}

func (p *Params) FromMap(params map[string]string) error {
	return nil
}

// FromParams parses Scrypt parameters from a map[string]string.
func FromParams(m map[string]string) (params *Params, err error) {
	params = &Params{}

	err = params.FromMap(m)
	if err != nil {
		return nil, err
	}

	return params, nil
}

// String returns the string representation of parameters
func (p *Params) String() string {
	return hashcodec.EncodeParams(p.ToMap())
}

// ToMap converts Params to a map[string]string
func (p *Params) ToMap() map[string]string {
	m := make(map[string]string)

	return m
}
func DefaultParams() *Params {
	return &Params{}
}
