package scrypt

import (
	"fmt"
	"strconv"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Params represents parameters for Scrypt algorithm
type Params struct {
	N      int
	R      int
	P      int
	KeyLen int
}

func (p *Params) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return errors.ErrSaltLengthTooShort
	}
	// N must be > 1 and a power of 2
	if p.N <= 1 || p.N&(p.N-1) != 0 {
		return fmt.Errorf("N must be > 1 and a power of 2")
	}
	if p.R <= 0 {
		return fmt.Errorf("r must be > 0")
	}
	if p.P <= 0 {
		return fmt.Errorf("p must be > 0")
	}
	if p.KeyLen <= 0 {
		return fmt.Errorf("key length must be > 0")
	}
	return nil
}

func (p *Params) FromMap(params map[string]string) error {
	// Parse N
	if v, ok := params["N"]; ok {
		n, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("invalid N: %v", err)
		}
		p.N = n
	}

	// Parse R
	if v, ok := params["r"]; ok {
		r, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("invalid r: %v", err)
		}
		p.R = r
	}

	// Parse P
	if v, ok := params["p"]; ok {
		pp, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("invalid p: %v", err)
		}
		p.P = pp
	}

	// Parse KeyLen
	if v, ok := params["k"]; ok {
		keyLen, err := strconv.Atoi(v)
		if err != nil {
			return fmt.Errorf("invalid KeyLen: %v", err)
		}
		p.KeyLen = keyLen
	}

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
	return codecPkg.EncodeParams(p.ToMap())
}

// ToMap converts Params to a map[string]string
func (p *Params) ToMap() map[string]string {
	m := make(map[string]string)
	if p.N > 0 {
		m["N"] = fmt.Sprintf("%d", p.N)
	}
	if p.R > 0 {
		m["r"] = fmt.Sprintf("%d", p.R)
	}
	if p.P > 0 {
		m["p"] = fmt.Sprintf("%d", p.P)
	}
	if p.KeyLen > 0 {
		m["k"] = fmt.Sprintf("%d", p.KeyLen)
	}
	return m
}
func DefaultParams() *Params {
	return &Params{
		N:      16384,
		R:      8,
		P:      1,
		KeyLen: 32,
	}
}
