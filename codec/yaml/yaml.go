/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package yaml provides the yaml functions
package yaml

import (
	"gopkg.in/yaml.v3"
)

var (
	Marshal    = yaml.Marshal
	Unmarshal  = yaml.Unmarshal
	NewDecoder = yaml.NewDecoder
	NewEncoder = yaml.NewEncoder
)

// MarshalToString returns json string, and ignores error
func MarshalToString(v any) string {
	b, err := Marshal(v)
	if err != nil {
		return ""
	}
	return string(b)
}

// MustToString returns json string, or panic
func MustToString(v any) string {
	data, err := Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(data)
}
