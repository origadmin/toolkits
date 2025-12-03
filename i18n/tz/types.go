/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package tz implements the functions, types, and interfaces for the module.
package tz

import (
	"time"
)

// TimeZoneInfo provides detailed timezone information
type TimeZoneInfo struct {
	Name         string    `json:"name"`
	Country      string    `json:"country"`
	Offset       int64     `json:"offset"`       // in seconds
	OffsetString string    `json:"offset_string"` // "+08:00" format
	IsDST        bool      `json:"is_dst"`
	Abbreviation string    `json:"abbreviation"`
	Location     *time.Location `json:"-"`
}

// TimeZoneFinder interface for finding timezones
type TimeZoneFinder interface {
	FindByZoneName(name string) (*TimeZone, bool)
	FindByCountryCode(code string) ([]*TimeZone, bool)
	FindByAbbreviation(abbrev string) ([]*TimeZone, bool)
	GetCurrentLocation() (*time.Location, error)
	GetAllTimeZones() []*TimeZone
}

// TimeZoneConverter interface for timezone conversion
type TimeZoneConverter interface {
	Convert(t time.Time, fromZone, toZone string) (time.Time, error)
	GetOffset(zone string) (int, error)
	IsDST(t time.Time, zone string) (bool, error)
}