// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package config is the config package for origen
package config

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"

	"github.com/origadmin/toolkits/codec"
)

// Config is the configuration of admin-cli.
type Config struct {
	Organization string
	Application  string
	Project      string
	Version      string
	Type         Type
	Mods         []Mod
	Platform     []Platform // TODO
	Web          Web
	Static       string
	Plugins      []Plugin
	Resources    Resources
	Document     Document
}

// Save saves config to file
func Save(path string, v any) error {
	dir, _ := filepath.Split(path)
	// check whether the directory exists
	stat, err := os.Stat(dir)
	if err != nil {
		// if the directory does not exist create it
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed to create directory: %w", err)
		}
	}
	if err == nil && !stat.IsDir() {
		// if the path exists but is not a directory, an error is returned
		return fmt.Errorf("%s is not a directory", dir)
	}

	return codec.EncodeFile(path, v)
}

// Load loads config from file
func Load(path string, v any) error {
	stat, err := os.Stat(path)
	if err != nil {
		return err
	}
	if stat.IsDir() {
		return fmt.Errorf("%s is a directory", path)
	}
	return codec.DecodeFile(path, v)
}

var (
	config Config
	once   sync.Once
)

// C is a singleton of Config
func C() Config {
	return config
}

// LoadGlobal loads config from file
func LoadGlobal(path string) {
	once.Do(func() {
		err := Load(path, &config)
		if err != nil {
			panic(err)
		}
	})
}
