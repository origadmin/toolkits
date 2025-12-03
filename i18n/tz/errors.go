/*
 * Copyright (c) 2024 OrigAdmin. All rights reserved.
 */

// Package tz implements the functions, types, and interfaces for the module.
package tz

import "errors"

// Common timezone errors
var (
	ErrTimeZoneNotFound        = errors.New("timezone not found")
	ErrInvalidTimeZone         = errors.New("invalid timezone format")
	ErrConversionFailed        = errors.New("timezone conversion failed")
	ErrLocationNotFound        = errors.New("location not found")
	ErrWindowsTimeZoneNotFound  = errors.New("Windows timezone not found")
	ErrIANATimeZoneNotFound    = errors.New("IANA timezone not found")
	ErrNotImplemented          = errors.New("feature not implemented")
	ErrCacheExpired            = errors.New("cache expired")
	ErrDetectionFailed         = errors.New("timezone detection failed")
)