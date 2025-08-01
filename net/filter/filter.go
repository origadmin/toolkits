/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package filter

import (
	"net/http"
)

const (
	AnySymbol        = "*"
	Terminator       = "*"
	DefaultDelimiter = ":"
)

type Filter interface {
	Skip(req *http.Request) bool
	Denied(method, path string) bool
	Allowed(method, path string) bool
	AllowStrings(method string) []string
	DenyStrings(method string) []string
}

type Setting = func(*Option)

type Option struct {
	Delimiter         string
	SkipAny           bool
	DenyOff           bool
	AllowOff          bool
	Allows            []string
	Denies            []string
	ExpectedElements  uint
	FalsePositiveRate float64
}

func WithDelimiter(delimiter string) Setting {
	return func(o *Option) {
		o.Delimiter = delimiter
	}
}

func WithAllows(allows ...string) Setting {
	return func(o *Option) {
		o.Allows = allows
	}
}

func WithDenies(denies ...string) Setting {
	return func(o *Option) {
		o.Denies = denies
	}
}

// SkipAny strong rule check. After this function is enabled, wildcard method match check is skipped
func SkipAny() Setting {
	return func(o *Option) {
		o.SkipAny = true
	}
}

func DisableAllow() Setting {
	return func(o *Option) {
		o.AllowOff = true
	}
}

func DisableDeny() Setting {
	return func(o *Option) {
		o.DenyOff = true
	}
}

func contains(roots []*Link, path string) bool {
	var root *Link
	if len(roots) == 0 {
		return false
	}
	paths := splitPath(path)
	for _, root = range roots {

		if root.Contains(paths) {
			return true
		}
	}
	return false
}

func showTree(prefix string, list []*Link) {
	if len(list) == 0 {
		return
	}
	for i := range list {
		showTree(prefix+"-"+list[i].Path, list[i].Subs)
	}
}
