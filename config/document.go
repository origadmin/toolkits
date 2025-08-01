// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package config implements the functions, types, and interfaces for the module.
package config

type Contact struct {
	Name string `json:"name" yaml:"name"  toml:"name"`
	URL  string `json:"URL" yaml:"URL" toml:"url"`
	Mail string `json:"mail" yaml:"mail" toml:"mail"`
}

type License struct {
	Name string `json:"name" yaml:"name" toml:"name"`
	URL  string `json:"URL" yaml:"URL" toml:"url"`
}

type Document struct {
	Title       string  `json:"title" yaml:"title" toml:"title"`
	Description string  `json:"description" yaml:"description" toml:"description"`
	Contact     Contact `json:"contact" yaml:"contact" toml:"contact"`
	License     License `json:"license" yaml:"license" toml:"license"`
}
