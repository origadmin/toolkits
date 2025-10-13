/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package net

import (
	"net"
	"reflect"
	"testing"
)

func TestDefaultIPStrategy(t *testing.T) {
	tests := []struct {
		name     string
		ips      []net.IP
		wantIP   net.IP
		wantErr  error
	}{
		{
			name:    "Prefer IPv4 over IPv6",
			ips:     []net.IP{net.ParseIP("2001:db8::1"), net.ParseIP("192.168.1.1")},
			wantIP:  net.ParseIP("192.168.1.1"),
			wantErr: nil,
		},
		{
			name:    "Only IPv6",
			ips:     []net.IP{net.ParseIP("2001:db8::1"), net.ParseIP("2001:db8::2")},
			wantIP:  net.ParseIP("2001:db8::1"), // Sorted by string representation
			wantErr: nil,
		},
		{
			name:    "Only IPv4",
			ips:     []net.IP{net.ParseIP("10.0.0.1"), net.ParseIP("192.168.1.1")},
			wantIP:  net.ParseIP("10.0.0.1"), // Sorted by string representation
			wantErr: nil,
		},
		{
			name:    "Empty list",
			ips:     []net.IP{},
			wantIP:  nil,
			wantErr: ErrNoIPFound,
		},
		{
			name:    "Mixed IPv4 and IPv6, IPv4 first in list",
			ips:     []net.IP{net.ParseIP("192.168.1.1"), net.ParseIP("2001:db8::1")},
			wantIP:  net.ParseIP("192.168.1.1"),
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIP, gotErr := defaultIPStrategy(tt.ips)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("defaultIPStrategy() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIP, tt.wantIP) {
				t.Errorf("defaultIPStrategy() gotIP = %v, wantIP %v", gotIP, tt.wantIP)
			}
		})
	}
}

func TestPreferIPv4Strategy(t *testing.T) {
	tests := []struct {
		name     string
		ips      []net.IP
		wantIP   net.IP
		wantErr  error
	}{
		{
			name:    "Prefer IPv4",
			ips:     []net.IP{net.ParseIP("2001:db8::1"), net.ParseIP("192.168.1.1"), net.ParseIP("10.0.0.1")},
			wantIP:  net.ParseIP("192.168.1.1"),
			wantErr: nil,
		},
		{
			name:    "Only IPv6",
			ips:     []net.IP{net.ParseIP("2001:db8::1"), net.ParseIP("2001:db8::2")},
			wantIP:  net.ParseIP("2001:db8::1"),
			wantErr: nil,
		},
		{
			name:    "Only IPv4",
			ips:     []net.IP{net.ParseIP("10.0.0.1"), net.ParseIP("192.168.1.1")},
			wantIP:  net.ParseIP("10.0.0.1"),
			wantErr: nil,
		},
		{
			name:    "Empty list",
			ips:     []net.IP{},
			wantIP:  nil,
			wantErr: ErrNoIPFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIP, gotErr := PreferIPv4Strategy(tt.ips)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("PreferIPv4Strategy() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIP, tt.wantIP) {
				t.Errorf("PreferIPv4Strategy() gotIP = %v, wantIP %v", gotIP, tt.wantIP)
			}
		})
	}
}

func TestPreferPublicIPStrategy(t *testing.T) {
	tests := []struct {
		name     string
		ips      []net.IP
		wantIP   net.IP
		wantErr  error
	}{
		{
			name:    "Prefer public over private",
			ips:     []net.IP{net.ParseIP("192.168.1.1"), net.ParseIP("8.8.8.8")},
			wantIP:  net.ParseIP("8.8.8.8"),
			wantErr: nil,
		},
		{
			name:    "Only private",
			ips:     []net.IP{net.ParseIP("192.168.1.1"), net.ParseIP("10.0.0.1")},
			wantIP:  net.ParseIP("192.168.1.1"), // Returns first private if no public
			wantErr: nil,
		},
		{
			name:    "Only public",
			ips:     []net.IP{net.ParseIP("8.8.8.8"), net.ParseIP("4.4.4.4")},
			wantIP:  net.ParseIP("8.8.8.8"),
			wantErr: nil,
		},
		{
			name:    "Empty list",
			ips:     []net.IP{},
			wantIP:  nil,
			wantErr: ErrNoIPFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIP, gotErr := PreferPublicIPStrategy(tt.ips)
			if !reflect.DeepEqual(gotErr, tt.wantErr) {
				t.Errorf("PreferPublicIPStrategy() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotIP, tt.wantIP) {
				t.Errorf("PreferPublicIPStrategy() gotIP = %v, wantIP %v", gotIP, tt.wantIP)
			}
		})
	}
}

func TestIsPrivateIP(t *testing.T) {
	tests := []struct {
		name string
		ip   string
		want bool
	}{
		{
			name: "Loopback IPv4",
			ip:   "127.0.0.1",
			want: true,
		},
		{
			name: "Loopback IPv6",
			ip:   "::1",
			want: true,
		},
		{
			name: "Private Class A",
			ip:   "10.0.0.1",
			want: true,
		},
		{
			name: "Private Class B",
			ip:   "172.16.0.1",
			want: true,
		},
		{
			name: "Private Class C",
			ip:   "192.168.1.1",
			want: true,
		},
		{
			name: "Public IPv4",
			ip:   "8.8.8.8",
			want: false,
		},
		{
			name: "Public IPv6",
			ip:   "2001:db8::1",
			want: false,
		},
		{
			name: "Link-Local Unicast IPv6",
			ip:   "fe80::1",
			want: true,
		},
		{
			name: "Link-Local Multicast IPv6",
			ip:   "ff02::1",
			want: true,
		},
		{
			name: "Invalid IP",
			ip:   "invalid",
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// net.ParseIP returns nil for invalid IP strings, which is handled by isPrivateIP
			if got := isPrivateIP(net.ParseIP(tt.ip)); got != tt.want {
				t.Errorf("isPrivateIP(%q) = %v, want %v", tt.ip, got, tt.want)
			}
		})
	}
}
