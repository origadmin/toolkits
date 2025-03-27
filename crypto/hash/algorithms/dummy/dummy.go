/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dummy

import (
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/base"
	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// DummyCrypto implements a dummy hashing algorithm
type DummyCrypto struct {
	config *types.Config
}

// NewDummyCrypto creates a new dummy crypto instance
func NewDummyCrypto(config *types.Config) (interfaces.Cryptographic, error) {
	return &DummyCrypto{
		config: config,
	}, nil
}

// Hash implements the hash method
func (c *DummyCrypto) Hash(password string) (string, error) {
	return "", fmt.Errorf("dummy algorithm not implemented")
}

// HashWithSalt implements the hash with salt method
func (c *DummyCrypto) HashWithSalt(password, salt string) (string, error) {
	return "", fmt.Errorf("dummy algorithm not implemented")
}

// Verify implements the verify method
func (c *DummyCrypto) Verify(hashed, password string) error {
	return fmt.Errorf("dummy algorithm not implemented")
}

// DummyHashEncoder implements the hash encoder interface
type DummyHashEncoder struct {
	*base.BaseHashCodec
}

// NewDummyHashEncoder creates a new dummy hash encoder
func NewDummyHashEncoder() interfaces.HashCodec {
	return &DummyHashEncoder{
		BaseHashCodec: base.NewBaseHashCodec(types.TypeCustom),
	}
}
