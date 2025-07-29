/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package crc

// Params holds the parameters for CRC algorithms
type Params struct {
	// No specific parameters for CRC beyond SaltLength in Config
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
