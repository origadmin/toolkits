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

// Var constructs a string by joining the given string slices with underscores and converting to uppercase.
func Var(vv ...string) string {
	// Join the string slices with underscores and convert to uppercase
	return strings.ToUpper(strings.Join(vv, "_"))
}
