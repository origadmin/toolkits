/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package interfaces implements the functions, types, and interfaces for the module.
package interfaces

// Params A common interface for algorithm parameters is defined
type Params interface {
	ConfigValidator
	// String returns the string representation of parameters
	String() string

	// ToMap converts parameters to map[string]string format for encoding.
	ToMap() map[string]string

	// FromMap parses parameters from map[string]string format.
	FromMap(params map[string]string) error
}
