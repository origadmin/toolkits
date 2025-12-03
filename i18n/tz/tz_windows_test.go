/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package tz implements the functions, types, and interfaces for the module.
package tz

import (
	"encoding/json"
	"os"
	"testing"
)

func TestWindowsZonesFromXML(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				filePath: "windows/windowsZones.xml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := WindowsZonesFromXMLToJSON(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("WindowsZonesFromXML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}

func TestWindowsZonesFromJSON(t *testing.T) {
	type args struct {
		filePath string
	}
	tests := []struct {
		name    string
		args    args
		want    WindowsZones
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				filePath: "windows/windows_zones.json",
			},
			wantErr: false, // Just check that it loads without error
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := WindowsZonesFromJSON(tt.args.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("WindowsZonesFromJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			// Just verify it loads and has expected structure
			if got.MapTimeZones.OtherVersion == "" {
				t.Errorf("WindowsZonesFromJSON() expected OtherVersion to be set")
			}
		})
	}
}

func TestSaveNewTimeZonesToJSON(t *testing.T) {
	var supplementalData SupplementalData
	file, err := os.Open("windows/windows_zones.json")
	if err != nil {
		t.Errorf("WindowsZonesFromJSON() error = %v", err)
		return
	}
	defer func() {
		if closeErr := file.Close(); closeErr != nil {
			t.Errorf("File close error = %v", closeErr)
		}
	}()
	if err := json.NewDecoder(file).Decode(&supplementalData); err != nil {
		t.Errorf("WindowsZonesFromJSON() error = %v", err)
		return
	}
	type args struct {
		timeZones []TimeZone
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				timeZones: FixTimeZoneFromWindowsZones(supplementalData.WindowsZones),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveNewTimeZonesToJSON(tt.args.timeZones); (err != nil) != tt.wantErr {
				t.Errorf("SaveNewTimeZonesToJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestTimeZoneToTimeZoneMap(t *testing.T) {
	type args struct {
		mapfile string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				mapfile: "map_zones.json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			TimeZoneToTimeZoneMap(tt.args.mapfile)
		})
	}
}
