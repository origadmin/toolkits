/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package config implements the functions, types, and interfaces for the module.
package config

import (
	configv1 "github.com/origadmin/toolkits/runtime/gen/go/config/v1"
)

const Type = "config"

type SourceOption struct {
	Options   []Option
	Customize *configv1.Customize
}

type SourceSetting = func(s *SourceOption)

// WithOptions sets the options field of the SourceOption struct.
func WithOptions(options ...Option) SourceSetting {
	return func(s *SourceOption) {
		s.Options = options
	}
}

// WithCustomize sets the customize field of the SourceOption struct.
func WithCustomize(customize *configv1.Customize) SourceSetting {
	return func(s *SourceOption) {
		s.Customize = customize
	}
}
