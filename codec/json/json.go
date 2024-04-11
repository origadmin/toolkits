// Copyright (c) 2024 OrigAdmin. All rights reserved.

//go:build !jsoniter

// Package json provides the json functions based on standard library
package json

import (
	"encoding/json"
)

var (
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
	MarshalIndent = json.MarshalIndent
)

// MarshalToString returns json string, and ignores error
func MarshalToString(v any) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}

// MustToString returns json string, or panic
func MustToString(v any) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
