package net

import (
	"net"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// mockAddr implements net.Addr for testing purposes.
type mockAddr struct {
	ip net.IP
}

func (m *mockAddr) Network() string { return "ip+net" }
func (m *mockAddr) String() string  { return m.ip.String() }

// mockInterfaceWithAddrs implements InterfaceWithAddrs for testing.
type mockInterfaceWithAddrs struct {
	net.Interface
	mockAddrs []net.Addr
}

func (m *mockInterfaceWithAddrs) GetInterface() net.Interface { return m.Interface }
func (m *mockInterfaceWithAddrs) GetAddrs() ([]net.Addr, error) { return m.mockAddrs, nil }

// newMockInterfaceWithAddrs is a helper to create a mock InterfaceWithAddrs.
func newMockInterfaceWithAddrs(name string, index int, flags net.Flags, ipStrs ...string) InterfaceWithAddrs {
	addrs := make([]net.Addr, len(ipStrs))
	for i, ipStr := range ipStrs {
		ip := net.ParseIP(ipStr)
		if ip == nil {
			panic("Invalid IP string in mock setup: " + ipStr)
		}
		addrs[i] = &net.IPNet{IP: ip, Mask: net.CIDRMask(24, 32)} // Use a common mask for simplicity
	}

	return &mockInterfaceWithAddrs{
		Interface: net.Interface{
			Index: index,
			MTU:   1500,
			Name:  name,
			HardwareAddr: nil,
			Flags: flags,
		},
		mockAddrs: addrs,
	}
}

// mockNetworkInterfaceProvider implements networkInterfaceProvider for testing.
type mockNetworkInterfaceProvider struct {
	interfaces []InterfaceWithAddrs
}

func (m *mockNetworkInterfaceProvider) Interfaces() ([]InterfaceWithAddrs, error) {
	return m.interfaces, nil
}

// ipInCIDR checks if the given IP is within any of the provided CIDR strings.
func ipInCIDR(t *testing.T, ipStr string, cidrStrs []string) bool {
	if ipStr == "" {
		return false
	}
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return false
	}

	for _, cidrStr := range cidrStrs {
		_, ipNet, err := net.ParseCIDR(cidrStr)
		if err != nil {
			t.Errorf("Invalid CIDR string in test: %s", cidrStr)
			continue
		}
		if ipNet.Contains(ip) {
			return true
		}
	}
	return false
}

// TestGetHostAddrFromEnvironmentVariable is updated to use GetHostAddr (real provider).
func TestGetHostAddrFromEnvironmentVariable(t *testing.T) {
	// Setup
	const envVarName = "TEST_HOST_IP"
	const expectedIP = "192.168.1.100"
	os.Setenv(envVarName, expectedIP)

	// Call function with environment variable option using the real provider
	result := GetHostAddr(WithEnvVar(envVarName))

	// Assert result matches expected IP
	assert.Equal(t, expectedIP, result, "Expected IP %s from environment variable, got %s", expectedIP, result)
}

// TestGetHostAddrReturnsEmptyWhenNoIPFound is updated to use GetHostAddr (real provider).
func TestGetHostAddrReturnsEmptyWhenNoIPFound(t *testing.T) {
	// Create a test configuration that should fail to find any IP
	// - Use non-existent environment variable
	// - Use non-matching interface patterns
	// - Disable fallback

	// Ensure environment variable is empty
	os.Unsetenv("NON_EXISTENT_ENV_VAR")

	// Call function with options that should result in no IP found using the real provider
	result := GetHostAddr(
		WithEnvVar("NON_EXISTENT_ENV_VAR"),
		WithFallback(false),
	)

	// Assert result is empty string
	assert.Empty(t, result, "Expected empty string when no IP can be found, got %s", result)
}

