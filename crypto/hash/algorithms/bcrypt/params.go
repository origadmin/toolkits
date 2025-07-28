package bcrypt

import (
	"fmt"
	"strconv"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/constants"
)

type Params struct {
	Cost int
}

func (p Params) String() string {
	paramsMap := make(map[string]string)
	paramsMap["c"] = fmt.Sprintf("%d", p.Cost)
	return codecPkg.EncodeParams(paramsMap)
}

func parseParams(paramsMap map[string]string) (result Params, err error) {
	if v, ok := paramsMap["c"]; ok {
		cost, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return result, fmt.Errorf("invalid cost: %v", err)
		}
		result.Cost = int(cost)
	}
	return result, nil
}

func DefaultParams() Params {
	return Params{
		Cost: constants.DefaultCost,
	}
}

// ToMap converts Params to a map[string]string
func (p Params) ToMap() map[string]string {
	paramsMap := make(map[string]string)
	paramsMap["c"] = fmt.Sprintf("%d", p.Cost)
	return paramsMap
}
