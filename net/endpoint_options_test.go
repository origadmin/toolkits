/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

package net

import (
	"net"
	"reflect"
	"testing"
)

func TestDefaultEndpointOptions(t *testing.T) {
	opts := DefaultOptions()

	if opts.EnvVar != "SERVICE_IP" {
		t.Errorf("DefaultOptions() EnvVar got %q, want %q", opts.EnvVar, "SERVICE_IP")
	}

	if opts.IPStrategy == nil {
		t.Errorf("DefaultOptions() IPStrategy is nil, want a default strategy")
	}

	if !opts.PreferPublicIP {
		t.Errorf("DefaultOptions() PreferPublicIP got %v, want %v", opts.PreferPublicIP, true)
	}

	if !opts.CloudMetadataEnabled {
		t.Errorf("DefaultOptions() CloudMetadataEnabled got %v, want %v", opts.CloudMetadataEnabled, true)
	}

	if opts.HostIP != "" {
		t.Errorf("DefaultOptions() HostIP got %q, want empty string", opts.HostIP)
	}

	if opts.EndpointFunc != nil {
		t.Errorf("DefaultOptions() EndpointFunc got non-nil, want nil")
	}

	if len(opts.AllowedInterfaces) != 0 {
		t.Errorf("DefaultOptions() AllowedInterfaces got %v, want empty slice", opts.AllowedInterfaces)
	}
}

func TestEndpointOptionsWithEnvVar(t *testing.T) {
	opts := &EndpointOptions{}
	opts.WithEnvVar("MY_SERVICE_IP")

	if opts.EnvVar != "MY_SERVICE_IP" {
		t.Errorf("WithEnvVar() got %q, want %q", opts.EnvVar, "MY_SERVICE_IP")
	}
}

func TestEndpointOptionsWithHostIP(t *testing.T) {
	opts := &EndpointOptions{}
	opts.WithHostIP("192.168.1.100")

	if opts.HostIP != "192.168.1.100" {
		t.Errorf("WithHostIP() got %q, want %q", opts.HostIP, "192.168.1.100")
	}
}

func TestEndpointOptionsWithIPStrategy(t *testing.T) {
	opts := &EndpointOptions{}
	customStrategy := func(ips []net.IP) (net.IP, error) { return nil, nil }
	opts.WithIPStrategy(customStrategy)

	if reflect.ValueOf(opts.IPStrategy).Pointer() != reflect.ValueOf(customStrategy).Pointer() {
		t.Errorf("WithIPStrategy() did not set the correct strategy")
	}
}

func TestEndpointOptionsWithCloudMetadata(t *testing.T) {
	opts := &EndpointOptions{CloudMetadataEnabled: true}
	opts.WithCloudMetadata(false)

	if opts.CloudMetadataEnabled {
		t.Errorf("WithCloudMetadata(false) got %v, want %v", opts.CloudMetadataEnabled, false)
	}

	opts.WithCloudMetadata(true)
	if !opts.CloudMetadataEnabled {
		t.Errorf("WithCloudMetadata(true) got %v, want %v", opts.CloudMetadataEnabled, true)
	}
}
