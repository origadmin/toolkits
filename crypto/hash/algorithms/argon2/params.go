package argon2

import (
	"fmt"
	"strconv"

	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// params represents parameters for Argon2 algorithm
type params struct {
	TimeCost   uint32
	MemoryCost uint32
	Threads    uint8
	KeyLength  uint32
}

func (p *params) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return fmt.Errorf("invalid salt length: %d, must be at least 8", config.SaltLength)
	}

	if p.TimeCost < 1 {
		return fmt.Errorf("invalid time cost: %d", p.TimeCost)
	}
	if p.MemoryCost < 1 {
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

func (p *params) FromMap(params map[string]string) error {
	if params == nil {
		return fmt.Errorf("params is nil")
	}
	// Parse time cost
	if v, ok := params["t"]; ok {
		timeCost, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid time cost: %v", err)
		}
		p.TimeCost = uint32(timeCost)
	}

	// Parse memory cost
	if v, ok := params["m"]; ok {
		memoryCost, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid memory cost: %v", err)
		}
		p.MemoryCost = uint32(memoryCost)
	}

	// Parse threads
	if v, ok := params["p"]; ok {
		threads, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return fmt.Errorf("invalid threads: %v", err)
		}
		p.Threads = uint8(threads)
	}

	// Parse key length
	if v, ok := params["k"]; ok {
		keyLength, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return fmt.Errorf("invalid key length: %v", err)
		}
		p.KeyLength = uint32(keyLength)
	}

	return nil
}

// parseParams parses Argon2 parameters from string
func parseParams(m map[string]string) (p *params, err error) {
	p = new(params)
	// Parse time cost
	if v, ok := m["t"]; ok {
		timeCost, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return p, fmt.Errorf("invalid time cost: %v", err)
		}
		p.TimeCost = uint32(timeCost)
	}

	// Parse memory cost
	if v, ok := m["m"]; ok {
		memoryCost, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return p, fmt.Errorf("invalid memory cost: %v", err)
		}
		p.MemoryCost = uint32(memoryCost)
	}

	// Parse threads
	if v, ok := m["p"]; ok {
		threads, err := strconv.ParseUint(v, 10, 8)
		if err != nil {
			return p, fmt.Errorf("invalid threads: %v", err)
		}
		p.Threads = uint8(threads)
	}

	// Parse key length
	if v, ok := m["k"]; ok {
		keyLength, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return p, fmt.Errorf("invalid key length: %v", err)
		}
		p.KeyLength = uint32(keyLength)
	}

	return p, nil
}

// String returns the string representation of parameters
func (p *params) String() string {
	return codecPkg.EncodeParams(p.ToMap())
}

// ToMap converts params to a map[string]string
func (p *params) ToMap() map[string]string {
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

func DefaultParams() types.Params {
	return &params{
		TimeCost:   constants.DefaultTimeCost,
		MemoryCost: constants.DefaultMemoryCost,
		Threads:    constants.DefaultThreads,
		KeyLength:  32,
	}
}
