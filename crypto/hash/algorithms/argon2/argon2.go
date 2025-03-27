/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package argon2

import (
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/crypto/argon2"

	"github.com/origadmin/toolkits/crypto/hash/base"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

func init() {
	base.RegisterAlgorithm(types.TypeArgon2, NewArgon2Crypto)
}

// Argon2Crypto implements the Argon2 hashing algorithm
type Argon2Crypto struct {
	config *types.Config
	codec  interfaces.HashCodec
}

// Argon2ConfigValidator implements the config validator for Argon2
type Argon2ConfigValidator struct{}

// Validate validates the Argon2 configuration
func (v *Argon2ConfigValidator) Validate(config *types.Config) error {
	if config.TimeCost < 1 {
		return fmt.Errorf("invalid time cost: %d", config.TimeCost)
	}
	if config.MemoryCost < 1 {
		return fmt.Errorf("invalid memory cost: %d", config.MemoryCost)
	}
	if config.Threads < 1 {
		return fmt.Errorf("invalid threads: %d", config.Threads)
	}
	if config.SaltLength < 1 {
		return fmt.Errorf("invalid salt length: %d", config.SaltLength)
	}
	return nil
}

// NewArgon2Crypto creates a new Argon2 crypto instance
func NewArgon2Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	validator := &Argon2ConfigValidator{}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid argon2 config: %v", err)
	}
	return &Argon2Crypto{
		config: config,
		codec:  base.GetCodec(types.TypeArgon2),
	}, nil
}

// Hash implements the hash method
func (c *Argon2Crypto) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *Argon2Crypto) HashWithSalt(password, salt string) (string, error) {
	hash := argon2.IDKey(
		[]byte(password),
		[]byte(salt),
		c.config.TimeCost,
		c.config.MemoryCost,
		c.config.Threads,
		32, // Key length
	)

	// 将配置参数编码到密码中
	params := fmt.Sprintf("%d,%d,%d", c.config.TimeCost, c.config.MemoryCost, c.config.Threads)
	return c.codec.Encode([]byte(salt), hash, params), nil
}

// Verify implements the verify method
func (c *Argon2Crypto) Verify(hashed, password string) error {
	parts, err := c.codec.Decode(hashed)
	if err != nil {
		return err
	}

	if parts.Algorithm != types.TypeArgon2 {
		return fmt.Errorf("algorithm mismatch")
	}

	// 从密码中解码配置参数
	params := strings.Split(parts.Params, ",")
	if len(params) != 3 {
		return fmt.Errorf("invalid argon2 params")
	}

	timeCost, err := strconv.ParseUint(params[0], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid time cost: %v", err)
	}

	memoryCost, err := strconv.ParseUint(params[1], 10, 32)
	if err != nil {
		return fmt.Errorf("invalid memory cost: %v", err)
	}

	threads, err := strconv.ParseUint(params[2], 10, 8)
	if err != nil {
		return fmt.Errorf("invalid threads: %v", err)
	}

	hash := argon2.IDKey(
		[]byte(password),
		parts.Salt,
		uint32(timeCost),
		uint32(memoryCost),
		uint8(threads),
		32, // Key length
	)

	if string(hash) != string(parts.Hash) {
		return fmt.Errorf("password not match")
	}

	return nil
}

// Argon2HashEncoder implements the hash encoder interface
type Argon2HashEncoder struct {
	*base.BaseHashCodec
}

// NewArgon2HashEncoder creates a new Argon2 hash encoder
func NewArgon2HashEncoder() interfaces.HashCodec {
	return &Argon2HashEncoder{
		BaseHashCodec: base.NewBaseHashCodec(types.TypeArgon2).(*base.BaseHashCodec),
	}
}
