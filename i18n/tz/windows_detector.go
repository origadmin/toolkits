/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package tz implements enhanced Windows timezone support with caching and fallback strategies.
package tz

import (
	"os/exec"
	"runtime"
	"sync"
	"time"
)

// WindowsTimeZoneDetector provides enhanced Windows timezone detection
type WindowsTimeZoneDetector struct {
	cache     map[string]string
	cacheTime time.Time
	ttl       time.Duration
	mutex     sync.RWMutex
}

// NewWindowsTimeZoneDetector creates a new detector with configurable TTL
func NewWindowsTimeZoneDetector(ttl time.Duration) *WindowsTimeZoneDetector {
	return &WindowsTimeZoneDetector{
		cache: make(map[string]string),
		ttl:   ttl,
	}
}

// GetSystemTimezone safely gets the system timezone with caching
func (d *WindowsTimeZoneDetector) GetSystemTimezone() (string, error) {
	d.mutex.RLock()
	if cached, exists := d.cache["system"]; exists && time.Since(d.cacheTime) < d.ttl {
		d.mutex.RUnlock()
		return cached, nil
	}
	d.mutex.RUnlock()

	// Acquire write lock for cache update
	d.mutex.Lock()
	defer d.mutex.Unlock()

	// Double-check after acquiring write lock
	if cached, exists := d.cache["system"]; exists && time.Since(d.cacheTime) < d.ttl {
		return cached, nil
	}

	timezone, err := d.detectSystemTimezone()
	if err != nil {
		return "", err
	}

	d.cache["system"] = timezone
	d.cacheTime = time.Now()
	return timezone, nil
}

// detectSystemTimezone implements the actual detection logic
func (d *WindowsTimeZoneDetector) detectSystemTimezone() (string, error) {
	if runtime.GOOS != "windows" {
		// Fallback for non-Windows systems
		return time.Now().Location().String(), nil
	}

	// Try tzutil first
	if timezone, err := d.getTimezoneFromTzutil(); err == nil {
		return timezone, nil
	}

	// Fallback to registry
	if timezone, err := d.getTimezoneFromRegistry(); err == nil {
		return timezone, nil
	}

	// Final fallback to Go's built-in detection
	return time.Now().Location().String(), nil
}

// getTimezoneFromTzutil uses Windows tzutil command
func (d *WindowsTimeZoneDetector) getTimezoneFromTzutil() (string, error) {
	path, err := exec.LookPath("tzutil")
	if err != nil {
		return "", &TimezoneDetectionError{
			Method: "tzutil",
			Cause:  err,
		}
	}

	cmd := exec.Command(path, "/g")
	output, err := cmd.Output()
	if err != nil {
		return "", &TimezoneDetectionError{
			Method: "tzutil",
			Cause:  err,
		}
	}

	zone := string(output)
	if len(zone) > 0 && zone[len(zone)-1] == '\n' {
		zone = zone[:len(zone)-1]
	}
	if len(zone) > 0 && zone[len(zone)-1] == '\r' {
		zone = zone[:len(zone)-1]
	}

	return zone, nil
}

// getTimezoneFromRegistry reads timezone from Windows registry
func (d *WindowsTimeZoneDetector) getTimezoneFromRegistry() (string, error) {
	// This would require Windows-specific registry access
	// For now, return an error to indicate this method is not implemented
	return "", &TimezoneDetectionError{
		Method: "registry",
		Cause:  ErrNotImplemented,
	}
}

// TimezoneDetectionError represents errors during timezone detection
type TimezoneDetectionError struct {
	Method string
	Cause  error
}

func (e *TimezoneDetectionError) Error() string {
	return "timezone detection failed using " + e.Method + ": " + e.Cause.Error()
}

func (e *TimezoneDetectionError) Unwrap() error {
	return e.Cause
}

// Global detector instance with 5-minute TTL
var globalDetector = NewWindowsTimeZoneDetector(5 * time.Minute)

// GetSystemTimezone gets the current system timezone using enhanced detection
func GetSystemTimezone() (string, error) {
	return globalDetector.GetSystemTimezone()
}