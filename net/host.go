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

func getFromEnv(envVar string) string {
	return os.Getenv(envVar)
}

// InterfaceWithAddrs abstracts a network interface and its addresses for mocking.
type InterfaceWithAddrs interface {
	GetInterface() net.Interface
	GetAddrs() ([]net.Addr, error)
}

// realInterfaceWithAddrs implements InterfaceWithAddrs for real net.Interface.
type realInterfaceWithAddrs struct {
	net.Interface
}

func (r *realInterfaceWithAddrs) GetInterface() net.Interface   { return r.Interface }
func (r *realInterfaceWithAddrs) GetAddrs() ([]net.Addr, error) { return r.Interface.Addrs() }

// getValidIP now takes InterfaceWithAddrs.
func getValidIP(ifaceWithAddrs InterfaceWithAddrs) net.IP {
	iface := ifaceWithAddrs.GetInterface()
	if (iface.Flags & net.FlagUp) == 0 {
		return nil
	}
	addrs, err := ifaceWithAddrs.GetAddrs()
	if err != nil {
		return nil
	}

	var firstIPv4 net.IP
	var firstIPv6 net.IP

	for _, rawAddr := range addrs {
		var ip net.IP
		switch addr := rawAddr.(type) {
		case *net.IPAddr:
			ip = addr.IP
		case *net.IPNet:
			ip = addr.IP
		default:
			continue
		}

		if IsUsableHostIP(ip.String()) { // Changed to IsUsableHostIP
			if ip.To4() != nil {
				if firstIPv4 == nil { // Store the first IPv4 found
					firstIPv4 = ip
				}
			} else {
				if firstIPv6 == nil { // Store the first IPv6 found
					firstIPv6 = ip
				}
			}
		}
	}

	if firstIPv4 != nil {
		return firstIPv4
	}
	return firstIPv6 // Return the first IPv6 if no IPv4 was found
}

// NetworkInterfaceProvider now returns InterfaceWithAddrs.
type NetworkInterfaceProvider interface {
	Interfaces() ([]InterfaceWithAddrs, error)
}

// realNetworkInterfaceProvider is the default implementation using net.Interfaces().
type realNetworkInterfaceProvider struct{}

func (r *realNetworkInterfaceProvider) Interfaces() ([]InterfaceWithAddrs, error) {
	realIfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	ifacesWithAddrs := make([]InterfaceWithAddrs, len(realIfaces))
	for i := range realIfaces {
		ifacesWithAddrs[i] = &realInterfaceWithAddrs{Interface: realIfaces[i]}
	}
	return ifacesWithAddrs, nil
}

// getByCIDR uses the provided NetworkInterfaceProvider.
func getByCIDR(provider NetworkInterfaceProvider, cidrFilters []*net.IPNet) string {
	interfaces, _ := provider.Interfaces()
	for _, ifaceWithAddrs := range interfaces {
		if ip := getValidIP(ifaceWithAddrs); ip != nil {
			for _, filter := range cidrFilters {
				if filter.Contains(ip) {
					return ip.String()
				}
			}
		}
	}
	return ""
}

// getByInterfacePattern uses the provided NetworkInterfaceProvider.
func getByInterfacePattern(provider NetworkInterfaceProvider, patterns []string) string {
	interfaces, _ := provider.Interfaces()
	for _, ifaceWithAddrs := range interfaces {
		iface := ifaceWithAddrs.GetInterface()
		// Match the name of the interface（like: eth*, eno*, wlan*, etc）
		for _, pattern := range patterns {
			if matched, _ := filepath.Match(pattern, iface.Name); matched {
				if ip := getValidIP(ifaceWithAddrs); ip != nil {
					//return ip
				}
			}
		}
	}
	return ""
}

// getFirstAvailableIP uses the provided NetworkInterfaceProvider.
func getFirstAvailableIP(provider NetworkInterfaceProvider) net.IP {
	interfaces, _ := provider.Interfaces()
	for _, ifaceWithAddrs := range interfaces {
		if ip := getValidIP(ifaceWithAddrs); ip != nil {
			return ip
		}
	}
	return nil
}

// HostAddr now accepts a NetworkInterfaceProvider.
func HostAddr(provider NetworkInterfaceProvider, opts ...Option) string {
	cfg := defaultConfig
	configure.Apply(&cfg, opts)

	if cfg.envVar != "" {
		if ip := getFromEnv(cfg.envVar); ip != "" {
			return ip
		}
	}

	if len(cfg.cidrFilters) > 0 {
		if ip := getByCIDR(provider, cfg.cidrFilters); ip != "" {
			return ip
		}
	}

	if len(cfg.patterns) > 0 {
		if ip := getByInterfacePattern(provider, cfg.patterns); ip != "" {
			return ip
		}
	}

	if cfg.fallbackToFirst {
		if ip := getFirstAvailableIP(provider); ip != nil {
			return ip.String()
		}
	}
	return ""
}

// RealHostAddr is a convenience function for external callers to use the real network provider.
func RealHostAddr(opts ...Option) string {
	return HostAddr(&realNetworkInterfaceProvider{}, opts...)
}
