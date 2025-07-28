/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hmac implements the functions, types, and interfaces for the module.
package hmac

import (
	"github.com/origadmin/toolkits/crypto/hash/codec"
)

// Params represents parameters for Argon2 algorithm
type Params struct {
	Type string `json:"type"`
}

// parseParams parses Argon2 parameters from string
func parseParams(paramsMap map[string]string) (result Params, err error) {
	// Parse time cost
	if v, ok := paramsMap["t"]; ok {
		result.Type = v
	}

	return result, nil
}

// String returns the string representation of parameters
func (p Params) String() string {
	paramsMap := make(map[string]string)
	if p.Type != "" {
		paramsMap["t"] = p.Type
	}
	return codec.EncodeParams(paramsMap)
}
