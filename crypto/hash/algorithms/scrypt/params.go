package scrypt

import (
	"fmt"
	"strconv"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
)

// Params represents parameters for Scrypt algorithm
type Params struct {
	N      int
	R      int
	P      int
	KeyLen int
}

// parseParams parses Scrypt parameters from a map[string]string.
func parseParams(paramsMap map[string]string) (*Params, error) {
	result := &Params{}

	// Parse N
	if v, ok := paramsMap["n"]; ok {
		n, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid N: %v", err)
		}
		result.N = n
	}

	// Parse R
	if v, ok := paramsMap["r"]; ok {
		r, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid R: %v", err)
		}
		result.R = r
	}

	// Parse P
	if v, ok := paramsMap["p"]; ok {
		p, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid P: %v", err)
		}
		result.P = p
	}

	// Parse KeyLen
	if v, ok := paramsMap["k"]; ok {
		keyLen, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid KeyLen: %v", err)
		}
		result.KeyLen = keyLen
	}

	return result, nil
}

// String returns the string representation of parameters
func (p *Params) String() string {
	paramsMap := p.ToMap()
	return codecPkg.EncodeParams(paramsMap)
}

func DefaultParams() *Params {
	return &Params{
		N:      16384,
		R:      8,
		P:      1,
		KeyLen: 32,
	}
}

// ToMap converts Params to a map[string]string
func (p *Params) ToMap() map[string]string {
	paramsMap := make(map[string]string)
	if p.N > 0 {
		paramsMap["n"] = fmt.Sprintf("%d", p.N)
	}
	if p.R > 0 {
		paramsMap["r"] = fmt.Sprintf("%d", p.R)
	}
	if p.P > 0 {
		paramsMap["p"] = fmt.Sprintf("%d", p.P)
	}
	if p.KeyLen > 0 {
		paramsMap["k"] = fmt.Sprintf("%d", p.KeyLen)
	}
	return paramsMap
}
