/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package net

import (
	"testing"
)

func TestIsPublicRoutableIP(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want bool
	}{
		{
			name: "Valid Global Unicast IPv4",
			ip:   "8.8.8.8",
			want: true,
		},
		{
			name: "Valid Global Unicast IPv6",
			ip:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			want: true,
		},
		{
			name: "Loopback IPv4",
			ip:   "127.0.0.1",
			want: false,
		},
		{
			name: "Loopback IPv6",
			ip:   "::1",
			want: false,
		},
		{
			name: "Private IPv4 (Class A)",
			ip:   "10.0.0.1",
			want: false,
		},
		{
			name: "Private IPv4 (Class B)",
			ip:   "172.16.0.1",
			want: false,
		},
		{
			name: "Private IPv4 (Class C)",
			ip:   "192.168.1.1",
			want: false,
		},
		{
			name: "Link-Local IPv6",
			ip:   "fe80::1",
			want: false,
		},
		{
			name: "Multicast IPv4",
			ip:   "224.0.0.1",
			want: false, // Changed from true to false
		},
		{
			name: "Unspecified IPv4",
			ip:   "0.0.0.0",
			want: false,
		},
		{
			name: "Unspecified IPv6",
			ip:   "::",
			want: false,
		},
		{
			name: "Invalid IP String",
			ip:   "invalid-ip",
			want: false,
		},
		{
			name: "Empty IP String",
			ip:   "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPublicRoutableIP(tt.ip); got != tt.want {
				t.Errorf("IsPublicRoutableIP(%q) = %v, want %v", tt.ip, got, tt.want)
			}
		})
	}
}

func TestIsUsableHostIP(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want bool
	}{
		{
			name: "Valid Global Unicast IPv4",
			ip:   "8.8.8.8",
			want: true,
		},
		{
			name: "Valid Global Unicast IPv6",
			ip:   "2001:0db8:85a3:0000:0000:8a2e:0370:7334",
			want: true,
		},
		{
			name: "Loopback IPv4",
			ip:   "127.0.0.1",
			want: false,
		},
		{
			name: "Loopback IPv6",
			ip:   "::1",
			want: false,
		},
		{
			name: "Private IPv4 (Class A)",
			ip:   "10.0.0.1",
			want: true,
		},
		{
			name: "Private IPv4 (Class B)",
			ip:   "172.16.0.1",
			want: true,
		},
		{
			name: "Private IPv4 (Class C)",
			ip:   "192.168.1.1",
			want: true,
		},
		{
			name: "Link-Local IPv6",
			ip:   "fe80::1",
			want: false,
		},
		{
			name: "Multicast IPv4",
			ip:   "224.0.0.1",
			want: false, // Changed from true to false
		},
		{
			name: "Unspecified IPv4",
			ip:   "0.0.0.0",
			want: false,
		},
		{
			name: "Unspecified IPv6",
			ip:   "::",
			want: false,
		},
		{
			name: "Invalid IP String",
			ip:   "invalid-ip",
			want: false,
		},
		{
			name: "Empty IP String",
			ip:   "",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsUsableHostIP(tt.ip); got != tt.want {
				t.Errorf("IsUsableHostIP(%q) = %v, want %v", tt.ip, got, tt.want)
			}
		})
	}
}
