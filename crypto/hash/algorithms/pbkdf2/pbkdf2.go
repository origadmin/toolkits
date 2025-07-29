/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package pbkdf2

import (
	"crypto/subtle"
	"fmt"

	"github.com/goexts/generic"
	"golang.org/x/crypto/pbkdf2"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/internal/stdhash"
	"github.com/origadmin/toolkits/crypto/hash/internal/validator"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/rand"
)

// PBKDF2 implements the PBKDF2 hashing algorithm
type PBKDF2 struct {
	p        types.Type
	params   *Params
	config   *types.Config
	hashHash stdhash.Hash
}

// Hash implements the hash method
func (c *PBKDF2) Hash(password string) (*types.HashParts, error) {
	salt, err := rand.RandomBytes(c.config.SaltLength)
	if err != nil {
		return nil, err
	}
	return c.HashWithSalt(password, salt)
}

// HashWithSalt implements the hash with salt method
func (c *PBKDF2) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	hashHash := pbkdf2.Key([]byte(password), salt, c.params.Iterations, int(c.params.KeyLength), c.hashHash.New)
	return types.NewHashPartsWithParams(c.Type(), salt, hashHash, c.params.ToMap()), nil
}

// Verify implements the verify method
func (c *PBKDF2) Verify(parts *types.HashParts, password string) error {
	algorithm, err := types.ParseType(parts.Algorithm)
	if err != nil {
		return err
	}
	if algorithm.Name != constants.PBKDF2 {
		return errors.ErrInvalidAlgorithm
	}

	// Parse parameters
	params, err := FromMap(parts.Params)
	if err != nil {
		return err
	}
	hashHash, err := stdhash.ParseHash(algorithm.Underlying)
	if err != nil {
		return err
	}
	newHash := pbkdf2.Key([]byte(password), parts.Salt, params.Iterations, int(params.KeyLength), hashHash.New)
	if subtle.ConstantTimeCompare(newHash, parts.Hash) != 1 {
		return errors.ErrPasswordNotMatch
	}
	return nil
}

func (c *PBKDF2) Type() types.Type {
	return c.p
}

// NewPBKDF2 creates a new PBKDF2 crypto instance
func NewPBKDF2(p types.Type, config *types.Config) (interfaces.Cryptographic, error) {
	if config == nil {
		config = DefaultConfig()
	}
	if config.ParamConfig == "" {
		config.ParamConfig = DefaultParams().String()
	}

	v := validator.WithParams(&Params{})
	if err := v.Validate(config); err != nil {
		return nil, fmt.Errorf("invalid pbkdf2 config: %v", err)
	}
	p = generic.Must(ResolveType(p))
	hashHash, err := stdhash.ParseHash(p.Underlying)
	if err != nil {
		return nil, err
	}
	return &PBKDF2{
		p:        p,
		params:   v.Params(),
		config:   config,
		hashHash: hashHash,
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength:  16,
		ParamConfig: DefaultParams().String(),
	}
}
