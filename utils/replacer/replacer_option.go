// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package replacer provides a string replacement mechanism with custom replace functions and keywords.
package replacer

// Setting is the setting of replacer.
type Setting = func(*replace)

// WithHook returns a new Replacer instance with the specified hooks.
func WithHook(hook ReplaceFunc) Setting {
	return func(o *replace) {
		o.hook = hook
	}
}

// WithStart returns a new Replacer instance with replacer start keyword.
func WithStart(keyword string) Setting {
	return func(o *replace) {
		o.sta = keyword
	}
}

// WithEnd returns a new Replacer instance with replacer end keyword.
func WithEnd(keyword string) Setting {
	return func(o *replace) {
		o.end = keyword
	}
}

// WithKeyword returns a new Replacer instance with the specified keyword.
func WithKeyword(keyword string) Setting {
	return func(o *replace) {
		o.sta = keyword
		o.end = keyword
	}
}

// WithFold returns a new Replacer instance with string case folding.
func WithFold() Setting {
	return func(o *replace) {
		o.fold = true
	}
}
