package pbkdf2

import (
	"fmt"
	"strconv"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
)

// Params represents parameters for PBKDF2 algorithm
type Params struct {
	Iterations int
	KeyLength  uint32
	HashType   string
}

// String returns the string representation of parameters
func (p *Params) String() string {
	paramsMap := p.ToMap()
	return codecPkg.EncodeParams(paramsMap)
}

// ToMap converts Params to a map[string]string
func (p *Params) ToMap() map[string]string {
	paramsMap := make(map[string]string)
	if p.Iterations > 0 {
		paramsMap["i"] = fmt.Sprintf("%d", p.Iterations)
	}
	if p.KeyLength > 0 {
		paramsMap["k"] = fmt.Sprintf("%d", p.KeyLength)
	}
	if p.HashType != "" {
		paramsMap["h"] = p.HashType
	}
	return paramsMap
}

// parseParams parses PBKDF2 parameters from a map[string]string.
func parseParams(paramsMap map[string]string) (*Params, error) {
	result := &Params{}

	// Parse iterations
	if v, ok := paramsMap["i"]; ok {
		iterations, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid iterations: %v", err)
		}
		result.Iterations = iterations
	}

	// Parse key length
	if v, ok := paramsMap["k"]; ok {
		keyLength, err := strconv.Atoi(v)
		if err != nil {
			return nil, fmt.Errorf("invalid key length: %v", err)
		}
		result.KeyLength = uint32(keyLength)
	}

	// Parse hash type
	if v, ok := paramsMap["h"]; ok {
		_, err := stdhash.ParseHash(v)
		if err == nil {
			result.HashType = v
		}
	}

	return result, nil
}

func DefaultParams() *Params {
	return &Params{
		Iterations: 10000,
		KeyLength:  32,
		HashType:   stdhash.SHA256.String(),
	}
}
