/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package bcrypt

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// Bcrypt implements the Bcrypt hashing algorithm
type Bcrypt struct {
	params Params
	config *types.Config
}

func (c *Bcrypt) Type() types.Type {
	return types.Type{Name: constants.BCRYPT}
}

// Hash implements the hash method
func (c *Bcrypt) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *Bcrypt) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	newpass := password + string(salt)
	hash, err := bcrypt.GenerateFromPassword([]byte(newpass), c.params.Cost)
	if err != nil {
		return nil, err
	}
	return types.NewHashPartsWithParams(c.Type(), salt, hash, c.params.ToMap()), nil
}

// Verify implements the verify method
func (c *Bcrypt) Verify(parts *types.HashParts, password string) error {
	algorithm, err := types.ParseAlgorithm(parts.Algorithm)
	if err != nil {
		return err
	}
	if constants.BCRYPT != algorithm.Name {
		return errors.ErrAlgorithmMismatch
	}
	newpass := password + string(parts.Salt)
	if err := bcrypt.CompareHashAndPassword(parts.Hash, []byte(newpass)); err != nil {
		return errors.ErrPasswordNotMatch
	}
	return nil
}

func NewBcrypt(config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	validator := ConfigValidator{
		config: config,
		params: DefaultParams(),
	}

	params, err := parseParams(config.ParamConfig)
	if err != nil {
		return nil, err
	}
	return &Bcrypt{
		params: params,
		config: config,
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: 16,
	}
}
