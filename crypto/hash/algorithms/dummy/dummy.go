/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package dummy

import (
	"fmt"

	"github.com/origadmin/toolkits/crypto/hash/interfaces"
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Crypto implements a dummy hashing algorithm
type Crypto struct {
	config *types.Config
}

// NewDummyCrypto creates a new dummy crypto instance
func NewDummyCrypto(config *types.Config) (interfaces.Cryptographic, error) {
	return &Crypto{
		config: config,
	}, nil
}

// Hash implements the hash method
func (c *Crypto) Hash(password string) (string, error) {
	return "", fmt.Errorf("dummy algorithm not implemented")
}

// HashWithSalt implements the hash with salt method
func (c *Crypto) HashWithSalt(password, salt string) (string, error) {
	return "", fmt.Errorf("dummy algorithm not implemented")
}

// Verify implements the verify method
func (c *Crypto) Verify(hashed, password string) error {
	return fmt.Errorf("dummy algorithm not implemented")
}
