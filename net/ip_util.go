/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package net

import (
	"errors"
	"net"
)

var ErrNoIPFound = errors.New("no ip found")

// IPSelectionStrategy defines how to select an IP from multiple available IPs
type IPSelectionStrategy func([]net.IP) (net.IP, error)

// IsPublicRoutableIP checks if the given IP address is a global unicast address and not a private or multicast address.
func IsPublicRoutableIP(addr string) bool {
	ip := net.ParseIP(addr)
	if ip == nil {
		return false
	}
	// A valid IP for our purpose should be global unicast AND NOT private AND NOT interface local multicast
	return ip.IsGlobalUnicast() && !isPrivateIP(ip) && !ip.IsInterfaceLocalMulticast()
}

// IsUsableHostIP checks if the given IP address is a usable host IP (not loopback, link-local, or unspecified).
func IsUsableHostIP(addr string) bool {
	ip := net.ParseIP(addr)
	if ip == nil {
		return false
	}
	// A usable host IP should not be loopback, link-local, or unspecified.
	return !ip.IsLoopback() && !ip.IsLinkLocalUnicast() && !ip.IsLinkLocalMulticast() && !ip.IsUnspecified()
}

// isPrivateIP checks if an IP is in a private network
func isPrivateIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}

	// Check for private address ranges
	ip4 := ip.To4()
	if ip4 == nil {
		// Not an IPv4 address
		return false
	}

	// 10.0.0.0/8
	if ip4[0] == 10 {
		return true
	}
	// 172.16.0.0/12
	if ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31 {
		return true
	}
	// 192.168.0.0/16
	if ip4[0] == 192 && ip4[1] == 168 {
		return true
	}

	return false
}
