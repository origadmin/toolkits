/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hmac implements the functions, types, and interfaces for the module.
package hmac

import (
	codecPkg "github.com/origadmin/toolkits/crypto/hash/codec"
	"github.com/origadmin/toolkits/crypto/hash/errors"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Params represents parameters for Argon2 algorithm
type Params struct {
}

func (p *Params) Validate(config *types.Config) error {
	if config.SaltLength < 8 {
		return errors.ErrSaltLengthTooShort
	}
	return nil
}

func (p *Params) ToMap() map[string]string {
	m := make(map[string]string)
	return m
}

func (p *Params) FromMap(params map[string]string) error {
	return nil
}

// String returns the string representation of parameters
func (p *Params) String() string {
	return codecPkg.EncodeParams(p.ToMap())
}

// FromMap parses Argon2 parameters from string
func FromMap(m map[string]string) (params *Params, err error) {
	params = &Params{}
	return params, nil
}

func DefaultParams() *Params {
	return &Params{}
}
