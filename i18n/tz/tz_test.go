/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package tz implements functions, types, and interfaces for module.
package tz

import (
	"testing"
	"time"
)

func TestGetLocation(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "get current location",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetLocation()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLocation() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetLocationString(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "get location string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetLocationString()
			if got == "" {
				t.Errorf("GetLocationString() should not return empty string")
			}
		})
	}
}

func TestFindTimeZoneByZoneName(t *testing.T) {
	// Initialize data first
	initialize()
	
	tests := []struct {
		name    string
		zone    string
		want    bool
	}{
		{
			name: "find existing timezone",
			zone: "Asia/Shanghai",
			want: true,
		},
		{
			name: "find non-existing timezone",
			zone: "Invalid/Timezone",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := FindTimeZoneByZoneName(tt.zone)
			if got != tt.want {
				t.Errorf("FindTimeZoneByZoneName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindTimeZonesByCountryCode(t *testing.T) {
	// Initialize data first
	initialize()
	
	tests := []struct {
		name    string
		code    string
		want    bool
	}{
		{
			name: "find existing country",
			code: "CN",
			want: true,
		},
		{
			name: "find non-existing country",
			code: "XX",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, got := FindTimeZonesByCountryCode(tt.code)
			if got != tt.want {
				t.Errorf("FindTimeZonesByCountryCode() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvertTime(t *testing.T) {
	tests := []struct {
		name      string
		fromZone  string
		toZone    string
		wantErr   bool
	}{
		{
			name:     "convert between valid timezones",
			fromZone: "Asia/Shanghai",
			toZone:   "America/New_York",
			wantErr:  false,
		},
		{
			name:     "convert with invalid from timezone",
			fromZone: "Invalid/Timezone",
			toZone:   "America/New_York",
			wantErr:  true,
		},
		{
			name:     "convert with invalid to timezone",
			fromZone: "Asia/Shanghai",
			toZone:   "Invalid/Timezone",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now()
			_, err := ConvertTime(now, tt.fromZone, tt.toZone)
			if (err != nil) != tt.wantErr {
				t.Errorf("ConvertTime() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetTimeZoneInfo(t *testing.T) {
	tests := []struct {
		name    string
		zone    string
		wantErr bool
	}{
		{
			name:    "get info for valid timezone",
			zone:    "Asia/Shanghai",
			wantErr: false,
		},
		{
			name:    "get info for invalid timezone",
			zone:    "Invalid/Timezone",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetTimeZoneInfo(tt.zone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTimeZoneInfo() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGetOffset(t *testing.T) {
	// Initialize data first
	initialize()
	
	tests := []struct {
		name    string
		zone    string
		wantErr bool
	}{
		{
			name:    "get offset for valid timezone",
			zone:    "Asia/Shanghai",
			wantErr: false,
		},
		{
			name:    "get offset for invalid timezone",
			zone:    "Invalid/Timezone",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetOffset(tt.zone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetOffset() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestIsDST(t *testing.T) {
	tests := []struct {
		name    string
		zone    string
		wantErr bool
	}{
		{
			name:    "check DST for valid timezone",
			zone:    "America/New_York",
			wantErr: false,
		},
		{
			name:    "check DST for invalid timezone",
			zone:    "Invalid/Timezone",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now := time.Now()
			_, err := IsDST(now, tt.zone)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDST() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}