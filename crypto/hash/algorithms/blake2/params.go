package blake2

import (
	"encoding/base64"
	"fmt"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
)

const (
	MinKeyLength = 16
	MaxKeyLength = 64
)

type Params struct {
	Key []byte
}

func (p Params) String() string {
	paramsMap := make(map[string]string)
	if len(p.Key) > 0 {
		paramsMap["k"] = base64.RawURLEncoding.EncodeToString(p.Key)
	}
	return codecPkg.EncodeParams(paramsMap)
}

func parseParams(paramsMap map[string]string) (result Params, err error) {
	if v, ok := paramsMap["k"]; ok {
		key, err := base64.RawURLEncoding.DecodeString(v)
		if err != nil {
			return result, fmt.Errorf("invalid key: %w", err)
		}
		result.Key = key
	}
	return result, nil
}

func DefaultParams() Params {
	return Params{}
}

// ToMap converts Params to a map[string]string
func (p Params) ToMap() map[string]string {
	paramsMap := make(map[string]string)
	if len(p.Key) > 0 {
		paramsMap["k"] = base64.RawURLEncoding.EncodeToString(p.Key)
	}
	return paramsMap
}
