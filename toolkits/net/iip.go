/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package net

import (
	"math"
)

func ParseIP(ip IP) int64 {
	var (
		one   byte
		intip int64
	)
	if v4 := ip.To4(); len(v4) == IPv4len {
		ip = v4
	}
	for _, one = range ip {
		intip = (intip << 8) | int64(one)
	}
	return intip
}

func ToIP(ip int64) IP {
	if ip > math.MaxUint32 {
		var ips []byte
		for i := 0; i < 8; i++ {
			ips = append(ips, byte(ip&0xff))
			ip = ip >> 8
		}
		return ips
	}
	return IPv4(byte(ip>>24&0xff), byte(ip>>16&0xff), byte(ip>>8&0xff), byte(ip&0xff))
}
