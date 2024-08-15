// Copyright (c) 2024 OrigAdmin. All rights reserved.

//go:build sonic

// Package json provides the json functions based on standard library
package json

import (
	"github.com/bytedance/sonic"
)

var (
	json          = sonic.ConfigStd
	Marshal       = json.Marshal
	Unmarshal     = json.Unmarshal
	NewDecoder    = json.NewDecoder
	NewEncoder    = json.NewEncoder
	MarshalIndent = json.MarshalIndent
)

// MarshalToString returns json string, and ignores error
func MarshalToString(v any) string {
	bytes, err := json.MarshalToString(v)
	if err != nil {
		return ""
	}
	return bytes
}

// MustToString returns json string, or panic
func MustToString(v any) string {
	bytes, err := json.MarshalToString(v)
	if err != nil {
		panic(err)
	}
	return bytes
}
