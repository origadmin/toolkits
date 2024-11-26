/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package net

import (
	"net"
	"reflect"
	"testing"
)

// Correctly parses valid IPv4 addresses
func TestParseValidIPv4(t *testing.T) {
	// Arrr! Let's see if this IP be valid!
	ip := ParseIP(net.ParseIP("192.168.1.1"))

	expected := int64(3232235777)
	if ip != expected {
		t.Fatalf("Expected %v, but got %v", expected, ip)
	}
}

// Converts each octet to an integer and combines them into a single int64
func TestConvertOctetsToUint32(t *testing.T) {
	// Ahoy! Let's combine these octets!
	ip := ParseIP(net.ParseIP("10.0.0.1"))

	expected := int64(167772161)
	if ip != expected {
		t.Fatalf("Expected %v, but got %v", expected, ip)
	}
}

// Handles leading zeros in octets correctly
func TestLeadingZerosInOctets(t *testing.T) {
	// Shiver me timbers! Leading zeros ahead!
	ip := ParseIP(net.ParseIP("192.168.1.1"))

	expected := int64(3232235777)
	if ip != expected {
		t.Fatalf("Expected %v, but got %v", expected, ip)
	}
}

// Returns zero for input with fewer or more than four octets
func TestInvalidNumberOfOctets(t *testing.T) {
	// Walk the plank! Invalid number of octets!
	ip := ParseIP(net.ParseIP("192.168.1"))

	if ip != 0 {
		t.Fatalf("Expected 0, but got %v", ip)
	}

	ip = ParseIP(net.ParseIP("192.168.1.1.1"))

	if ip != 0 {
		t.Fatalf("Expected 0, but got %v", ip)
	}
}

// Handles empty string input
func TestEmptyStringInput(t *testing.T) {
	// Empty as Davy Jones' locker!
	ip := ParseIP(net.ParseIP(""))
	if ip != 0 {
		t.Fatalf("Expected 0, but got %v", ip)
	}
}

// Converts a valid int64 IP address to its correct IPv4 string representation
func TestConvertsValidUint32ToIP(t *testing.T) {
	// Arrr! Let's see if this IP be converted correctly!
	ip := int64(3232235777) // 192.168.1.1
	expected := net.ParseIP("192.168.1.1")
	result := ToIP(ip)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
func TestHandlesTypicalIPAddress(t *testing.T) {
	// Ahoy! Let's test the loopback address!
	ip := int64(2130706433) // 127.0.0.1
	expected := net.ParseIP("127.0.0.1")
	result := ToIP(ip)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestProcessesIPWithLeadingZeros(t *testing.T) {
	// Shiver me timbers! Let's check those leading zeros!
	ip := int64(16843009) // 1.1.1.1
	expected := net.ParseIP("1.1.1.1")
	result := ToIP(ip)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestHandlesMinimumUint32Value(t *testing.T) {
	// Avast! Let's see if it handles the lowest of the low!
	ip := int64(0) // 0.0.0.0
	expected := net.ParseIP("0.0.0.0")
	result := ToIP(ip)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}

func TestHandlesMaximumUint32Value(t *testing.T) {
	// Yo ho ho! Let's see if it handles the highest of the high!
	ip := int64(4294967295) // 255.255.255.255
	expected := net.ParseIP("255.255.255.255")
	result := ToIP(ip)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
