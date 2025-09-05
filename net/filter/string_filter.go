/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package filter

import (
	"net/http"
	"strings"

	"github.com/goexts/generic/configure"
)

type stringFilter struct {
	skipAllow bool
	allows    [MethodTypeMax][]string
	skipDeny  bool
	denies    [MethodTypeMax][]string
}

func (s stringFilter) AllowStrings(method string) []string {
	//TODO implement me
	panic("implement me")
}

func (s stringFilter) DenyStrings(method string) []string {
	//TODO implement me
	panic("implement me")
}

func (s stringFilter) Skip(req *http.Request) bool {
	path := req.URL.Path
	method := req.Method
	return s.Denied(method, path) || !s.Allowed(method, path)
}

func (s stringFilter) Denied(method, path string) bool {
	if s.skipDeny {
		return false
	}

	idx := MethodIndex(method)
	if roots := s.denies[idx]; len(roots) > 0 {
		if containsString(roots, path) {
			// slog.Info("Deny Check on Contains *", "method", method, "path", path)
			return true
		}
	}

	if idx == MethodAny {
		return false
	}

	if roots := s.denies[idx]; len(roots) > 0 {
		if containsString(roots, path) {
			// slog.Info("Deny Check on Contains Method", "method", method, "path", path)
			return true
		}
	}
	return false
}

func (s stringFilter) Allowed(method, path string) bool {
	if s.skipAllow {
		return true
	}

	idx := MethodIndex(method)
	if roots := s.allows[idx]; len(roots) > 0 {
		if containsString(roots, path) {
			// slog.Info("Deny Check on Contains *", "method", method, "path", path)
			return true
		}
	}

	if idx == MethodAny {
		return false
	}

	if roots := s.allows[idx]; len(roots) > 0 {
		if containsString(roots, path) {
			return true
		}
	}
	return false

}

func NewStringFilter(ss ...Setting) Filter {
	opt := settings.Apply(&Option{
		Delimiter: DefaultDelimiter,
		SkipAny:   false,
		DenyOff:   false,
		AllowOff:  false,
	}, ss)

	f := &stringFilter{
		skipAllow: true,
		allows:    [MethodTypeMax][]string{},
		skipDeny:  true,
		denies:    [MethodTypeMax][]string{},
	}

	for _, s := range opt.Allows {
		if s == "" {
			continue
		}
		methods, path := splitURI(s, opt.Delimiter)
		for i := range methods {
			idx := MethodIndex(methods[i])
			f.allows[idx] = append(f.allows[idx], path)
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
		}
	}

	for _, s := range opt.Denies {
		if s == "" {
			continue
		}
		methods, path := splitURI(s, opt.Delimiter)
		for i := range methods {
			idx := MethodIndex(methods[i])
			if !strings.HasSuffix(path, "/") {
				path += "/"
			}
			f.denies[idx] = append(f.denies[idx], path)
		}
	}
	return f
}

func containsString(roots []string, path string) bool {
	pathLen := len(path)
	if strings.HasSuffix(path, "/") {
		path += "/"
	}
	for _, rule := range roots {
		if rs := len(rule); pathLen >= rs && path[:rs] == rule {
			return true
		}
	}
	return false
}
