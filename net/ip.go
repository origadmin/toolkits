package net

import (
	"net"
)

const (
	IPv4len = net.IPv4len
	IPv6len = net.IPv6len
)

type (
	IP = net.IP
)

func IPv4(a, b, c, d byte) IP {
	return net.IPv4(a, b, c, d)
}
