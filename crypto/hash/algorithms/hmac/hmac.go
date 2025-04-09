/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package hmac

import (
	"crypto/hmac"
	"crypto/subtle"
	"fmt"
	"strings"

	"github.com/origadmin/toolkits/crypto/hash/core"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/utils"
)

// HMAC implements the HMAC hashing algorithm
type HMAC struct {
	config   *types.Config
	codec    interfaces.Codec
	hashHash core.Hash
}

func (c *HMAC) Type() string {
	return fmt.Sprintf("hmac-%s", c.hashHash.String())
}

type ConfigValidator struct {
	SaltLength int
}

func (v ConfigValidator) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return fmt.Errorf("salt length must be at least 8 bytes")
	}
	return nil
}

// NewHMACCrypto creates a new HMAC crypto instance
func NewHMACCrypto(hashType types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	validator := &ConfigValidator{}
	if err := validator.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid hmac256 config: %v", err)
	}
	if !strings.HasPrefix(hashType.String(), "hmac") {
		return nil, fmt.Errorf("unsupported hash type: %s", hashType.String())
	}
	switch hashType {
	case types.TypeHMAC256:
		hashType = "hmac-sha256"
	case types.TypeHMAC512:
		hashType = "hmac-sha512"
	}
	hash := strings.TrimLeft(hashType.String(), "hmac-")
	hashHash, err := core.ParseHash(hash)
	if err != nil {
		return nil, err
	}
	return &HMAC{
		config:   config,
		codec:    core.NewCodec(hashType),
		hashHash: hashHash,
	}, nil
}

func NewHMAC256Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return NewHMACCrypto(types.TypeHMAC256, config)
}
func NewHMAC512Crypto(config *types.Config) (interfaces.Cryptographic, error) {
	return NewHMACCrypto(types.TypeHMAC512, config)
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}

// Hash implements the hash method
func (c *HMAC) Hash(password string) (string, error) {
	salt, err := utils.GenerateSalt(c.config.SaltLength)
	if err != nil {
		return "", err
	}
	return c.HashWithSalt(password, string(salt))
}

// HashWithSalt implements the hash with salt method
func (c *HMAC) HashWithSalt(password, salt string) (string, error) {
	h := hmac.New(c.hashHash.New, []byte(salt))
	h.Write([]byte(password))
	hash := h.Sum(nil)
	return c.codec.Encode([]byte(salt), hash, ""), nil
}

// Verify implements the verify method
func (c *HMAC) Verify(parts *types.HashParts, password string) error {
	if !strings.HasPrefix(parts.Algorithm.String(), types.TypeHMAC.String()) {
		return core.ErrAlgorithmMismatch
	}
	switch parts.Algorithm {
	case types.TypeHMAC256:
		parts.Algorithm = "hmac-sha256"
	case types.TypeHMAC512:
		parts.Algorithm = "hmac-sha512"
	}

	hash := strings.TrimLeft(parts.Algorithm.String(), "hmac-")
	hashHash, err := core.ParseHash(hash)
	if err != nil {
		return err
	}

	h := hmac.New(hashHash.New, parts.Salt)
	h.Write([]byte(password))
	newHash := h.Sum(nil)
	if subtle.ConstantTimeCompare(newHash, parts.Hash) != 1 {
		return core.ErrPasswordNotMatch
	}

	return nil
}
