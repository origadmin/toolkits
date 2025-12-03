/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package tz implements functions, types, and interfaces for module.
package tz

import (
	"testing"
)

func TestGenerateJSON(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "generate json files",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := GenerateJSON(); (err != nil) != tt.wantErr {
				t.Errorf("GenerateJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCountriesFromCSV(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "load countries from csv",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := CountriesFromCSV("data/country.csv")
			if (err != nil) != tt.wantErr {
				t.Errorf("CountriesFromCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTimeZonesFromCSV(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "load timezones from csv",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := TimeZonesFromCSV("data/time_zone.csv")
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeZonesFromCSV() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}