// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package replacer provides a string Replacement mechanism with custom Replacement functions and keywords.
package replacer

import (
	"bytes"
	"strings"

	"github.com/goexts/ggb/settings"
)

const (
	// DefaultStartKeyword defines the default keyword.
	DefaultStartKeyword          = "${"
	DefaultHostStartKeyword      = "@{"
	DefaultMatchStartKeyword     = "@"
	DefaultMatchSeparatorKeyword = "="
	DefaultEndKeyword            = "}"
	DefaultHostEndKeyword        = ":"
	DefaultSeparatorKeyword      = ":"
	DefaultMatchEndKeyword       = DefaultHostEndKeyword
)

// ReplaceFunc is a function type that accepts a string and returns a replaced string.
type ReplaceFunc func(src, key, value string, fold bool) (string, bool)

// Replacer interface defines methods for setting keywords and performing replacements.
type Replacer interface {
	Replace(content []byte, replacements map[string]string) []byte // Replaces substrings based on provided key-value pairs.
	ReplaceString(content string, replacements map[string]string) string
}

// Replace replaces substrings within the provided content byte slice using the provided key-value pairs.
func Replace(content []byte, replacements map[string]string) []byte {
	return _globalReplacer.Replace(content, replacements)
}

// ReplaceString replaces substrings within the provided content string using the provided key-value pairs.
func ReplaceString(content string, replacements map[string]string) string {
	return _globalReplacer.ReplaceString(content, replacements)
}

// Replacement struct implements the Replacer interface, storing Replacement hooks and the current keyword.
type Replacement struct {
	offset int
	sta    string
	end    string
	sep    string
	hook   ReplaceFunc
	fold   bool
}

func (r Replacement) ToMatch(replacements map[string]string) Matcher {
	return NewMatch(replacements, WithMatchSta(r.sta), WithMatchEnd(r.end), WithMatchFold(r.fold))
}

// ReplaceString replaces substrings within the provided content string using the provided key-value pairs.
func (r Replacement) ReplaceString(content string, replacements map[string]string) string {
	return string(r.Replace([]byte(content), replacements))
}

// Replace within the Replacement struct iterates through the values map, applying custom hooks and
// replacing placeholders found in the source string.
func (r Replacement) Replace(content []byte, replacements map[string]string) []byte {
	// Create a buffer to hold the modified content
	var result bytes.Buffer
	contentStr := string(content)

	// Iterate through the content to find ${name} patterns
	cursor := 0
	for {
		// Find the next occurrence of `${`
		sta := strings.Index(contentStr[cursor:], r.sta)
		if sta == -1 {
			// No more occurrences, write the remaining content and break
			result.WriteString(contentStr[cursor:])
			break
		}

		// Write the content before the found pattern
		result.WriteString(contentStr[cursor : cursor+sta])

		// Find the closing `}`
		end := strings.Index(contentStr[cursor+sta:], r.end)
		if end == -1 {
			// No closing brace found, write the remaining content and break
			result.WriteString(contentStr[cursor+sta:])
			break
		}

		// Extract the variable name
		varName := contentStr[cursor+sta+r.offset : cursor+sta+end]
		vars := strings.Split(varName, r.sep)
		srcKey := varName
		if len(vars) > 0 {
			srcKey = vars[0]
		}
		// Check for Replacement in the map (case-insensitive)
		found := false
		for key, value := range replacements {
			newValue, ok := r.hook(srcKey, key, value, r.fold)
			if ok {
				result.WriteString(newValue)
				found = true
				break
			}
		}

		// If no Replacement was found, write the original pattern
		if !found {
			srcValue := varName
			if len(vars) > 1 {
				srcValue = vars[1]
				result.WriteString(srcValue)
			} else {
				result.WriteString(r.sta + srcValue + r.end)
			}

		}

		// Update the cursor position for the next iteration
		cursor = cursor + sta + end + 1
	}

	return result.Bytes()
}

func defaultMatchFunc(src, key string, fold bool) bool {
	return fold && strings.EqualFold(key, src) || key == src
}

func defaultReplacer(src, key, value string, fold bool) (string, bool) {
	switch {
	case fold && value != "" && strings.EqualFold(key, src):
		return value, true
	case value != "" && key == src:
		return value, true
	default:
		return "", false
	}
}

// New returns a new Replacer instance with default settings.
func New(ss ...Setting) *Replacement {
	r := settings.Apply(&Replacement{
		sta:  DefaultStartKeyword,
		end:  DefaultEndKeyword,
		sep:  DefaultSeparatorKeyword,
		hook: defaultReplacer,
	}, ss)

	r.offset = len(r.sta)
	return r
}

// NewHost returns a new Replacer instance with default host settings.
func NewHost(ss ...Setting) *Replacement {
	r := settings.Apply(&Replacement{
		sta:  DefaultHostStartKeyword,
		end:  DefaultEndKeyword,
		sep:  DefaultSeparatorKeyword,
		hook: defaultReplacer,
	}, ss)

	r.offset = len(r.sta)
	return r
}

var _globalReplacer = New(WithFold())