func TestCIDRFilter(t *testing.T) {
	// Define a common set of mock interfaces for all tests
	mockInterfaces := []InterfaceWithAddrs{
		newMockInterfaceWithAddrs("lo0", 1, net.FlagUp|net.FlagLoopback, "127.0.0.1"),
		newMockInterfaceWithAddrs("eth0", 2, net.FlagUp, "192.168.28.81", "fe80::1"),
		newMockInterfaceWithAddrs("eth1", 3, net.FlagUp, "10.1.0.30"),
		newMockInterfaceWithAddrs("eth2", 4, 0), // Down interface
		newMockInterfaceWithAddrs("docker0", 5, net.FlagUp, "172.17.0.1"),
	}

	// Create a mock provider with these interfaces
	mockProvider := &mockNetworkInterfaceProvider{interfaces: mockInterfaces}

	tests := []struct {
		name         string
		cidrs        []string
		expectedCIDR []string // Expected IP should be within this CIDR, or empty if no IP expected
		expectEmpty  bool     // True if an empty string is expected
	}{
		{
			name:         "Should find IP in 192.168.28.0/24",
			cidrs:        []string{"192.168.28.0/24"},
			expectedCIDR: []string{"192.168.28.0/24"},
			expectEmpty:  false,
		},
		{
			name:         "Should find IP in 10.0.0.0/8",
			cidrs:        []string{"10.0.0.0/8"},
			expectedCIDR: []string{"10.0.0.0/8"},
			expectEmpty:  false,
		},
		{
			name:         "Should find IP in 172.16.0.0/16 or 192.168.0.0/16",
			cidrs:        []string{"172.16.0.0/16", "192.168.0.0/16"},
			expectedCIDR: []string{"172.16.0.0/16", "192.168.0.0/16"},
			expectEmpty:  false,
		},
		{
			name:         "Should not find IP in non-existent CIDR",
			cidrs:        []string{"1.1.1.0/24"},
			expectedCIDR: nil,
			expectEmpty:  true,
		},
		{
			name:         "Should not find IP if no matching CIDR",
			cidrs:        []string{"192.168.1.0/24"},
			expectedCIDR: nil,
			expectEmpty:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ip := HostAddr(
				mockProvider,
				WithCIDRFilters(tt.cidrs),
				WithFallback(false),
			)

			if tt.expectEmpty {
				assert.Empty(t, ip, "Expected empty IP but got %s", ip)
			} else {
				assert.NotEmpty(t, ip, "Expected a non-empty IP but got empty")
				if ip != "" {
					assert.True(t, ipInCIDR(t, ip, tt.expectedCIDR), "IP %s not in expected CIDRs %v", ip, tt.expectedCIDR)
				}
			}
		})
	}
}

// TestGetHostAddrFromEnvironmentVariable and TestGetHostAddrReturnsEmptyWhenNoIPFound
// also need to be updated to use the new HostAddr signature.
func TestGetHostAddrFromEnvironmentVariableWithMock(t *testing.T) {
	// Setup
	const envVarName = "TEST_HOST_IP"
	const expectedIP = "192.168.1.100"
	os.Setenv(envVarName, expectedIP)

	// Create a dummy mock provider, as env var takes precedence
	mockProvider := &mockNetworkInterfaceProvider{interfaces: []InterfaceWithAddrs{}}

	// Call function with environment variable option
	result := HostAddr(mockProvider, WithEnvVar(envVarName))

	// Assert result matches expected IP
	assert.Equal(t, expectedIP, result, "Expected IP %s from environment variable, got %s", expectedIP, result)
}

func TestGetHostAddrReturnsEmptyWhenNoIPFoundWithMock(t *testing.T) {
	// Create a test configuration that should fail to find any IP
	// - Use non-existent environment variable
	// - Use non-matching interface patterns
	// - Disable fallback

	// Ensure environment variable is empty
	os.Unsetenv("NON_EXISTENT_ENV_VAR")

	// Create a mock provider with no valid interfaces
	mockProvider := &mockNetworkInterfaceProvider{interfaces: []InterfaceWithAddrs{
		newMockInterfaceWithAddrs("lo0", 1, net.FlagUp|net.FlagLoopback, "127.0.0.1"), // Loopback is not global unicast
		newMockInterfaceWithAddrs("eth0", 2, 0),                                     // Down interface
		newMockInterfaceWithAddrs("eth1", 3, net.FlagUp, "169.254.0.1"),             // Link-local is not global unicast
	}}

	// Call function with options that should result in no IP found
	result := HostAddr(
		mockProvider,
		WithEnvVar("NON_EXISTENT_ENV_VAR"),
		WithFallback(false),
	)

	// Assert result is empty string
	assert.Empty(t, result, "Expected empty string when no IP can be found, got %s", result)
}
