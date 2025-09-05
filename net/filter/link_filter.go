/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package filter

import (
	"net/http"

	"github.com/goexts/generic/configure"
)

type linkFilter struct {
	delimiter string
	AllowOff  bool
	allows    [MethodTypeMax][]*Link
	DenyOff   bool
	denies    [MethodTypeMax][]*Link
	skipAny   bool
}

func (f *linkFilter) AllowStrings(method string) []string {
	links := f.allows[MethodIndex(method)]
	var result []string
	for i := range links {
		ret := links[i].StringList("")
		if len(ret) > 0 {
			result = append(result, ret...)
		}
	}
	return result
}

func (f *linkFilter) DenyStrings(method string) []string {
	links := f.denies[MethodIndex(method)]
	var result []string
	for i := range links {
		ret := links[i].StringList("")
		if len(ret) > 0 {
			result = append(result, ret...)
		}
	}
	return result
}

func (f *linkFilter) Skip(req *http.Request) bool {
	path := req.URL.Path
	method := req.Method
	return f.Denied(method, path) || !f.Allowed(method, path)
}

func (f *linkFilter) Allowed(method, path string) bool {
	if f.AllowOff {
		return true
	}

	idx := MethodIndex(method)
	if f.skipAny && (idx == MethodAny || len(f.allows[idx]) == 0) {
		return false
	}

	if contains(f.allows[idx], path) {
		return true
	}

	if f.skipAny || idx == MethodAny {
		return false
	}

	return contains(f.allows[MethodAny], path)
}

func (f *linkFilter) Denied(method string, path string) bool {
	if f.DenyOff {
		return false
	}

	idx := MethodIndex(method)
	if f.skipAny && (idx == MethodAny || len(f.denies[idx]) == 0) {
		return false
	}

	if contains(f.denies[idx], path) {
		return true
	}

	if f.skipAny || idx == MethodAny {
		return false
	}

	return contains(f.denies[MethodAny], path)
}

func hasNext(paths []string) ([]string, bool) {
	if len(paths) == 0 {
		return nil, false
	}
	return paths[1:], true
}

func NewFilter(ss ...Setting) Filter {
	o := settings.Apply(&Option{
		Delimiter: DefaultDelimiter,
		SkipAny:   false,
		DenyOff:   false,
		AllowOff:  false,
	}, ss)

	f := &linkFilter{
		delimiter: o.Delimiter,
		skipAny:   o.SkipAny,
		AllowOff:  o.AllowOff,
		allows:    [MethodTypeMax][]*Link{},
		DenyOff:   o.DenyOff,
		denies:    [MethodTypeMax][]*Link{},
	}

	for _, s := range o.Allows {
		if s == "" {
			continue
		}
		methods, path := splitURI(s, o.Delimiter)
		for i := range methods {
			idx := MethodIndex(methods[i])
			if f.skipAny && idx == MethodAny {
				continue
			}
			f.allows[idx] = buildLinkRoot(f.allows[idx], path)
		}
		f.AllowOff = false
	}

	for _, s := range o.Denies {
		if s == "" {
			continue
		}
		methods, path := splitURI(s, o.Delimiter)
		for i := range methods {
			idx := MethodIndex(methods[i])
			if f.skipAny && idx == MethodAny {
				continue
			}
			f.denies[idx] = buildLinkRoot(f.denies[idx], path)
		}
		f.DenyOff = false
	}

	return f
}
