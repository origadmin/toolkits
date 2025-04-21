/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package hmac implements the functions, types, and interfaces for the module.
package hmac

import (
	"fmt"
	"strings"

	"github.com/origadmin/toolkits/crypto/hash/core"
)

// Params represents parameters for Argon2 algorithm
type Params struct {
	Type string `json:"type"`
}

// parseParams parses Argon2 parameters from string
func parseParams(params string) (result Params, err error) {
	// Handle empty string case
	if params == "" {
		return result, nil
	}

	kv, err := core.ParseParams(params)
	if err != nil {
		return result, err
	}
	// Parse time cost
	if v, ok := kv["t"]; ok {
		result.Type = v
	}

	return result, nil
}

// String returns the string representation of parameters
func (p Params) String() string {
	var parts []string
	if p.Type != "" {
		parts = append(parts, fmt.Sprintf("t:%s", p.Type))
	}

	return strings.Join(parts, ",")
}
