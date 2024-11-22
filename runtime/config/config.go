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
	Options []Option
	Custom  *configv1.Custom
}

type SettingFunc = func(s *Setting)

func WithOptions(options ...Option) SettingFunc {
	return func(s *Setting) {
		s.Options = options
	}
}

func WithCustom(custom *configv1.Custom) SettingFunc {
	return func(s *Setting) {
		s.Custom = custom
	}
}
