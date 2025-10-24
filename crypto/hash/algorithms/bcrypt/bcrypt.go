/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package bcrypt

import (
	"golang.org/x/crypto/bcrypt"

	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
	"github.com/origadmin/toolkits/crypto/hash/validator"
	"github.com/origadmin/toolkits/crypto/rand"
)

// Bcrypt implements the Bcrypt hashing algorithm
type Bcrypt struct {
	params *Params
	config *types.Config
}

var bcryptSpec = types.Spec{Name: types.BCRYPT}

func (c *Bcrypt) Spec() types.Spec {
	return bcryptSpec
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
// WARNING: Manually concatenating salt for Bcrypt is INSECURE as Bcrypt handles salt internally.
// This implementation is for framework consistency, but should be used with caution.
func (c *Bcrypt) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	var data []byte
	if len(salt) > 0 {
		data = append([]byte(password), salt...)
	} else {
		data = []byte(password)
	}

	hashBytes, err := bcrypt.GenerateFromPassword(data, c.params.Cost)
	if err != nil {
		return nil, err
	}
	return types.NewHashPartsFull(c.Spec(), hashBytes, salt, c.params.ToMap()), nil
}

// Verify implements the verify method
// WARNING: Manually concatenating salt for Bcrypt is INSECURE as Bcrypt handles salt internally.
// This implementation is for framework consistency, but should be used with caution.
func (c *Bcrypt) Verify(parts *types.HashParts, password string) error {
	// parts.Algorithm is already of type types.Spec, so no need to parse it again.
	// We can directly use parts.Algorithm.Name for comparison.
	if types.BCRYPT != parts.Algorithm.Name {
		return errors.ErrAlgorithmMismatch
	}

	var data []byte
	if len(parts.Salt) > 0 {
		data = append([]byte(password), parts.Salt...)
	} else {
		data = []byte(password)
	}

	if err := bcrypt.CompareHashAndPassword(parts.Hash, data); err != nil {
		return errors.ErrPasswordNotMatch
	}
	return nil
}

func NewBcrypt(config *types.Config) (interfaces.Cryptographic, error) {
	// Ensure algorithm-specific default config is applied when caller passes nil.
	if config == nil {
		config = DefaultConfig()
	}

	v, err := validator.ValidateParams(config, DefaultParams())
	if err != nil {
		return nil, err
	}

	return &Bcrypt{
		params: v.Params,
		config: v.Config,
	}, nil
}

func DefaultConfig() *types.Config {
	return &types.Config{
		SaltLength: types.DefaultSaltLength,
	}
}
