/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package linkFilter for Toolkits
package filter

import (
	"net/http"
	"strings"
)

//go:generate go run -mod=mod golang.org/x/tools/cmd/stringer@latest -type=MethodType -trimprefix=Method
type MethodType uint8

const (
	MethodAny MethodType = iota
	MethodGet
	MethodPut
	MethodDelete
	MethodConnect
	MethodHead
	MethodPost
	MethodPatch
	MethodTrace
	MethodOptions
	MethodTypeMax
)

func MethodIndex(method string) MethodType {
	switch strings.ToUpper(method) {
	case http.MethodGet:
		return MethodGet
	case http.MethodPut:
		return MethodPut
	case http.MethodDelete:
		return MethodDelete
	case http.MethodConnect:
		return MethodConnect
	case http.MethodHead:
		return MethodHead
	case http.MethodPost:
		return MethodPost
	case http.MethodPatch:
		return MethodPatch
	case http.MethodTrace:
		return MethodTrace
	case http.MethodOptions:
		return MethodOptions
	default:
		return MethodAny
	}
}
