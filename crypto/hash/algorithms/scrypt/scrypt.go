/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package scrypt

import (
	"crypto/subtle"
	"fmt"

	"golang.org/x/crypto/scrypt"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/scheme"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/validator"
	"github.com/origadmin/toolkits/crypto/rand"
)

// Scrypt implements the Scrypt hashing algorithm
type Scrypt struct {
	params *Params
	config *types.Config
}

var specScrypt = types.Spec{
	Name: types.SCRYPT,
}

func (c *Scrypt) Spec() types.Spec {
	return specScrypt
}

// NewScrypt creates a new Scrypt crypto instance
func NewScrypt(config *types.Config) (scheme.Scheme, error) {
	// Ensure algorithm-specific default config is applied when caller passes nil.
	if config == nil {
		config = DefaultConfig()
	}

	v, err := validator.ValidateParams(config, DefaultParams())
	if err != nil {
		return nil, fmt.Errorf("invalid scrypt param config: %v", err)
	}

	return &Scrypt{
		params: v.Params,
		config: v.Config,
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
		Params:     DefaultParams().ToMap(),
	}
}

// Hash implements the hash method
func (c *Scrypt) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *Scrypt) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	hash, err := scrypt.Key([]byte(password), salt, c.params.N, c.params.R, c.params.P, c.params.KeyLen)
	if err != nil {
		return nil, err
	}
	return types.NewHashParts(c.Spec()).WithHashSalt(hash, salt).WithParams(c.params.ToMap()), nil
}

// Verify implements the verify method
func (c *Scrypt) Verify(parts *types.HashParts, password string) error {
	// parts.Spec is already of type types.Spec, so no need to parse it again.
	// We can directly use parts.Spec.Name for comparison.
	if parts.Spec.Name != types.SCRYPT {
		return errors.ErrAlgorithmMismatch
	}
	// Parse parameters
	params, err := FromParams(parts.Params)
	if err != nil {
		return err
	}
	hash, err := scrypt.Key([]byte(password), parts.Salt, params.N, params.R, params.P, params.KeyLen)
	if err != nil {
		return err
	}
	if subtle.ConstantTimeCompare(parts.Hash, hash) != 1 {
		return errors.ErrPasswordNotMatch
	}

	return nil
}
