// Copyright (c) 2024 OrigAdmin. All rights reserved.

// Package utils implements the functions, types, and interfaces for the module.
package utils

import (
	"net/netip"
	"strconv"
	"strings"
)

const (
	defaultDiscoveryPrefix = "discovery:///"
)

func NameDiscovery(serviceName string) string {
	return defaultDiscoveryPrefix + serviceName
}

func DiscoveryEndpoint(endpoint, scheme, host, addr string) string {
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
