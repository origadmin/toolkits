/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package crc

import (
	"github.com/origadmin/toolkits/crypto/hash/types"
)

// Params holds the parameters for CRC algorithms
type Params struct {
	// No specific parameters for CRC beyond SaltLength in Config
}

func (p *Params) IsNil() bool {
	return p == nil
}

func (p *Params) Validate(config *types.Config) error {
	return nil
}

func (p *Params) String() string {
	return ""
}

func (p *Params) ToMap() map[string]string {
	return map[string]string{}
}

func (p *Params) FromMap(params map[string]string) error {
	return nil
}
