/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package net

import (
	"reflect"
	"testing"
)

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	expectedPatterns := []string{"eth*", "eno*", "wlan*"}
	if !reflect.DeepEqual(cfg.patterns, expectedPatterns) {
		t.Errorf("DefaultConfig() patterns got %v, want %v", cfg.patterns, expectedPatterns)
	}

	if !cfg.fallbackToFirst {
		t.Errorf("DefaultConfig() fallbackToFirst got %v, want %v", cfg.fallbackToFirst, true)
	}

	if cfg.envVar != "" {
		t.Errorf("DefaultConfig() envVar got %q, want empty string", cfg.envVar)
	}

	if len(cfg.cidrFilters) != 0 {
		t.Errorf("DefaultConfig() cidrFilters got %v, want empty slice", cfg.cidrFilters)
	}
}

func TestWithEnvVar(t *testing.T) {
	cfg := HostConfig{}
	WithEnvVar("TEST_IP")(&cfg)

	if cfg.envVar != "TEST_IP" {
		t.Errorf("WithEnvVar() got %q, want %q", cfg.envVar, "TEST_IP")
	}
}

func TestWithFallback(t *testing.T) {
	cfg := HostConfig{fallbackToFirst: false}
	WithFallback(true)(&cfg)

	if !cfg.fallbackToFirst {
		t.Errorf("WithFallback(true) got %v, want %v", cfg.fallbackToFirst, true)
	}

	WithFallback(false)(&cfg)
	if cfg.fallbackToFirst {
		t.Errorf("WithFallback(false) got %v, want %v", cfg.fallbackToFirst, false)
	}
}

func TestWithCIDRFilters(t *testing.T) {
	cfg := HostConfig{}
	cidrs := []string{"192.168.1.0/24", "10.0.0.0/8"}
	WithCIDRFilters(cidrs)(&cfg)

	if len(cfg.cidrFilters) != 2 {
		t.Fatalf("WithCIDRFilters() len got %d, want %d", len(cfg.cidrFilters), 2)
	}

	if cfg.cidrFilters[0].String() != "192.168.1.0/24" || cfg.cidrFilters[1].String() != "10.0.0.0/8" {
		t.Errorf("WithCIDRFilters() got %v, want %v", cfg.cidrFilters, cidrs)
	}

	// Test with invalid CIDR
	cfg = HostConfig{}
	WithCIDRFilters([]string{"invalid-cidr"})(&cfg)
	if len(cfg.cidrFilters) != 0 {
		t.Errorf("WithCIDRFilters() with invalid CIDR got %v, want empty slice", cfg.cidrFilters)
	}
}

func TestWithInterfacePatterns(t *testing.T) {
	cfg := HostConfig{}
	patterns := []string{"eth0", "enp*"}
	WithInterfacePatterns(patterns)(&cfg)

	if !reflect.DeepEqual(cfg.patterns, patterns) {
		t.Errorf("WithInterfacePatterns() got %v, want %v", cfg.patterns, patterns)
	}
}
