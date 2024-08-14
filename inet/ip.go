package inet

import (
	"math"
	"net"
)

func ParseIP(ip net.IP) int64 {
	var (
		one   byte
		intip int64
	)
	if v4 := ip.To4(); len(v4) == net.IPv4len {
		return int64(v4[0])<<24 | int64(v4[1])<<16 | int64(v4[2])<<8 | int64(v4[3])
	}
	for _, one = range ip {
		intip = (intip << 8) | int64(one)
	}
	return intip
}

func ToIP(ip int64) net.IP {
	if ip > math.MaxUint32 {
		var ips []byte
		for i := 0; i < 8; i++ {
			ips = append(ips, byte(ip&0xff))
			ip = ip >> 8
		}
		return ips
	}
	return net.IPv4(byte(ip>>24&0xff), byte(ip>>16&0xff), byte(ip>>8&0xff), byte(ip&0xff))
}
