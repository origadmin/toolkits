package pbkdf2

import (
	"fmt"
	"strconv"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Params represents parameters for PBKDF2 algorithm
type Params struct {
	Iterations int
	KeyLength  uint32
}

func (p *Params) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return errors.ErrSaltLengthTooShort
	}
	if p.Iterations < 1000 {
		return errors.ErrCostOutOfRange
	}
	if p.KeyLength < 8 {
		return errors.ErrKeyLengthTooShort
	}
	return nil
}

func (p *Params) FromMap(params map[string]string) error {
	// Parse iterations
	if v, ok := params["i"]; ok {
		iterations, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("invalid iterations: %v", err)
		}
		p.Iterations = iterations
	}

	// Parse key length
	if v, ok := params["k"]; ok {
		keyLength, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("invalid key length: %v", err)
		}
		p.KeyLength = uint32(keyLength)
	}

	return nil
}

// String returns the string representation of parameters
func (p *Params) String() string {
	return codecPkg.EncodeParams(p.ToMap())
}

// ToMap converts Params to a map[string]string
func (p *Params) ToMap() map[string]string {
	m := make(map[string]string)
	if p.Iterations > 0 {
		m["i"] = fmt.Sprintf("%d", p.Iterations)
	}
	if p.KeyLength > 0 {
		m["k"] = fmt.Sprintf("%d", p.KeyLength)
	}
	return m
}

// FromMap parses PBKDF2 parameters from a map[string]string.
func FromMap(m map[string]string) (params *Params, err error) {
	params = &Params{}

	// Parse iterations
	if v, ok := m["i"]; ok {
		iterations, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid iterations: %v", err)
		}
		params.Iterations = iterations
	}

	// Parse key length
	if v, ok := m["k"]; ok {
		keyLength, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid key length: %v", err)
		}
		params.KeyLength = uint32(keyLength)
	}

	return params, nil
}

func DefaultParams() *Params {
	return &Params{
		Iterations: 10000,
		KeyLength:  32,
	}
}
