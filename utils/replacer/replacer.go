// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package replacer provides a string replacement mechanism with custom replace functions and keywords.
package replacer

import (
	"bytes"
	"strings"

	"github.com/goexts/ggb/settings"
)

const (
	// DefaultStartKeyword defines the default keyword.
	DefaultStartKeyword      = "${"
	DefaultHostStartKeyword  = "@{"
	DefaultMatchStartKeyword = "@"

	DefaultEndKeyword      = "}"
	DefaultHostEndKeyword  = ":"
	DefaultMatchEndKeyword = DefaultHostEndKeyword
)

// ReplaceFunc is a function type that accepts a string and returns a replaced string.
type ReplaceFunc func(src, key, value string, fold bool) (string, bool)

// Replacer interface defines methods for setting keywords and performing replacements.
type Replacer interface {
	Matcher(replacements map[string]string) Matcher
	Replace(content []byte, replacements map[string]string) []byte // Replaces substrings based on provided key-value pairs.
	ReplaceString(content string, replacements map[string]string) string
}

// replace struct implements the Replacer interface, storing replace hooks and the current keyword.
type replace struct {
	offset int
	sta    string
	end    string
	hook   ReplaceFunc
	fold   bool
}

func (r replace) Matcher(replacements map[string]string) Matcher {
	return &matcher{
		sta:         r.sta,
		end:         r.end,
		fold:        r.fold,
		replacement: replacements,
	}
}

// ReplaceString replaces substrings within the provided content string using the provided key-value pairs.
func (r replace) ReplaceString(content string, replacements map[string]string) string {
	return string(r.Replace([]byte(content), replacements))
}

// Replace within the replacement struct iterates through the values map, applying custom hooks and
// replacing placeholders found in the source string.
func (r replace) Replace(content []byte, replacements map[string]string) []byte {
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

		// Check for replacement in the map (case-insensitive)
		found := false
		for key, value := range replacements {
			newValue, ok := r.hook(varName, key, value, r.fold)
			if ok {
				result.WriteString(newValue)
				found = true
				break
			}
		}

		// If no replacement was found, write the original pattern
		if !found {
			result.WriteString(r.sta + varName + r.end)
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
func New(ss ...Setting) Replacer {
	r := settings.Apply(&replace{
		sta:  DefaultStartKeyword,
		end:  DefaultEndKeyword,
		hook: defaultReplacer,
	}, ss)

	r.offset = len(r.sta)
	return r
}

func NewHost(ss ...Setting) Replacer {
	r := settings.Apply(&replace{
		sta:  DefaultHostStartKeyword,
		end:  DefaultEndKeyword,
		hook: defaultReplacer,
	}, ss)

	r.offset = len(r.sta)
	return r

}

var _globalReplacer = New(WithFold())
