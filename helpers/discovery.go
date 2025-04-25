/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package helpers implements the functions, types, and interfaces for the module.
package helpers

import (
	"net"
	"net/netip"
	"strconv"
	"strings"

	"github.com/origadmin/toolkits/errors"
)

const (
	defaultDiscoveryPrefix = "discovery:///"
)

// ServiceDiscovery ...
func ServiceDiscovery(serviceName string) string {
	return defaultDiscoveryPrefix + serviceName
}

// ServiceName ...
// Deprecated: use ServiceDiscovery
func ServiceName(serviceName string) string {
	return defaultDiscoveryPrefix + serviceName
}

// ServiceDiscoveryEndpoint ...
// Deprecated: use ServiceEndpoint
func ServiceDiscoveryEndpoint(endpoint, scheme, host, addr string) string {
	naip, _ := netip.ParseAddrPort(addr)
	if endpoint == "" {
		endpoint = scheme + "://" + host + ":" + strconv.Itoa(int(naip.Port()))
	} else {
		prefix, suffix, ok := strings.Cut(endpoint, "://")
		if !ok {
			endpoint = scheme + "://" + prefix
		} else {
			args := strings.SplitN(suffix, ":", 2)
			if len(args) == 2 {
				args[1] = strconv.Itoa(int(naip.Port()))
			} else if len(args) == 1 {
				args = append(args, strconv.Itoa(int(naip.Port())))
			} else {
				// unknown
			}
			endpoint = prefix + "://" + strings.Join(args, ":")
		}
	}
	return endpoint
}

// ServiceEndpoint get the local ip address and port
// Deprecated: move to runtime.service
func ServiceEndpoint(scheme, host, hostPort string) (string, error) {
	_, port, err := net.SplitHostPort(hostPort)
	if err != nil && host == "" {
		return "", errors.Wrap(err, "invalid host")
	}
	if len(host) > 0 && (host != "0.0.0.0" && host != "[::]" && host != "::") {
		return schemeHost(scheme, net.JoinHostPort(host, port)), nil
	}
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", errors.Wrap(err, "failed to get local ip")
	}
	minIndex := int(^uint(0) >> 1)
	ips := make([]net.IP, 0)
	for _, iface := range ifaces {
		if (iface.Flags & net.FlagUp) == 0 {
			continue
		}
		if iface.Index >= minIndex && len(ips) != 0 {
			continue
		}
		addrs, err := iface.Addrs()
		if err != nil {
			continue
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
	}
	if len(ips) != 0 {
		return schemeHost(scheme, net.JoinHostPort(ips[len(ips)-1].String(), port)), nil
	}
	return "", errors.New("no local ip found")
}

func schemeHost(scheme, host string) string {
	return scheme + "://" + host
}

func isValidIP(addr string) bool {
	ip := net.ParseIP(addr)
	return ip.IsGlobalUnicast() && !ip.IsInterfaceLocalMulticast()
}
