/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package noop

import (
	"errors"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/constants"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

var unknownType = types.Type{Name: constants.UNKNOWN}

// noop implements a noop hashing algorithm
type noop struct {
}

func (n *noop) Type() types.Type {
	return unknownType
}

// New creates a new noop crypto instance
func New(config *types.Config) (interfaces.Cryptographic, error) {
	return nil, errors.New("algorithm not implemented")
}

// Hash implements the hash method
func (n *noop) Hash(password string) (*types.HashParts, error) {
	return nil, fmt.Errorf("algorithm not implemented")
}

// HashWithSalt implements the hash with salt method
func (n *noop) HashWithSalt(password string, salt []byte) (*types.HashParts, error) {
	return nil, fmt.Errorf("noop algorithm not implemented")
}

// Verify implements the verify method
func (n *noop) Verify(parts *types.HashParts, password string) error {
	return fmt.Errorf("algorithm not implemented")
}

func DefaultConfig() *types.Config {
	return &types.Config{}
}

func Noop() (interfaces.Cryptographic, error) {
	return &noop{}, nil
}
