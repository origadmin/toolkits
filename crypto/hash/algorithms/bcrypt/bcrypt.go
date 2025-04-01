/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package bcrypt

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// Bcrypt implements the Bcrypt hashing algorithm
type Bcrypt struct {
	params *Params
	config *types.Config
	codec  interfaces.Codec
}

func (c *Bcrypt) Type() string {
	return types.TypeBcrypt.String()
}

type Params struct {
	Cost int
}

func (p *Params) String() string {
	return fmt.Sprintf("c:%d", p.Cost)
}

func parseParams(params string) (*Params, error) {
	result := &Params{}

	if params == "" {
		return result, nil
	}
	for _, param := range strings.Split(params, ",") {
		parts := strings.Split(param, ":")
		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid bcrypt param format: %s", param)
		}
		switch parts[0] {
		case "c":
			cost, err := strconv.ParseInt(parts[1], 10, 32)
			if err != nil {
				return nil, fmt.Errorf("invalid cost: %v", err)
			}
			result.Cost = int(cost)
		}
	}
	return result, nil
}

type ConfigValidator struct {
	params *Params
}

func (v ConfigValidator) Validate(config *types.Config) interface{} {
	if config.SaltLength < 8 {
		return core.ErrSaltLengthTooShort
	}
	if v.params.Cost < 4 || v.params.Cost > 31 {
		return core.ErrCostOutOfRange
	}
	return nil
}

// NewBcryptCrypto creates a new Bcrypt crypto instance
func NewBcryptCrypto(config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	if config.ParamConfig == "" {
		config.ParamConfig = DefaultParams().String()
	}
	params, err := parseParams(config.ParamConfig)
	if err != nil {
		return nil, fmt.Errorf("invalid bcrypt param config: %v", err)
	}

	validator := &ConfigValidator{
		params: params,
	}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid bcrypt config: %v", err)
	}
	return &Bcrypt{
		config: config,
		params: params,
		codec:  core.NewCodec(types.TypeBcrypt),
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength:  16,
		ParamConfig: DefaultParams().String(),
	}
}

func DefaultParams() *Params {
	return &Params{
		Cost: bcrypt.DefaultCost,
	}
}

// Hash implements the hash method
func (c *Bcrypt) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, string(salt))
}

// HashWithSalt implements the hash with salt method
func (c *Bcrypt) HashWithSalt(password, salt string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password+salt), c.params.Cost)
	if err != nil {
		return "", err
	}
	return c.codec.Encode([]byte(salt), hash), nil
}

// Verify implements the verify method
func (c *Bcrypt) Verify(hashed, password string) error {
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}
	if parts.Algorithm != types.TypeBcrypt {
		return core.ErrAlgorithmMismatch
	}
	err = bcrypt.CompareHashAndPassword(parts.Hash, []byte(password+string(parts.Salt)))
	if err != nil {
		return core.ErrPasswordNotMatch
	}
	return nil
}
