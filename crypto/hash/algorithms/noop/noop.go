/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package noop

import (
	"errors"
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// noop implements a dummy hashing algorithm
type noop struct {
}

func (n *noop) Type() string {
	return "dummy"
}

// NewNoopCrypto creates a new noop crypto instance
func NewNoopCrypto(config *types.Config) (interfaces.Cryptographic, error) {
	return nil, errors.New("algorithm not implemented")
}

// Hash implements the hash method
func (n *noop) Hash(password string) (string, error) {
	return "", fmt.Errorf("algorithm not implemented")
}

// HashWithSalt implements the hash with salt method
func (n *noop) HashWithSalt(password, salt string) (string, error) {
	return "", fmt.Errorf("dummy algorithm not implemented")
}

// Verify implements the verify method
func (n *noop) Verify(parts *types.HashParts, password string) error {
	return fmt.Errorf("algorithm not implemented")
}

func DefaultConfig() *types.Config {
	return &types.Config{}
}
