// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package replacer provides a string replacement mechanism with custom replace functions and keywords.
package replacer

import (
	"github.com/goexts/ggb/settings"
)

// Setting is the setting of replacer.
type Setting = settings.Setting[replace]

// WithHooks returns a new Replacer instance with the specified hooks.
func WithHooks(hooks map[string]ReplaceFunc) Setting {
	return func(o *replace) {
		o.hooks = hooks
	}
}

// WithKeyword returns a new Replacer instance with the specified keyword.
func WithKeyword(keyword string) Setting {
	return func(o *replace) {
		o.keyword = keyword
	}
}
