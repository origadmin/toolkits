/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package net implements the functions, types, and interfaces for the module.
package net

import (
	"net"
	"os"
	"path/filepath"

	"github.com/goexts/generic/configure"
)

// HostConfig defines the configuration for the host.
type HostConfig struct {
	envVar          string       // The name of the environment variable
	patterns        []string     // Interface name patterns
	cidrFilters     []*net.IPNet // Segment filters
	fallbackToFirst bool         // Whether to roll back to the first NIC
}

func WithEnvVar(name string) Option {
	return func(c *HostConfig) {
		c.envVar = name
	}
}

func WithFallback(enable bool) Option {
	return func(c *HostConfig) {
		c.fallbackToFirst = enable
	}
}

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

func WithInterfacePatterns(patterns []string) Option {
	return func(c *HostConfig) {
		c.patterns = patterns
	}
}

type Option = func(*HostConfig)

var defaultConfig = DefaultConfig()

func DefaultConfig() HostConfig {
	return HostConfig{
		patterns:        []string{"eth*", "eno*", "wlan*"},
		fallbackToFirst: true,
	}
}

func getFromEnv(envVar string) string {
	return os.Getenv(envVar)
}

func isValidIP(addr string) bool {
	ip := net.ParseIP(addr)
	return ip.IsGlobalUnicast() && !ip.IsInterfaceLocalMulticast()
}

func getValidIP(iface net.Interface) net.IP {
	addrs, _ := iface.Addrs()
	minIndex := int(^uint(0) >> 1)
	ips := make([]net.IP, 0)
	if (iface.Flags & net.FlagUp) == 0 {
		return nil
	}
	if iface.Index >= minIndex && len(ips) != 0 {
		return nil
	}
	addrs, err := iface.Addrs()
	if err != nil {
		return nil
	}
	for i, rawAddr := range addrs {
		var ip net.IP
		switch addr := rawAddr.(type) {
		case *net.IPAddr:
			ip = addr.IP
		case *net.IPNet:
			ip = addr.IP
		default:
			continue
		}
		if isValidIP(ip.String()) {
			minIndex = iface.Index
			if i == 0 {
				ips = make([]net.IP, 0, 1)
			}
			ips = append(ips, ip)
			if ip.To4() != nil {
				break
			}
		}
	}
	if len(ips) != 0 {
		return ips[len(ips)-1]
	}
	return nil
}

func getByCIDR(cidrFilters []*net.IPNet) string {
	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		if ip := getValidIP(iface); ip != nil {
			for _, filter := range cidrFilters {
				if filter.Contains(ip) {
					return ip.String()
				}
			}
		}
	}
	return ""
}

func getByInterfacePattern(patterns []string) string {
	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		// Match the name of the interface（like: eth*, eno*, wlan*, etc）
		for _, pattern := range patterns {
			if matched, _ := filepath.Match(pattern, iface.Name); matched {
				if ip := getValidIP(iface); ip != nil {
					//return ip
				}
			}
		}
	}
	return ""
}

func getFirstAvailableIP() net.IP {
	interfaces, _ := net.Interfaces()
	for _, iface := range interfaces {
		if ip := getValidIP(iface); ip != nil {
			return ip
		}
	}
	return nil
}

func HostAddr(opts ...Option) string {
	cfg := defaultConfig
	configure.Apply(&cfg, opts)

	if cfg.envVar != "" {
		if ip := getFromEnv(cfg.envVar); ip != "" {
			return ip
		}
	}

	if len(cfg.cidrFilters) > 0 {
		if ip := getByCIDR(cfg.cidrFilters); ip != "" {
			return ip
		}
	}

	if len(cfg.patterns) > 0 {
		if ip := getByInterfacePattern(cfg.patterns); ip != "" {
			return ip
		}
	}

	if cfg.fallbackToFirst {
		if ip := getFirstAvailableIP(); ip != nil {
			return ip.String()
		}
	}
	return ""
}
