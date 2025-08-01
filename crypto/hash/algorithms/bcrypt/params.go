package bcrypt

import (
	"fmt"
	"strconv"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

type Params struct {
	Cost int
}

func (p *Params) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return fmt.Errorf("salt length must be at least 8 bytes")
	}
	if p.Cost < 4 {
		return fmt.Errorf("cost must be at least 4")
	}
	if p.Cost > 31 {
		return fmt.Errorf("cost must be at most 31")
	}
	return nil
}

func (p *Params) FromMap(params map[string]string) error {
	if v, ok := params["c"]; ok {
		cost, err := strconv.ParseInt(v, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid cost: %v", err)
		}
		p.Cost = int(cost)
	}
	return nil
}

// ToMap converts Params to a map[string]string
func (p *Params) ToMap() map[string]string {
	m := make(map[string]string)
	m["c"] = fmt.Sprintf("%d", p.Cost)
	return m
}

// String returns the string representation of parameters
func (p *Params) String() string {
	return codecPkg.EncodeParams(p.ToMap())
}

// FromMap parses Argon2 parameters from string
func FromMap(m map[string]string) (params *Params, err error) {
	params = &Params{}
	if err = params.FromMap(m); err != nil {
		return nil, err
	}
	return params, nil
}

func DefaultParams() interfaces.Params {
	return &Params{
		Cost: constants.DefaultCost,
	}
}
