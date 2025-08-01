package net

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Returns IP address from environment variable when specified and available
func TestGetHostAddrFromEnvironmentVariable(t *testing.T) {
	// Setup
	const envVarName = "TEST_HOST_IP"
	const expectedIP = "192.168.1.100"
	os.Setenv(envVarName, expectedIP)

	// Call function with environment variable option
	result := HostAddr(WithEnvVar(envVarName))

	// Assert result matches expected IP
	if result != expectedIP {
		t.Errorf("Expected IP %s from environment variable, got %s", expectedIP, result)
	}
}

// Returns empty string when no IP address can be found through any method
func TestGetHostAddrReturnsEmptyWhenNoIPFound(t *testing.T) {
	// Create a test configuration that should fail to find any IP
	// - Use non-existent environment variable
	// - Use non-matching interface patterns
	// - Disable fallback

	// Ensure environment variable is empty
	os.Unsetenv("NON_EXISTENT_ENV_VAR")

	// Call function with options that should result in no IP found
	result := HostAddr(
		WithEnvVar("NON_EXISTENT_ENV_VAR"),
		WithFallback(false),
	)

	// Assert result is empty string
	if result != "" {
		t.Errorf("Expected empty string when no IP can be found, got %s", result)
	}
}

func TestCIDRFilter(t *testing.T) {
	// 准备测试接口
	tests := []struct {
		cidrs    []string
		expected string
	}{
		{
			cidrs:    []string{"192.168.28.0/24"},
			expected: "192.168.28.81",
		},
		{
			cidrs:    []string{"10.0.0.0/8"},
			expected: "",
		},
	}

	for _, tt := range tests {
		ip := HostAddr(
			WithCIDRFilters(tt.cidrs),
			WithFallback(false),
		)
		assert.Equal(t, tt.expected, ip)
	}
}
