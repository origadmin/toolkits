/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package identifier provides a unified interface for generating and validating unique identifiers.
package identifier

import (
	"sync"
)

// Registry manages the registration and retrieval of string and number identifier generators.
type Registry struct {
	sync.RWMutex
	stringGenerators map[string]StringIdentifier
	numberGenerators map[string]NumberIdentifier
	defaultStringGen StringIdentifier
	defaultNumberGen NumberIdentifier
}

// NewRegistry creates and returns a new Registry instance.
func NewRegistry() *Registry {
	return &Registry{
		stringGenerators: make(map[string]StringIdentifier),
		numberGenerators: make(map[string]NumberIdentifier),
	}
}

// RegisterString registers a string identifier generator.
func (r *Registry) RegisterString(gen StringIdentifier) {
	r.Lock()
	defer r.Unlock()
	r.stringGenerators[gen.Name()] = gen
}

// RegisterNumber registers a number identifier generator.
func (r *Registry) RegisterNumber(gen NumberIdentifier) {
	r.Lock()
	defer r.Unlock()
	r.numberGenerators[gen.Name()] = gen
}

// SetDefaultString sets the default string identifier generator.
func (r *Registry) SetDefaultString(gen StringIdentifier) {
	r.Lock()
	defer r.Unlock()
	r.defaultStringGen = gen
}

// SetDefaultNumber sets the default number identifier generator.
func (r *Registry) SetDefaultNumber(gen NumberIdentifier) {
	r.Lock()
	defer r.Unlock()
	r.defaultNumberGen = gen
}

// GetString retrieves a string identifier generator by name.
func (r *Registry) GetString(name string) StringIdentifier {
	r.RLock()
	defer r.RUnlock()
	return r.stringGenerators[name]
}

// GetNumber retrieves a number identifier generator by name.
func (r *Registry) GetNumber(name string) NumberIdentifier {
	r.RLock()
	defer r.RUnlock()
	return r.numberGenerators[name]
}

// DefaultString returns the default string identifier generator.
func (r *Registry) DefaultString() StringIdentifier {
	r.RLock()
	defer r.RUnlock()
	return r.defaultStringGen
}

// DefaultNumber returns the default number identifier generator.
func (r *Registry) DefaultNumber() NumberIdentifier {
	r.RLock()
	defer r.RUnlock()
	return r.defaultNumberGen
}

// RegisterStringIdentifier registers a string identifier generator.
func RegisterStringIdentifier(gen StringIdentifier) {
	registry.RegisterString(gen)
}

// RegisterNumberIdentifier registers a number identifier generator.
func RegisterNumberIdentifier(gen NumberIdentifier) {
	registry.RegisterNumber(gen)
}

// RegisterMultiTypeIdentifier registers a multi-type identifier generator.
func RegisterMultiTypeIdentifier(gen MultiTypeIdentifier) {
	RegisterStringIdentifier(gen)
	RegisterNumberIdentifier(gen)
}

func SetDefaultString(identifier StringIdentifier) {
	registry.SetDefaultString(identifier)
}

func SetDefaultNumber(identifier NumberIdentifier) {
	registry.SetDefaultNumber(identifier)
}

func DefaultString() StringIdentifier {
	return registry.DefaultString()
}

func DefaultNumber() NumberIdentifier {
	return registry.DefaultNumber()
}
