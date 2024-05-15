// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package replacer provides a string replacement mechanism with custom replace functions and delimiters.
package replacer

import "strings"

// DefaultDelimiter defines the default delimiter.
const DefaultDelimiter = "$$"

// ReplaceFunc is a function type that accepts a string and returns a replaced string.
type ReplaceFunc func(string) string

// Replacer interface defines methods for setting delimiters and performing replacements.
type Replacer interface {
	SetDelimiter(string)                      // Sets the delimiter.
	Replace(string, map[string]string) string // Replaces substrings based on provided key-value pairs.
}

// replace struct implements the Replacer interface, storing replace hooks and the current delimiter.
type replace struct {
	Hooks     map[string]ReplaceFunc // Mapping of keys to replace functions.
	Delimiter string                 // The currently used delimiter.
}

var replacer = New()

// Replace replaces substrings in the source string based on provided key-value pairs.
func Replace(src string, values map[string]string) string {
	return replacer.Replace(src, values)
}

// SetDelimiter sets the delimiter used by the replacer.
func (r *replace) SetDelimiter(s string) {
	r.Delimiter = s
}

// getKey returns the key wrapped with the current delimiter.
func (r *replace) getKey(key string) string {
	return getDelimiterKey(key, r.Delimiter)
}

// Replace within the replacement struct iterates through the values map, applying custom hooks and
// replacing placeholders found in the source string.
func (r *replace) Replace(src string, values map[string]string) string {
	for key, val := range values {
		if key != "" && strings.Contains(src, r.getKey(key)) {
			if hook, exists := r.Hooks[key]; exists {
				val = hook(val)
			}
			src = strings.ReplaceAll(src, r.getKey(key), val)
		}
	}
	return src
}

// New returns a new Replacer instance with default settings.
func New() Replacer {
	return &replace{
		Hooks:     map[string]ReplaceFunc{},
		Delimiter: DefaultDelimiter,
	}
}

// WithHooks returns a new Replacer instance with the specified hooks.
func WithHooks(hooks map[string]ReplaceFunc) Replacer {
	return &replace{
		Hooks:     hooks,
		Delimiter: DefaultDelimiter,
	}
}

// getDelimiterKey wraps the key with the provided delimiter.
func getDelimiterKey(key string, delimiter string) string {
	return delimiter + key + delimiter
}
