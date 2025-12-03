/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package tz implements the functions, types, and interfaces for the module.
package tz

import (
	_ "embed"
	"encoding/json"
	"sync"
	"time"
)

const (
	defaultTimeZone = "Asia/Shanghai"
)

var (
	countries []Country
	timeZones []*TimeZone
	
	// Cache for performance
	timeZoneCache = make(map[string]*time.Location)
	cacheMutex    sync.RWMutex
	
	// Lazy loading flag
	initialized bool
	initMutex   sync.Once
)

// initialize data structures with lazy loading
func initialize() {
	initMutex.Do(func() {
		// Load countries
		var countryList []Country
		if err := json.Unmarshal(jsonCountries, &countryList); err == nil {
			countries = countryList
		}
		
		// Load and convert timezones
		var rawTimeZones []RawTimeZone
		if err := json.Unmarshal(jsonTimeZones, &rawTimeZones); err == nil {
			timeZones = make([]*TimeZone, 0, len(rawTimeZones))
			seen := make(map[string]bool) // Avoid duplicates
			
			for _, raw := range rawTimeZones {
				key := raw.ZoneName + ":" + raw.CountryCode
				if !seen[key] {
					dst := int64(0)
					if raw.Dst == 1 {
						dst = 1
					}
					tz := &TimeZone{
						ZoneName:     raw.ZoneName,
						CountryCode:  raw.CountryCode,
						Abbreviation: raw.Abbreviation,
						GmtOffset:    raw.GmtOffset,
						Dst:          dst,
					}
					
					timeZones = append(timeZones, tz)
					seen[key] = true
				}
			}
		}
		
		initialized = true
	})
}

// GetCountries returns all countries
func GetCountries() []Country {
	initialize()
	return countries
}

// GetTimeZones returns all timezones
func GetTimeZones() []*TimeZone {
	initialize()
	return timeZones
}

// GetLocation returns the current system timezone
func GetLocation() (*time.Location, error) {
	timezoneName := location()
	
	// Try to load the timezone location
	if loc, err := time.LoadLocation(timezoneName); err == nil {
		return loc, nil
	}
	
	// Fallback to default timezone
	return time.LoadLocation(defaultTimeZone)
}

// GetLocationString returns the current system timezone as string
func GetLocationString() string {
	return location()
}

// FindTimeZoneByZoneName finds timezone by zone name
func FindTimeZoneByZoneName(name string) (*TimeZone, bool) {
	initialize()
	for _, tz := range timeZones {
		if tz.ZoneName == name {
			return tz, true
		}
	}
	return nil, false
}

// FindTimeZonesByCountryCode finds timezones by country code
func FindTimeZonesByCountryCode(code string) ([]*TimeZone, bool) {
	initialize()
	var result []*TimeZone
	for _, tz := range timeZones {
		if tz.CountryCode == code {
			result = append(result, tz)
		}
	}
	if len(result) > 0 {
		return result, true
	}
	return nil, false
}

// FindTimeZonesByAbbreviation finds timezones by abbreviation
func FindTimeZonesByAbbreviation(abbrev string) ([]*TimeZone, bool) {
	initialize()
	var result []*TimeZone
	for _, tz := range timeZones {
		if tz.Abbreviation == abbrev {
			result = append(result, tz)
		}
	}
	if len(result) > 0 {
		return result, true
	}
	return nil, false
}

// GetTimeZoneInfo returns detailed timezone information
func GetTimeZoneInfo(name string) (*TimeZoneInfo, error) {
	tz, found := FindTimeZoneByZoneName(name)
	if !found {
		return nil, ErrTimeZoneNotFound
	}
	
	loc, err := getTimeZoneLocation(name)
	if err != nil {
		return nil, err
	}
	
	now := time.Now().In(loc)
	offset := now.Format("-07:00")
	
	return &TimeZoneInfo{
		Name:         tz.ZoneName,
		Country:      tz.CountryCode,
		Offset:       tz.GmtOffset,
		OffsetString: offset,
		IsDST:        now.IsDST(),
		Abbreviation: tz.Abbreviation,
		Location:     loc,
	}, nil
}

// ConvertTime converts time from one timezone to another
func ConvertTime(t time.Time, fromZone, toZone string) (time.Time, error) {
	fromLoc, err := getTimeZoneLocation(fromZone)
	if err != nil {
		return time.Time{}, err
	}
	
	toLoc, err := getTimeZoneLocation(toZone)
	if err != nil {
		return time.Time{}, err
	}
	
	// Convert to UTC first, then to target timezone
	utc := t.In(fromLoc).UTC()
	return utc.In(toLoc), nil
}

// GetOffset returns timezone offset in seconds
func GetOffset(zone string) (int64, error) {
	tz, found := FindTimeZoneByZoneName(zone)
	if !found {
		return 0, ErrTimeZoneNotFound
	}
	return tz.GmtOffset, nil
}

// IsDST checks if given time is DST in the specified timezone
func IsDST(t time.Time, zone string) (bool, error) {
	loc, err := getTimeZoneLocation(zone)
	if err != nil {
		return false, err
	}
	return t.In(loc).IsDST(), nil
}

// getTimeZoneLocation gets cached timezone location
func getTimeZoneLocation(name string) (*time.Location, error) {
	cacheMutex.RLock()
	if loc, exists := timeZoneCache[name]; exists {
		cacheMutex.RUnlock()
		return loc, nil
	}
	cacheMutex.RUnlock()
	
	loc, err := time.LoadLocation(name)
	if err != nil {
		return nil, err
	}
	
	cacheMutex.Lock()
	timeZoneCache[name] = loc
	cacheMutex.Unlock()
	
	return loc, nil
}

// contains checks if string exists in slice
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}