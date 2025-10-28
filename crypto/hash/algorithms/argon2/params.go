package argon2

import (
	"fmt"
	"strconv"

	hashcodec "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/validator"
)

// Params represents parameters for Argon2 algorithm
type Params struct {
	TimeCost   uint32
	MemoryCost uint32
	Threads    uint8
	KeyLength  uint32
}

func (p *Params) IsNil() bool {
	return p == nil
}

func (p *Params) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return fmt.Errorf("invalid salt length: %d, must be at least 8", config.SaltLength)
	}

	if p.TimeCost < 3 {
		return fmt.Errorf("invalid time cost: %d", p.TimeCost)
	}
	if p.MemoryCost < 64*1024 {
		return fmt.Errorf("invalid memory cost: %d", p.MemoryCost)
	}
	if p.Threads < 1 {
		return fmt.Errorf("invalid threads: %d", p.Threads)
	}
	if p.KeyLength < 4 || p.KeyLength > 1024 {
		return fmt.Errorf("invalid key length: %d, must be between 4 and 1024", p.KeyLength)
	}
	return nil
}

func (p *Params) FromMap(params map[string]string) error {
	if params == nil {
		return fmt.Errorf("params is nil")
	}
	p.TimeCost = 0
	if v, ok := params["t"]; ok {
		timeCost, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid time cost: %v", err)
		}
		p.TimeCost = uint32(timeCost)
	}

	p.MemoryCost = 0
	if v, ok := params["m"]; ok {
		memoryCost, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid memory cost: %v", err)
		}
		p.MemoryCost = uint32(memoryCost)
	}

	p.Threads = 0
	if v, ok := params["p"]; ok {
		threads, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return fmt.Errorf("invalid threads: %v", err)
		}
		p.Threads = uint8(threads)
	}

	p.KeyLength = 0
	if v, ok := params["k"]; ok {
		keyLength, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid key length: %v", err)
		}
		p.KeyLength = uint32(keyLength)
	}

	return nil
}

// String returns the string representation of parameters
func (p *Params) String() string {
	return hashcodec.EncodeParams(p.ToMap())
}

// ToMap converts Params to a map[string]string
func (p *Params) ToMap() map[string]string {
	params := make(map[string]string)
	if p.TimeCost > 0 {
		params["t"] = fmt.Sprintf("%d", p.TimeCost)
	}
	if p.MemoryCost > 0 {
		params["m"] = fmt.Sprintf("%d", p.MemoryCost)
	}
	if p.Threads > 0 {
		params["p"] = fmt.Sprintf("%d", p.Threads)
	}
	if p.KeyLength > 0 {
		params["k"] = fmt.Sprintf("%d", p.KeyLength)
	}
	return params
}

// FromMap parses Argon2 parameters from string
func FromMap(m map[string]string) (params *Params, err error) {
	params = &Params{}
	if err = params.FromMap(m); err != nil {
		return nil, err
	}
	return params, nil
}

func WithTimeCost(timeCost uint32) func(*types.Config) {
	return func(c *types.Config) {
		c.Params["t"] = fmt.Sprintf("%d", timeCost)
	}
}

func WithMemoryCost(memoryCost uint32) func(*types.Config) {
	return func(c *types.Config) {
		c.Params["m"] = fmt.Sprintf("%d", memoryCost)
	}
}

func WithThreads(threads uint8) func(*types.Config) {
	return func(c *types.Config) {
		c.Params["p"] = fmt.Sprintf("%d", threads)
	}
}

func WithKeyLength(keyLength uint32) func(*types.Config) {
	return func(c *types.Config) {
		c.Params["k"] = fmt.Sprintf("%d", keyLength)
	}
}

func WithParams(params *Params) func(*types.Config) {
	return func(c *types.Config) {
		c.Params = params.ToMap()
	}
}

func DefaultParams() *Params {
	return &Params{
		TimeCost:   types.DefaultTimeCost,
		MemoryCost: types.DefaultMemoryCost,
		Threads:    types.DefaultThreads,
		KeyLength:  32,
	}
}

var _ validator.Parameters = (*Params)(nil)
