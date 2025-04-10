/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package env implements the functions, types, and interfaces for the module.
package env

import (
	"os"
	"strings"
)

// SetEnv sets an environment variable with the given key and value.
// The key is converted to uppercase before setting the environment variable.
func SetEnv(key, value string) error {
	// Convert the key to uppercase to ensure consistency in environment variable names
	return os.Setenv(strings.ToUpper(key), value)
}

// GetEnv retrieves the value of an environment variable with the given key.
func GetEnv(key string) string {
	// Return the value of the environment variable, or an empty string if not set
	return os.Getenv(key)
}

// LookupEnv retrieves the value of an environment variable with the given key.
func LookupEnv(key string) (string, bool) {
	return os.LookupEnv(key)
}

// Var constructs a string by joining the given string slices with underscores and converting to uppercase.
func Var(vv ...string) string {
	// Join the string slices with underscores and convert to uppercase
	return strings.ToUpper(strings.Join(vv, "_"))
}

func prefixedVar(prefix string, vv ...string) string {
	return Var(append([]string{prefix}, vv...)...)
}

// Env is an interface that provides methods for working with environment variables.
// It includes methods for looking up, setting, and getting environment variables.
type Env interface {
	// Var constructs a string by joining the given string slices with underscores and converting to uppercase.
	Var(...string) string
	// LookupEnv retrieves the value of an environment variable with the given key.
	LookupEnv(string) (string, bool)
	// SetEnv sets an environment variable with the given key and value.
	SetEnv(string, string) error
	// GetEnv retrieves the value of an environment variable with the given key.
	GetEnv(string) string
}

type env struct {
	prefix string
}

func (e env) LookupEnv(s string) (string, bool) {
	return os.LookupEnv(e.Var(s))
}

func (e env) Var(s ...string) string {
	return prefixedVar(e.prefix, s...)
}

func (e env) SetEnv(s string, v string) error {
	return SetEnv(e.Var(s), v)
}

func (e env) GetEnv(s string) string {
	return GetEnv(e.Var(s))
}

// WithPrefix creates a new Env instance with the given prefix.
func WithPrefix(prefix string) Env {
	if prefix == "" {
		panic("prefix cannot be empty")
	}
	return &env{
		prefix: prefix,
	}
}
