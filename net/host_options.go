/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package net

import "net"

// HostConfig defines the configuration for the host.
type HostConfig struct {
	envVar          string       // The name of the environment variable
	patterns        []string     // Interface name patterns
	cidrFilters     []*net.IPNet // Segment filters
	fallbackToFirst bool         // Whether to roll back to the first NIC
}

// Option is a function that configures a HostConfig.
type Option func(*HostConfig)

// WithEnvVar sets the environment variable name to look for an IP address.
func WithEnvVar(name string) Option {
	return func(c *HostConfig) {
		c.envVar = name
	}
}

// WithFallback enables or disables falling back to the first available IP if no other method succeeds.
func WithFallback(enable bool) Option {
	return func(c *HostConfig) {
		c.fallbackToFirst = enable
	}
}

// WithCIDRFilters sets CIDR ranges to filter interfaces by.
func WithCIDRFilters(cidrs []string) Option {
	return func(c *HostConfig) {
		for _, cidrStr := range cidrs {
			_, ipNet, err := net.ParseCIDR(cidrStr)
			if err == nil {
				c.cidrFilters = append(c.cidrFilters, ipNet)
			}
		}
	}
}

// WithInterfacePatterns sets patterns to match interface names against.
func WithInterfacePatterns(patterns []string) Option {
	return func(c *HostConfig) {
		c.patterns = patterns
	}
}

// defaultConfig is the default configuration for host address resolution.
var defaultConfig = DefaultConfig()

// DefaultConfig returns the default host configuration.
func DefaultConfig() HostConfig {
	return HostConfig{
		patterns:        []string{"eth*", "eno*", "wlan*"},
		fallbackToFirst: true,
	}
}
