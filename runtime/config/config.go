/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package config implements the functions, types, and interfaces for the module.
package config

import (
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
)

const Type = "config"

type Setting struct {
	Options   []Option
	Customize *configv1.Customize
}

type SettingFunc = func(s *Setting)

// WithOptions sets the options field of the Setting struct.
func WithOptions(options ...Option) SettingFunc {
	return func(s *Setting) {
		s.Options = options
	}
}

// WithCustomize sets the customize field of the Setting struct.
func WithCustomize(custom *configv1.Customize) SettingFunc {
	return func(s *Setting) {
		s.Customize = custom
	}
}
