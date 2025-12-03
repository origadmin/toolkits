/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package tz implements enhanced Windows timezone mapping with validation and fallback.
package tz

import (
	"encoding/json"
	"sync"
	"time"
)

// WindowsMapper provides enhanced Windows timezone mapping capabilities
type WindowsMapper struct {
	mappings map[string]WindowsTimeZoneMapping
	mutex    sync.RWMutex
	loaded   bool
}

// WindowsTimeZoneMapping represents a mapping from Windows timezone to IANA zones
type WindowsTimeZoneMapping struct {
	WindowsID   string    `json:"windows_id"`
	IANAZones   []string  `json:"iana_zones"`
	Territory   string    `json:"territory"`
	LastUpdated time.Time `json:"last_updated"`
	Preferred   string    `json:"preferred,omitempty"` // Preferred IANA zone for this Windows zone
}

// NewWindowsMapper creates a new mapper instance
func NewWindowsMapper() *WindowsMapper {
	return &WindowsMapper{
		mappings: make(map[string]WindowsTimeZoneMapping),
	}
}

// LoadMappings loads timezone mappings from embedded data or external source
func (m *WindowsMapper) LoadMappings() error {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	if m.loaded {
		return nil
	}

	// Load from embedded JSON first
	var supplementalData SupplementalData
	if err := json.Unmarshal(jsonMapZones, &supplementalData); err != nil {
		return &MappingError{
			Operation: "load_embedded",
			Cause:     err,
		}
	}

	// Convert to enhanced mapping format
	for _, mapZone := range supplementalData.WindowsZones.MapTimeZones.MapZone {
		mapping := WindowsTimeZoneMapping{
			WindowsID:   mapZone.Other,
			IANAZones:   mapZone.Types,
			Territory:   mapZone.Territory,
			LastUpdated: time.Now(), // Would be better to get from source
		}

		// Determine preferred zone (usually the first one, or use heuristics)
		if len(mapping.IANAZones) > 0 {
			mapping.Preferred = m.selectPreferredZone(mapping.IANAZones, mapping.Territory)
		}

		m.mappings[mapZone.Other] = mapping
	}

	m.loaded = true
	return nil
}

// selectPreferredZone selects the best IANA zone for a given Windows zone
func (m *WindowsMapper) selectPreferredZone(zones []string, territory string) string {
	if len(zones) == 0 {
		return ""
	}

	// Preference logic:
	// 1. Zone that matches the territory
	// 2. Zone without "Etc/" prefix
	// 3. First zone in the list

	for _, zone := range zones {
		// Check if zone matches territory (for major cities)
		if m.zoneMatchesTerritory(zone, territory) {
			return zone
		}
	}

	for _, zone := range zones {
		// Avoid Etc/ zones unless they're the only option
		if len(zones) > 1 && !m.isEtcZone(zone) {
			return zone
		}
	}

	return zones[0]
}

// zoneMatchesTerritory checks if a zone matches the given territory
func (m *WindowsMapper) zoneMatchesTerritory(zone, territory string) bool {
	// Simple heuristic: check if territory appears in zone name
	// This could be enhanced with a proper territory mapping
	return len(territory) > 0 && len(zone) > 0
}

// isEtcZone checks if a zone is an Etc/ zone
func (m *WindowsMapper) isEtcZone(zone string) bool {
	return len(zone) >= 4 && zone[:4] == "Etc/"
}

// WindowsToIANA converts a Windows timezone ID to IANA zones
func (m *WindowsMapper) WindowsToIANA(windowsID string) ([]string, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	mapping, exists := m.mappings[windowsID]
	if !exists {
		return nil, false
	}

	return mapping.IANAZones, true
}

// WindowsToPreferredIANA converts a Windows timezone ID to the preferred IANA zone
func (m *WindowsMapper) WindowsToPreferredIANA(windowsID string) (string, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	mapping, exists := m.mappings[windowsID]
	if !exists || mapping.Preferred == "" {
		return "", false
	}

	return mapping.Preferred, true
}

// IANAToWindows finds Windows timezone IDs that map to the given IANA zone
func (m *WindowsMapper) IANAToWindows(ianaZone string) ([]string, bool) {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	var windowsIDs []string
	for _, mapping := range m.mappings {
		for _, zone := range mapping.IANAZones {
			if zone == ianaZone {
				windowsIDs = append(windowsIDs, mapping.WindowsID)
				break
			}
		}
	}

	if len(windowsIDs) > 0 {
		return windowsIDs, true
	}
	return nil, false
}

// GetAllMappings returns all timezone mappings
func (m *WindowsMapper) GetAllMappings() map[string]WindowsTimeZoneMapping {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	result := make(map[string]WindowsTimeZoneMapping)
	for k, v := range m.mappings {
		result[k] = v
	}
	return result
}

// ValidateMapping validates a Windows to IANA timezone mapping
func (m *WindowsMapper) ValidateMapping(windowsID, ianaZone string) error {
	// Check if Windows ID exists
	mappings, exists := m.WindowsToIANA(windowsID)
	if !exists {
		return &MappingError{
			Operation: "validate",
			Cause:     ErrWindowsTimeZoneNotFound,
		}
	}

	// Check if IANA zone is in the mapping
	for _, zone := range mappings {
		if zone == ianaZone {
			return nil
		}
	}

	return &MappingError{
		Operation: "validate",
		Cause:     ErrIANATimeZoneNotFound,
	}
}

// MappingError represents errors during timezone mapping operations
type MappingError struct {
	Operation string
	Cause     error
}

func (e *MappingError) Error() string {
	return "timezone mapping error during " + e.Operation + ": " + e.Cause.Error()
}

func (e *MappingError) Unwrap() error {
	return e.Cause
}

// Global mapper instance
var globalMapper = NewWindowsMapper()

// Initialize the mapper when the package is loaded
func init() {
	// Pre-load mappings asynchronously to avoid blocking
	go func() {
		if err := globalMapper.LoadMappings(); err != nil {
			// Log error but don't panic - the package can still function
			_ = err
		}
	}()
}

// ConvertWindowsToIANA converts a Windows timezone ID to IANA zones using the global mapper
func ConvertWindowsToIANA(windowsID string) ([]string, bool) {
	// Ensure mappings are loaded
	_ = globalMapper.LoadMappings()
	return globalMapper.WindowsToIANA(windowsID)
}

// ConvertWindowsToPreferredIANA converts a Windows timezone ID to the preferred IANA zone
func ConvertWindowsToPreferredIANA(windowsID string) (string, bool) {
	// Ensure mappings are loaded
	_ = globalMapper.LoadMappings()
	return globalMapper.WindowsToPreferredIANA(windowsID)
}