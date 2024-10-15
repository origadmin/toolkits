// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package replacer provides a string replacement mechanism with custom replace functions and keywords.
package replacer

import (
	"strings"

	"github.com/goexts/ggb/settings"
)

// DefaultKeyword defines the default keyword.
const DefaultKeyword = "$$"

// ReplaceFunc is a function type that accepts a string and returns a replaced string.
type ReplaceFunc func(key, value string) string

// Replacer interface defines methods for setting keywords and performing replacements.
type Replacer interface {
	AddHook(string, ReplaceFunc)              // Adds a new replace hook.
	Replace(string, map[string]string) string // Replaces substrings based on provided key-value pairs.
}

// replace struct implements the Replacer interface, storing replace hooks and the current keyword.
type replace struct {
	keyword string
	hooks   map[string]ReplaceFunc
}

// hitKey returns the key wrapped with the current keyword.
func (r *replace) hitKey(key string) string {
	return r.keyword + key + r.keyword
}

// AddHook adds a new replace hook.
func (r *replace) AddHook(key string, replaceFunc ReplaceFunc) {
	r.hooks[key] = replaceFunc
}

// Replace within the replacement struct iterates through the values map, applying custom hooks and
// replacing placeholders found in the source string.
func (r *replace) Replace(src string, values map[string]string) string {
	for key, val := range values {
		if key != "" && strings.Contains(src, r.hitKey(key)) {
			if hook, exists := r.hooks[key]; exists {
				val = hook(key, val)
			}
			src = strings.ReplaceAll(src, r.hitKey(key), val)
		}
	}
	return src
}

// New returns a new Replacer instance with default settings.
func New(ss ...Setting) Replacer {
	op := settings.Apply(&replace{
		keyword: DefaultKeyword,
		hooks:   make(map[string]ReplaceFunc),
	}, ss)
	return op
}
