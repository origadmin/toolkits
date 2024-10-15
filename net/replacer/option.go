// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package replacer provides a string replacement mechanism with custom replace functions and keywords.
package replacer

// Setting is the setting of replacer.
type Setting = func(*replace)

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
